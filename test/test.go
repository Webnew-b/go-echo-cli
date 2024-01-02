package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ChangeDir() {
	pwd, _ := os.Getwd()
	fmt.Println("开始工作目录", pwd)
	// 程序所在目录
	execDir, err := filepath.Abs(filepath.Dir("G:\\后端\\WSC\\"))
	if err != nil {
		log.Fatal(err)
	}
	if pwd == execDir {
		fmt.Println("不需要切换工作目录")
		return
	}
	fmt.Println("切换工作目录到", execDir)
	if err := os.Chdir(execDir); err != nil {
		log.Fatal(err)
	}
	pwd, _ = os.Getwd()
	fmt.Println("切换后工作目录:", pwd)
}
