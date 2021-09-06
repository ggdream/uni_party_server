package group

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

type test struct {
	Data	*[]string
}

func TestName(t *testing.T) {
	data := make([]string, 0)
	te := test{&data}
	for i := 0; i < 100000; i++ {
		data = append(data, string(rune(i)))
	}
	fmt.Println(te.Data)
}

func TestNewUniversity(t *testing.T) {
	file, err := os.Open("./sample.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var class []SingleStruct
	br := bufio.NewScanner(file)
	for {
		if !br.Scan() {
			break
		}
		line := br.Text()
		fmt.Println(line)
		temp := strings.Split(line, ",")
		i, err := strconv.Atoi(temp[3])
		if err != nil {
			panic(err)
		}
		class = append(class, SingleStruct{
			Campus:  temp[1],
			College: temp[2],
			Grade:   uint8(i),
			Major:   temp[4],
			Class:   temp[5],
			UID:     0,
		})
	}

	//value := NewUniversity("四川师范大学", 520, [][2]interface{}{{uint(452), "商务部"}}, map[int][]SingleStruct{4: class})
}
