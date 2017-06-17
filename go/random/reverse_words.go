package main

import "fmt"

func main() {
	sample := []string{"My", "name", "is", "Chris", "Patt"}

	for i, _ := range sample {
		if i == 2 {
			fmt.Println(i, sample)
			break
		}
		a := i + 1
		sample[i], sample[len(sample)-a] = sample[len(sample)-a], sample[i]
	}
	fmt.Println(sample)

}
