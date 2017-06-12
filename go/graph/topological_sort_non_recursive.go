package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type samples struct {
	Sample []sample `json:"samples"`
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
		fmt.Println(err)
	}

	for _, v := range s.Sample {
		g := Graph{}
		g.edgeMap = make(map[string][]string)

		fmt.Println("Given: ", v.WordList)
		fmt.Println("Want:\tis sorted: ", v.IsSorted)
		buildGraph(g, v.WordList)
		sorted, isDAG := g.topologicalSort()
		fmt.Println("Got:\tis sorted:", isDAG)
		if v.IsSorted {
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

// Non-recursive DFS
func (g *Graph) topologicalSort() ([]string, bool) {
	sorted := []string{}
	isDAG := true

	dfsStack := dfsStack{}
	color := make(map[string]string)

	for k, _ := range g.edgeMap {
		dfsStack.push(k)

		for len(dfsStack) != 0 {
			top := dfsStack.peek()

			// "black" if node has been processed
			// "grey" if discovered
			// "" if undiscovered
			if color[top] == "black" {
				dfsStack = dfsStack.pop()
				continue
			} else if color[top] == "grey" {
				sorted = pushFront(sorted, top)
				color[top] = "black"
				dfsStack = dfsStack.pop()
				continue
			}

			color[top] = "grey"
			followers := g.edgeMap[top]

			if len(followers) == 0 {
				sorted = pushFront(sorted, top)
				color[top] = "black"
				dfsStack = dfsStack.pop()
			} else {
				for _, node := range followers {
					if color[node] == "grey" {
						isDAG = false
						return sorted, isDAG
					} else if color[node] == "" {
						dfsStack.push(node)
					}
				}

			}
		}

	}

	return sorted, isDAG
}

type dfsStack []string

func (s *dfsStack) push(c string) {
	*s = append(*s, c)
}

func (s *dfsStack) pop() dfsStack {
	stack := *s
	stack = stack[:len(stack)-1]
	return stack
}

func (s *dfsStack) peek() string {
	stack := *s
	return stack[len(stack)-1]
}

func pushFront(sorted []string, node string) []string {
	newSorted := make([]string, len(sorted)+1)
	newSorted[0] = node
	copy(newSorted[1:], sorted)
	return newSorted
}
