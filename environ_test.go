package environ_test

import (
	"testing"
	environ "github.com/gentcod/dotenv"
	"github.com/stretchr/testify/require"
)

func TestEnviron(t *testing.T) {
	filepath := "./env/test.env"
	env, err := environ.Init(filepath)

	require.NotEmpty(t, env)
	require.NoError(t, err)
}