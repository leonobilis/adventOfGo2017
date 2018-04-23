package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetInput(day int) string {
	input, e := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), "src/adventOfGo2017", fmt.Sprintf("day%02d", day), "input.txt"))
	if e != nil {
		panic(e)
	}
	return string(input)
}

func Int(s string, i int) int {
	return int(s[i] - '0')
}
