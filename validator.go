package urlvalidator

import (
	"net"
	"net/netip"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	maxURLRuneCount = 2083
	minURLRuneCount = 3
)

// ValidateURL validates absolute URLs
func ValidateURL(str string) error {
	var scheme, host, port, path string

	if str == "" || utf8.RuneCountInString(str) >= maxURLRuneCount || len(str) <= minURLRuneCount {
		return errors.New("invalid URL length")
	}
	if strings.HasPrefix(str, ".") {
		return errors.New("URL cannot start with a period")
	}
	// Extract the scheme
	next := str
	if before, rest, schemeFound := strings.Cut(str, "://"); schemeFound {
		scheme = before
		next = rest
	}
	// Separate the host part form the path
	hostCandidate, path, _ := strings.Cut(next, "/")
	// ignore the path, but save it

	host, port, err := net.SplitHostPort(hostCandidate)
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "missing port in address"):
		host = hostCandidate
		port = ""
	case strings.Contains(err.Error(), "too many colons in address"):
		// IPv6
		addr, err6 := netip.ParseAddr(hostCandidate)
		if err6 != nil {
			return errors.Wrap(err6, "invalid URL")
		}
		if !addr.IsValid() {
			return errors.Wrap(err, "invalid URL")
		}
		host = addr.String()
	default:
		return errors.Wrap(err, "invalid URL")
	}

	if strings.Contains(host, " ") {
		return errors.New("hostname contains invalid characters")
	}
	if host == "" {
		return errors.New("hostname cannot be empty")
	}
	_ = scheme
	_ = path
	_ = port
	return nil
}
