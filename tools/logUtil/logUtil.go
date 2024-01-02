package logUtil

import (
	"log"
	"os"
	"path/filepath"
)

func makeLogDir(path string) string {
	dir := filepath.Join(path, "log")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("创建log目录失败: %v", err)
	}
	return dir
}

func openLogFile(path string, fileName string) *os.File {
	logFile := filepath.Join(path, fileName)
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}
	return file
}

func closeLogger() {
	if file, ok := Log.Writer().(*os.File); ok {
		file.Close()
	}
}

func closeDbLogger() {
	if file, ok := DBLog.Writer().(*os.File); ok {
		file.Close()
	}
}

func closeRouteLogger() {
	if file, ok := RouteLog.Writer().(*os.File); ok {
		file.Close()
	}
}

func Close() {
	closeDbLogger()
	closeRouteLogger()
	closeLogger()
}
