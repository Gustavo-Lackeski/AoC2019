package main

import (
	"bufio"
	"container/list"
	"os"
	"strings"
)

type Planet struct {
	name       string
	neighbors  []*Planet
	depth      int
	discovered bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	planetMapping := make(map[string]*Planet)
	for scanner.Scan() {
		addPlanets(scanner, planetMapping)
	}
	BFS(planetMapping["YOU"])
	total := 0
	for _, m := range planetMapping {
		total += m.depth
	}
	println(planetMapping["SAN"].depth - 2)

}

func addPlanets(scanner *bufio.Scanner, planetMapping map[string]*Planet) {
	tokens := strings.Split(scanner.Text(), ")")
	var leftPlanet *Planet
	if _, ok := planetMapping[tokens[0]]; !ok {
		leftPlanet = &Planet{
			name: tokens[0],
		}
		planetMapping[tokens[0]] = leftPlanet
	} else {
		leftPlanet = planetMapping[tokens[0]]
	}

	var rightPlanet *Planet
	if _, ok := planetMapping[tokens[1]]; !ok {
		rightPlanet = &Planet{
			name: tokens[1],
		}
		planetMapping[tokens[1]] = rightPlanet
	} else {
		rightPlanet = planetMapping[tokens[1]]
	}

	leftPlanet.neighbors = append(leftPlanet.neighbors, rightPlanet)
	rightPlanet.neighbors = append(rightPlanet.neighbors, leftPlanet)

}

func BFS(p *Planet) {
	queue := list.New()
	p.discovered = true
	p.depth = 0
	queue.PushBack(p)
	for queue.Len() > 0 {
		v := queue.Front()
		queue.Remove(v)
		value, ok := v.Value.(*Planet)
		if !ok {
			panic(v.Value)
		}
		for _, w := range value.neighbors {
			if !w.discovered {
				w.discovered = true
				w.depth = value.depth + 1
				queue.PushBack(w)
			}
		}
	}
}
