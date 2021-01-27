package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func init_env(env_values *map[interface{}]map[interface{}]interface{}) error {
	secrets_yaml, err := ioutil.ReadFile("secrets.yaml")
	if err != nil {
		return err
	}

	return yaml.Unmarshal(secrets_yaml, &env_values)
}
