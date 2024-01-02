package envConfig

import (
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"log"
	"wscmakebygo.com/config"
)

var (
	Config *config.Config
)

func GetConfig() config.Config {
	if Config == nil {
		panic("config not initialized")
	}
	return *Config
}

func unmarshalConfigYaml(yamlFile []byte) *config.Config {
	var _config *config.Config
	err := yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		log.Println(err.Error())
		panic("yaml file can't unmarshal")
	}
	return _config
}

func validateStruct(i interface{}) error {
	var valid = validator.New()
	err := valid.Struct(i)
	if err != nil {
		return err
	}
	return nil
}

func InitVal() {
	log.Println("get Config")
	yamlFile := config.ReadYamlFile()
	_config := unmarshalConfigYaml(yamlFile)
	err := validateStruct(_config)
	if err != nil {
		panic(err.Error())
	}
	Config = _config
}
