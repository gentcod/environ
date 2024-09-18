package environ

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//Init initializes the environment configurations and returns an error if it occurs.
func Init(path string, conc any) (error) {
	err := loadConfig(path, &conc)
	if err != nil {
		return err
	}

	return nil
}

// loadConfig takes the file path of the .env file and parses the values to the config struct 
func loadConfig(filepath string, conc *any) (error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error encoutered reading env file, %v", err)
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
			return fmt.Errorf("error related to env variable, make sure it is properly set, %v", err)
		}
		configMap[keyVal[i][0]] = keyVal[i][1]
	}

	keystring, err := json.Marshal(configMap)
	if err != nil {
		return fmt.Errorf("error encoutered trying to marshal configMap, %v", err)
	}
	err = json.Unmarshal(keystring, &conc)
	if err != nil {
		return fmt.Errorf("error encoutered trying to unmarshall config struct, %v", err)
	}

	return nil
}