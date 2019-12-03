package hello

import "fmt"

func hello(id int, name string) {
	fmt.Printf("Hello, %s! - %d\n", name, id)
}

func people(people []string) {
	for k, p := range people {
		go hello(k+1, p)
	}
}

// People will say hello to all the people in random order
func People(names []string) {
	people(names)
}
