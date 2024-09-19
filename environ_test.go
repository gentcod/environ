package environ_test

import (
	"testing"
	"time"

	environ "github.com/gentcod/environ"
	"github.com/stretchr/testify/require"
)

type Config struct {
	PortAddress string
	GrpcAddress string
	DBDriver string
	DBUrl string
	TokenSymmetricKey string
	TokenSecretKey string
	TokenDuration time.Duration
	RefreshTokenDuration time.Duration
	MigrationUrl string
}

func TestEnviron(t *testing.T) {
	filepath := "test.env"
	var conc Config
	err := environ.Init(filepath, &conc)

	require.NotEmpty(t, conc)
	require.NoError(t, err)
}