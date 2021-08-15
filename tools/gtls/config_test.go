package gtls

import (
	"fmt"
	"testing"
)

func TestParseConfig(t *testing.T) {
	config, err := ParseConfig("./example.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
