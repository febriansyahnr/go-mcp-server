package util_test

import (
	"errors"
	"testing"
	"time"

	"github.com/paper-indonesia/pg-mcp-server/pkg/util"
	"github.com/stretchr/testify/require"
)

func mockLocationLoader(name string) (*time.Location, error) {
	return nil, errors.New("mock location error")
}

func TestGetTimeWithLoader(t *testing.T) {
	tests := []struct {
		name       string
		loaderFunc util.LocationLoader
		wantErr    bool
	}{
		{
			name:       "Successful Timezone Retrieval",
			loaderFunc: time.LoadLocation,
			wantErr:    false,
		},
		{
			name:       "Error Handling",
			loaderFunc: mockLocationLoader,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := util.GetJakartaTimeWithLoader(tt.loaderFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: GetJakartaTimeWithLoader() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestGetJakartaTime(t *testing.T) {
	jakartaTime, err := util.GetJakartaTime()
	if err != nil {
		t.Errorf("GetJakartaTime() returned an error: %v", err)
	}

	_, offset := jakartaTime.Zone()
	if offset != 7*60*60 {
		t.Errorf("GetJakartaTime() did not return time in 'Asia/Jakarta' timezone: got offset %d", offset)
	}
}

func TestSnapCompatible(t *testing.T) {
	tests := []struct {
		name           string
		inputTime      time.Time
		expectedOutput string
	}{
		{
			name:           "Valid input time",
			inputTime:      time.Date(2022, time.January, 1, 12, 0, 0, 0, time.UTC),
			expectedOutput: "2022-01-01T19:00:00+07:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SnapFormat(tt.inputTime)
			if result != tt.expectedOutput {
				t.Errorf("SnapCompatible(%v) = %s; want %s", tt.inputTime, result, tt.expectedOutput)
			}
		})
	}
}

func TestC2aDateFormat(t *testing.T) {
	now := time.Date(2024, time.January, 1, 12, 4, 5, 6, time.UTC)
	resultDate, resultTime := util.C2aDateFormat(now)

	require.Equal(t, resultDate, "20240101")
	require.Equal(t, resultTime, "190405")
}
