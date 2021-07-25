package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// mk 操作文件、文件夹
type mk struct {
	projectPath string
	dirs        []*dir
}

func New(target string) *mk {
	projectPath := filepath.Join("./", target)
	_, err := os.Stat(projectPath)
	if err == nil || os.IsExist(err) {
		fmt.Println("The dir is not empty!! So we stop it")
		os.Exit(1)
	}

	projPathList := strings.Split(target, "/")
	projName := projPathList[len(projPathList)-1]

	return &mk{
		projectPath,
		[]*dir{
			binDir,
			configDir,
			handlerDir,
			modelDir,
			protoDir(projName),
			currentDir(projName),
		},
	}
}

// Create 创建相关目录
func (m *mk) Create() error {
	for _, v := range m.dirs {
		if err := m.create(v); err != nil {
			return err
		}
	}

	m.print()
	return nil
}

func (m *mk) create(d *dir) error {
	thePath := filepath.Join(m.projectPath, d.name)
	if err := os.MkdirAll(thePath, 0755); err != nil {
		return err
	}

	for k, v := range d.files {
		filePath := filepath.Join(thePath, k)
		if err := os.WriteFile(filePath, []byte(v), 0644); err != nil {
			return err
		}
	}

	return nil
}

func (m *mk) print() {
	value := "Please download the compile tools:\n\t$ go install google.golang.org/protobuf/cmd/protoc-gen-go\n\t$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc\nget protoc: https://github.com/protocolbuffers/protobuf/releases\n\nCreate it successfully!"

	fmt.Println(value)
}
