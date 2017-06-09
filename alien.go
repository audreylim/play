package main

import "fmt"

func main() {
	//sorted1 := []string{"xza", "ayh", "ples", "plares", "bhaaz", "bnc"}
	unsorted1 := []string{"xza", "ayh", "eyes", "ples", "plares", "bhaaz", "bnc"}

	fmt.Println(unsorted1)

	g := Graph{}
	g.edgeMap = make(map[string][]string)
	buildGraph(g, unsorted1)
	fmt.Println(g)
}

func buildGraph(g Graph, words []string) {
	prev := words[0]
	g.addVertex(string(prev[0]))

	for i := 1; i < len(words); i++ {
		word := words[i]
		g.addVertex(string(word[0]))
		generateRelationship(g, string(prev), words[i])
		prev = words[i]
	}
}

func generateRelationship(g Graph, prev, cur string) {
	for i := 0; i < len(prev) && i < len(cur); i++ {
		if prev[i] != cur[i] {
			g.addVertex(string(cur[i]))
			g.addVertex(string(prev[i]))
			g.addEdge(string(prev[i]), string(cur[i]))
			break
		}
	}
}

type Graph struct {
	edgeMap map[string][]string
}

func (g *Graph) addVertex(v string) {
	if _, ok := g.edgeMap[v]; !ok {
		g.edgeMap[v] = []string{}
	}
}

func (g *Graph) addEdge(from, to string) {
	relations := g.edgeMap[from]
	relations = append(relations, to)
	g.edgeMap[from] = relations
}
