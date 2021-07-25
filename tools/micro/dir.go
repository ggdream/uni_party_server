package main

import "fmt"

// dir 项目根路径下需要创建的单个文件夹
type dir struct {
	name  string
	files map[string]string
}

const (
	configYaml = ``
	mainGo     = "package main\n\nfunc main() {\n\t\n}\n"

	configGo = `package config\n\n`

	handlerGo = `package handler\n\n`

	modelGo = `package model\n\n`

	protoFile = `syntax = "proto3";

option go_package = "./proto";
package proto;

service Service {
  rpc Greet (Request) returns (Response) {};
}

message Request {
  string name = 1;
}

message Response {
  string reply = 1;
}
`

	READMEFile = `# %s\n\n`

	Makefile = `.PHONY: compile

compile:
	protoc --go_out=. --go-grpc_out=. proto/%s.proto
`
)

var (
	binDir = &dir{
		name: "bin",
		files: map[string]string{
			"config.yaml": configYaml,
			"main.go":     mainGo,
		},
	}

	configDir = &dir{
		name: "config",
		files: map[string]string{
			"config.go": configGo,
		},
	}

	handlerDir = &dir{
		name: "handler",
		files: map[string]string{
			"handler.go": handlerGo,
		},
	}

	modelDir = &dir{
		name: "model",
		files: map[string]string{
			"model.go": modelGo,
		},
	}

	protoDir = func(projName string) *dir {
		return &dir{
			name: "proto",
			files: map[string]string{
				fmt.Sprintf("%s.proto", projName): protoFile,
			},
		}
	}

	currentDir = func(projName string) *dir {
		return &dir{
			name: "./",
			files: map[string]string{
				"README.md": fmt.Sprintf(READMEFile, projName),
				"Makefile":  fmt.Sprintf(Makefile, projName),
			},
		}
	}
)
