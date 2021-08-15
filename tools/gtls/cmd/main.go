package main

import (
	"fmt"
	"gateway/tools/gtls"
	"os"
	"time"
)

func main() {
	filename := argv()
	config, err := gtls.ParseConfig(filename)
	if err != nil {
		panic(err)
	}

	options, err := config.Tran2Options()
	if err != nil {
		panic(err)
	}

	var handler *gtls.Cert
	err = config.CheckFile()
	if err != nil {
		option := &gtls.Options{
			KeyType:      gtls.KeyTypeRSA,
			Duration:     365 * 24 * time.Hour,
		}
		handler, err = gtls.New(option)
		caCrtWrapper, caKeyWrapper, err := handler.CA()
		if err != nil {
			panic(err)
		}
		if err = caCrtWrapper.File(config.Crt); err != nil {
			panic(err)
		}
		if err = caKeyWrapper.File(config.Key); err != nil {
			panic(err)
		}

	} else {
		handler, err = gtls.File(config.Crt, config.Key)
	}

	if err != nil {
		panic(err)
	}

	for _, option := range options {
		crtWrapper, keyWrapper, err := handler.GenerateCrtAndKey(option)
		if err != nil {
			panic(err)
		}
		if err = crtWrapper.File(fmt.Sprintf("./%s.crt", option.Name)); err != nil {
			panic(err)
		}
		if err = keyWrapper.File(fmt.Sprintf("./%s.key", option.Name)); err != nil {
			panic(err)
		}
	}
}

func argv() (filename string) {
	args := os.Args
	if len(args) == 1 {
		filename = "gtls.yaml"
	} else {
		filename = args[1]
	}
	return
}
