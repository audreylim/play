package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type samples struct {
	Samples []sample `json:"samples"`
}

type sample struct {
	WordList []string `json:"word_list"`
	IsSorted bool     `json:"is_sorted"`
}

func main() {
	f, err := ioutil.ReadFile("go/graph/graph.json")
	if err != nil {
		panic(err)
	}

	var s samples
	if err := json.Unmarshal(f, &s); err != nil {
		panic(err)
	}

	for _, sample := range s.Samples {
		g := Graph{}
		g.edgeMap = make(map[string][]string)

		fmt.Println("Given: ", sample.WordList)
		fmt.Println("Want:\tis sorted: ", sample.IsSorted)
		buildGraph(g, sample.WordList)
		sorted, isDAG := g.topologicalSort()
		fmt.Println("Got:\tis sorted:", isDAG)
		if sample.IsSorted {
			fmt.Println("\tPossible sort order: ", sorted)
		}
		fmt.Println("")
	}
}

type Graph struct {
	edgeMap map[string][]string
}

func buildGraph(g Graph, words []string) {
	prev := words[0]

	for i := 1; i < len(words); i++ {
		word := words[i]
		g.addNode(string(word[0]))
		generateRelationship(g, string(prev), words[i])
		prev = words[i]
	}
}

func generateRelationship(g Graph, prev, cur string) {
	for i := 0; i < len(prev) && i < len(cur); i++ {
		if prev[i] != cur[i] {
			from := string(prev[i])
			to := string(cur[i])

			g.addNode(from)
			g.addNode(to)
			g.addEdge(from, to)
			break
		}
	}
}

func (g *Graph) addNode(n string) {
	if _, ok := g.edgeMap[n]; !ok {
		g.edgeMap[n] = []string{}
	}
}

func (g *Graph) addEdge(from, to string) {
	relations := g.edgeMap[from]
	relations = append(relations, to)
	g.edgeMap[from] = relations
}

// Recursive DFS
func (g *Graph) topologicalSort() ([]string, bool) {
	sorted := []string{}
	isDAG := true

	color := make(map[string]string)

	for k, _ := range g.edgeMap {
		if color[k] == "" {
			color, isDAG = sort(k, color, &sorted, g)
			if !isDAG {
				break
			}
		}

	}
	return sorted, isDAG
}

func sort(node string, color map[string]string, sorted *[]string, g *Graph) (map[string]string, bool) {
	if color[node] == "grey" {
		return color, false
	}
	if color[node] == "black" {
		return color, true
	}

	followers := g.edgeMap[node]
	if len(followers) == 0 {
		pushFront(sorted, node)
		color[node] = "black"
		return color, true
	}

	for _, n := range followers {
		color[node] = "grey"
		color, isDAG := sort(n, color, sorted, g)
		if !isDAG {
			return color, false
		}
	}
	color[node] = "black"
	pushFront(sorted, node)

	return color, true
}

func pushFront(sorted *[]string, node string) {
	newSorted := make([]string, len(*sorted)+1)
	newSorted[0] = node
	copy(newSorted[1:], *sorted)
	*sorted = newSorted
}
