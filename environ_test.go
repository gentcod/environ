package environ_test

import (
	"testing"
	"time"

	environ "github.com/gentcod/environ"
	"github.com/stretchr/testify/require"
)

type Config struct {
	PortAddress          string        `json:"PORT_ADDRESS"`
	GrpcAddress          string        `json:"GRPC_ADDRESS"`
	DBDriver             string        `json:"DB_DRIVER"`
	DBUrl                string        `json:"DB_URL"`
	TokenSymmetricKey    string        `json:"TOKEN_SYMMETRIC_KEY"`
	TokenSecretKey       string        `json:"TOKEN_SECRET_KEY"`
	TokenDuration        time.Duration `json:"TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `json:"REFRESH_TOKEN_DURATION"`
	MigrationUrl         string        `json:"MIGRATION_URL"`
	TestInt              int64        `json:"TEST_INT"`
	TestBool             bool        `json:"TEST_BOOL"`
}

func TestEnivron(t *testing.T) {
	filepath := "test.env"
	var conc Config
	err := environ.Init(filepath, &conc)

	require.NoError(t, err)
	require.NotEmpty(t, conc)
	require.Equal(t, "0.0.0.0:5000", conc.PortAddress)
	require.Equal(t, "postgres", conc.DBDriver)
	require.Equal(t, "postgres://root:secret@localhost:5431/envrion?sslmode=disable", conc.DBUrl)
	require.Equal(t, 15*time.Minute, conc.TokenDuration)
	require.Equal(t, 24*time.Hour, conc.RefreshTokenDuration)
}

func TestInit_NotPointer(t *testing.T) {
	var cfg Config
	err := environ.Init("test.env", cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "config must be a non-nil pointer to a struct")
}

func TestInit_NilPointer(t *testing.T) {
	var cfg *Config
	err := environ.Init("test.env", cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "config must be a non-nil pointer to a struct")
}

func TestInit_InvalidConfigType(t *testing.T) {
	var cfg string
	err := environ.Init("test.env", cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "config must be a non-nil pointer to a struct")
}
