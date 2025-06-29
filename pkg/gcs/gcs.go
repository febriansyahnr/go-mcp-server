package gcs

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type GCSService interface {
	Close() error

	SetClient(ctx context.Context) (*storage.Client, error)
	SetBucketWriter(ctx context.Context, objectName string) (*storage.Writer, error)
	UploadFileToGCS(ctx context.Context, objectName, srcFile string, sync bool, ttl *time.Duration) (*Response, error)

	// Improvement
	UploadFileFromMultipart(ctx context.Context, objectName string, file *multipart.FileHeader, sync bool) (*UploadMultipart, error)
	CreateSignedURL(ctx context.Context, object string, expires time.Duration) (url string, err error)
	ReadAll(ctx context.Context, bucket, object string) ([]byte, error)
	UploadFile(ctx context.Context, objectName string, src io.Reader, sync bool) (*UploadMultipart, error)
}

type gcsService struct {
	config Config

	gcsClient *storage.Client
	gcsWriter *storage.Writer
}

const publicURL = "https://storage.googleapis.com"

const (
	PrivateCache = "private"

	ErrStorageNewClientFormat = "storage.NewClient: %v"
	ErrBucketSignedURLFormat  = "Bucket(%q).SignedURL: %w"
	ErrCopyFormat             = "ERROR GCS: io.Copy: %v"
)

func NewGCSService(config Config) GCSService {
	once.Do(func() {
		client, rootErr = storage.NewClient(context.Background())
	})

	return &gcsService{
		config: config,
	}
}

func (g *gcsService) SetClient(ctx context.Context) (*storage.Client, error) {
	segment := newrelic.
		FromContext(ctx).
		StartSegment("pkg/gcs/InitClient")
	defer segment.End()

	if g.gcsClient != nil {
		return g.gcsClient, nil
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrStorageNewClientFormat, err)
	}

	// Assign gcsClient
	g.gcsClient = client

	return g.gcsClient, nil
}

func (g *gcsService) SetBucketWriter(ctx context.Context, objectName string) (*storage.Writer, error) {
	segment := newrelic.
		FromContext(ctx).
		StartSegment("pkg/gcs/InitBucketWriter")
	defer segment.End()

	if g.gcsClient == nil {
		if _, err := g.SetClient(ctx); err != nil {
			return nil, err
		}
	}

	// Create and assign writer
	g.gcsWriter = g.gcsClient.Bucket(g.config.ServiceBucketName).Object(objectName).NewWriter(ctx)
	g.gcsWriter.ContentType = g.mimeTypes(strings.Replace(filepath.Ext(objectName), ".", "", 1))

	return g.gcsWriter, nil
}

func (g *gcsService) UploadFileToGCS(ctx context.Context, objectName, srcFile string, sync bool, ttl *time.Duration) (*Response, error) {
	segment := newrelic.
		FromContext(ctx).
		StartSegment("pkg/gcs/UploadFileToGCS")
	defer segment.End()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrStorageNewClientFormat, err)
	}

	chanErr := make(chan error, 1)
	go func() {
		defer client.Close()

		f, err := os.Open(srcFile)
		if err != nil {
			chanErr <- fmt.Errorf("file.Open: %v", err)
			fmt.Printf("ERROR: file.Open: %v\n", err)
			return
		}
		defer f.Close()

		if !sync {
			ctx = newrelic.NewContext(context.Background(), newrelic.FromContext(ctx))
		}
		wr := client.Bucket(g.config.ServiceBucketName).Object(objectName).NewWriter(ctx)
		wr.CacheControl = "private"
		wr.Created = time.Now().UTC()
		wr.ContentType = g.mimeTypes(strings.Replace(filepath.Ext(objectName), ".", "", 1))

		if _, err := io.Copy(wr, f); err != nil {
			chanErr <- fmt.Errorf(ErrCopyFormat, err)
			fmt.Printf("ERROR: io.Copy: %v\n", err)
			return
		}

		chanErr <- wr.Close()
	}()

	if sync {
		if err := <-chanErr; err != nil {
			return nil, err
		}
	}

	expireDuration := 15 * time.Minute
	if ttl != nil {
		expireDuration = *ttl
	}

	signedUrl, err := client.Bucket(g.config.ServiceBucketName).SignedURL(objectName, &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(expireDuration),
	})
	if err != nil {
		return nil, fmt.Errorf(ErrBucketSignedURLFormat, g.config.ServiceBucketName, err)
	}

	return &Response{
		PublicUrl: publicURL,
		SignedUrl: signedUrl,
	}, nil
}

func (s *gcsService) UploadFileFromMultipart(ctx context.Context, objectName string, file *multipart.FileHeader, sync bool) (*UploadMultipart, error) {
	segment := newrelic.
		FromContext(ctx).StartSegment("pkg/gcs/UploadFileFromMultipart")
	defer segment.End()

	if client == nil {
		return nil, rootErr
	}
	chanErr := make(chan error, 1)

	go func() {
		f, err := file.Open()
		if err != nil {
			chanErr <- fmt.Errorf("file.Open: %v", err)
			fmt.Printf("ERROR: file.Open: %v\n", err)
			return
		}
		defer f.Close()

		if !sync {
			ctx = newrelic.NewContext(context.Background(), newrelic.FromContext(ctx))
		}
		wr := client.Bucket(s.config.ServiceBucketName).Object(objectName).NewWriter(ctx)
		wr.CacheControl = "private"
		wr.Created = time.Now().UTC()
		wr.ContentType = s.mimeTypes(strings.Replace(filepath.Ext(file.Filename), ".", "", 1))

		if _, err := io.Copy(wr, f); err != nil {
			chanErr <- fmt.Errorf(ErrCopyFormat, err)
			fmt.Printf("ERROR: io.Copy: %v\n", err)
			return
		}

		chanErr <- wr.Close()
	}()

	if sync {
		if err := <-chanErr; err != nil {
			return nil, err
		}
	}
	return &UploadMultipart{
		PublicURL:  publicURL,
		Bucket:     s.config.ServiceBucketName,
		ObjectName: objectName,
	}, nil
}

func (s *gcsService) mimeTypes(ext string) string {
	// Define a map with file extensions as keys and MIME types as values
	mimeTypeMap := map[string]string{
		"xps":  "application/oxps",
		"doc":  "application/msword",
		"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"xls":  "application/vnd.ms-excel",
		"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"csv":  "text/csv",
		"pdf":  "application/pdf",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"gif":  "image/gif",
		"tiff": "image/tiff",
		"png":  "image/png",
		"bmp":  "image/bmp",
		"zip":  "application/zip",
		"7z":   "application/x-7z-compressed",
		"rar":  "application/x-rar-compressed",
	}

	// Return the MIME type if the extension is found, otherwise return an empty string
	if mimeType, exists := mimeTypeMap[ext]; exists {
		return mimeType
	}

	return ""
}
