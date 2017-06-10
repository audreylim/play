package main

import "fmt"

func main() {
	sorted1 := []string{"xza", "ayh", "ples", "plares", "bhaaz", "bnc"}
	//	unsorted1 := []string{"xza", "ayh", "eyes", "ples", "plares", "bhaaz", "bnc"}

	fmt.Println(sorted1)

	g := Graph{}
	g.edgeMap = make(map[string][]string)
	graph := buildGraph(g, sorted1)
	fmt.Println(graph)
	sorted := g.topologicalSort()
	fmt.Println("sorted", sorted)
}

func buildGraph(g Graph, words []string) Graph {
	prev := words[0]
	g.addVertex(string(prev[0]))
	g.firstNode = string(words[0][0])

	for i := 1; i < len(words); i++ {
		word := words[i]
		g.addVertex(string(word[0]))
		generateRelationship(g, string(prev), words[i])
		prev = words[i]
	}

	return g
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
	edgeMap   map[string][]string
	firstNode string
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

func (g *Graph) topologicalSort() []string {
	sortedChar := []string{}

	//	discovered := make(map[string]bool)
	processed := make(map[string]bool)
	dfsStack := dfsStack{}

	//root := g.firstNode
	root := "x"
	dfsStack.push(root)

	for len(dfsStack) != 0 {
		top := dfsStack.peek()

		if _, ok := g.edgeMap[top]; !ok {
			processed[top] = true
			sortedChar = append(sortedChar, top)
		} else if len(g.edgeMap[top]) == 0 {
			sortedChar = append(sortedChar, top)
			processed[top] = true
			dfsStack = dfsStack.pop()
		} else {
			for _, followers := range g.edgeMap[top] {
				for _, v := range followers {
					if _, ok := processed[string(v)]; !ok {
						dfsStack.push(string(v))
					} else {
						sortedChar = append(sortedChar, top)
						processed[top] = true
						dfsStack = dfsStack.pop()
					}
				}
			}
		}

	}

	return sortedChar
}

type dfsStack []string

func (s *dfsStack) push(c string) {
	*s = append(*s, c)
}

func (s *dfsStack) pop() dfsStack {
	sa := *s
	sa = sa[:len(sa)-1]
	return sa
}

func (s *dfsStack) peek() string {
	sa := *s
	return sa[len(sa)-1]
}
