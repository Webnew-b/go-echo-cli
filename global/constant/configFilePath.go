package constant

import "fmt"

const (
	CONFIG_FILE_PATH = "config.yml"
	CONFIG_DIR_PATH  = "config/"
)

func GetConfigPath() string {
	return fmt.Sprintf("%s%s", CONFIG_DIR_PATH, CONFIG_FILE_PATH)
}
