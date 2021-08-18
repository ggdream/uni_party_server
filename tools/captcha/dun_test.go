package captcha

import (
	"fmt"
	"testing"
)

func TestNewDun(t *testing.T) {
	d, err := NewDun("ss", "ss", "ss")
	if err != nil {
		panic(err)
	}
	res, err := d.Verify("sfsfsaf", "moca")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", res)
}
