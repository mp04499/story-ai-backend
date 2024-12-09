package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	LlmUrl string `yaml:"llm_url"`
}

func LoadConfig() *Config {
	config := &Config{}
	env := os.Getenv("STORY_ENV")
	if env == "" {
		env = "local"
	}
	fileName := "./configs/config_" + env + ".yml"
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if err != nil {
		panic(err)
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}

	return config
}
