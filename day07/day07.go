package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type tower map[string]node

type node struct {
	weight      int
	children    []string
	totalWeight int
}

func (n node) HasChild(name string) bool {
	for _, c := range n.children {
		if c == name {
			return true
		}
	}
	return false
}

func (t tower) FindRoot() string {
	root, stop := make(chan string, 1), make(chan bool, 1)
	var wg sync.WaitGroup
	for k, v := range t {
		if v.children != nil {
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				for _, n := range t {
					select {
					case <-stop:
						stop <- true
						return
					default:
						if n.HasChild(name) {
							return
						}
					}
				}
				stop <- true
				root <- name
			}(k)
		}
	}
	wg.Wait()
	return <-root
}

func (t tower) FindUnbalanced(name string) (string, int, bool) {
	n := t[name]
	if n.children != nil || len(n.children) >= 2 {
		type childWeight struct {
			name        string
			totalWeight int
		}
		tmp := make([]childWeight, 0)
		for _, c := range n.children {
			tmp = append(tmp, childWeight{c, t[c].totalWeight})
		}
		sort.Slice(tmp, func(i, j int) bool { return tmp[i].totalWeight < tmp[j].totalWeight })
		if diff := tmp[1].totalWeight - tmp[0].totalWeight; diff != 0 {
			return tmp[0].name, diff, true
		}
		if diff := tmp[len(tmp)-2].totalWeight - tmp[len(tmp)-1].totalWeight; diff != 0 {
			return tmp[len(tmp)-1].name, diff, true
		}
	}

	return "", 0, false
}

func (t tower) CalcTotalWeight(name string) {
	n := t[name]
	totalWeight := n.weight

	for _, c := range n.children {
		t.CalcTotalWeight(c)
		totalWeight += t[c].totalWeight
	}

	n.totalWeight = totalWeight
	t[name] = n
}

func parseInput(input string) tower {
	nodes := make(tower)
	re := regexp.MustCompile("([a-z]+) \\(([\\d]+)\\)( -> ([a-z, ]+)|())")
	for _, s := range strings.Split(input, "\n") {
		groups := re.FindStringSubmatch(s)
		name := groups[1]
		weight, err := strconv.Atoi(groups[2])
		if err != nil {
			panic(err)
		}
		var children []string = nil
		if groups[4] != "" {
			children = strings.Split(groups[4], ", ")
		}
		nodes[name] = node{weight, children, 0}
	}
	return nodes
}

func p1(input string) string {
	nodes := parseInput(input)
	return nodes.FindRoot()

}

func p2(input string) int {
	nodes := parseInput(input)
	root := nodes.FindRoot()
	nodes.CalcTotalWeight(root)
	ubName, ubDiff := root, 0
	for name, diff, ub := nodes.FindUnbalanced(root); ub; name, diff, ub = nodes.FindUnbalanced(ubName) {
		ubName, ubDiff = name, diff
	}
	return nodes[ubName].weight + ubDiff
}

func main() {
	input := utils.GetInput(7)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
