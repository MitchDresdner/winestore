package util

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Db struct {
		Pg map[string]string `yaml:"postgress"`
	}
}

func FetchYAML () (string, error) {

	dat, err := ioutil.ReadFile("./properties/dev-properties.yaml")
	if err != nil {
		return "", err
	}
	//fmt.Print(string(dat))

	var config Config
	err = yaml.Unmarshal([]byte(dat), &config)
	if err != nil {
		return "", err
	}

	fmt.Println("Shhh:", os.Getenv("SECRET"))
	for k, v := range config.Db.Pg {
		// extract and decode strings wrapped in ![]
		if strings.HasPrefix(v,"![") && strings.HasSuffix(v,"]") {
			secret := strings.TrimRight(strings.TrimLeft(v, "!["), "]")
			data, err := hex.DecodeString(secret)
			if err != nil {
				return "", err
			}
			decoded, e := decrypt(data,[]byte(os.Getenv("SECRET")))
			if e != nil {
				return "", fmt.Errorf("error: %v", e)
			}

			fmt.Printf("found %s\n", decoded)
			// replace value with decoded secret in map
			v = string(decoded)
			config.Db.Pg[k] = v
		}
		fmt.Printf("key %s: %s\n", k, v)
	}

	// connStr := "user=goland dbname=goland password=goland host=localhost port=30432 sslmode=disable"
	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.Db.Pg["user"], config.Db.Pg["dbname"], config.Db.Pg["password"],
		config.Db.Pg["host"], config.Db.Pg["port"]), err
}
