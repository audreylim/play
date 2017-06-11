package main

import "fmt"

func main() {
	sorted1 := []string{"xza", "ayh", "ples", "plares", "bhaaz", "bnc"}
	unsorted1 := []string{"xza", "ayh", "eyes", "ples", "plares", "bhaaz", "bnc"}

	g := Graph{}
	g.edgeMap = make(map[string][]string)

	fmt.Println("sorted", sorted1)
	graph1 := buildGraph(g, sorted1)
	fmt.Println(graph1)
	sorted, isDAG := g.topologicalSort()
	fmt.Println("sorted", sorted)
	fmt.Println("isDAG", isDAG)

	g2 := Graph{}
	g2.edgeMap = make(map[string][]string)

	fmt.Println("unsorted", unsorted1)
	graph2 := buildGraph(g2, unsorted1)
	fmt.Println(graph2)
	unsorted, isDAG := g2.topologicalSort()
	fmt.Println("unsorted", unsorted)
	fmt.Println("isDAG", isDAG)
}

func buildGraph(g Graph, words []string) Graph {
	prev := words[0]

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

func (g *Graph) topologicalSort() ([]string, bool) {
	sortedChar := []string{}

	discovered := make(map[string]bool)
	processed := make(map[string]bool)
	dfsStack := dfsStack{}

	for k, _ := range g.edgeMap {
		dfsStack.push(k)

		for len(dfsStack) != 0 {
			top := dfsStack.peek()
			discovered[top] = true

			if _, ok := g.edgeMap[top]; !ok {
				processed[top] = true
				discovered[top] = false
				sortedChar = pushFront(sortedChar, top)
				dfsStack = dfsStack.pop()
			} else if len(g.edgeMap[top]) == 0 {
				sortedChar = pushFront(sortedChar, top)
				processed[top] = true
				discovered[top] = false
				dfsStack = dfsStack.pop()
			} else {
				for _, followers := range g.edgeMap[top] {
					for _, v := range followers {
						if _, ok := processed[string(v)]; !ok {

							if discovered[string(v)] {
								return sortedChar, false
							}

							dfsStack.push(string(v))
						} else {
							sortedChar = pushFront(sortedChar, top)
							processed[top] = true
							discovered[top] = false
							dfsStack = dfsStack.pop()
						}
					}
				}
			}
		}

	}

	return sortedChar, true
}

func pushFront(sortedChar []string, node string) []string {
	newSortedChar := make([]string, len(sortedChar)+1)
	newSortedChar[0] = node
	copy(newSortedChar[1:], sortedChar)
	return newSortedChar
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
