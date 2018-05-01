package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"regexp"
	"strings"
)

type coordinates struct {
	x, y, z int
}

type particle struct {
	p, v, a coordinates
}

func (c coordinates) dist() int {
	return utils.Abs(c.x) + utils.Abs(c.y) + utils.Abs(c.z)
}

func (c *coordinates) add(arg coordinates) {
	c.x += arg.x
	c.y += arg.y
	c.z += arg.z
}

func (p *particle) update() {
	p.v.add(p.a)
	p.p.add(p.v)
}

func ParseInput(input string) map[int]*particle {
	particles := make(map[int]*particle)
	re := regexp.MustCompile("p=<([-\\d]+),([-\\d]+),([-\\d]+)>, v=<([-\\d]+),([-\\d]+),([-\\d]+)>, a=<([-\\d]+),([-\\d]+),([-\\d]+)>")
	for i, s := range strings.Split(input, "\n") {
		groups := re.FindStringSubmatch(s)
		particles[i] = &particle{
			coordinates{utils.Atoi(groups[1]), utils.Atoi(groups[2]), utils.Atoi(groups[3])},
			coordinates{utils.Atoi(groups[4]), utils.Atoi(groups[5]), utils.Atoi(groups[6])},
			coordinates{utils.Atoi(groups[7]), utils.Atoi(groups[8]), utils.Atoi(groups[9])}}
	}
	return particles
}

func p1(input string) int {
	particles := ParseInput(input)
	minA := 1<<31 - 1
	for _, p := range particles {
		if dist := p.a.dist(); dist < minA {
			minA = dist
		}
	}
	minAParticles := make([]int, 0)
	for i, p := range particles {
		if p.a.dist() == minA {
			minAParticles = append(minAParticles, i)
		}
	}
	minV := 1<<31 - 1
	minVParticle := 0
	for _, i := range minAParticles {
		if dist := particles[i].v.dist(); dist < minV {
			minV = dist
			minVParticle = i
		}
	}
	return minVParticle
}

func p2(input string) int {
	particles := ParseInput(input)
	for noChangeCounter:=0; noChangeCounter < 100; {
		positions := make(map[coordinates][]int)
		for i, p := range particles {
			p.update()
			positions[p.p] = append(positions[p.p], i)
		}
		for _, p := range positions {
			if len(p) > 1 {
				for _, i := range p {
					delete(particles, i)
				}
				noChangeCounter = 0
			}
		}
		noChangeCounter++
	}
	return len(particles)
}

func main() {
	input := utils.GetInput(20)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
