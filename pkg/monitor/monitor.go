package pkgMonitor

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/paper-indonesia/pdk/go/monitoring"

	"github.com/go-chi/chi/v5/middleware"
)

type Monitoring = monitoring.Monitor

var (
	once    sync.Once
	monitor *Monitoring
)

func New(repoName string, statsdIP, statsdPort string) (*Monitoring, error) {
	var err error

	once.Do(func() {
		monitor, err = monitoring.New(repoName, statsdIP, statsdPort)
	})
	return monitor, err
}

func SetGlobalMonitoring(inst *Monitoring) {
	monitor = inst
}

func WriteAndSend(ctx context.Context, funcName string, timeStart time.Time, w middleware.WrapResponseWriter, err error, f func() []string) {
	statusCode := 0
	if w != nil {
		statusCode = w.Status()
	}
	if f == nil {
		f = func() []string { return []string{} }
	}

	tags := monitor.GetDefaultTagsWithNormalErr(ctx, statusCode, err)

	monitor.Send(funcName, timeStart, append(tags, f()...))
}

func WrapResponse(w http.ResponseWriter, r *http.Request) middleware.WrapResponseWriter {
	return middleware.NewWrapResponseWriter(w, r.ProtoMajor)
}
