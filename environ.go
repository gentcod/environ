package environ

import (
	"fmt"
	"os"
	"strings"
)

type config struct {
	env map[string]any
}

//Init initializes the environment configurations and returns a map of key value
func Init(path string) (map[string]any, error) {
	config, err := loadConfig(path)
	if err != nil || config == nil {
		return nil, err
	}

	fmt.Println(config.env)

	return config.env, nil
}

// loadConfig takes the file path of the .env file and generates a
func loadConfig(filepath string) (*config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	dataString := string(data)
	fields := strings.Fields(dataString)
	keyVal := [][]string{}
	for i := 0; i < len(fields); i++ {
		part := strings.Split(fields[i], "=")
		keyVal = append(keyVal, part)
	}

	configMap := make(map[string]any)
	for i := 0; i < len(keyVal); i++ {
		if len(keyVal[i]) != 2 {
			return nil, err
		}
		configMap[keyVal[i][0]] = keyVal[i][1]
	}

	config := &config{
		env: configMap,
	}

	return config, nil
}