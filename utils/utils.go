package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func GetInput(day int) string {
	input, e := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), "src/adventOfGo2017", fmt.Sprintf("day%02d", day), "input.txt"))
	if e != nil {
		panic(e)
	}
	return string(input)
}

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
