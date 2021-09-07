package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	EnvDev = runEnv(iota)
	EnvTest
	EnvProd
)

const configFileName = "config.yaml"

type runEnv int8

// getRunEnv 从runenv环境变量里获取运行环境
func getRunEnv(e []runEnv) (runEnv, error) {
	if len(e) > 0 {
		return e[0], nil
	}

	value := os.Getenv("runenv")
	env, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return EnvDev, nil
	}

	if env < 0 || env > 2 {
		return 0, errors.New("invalid value")
	}

	return runEnv(env), nil
}
