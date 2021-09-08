package ipip

import (
	"fmt"
	"testing"
)

func TestIpIp_Query(t *testing.T) {
	ipip, err := New()
	if err != nil {
		panic(err)
	}

	province, city, err := ipip.Query("117.173.86.7")
	if err != nil {
		panic(err)
	}

	fmt.Println(province, city)	
}
