package hello

import (
	"fmt"
	"sync"
)

func helloWG(wg *sync.WaitGroup, id int, name string) {
	defer wg.Done()
	fmt.Printf("Hello, %s! - %d\n", name, id)
}
func helloPeopleWG(wg *sync.WaitGroup, people []string) {
	for k, p := range people {
		wg.Add(1)
		go helloWG(wg, k+1, p)
	}
}

// PeopleWG says hello to people in random order, takes a sync.WaitGroup
func PeopleWG(wg *sync.WaitGroup, people []string) {
	helloPeopleWG(wg, people)
}
