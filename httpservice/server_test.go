package httpservice_test

import (
	"context"
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/httpservice"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestHTTPServer(t *testing.T) {
	cache := crypto.NewCache()
	srv := httpservice.NewServer(cache)

	// Kill after some time
	go func() {
		time.Sleep(1 * time.Second)
		srv.Shutdown(context.Background())
	}()

	err := srv.Start()
	require.NoError(t, err)
}
