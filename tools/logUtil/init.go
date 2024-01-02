package logUtil

import (
	"io"
	"log"
	"os"
	"wscmakebygo.com/global/constant"
	"wscmakebygo.com/global/envConfig"
	"wscmakebygo.com/tools/fileUtil"
)

var (
	Log      *log.Logger
	RouteLog *log.Logger
	DBLog    *log.Logger

	logDir string
)

func CreateLogger() {
	log.Println("creating Log")
	cwd, _ := fileUtil.GetWorkingDir()
	logDir = makeLogDir(cwd)
	createLog(logDir)
	createAccessLog(logDir)
	createSqlLog(logDir)
	log.Println("created Log Success")
}

func GetEchoLogFile() io.Writer {
	return createMultiWriter(openLogFile(logDir, constant.ECHO_LOG_FILE_NAME))
}

func createLog(logDir string) {
	file := openLogFile(logDir, constant.APP_LOG_FILE_NAME)
	multiWriter := createMultiWriter(file)
	Log = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func createAccessLog(logDir string) {
	file := openLogFile(logDir, constant.ROUTER_LOG_FILE_NAME)
	multiWriter := createMultiWriter(file)
	RouteLog = log.New(multiWriter, "ACCESS: ", log.Ldate|log.Ltime)
}

func createSqlLog(logDir string) {
	file := openLogFile(logDir, constant.SQL_LOG_FILE_NAME)
	multiWriter := createMultiWriter(file)
	DBLog = log.New(multiWriter, "SQL: ", log.Ldate|log.Ltime)
}
func createMultiWriter(file *os.File) io.Writer {
	if envConfig.GetConfig().Env == "develop" {
		return io.MultiWriter(file, os.Stdout)
	} else {
		return io.MultiWriter(file)
	}
}
