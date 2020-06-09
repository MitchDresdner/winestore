package util

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Db struct {
		Pg map[string]string `yaml:"postgress"`
	}
}

func FetchYAML() (string, error) {

	dat, err := ioutil.ReadFile("./properties/dev-properties.yaml")
	if err != nil {
		return "", err
	}

	// Unmarshal YAML file into Config struct
	var config Config
	err = yaml.Unmarshal([]byte(dat), &config)
	if err != nil {
		return "", err
	}

	// Iterate over properties, decoding any secrets
	for k, v := range config.Db.Pg {
		// Extract and decode secrets which are wrapped in ![] (like Mule)
		if strings.HasPrefix(v, "![") && strings.HasSuffix(v, "]") {
			secret := strings.TrimRight(strings.TrimLeft(v, "!["), "]")
			// Convert hex representation of secret to uint8 byte array
			data, err := hex.DecodeString(secret)
			if err != nil {
				return "", err
			}
			decoded, e := decrypt(data, []byte(os.Getenv("SECRET")))
			if e != nil {
				return "", fmt.Errorf("error: %v", e)
			}

			// fmt.Printf("found %s\n", decoded)
			// Replace encoded map value with the decoded secret
			v = string(decoded)
			config.Db.Pg[k] = v
		}
		// fmt.Printf("key %s: %s\n", k, v)
	}

	// Create the database connection string, this one is for Postgres
	// ToDo: Create a general purpose function that handles other DB's
	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.Db.Pg["user"], config.Db.Pg["dbname"], config.Db.Pg["password"],
		config.Db.Pg["host"], config.Db.Pg["port"]), err
}
