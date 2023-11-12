package util

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	config, err := LoadConfig("..")
	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, config.AccessTokenDuration, 15*time.Minute)
	fmt.Printf("config: %v", config)
}
