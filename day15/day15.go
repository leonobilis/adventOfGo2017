package main

import (
	"fmt"
)

type generator interface {
	start(int, func (int) int)
	next() int
	stop()
}

type gen struct {
	getNext chan bool
	nextValue chan int
}

func (g *gen) start(initial int, nextFunction func (int) int) {
	g.getNext, g.nextValue = make(chan bool), make(chan int)
	go func () {
		i := initial
		for range g.getNext {
			i = nextFunction(i)
			g.nextValue <- i
		}
		close(g.nextValue)
	}()
}

func (g gen) next() int{
	g.getNext <- true
	return <-g.nextValue
}

func (g gen) stop() {
	close(g.getNext)
}

func p1(a, b int) int {
	sum := 0
	genA, genB := gen{}, gen{}
	genA.start(a, func (i int) int{
		return 16807 * i % 2147483647
	})
	genB.start(b, func (i int) int{
		return 48271 * i % 2147483647
	})

	for i:=0; i < 40000000; i++ {
		if genA.next() & 65535 == genB.next() & 65535 {
			sum++
		}
	}
	genA.stop()
	genB.stop()
	return sum
}

func p2(a, b int) int {
	sum := 0
	genA, genB := gen{}, gen{}
	genA.start(a, func (i int) int{
		for i = 16807 * i % 2147483647; i%4 != 0; i = 16807 * i % 2147483647{}
		return i
	})
	genB.start(b, func (i int) int{
		for i = 48271 * i % 2147483647; i%8 != 0; i = 48271 * i % 2147483647{}
		return i
	})

	for i:=0; i < 5000000; i++ {
		genA.getNext <- true
		genB.getNext <- true
		if <-genA.nextValue & 65535 == <-genB.nextValue & 65535 {
			sum++
		}
	}
	genA.stop()
	genB.stop()
	return sum
}

func main() {
	a, b := 116, 299
	fmt.Printf("Part 1: %v\n", p1(a, b))
	fmt.Printf("Part 2: %v\n", p2(a, b))
}
