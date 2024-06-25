package urlvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateURL(t *testing.T) {
	tests := map[string]struct {
		URL     string
		wantErr bool
	}{
		"Valid URL with scheme, host and port": {
			URL:     "http://example.com:80/abc",
			wantErr: false,
		},
		"Valid URL with host and port": {
			URL:     "example.com:80/abc",
			wantErr: false,
		},
		"Valid URL with host": {
			URL:     "example.com",
			wantErr: false,
		},
		"Valid URL with IP as host": {
			URL:     "192.168.178.1",
			wantErr: false,
		},
		"Valid URL with IP as host and port": {
			URL:     "127.0.0.1:61273",
			wantErr: false,
		},
		"Valid URL with IPv6 as host": {
			URL:     "2001:0db8:0000:0000:0000:ff00:0042:8329",
			wantErr: false,
		},
		"Valid URL with IPv6 as host and port": {
			URL:     "[2001:0db8:0000:0000:0000:ff00:0042:8329]:61273",
			wantErr: false,
		},
		// Invalid URLs
		// IPv6 address with port must include square brackets
		"Invalid URL with IPv6 as host and port": {
			URL:     "2001:0db8:0000:0000:0000:ff00:0042:8329:61273",
			wantErr: true,
		},
		"Invalid URL with scheme, port and space in host": {
			URL:     "http://exam ple.com:80/abc",
			wantErr: true,
		},
		"Invalid URL with port and space in host": {
			URL:     "exam ple.com:80/abc",
			wantErr: true,
		},
		"Invalid URL with scheme, and space in host": {
			URL:     "tcp://exam ple.com/abc",
			wantErr: true,
		},
		"Invalid URL with leading space in host": {
			URL:     " example.com",
			wantErr: true,
		},
		"Invalid URL with trailing space in host": {
			URL:     "example.com ",
			wantErr: true,
		},
		"Invalid URL with space in host": {
			URL:     "exam ple.com/abc",
			wantErr: true,
		},
	}
	for tname, tt := range tests {
		t.Run(tname, func(t *testing.T) {
			gotErr := ValidateURL(tt.URL)
			if !tt.wantErr {
				assert.NoError(t, gotErr)
				//assert.Equal(t, tt.wantURL, gotAddr)
			} else {
				assert.Error(t, gotErr)
			}

			//host, zone, port, err := netutil.ParseEndpoint(gotAddr)
			//if tt.wantErrToContain == "" {
			//	require.NoError(t, err)
			//	assert.Equal(t, u.Host, host, "host does not match")
			//	assert.Equal(t, u.Scheme, zone, "zone does not match")
			//	assert.Equal(t, u.Port(), port, "port does not match")
			//}
		})
	}
}
