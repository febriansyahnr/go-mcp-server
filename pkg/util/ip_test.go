package util

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOutboundIP(t *testing.T) {
	testCases := []struct {
		desc    string
		host    string
		wantErr bool
	}{
		{
			desc:    "should get outbound IP for valid host",
			host:    "8.8.8.8:80",
			wantErr: false,
		},
		{
			desc:    "should get outbound IP for localhost",
			host:    "127.0.0.1:80",
			wantErr: false,
		},
		{
			desc:    "should return error for invalid host format",
			host:    "invalid-host",
			wantErr: true,
		},
		{
			desc:    "should return error for non-existent host",
			host:    "999.999.999.999:80",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ip, err := GetOutboundIP(tc.host)
			
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, ip)
				assert.IsType(t, net.IP{}, ip)
			}
		})
	}
}