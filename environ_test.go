package environ_test

import (
	"testing"

	environ "github.com/gentcod/environ"
	"github.com/stretchr/testify/require"
)


type Config struct {
	Port string `mapstructure:"PORT"`
	Url string `mapstructure:"URL"`
	Key string `mapstructure:"KEY"`
}

func TestEnviron(t *testing.T) {
	filepath := "./env/test.env"
	var conc Config
	err := environ.Init(filepath, &conc)

	require.NotEmpty(t, conc)
	require.NoError(t, err)
}