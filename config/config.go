package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"wscmakebygo.com/global/constant"
	"wscmakebygo.com/tools/fileUtil"
)

type Config struct {
	Env string `yaml:"env" validate:"required"`

	App   *App   `yaml:"app"`
	Db    *Db    `yaml:"db"`
	Redis *Redis `yaml:"redis"`
}

// 以下是2层，三层往下添加对应结构体
type App struct {
	Host string `yaml:"host" validate:"required"`
	Port int    `yaml:"port" validate:"required"`
}

type Db struct {
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name" validate:"required"`
	Charset  string `yaml:"charset" validate:"required"`
	Loc      string `yaml:"loc" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
}

type Redis struct {
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db" validate:"gte=0"`
}

func ReadYamlFile() []byte {
	workingPath, err := fileUtil.GetWorkingDir()
	if err != nil {

		panic(err.Error())
	}
	path := filepath.Join(workingPath, constant.GetConfigPath())
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		checkError(err)
	}
	return yamlFile
}

func checkError(err error) {
	if err != nil {
		// 检查文件是否不存在
		if os.IsNotExist(err) {
			createYamlFile()
			panic("please make sure that the " +
				constant.GetConfigPath() +
				" file and directory are available for reading!")
		} else if os.IsPermission(err) {
			panic("please make sure that the " +
				constant.GetConfigPath() +
				" file and directory are available for reading!")
		} else {
			panic(fmt.Sprintf("Failed to open file: %v\n", err))
		}
		return
	}

}

func createYamlFile() {
	isExistDir()

	data := createConfig()

	err := os.WriteFile(constant.GetConfigPath(), data, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	panic("File is created,please field the serve value")
}

func createConfig() []byte {
	config := Config{
		Env: "product",
		App: &App{
			Host: "0.0.0.0",
			Port: 8080,
		},
		Db: &Db{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "123456",
			DbName:   "WSC_api",
			Charset:  "utf8mb4",
			Loc:      "Local",
		},
		Redis: &Redis{
			Host: "127.0.0.1",
			Port: 6379,
			DB:   0,
		},
	}
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return data
}

func isExistDir() {
	if _, err := os.Stat(constant.CONFIG_DIR_PATH); os.IsNotExist(err) {
		err = os.Mkdir(constant.CONFIG_DIR_PATH, 0755) // 使用适当的权限
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
		log.Printf("Folder %v created successfully.\n", constant.CONFIG_DIR_PATH)
	}
}
