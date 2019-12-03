package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jaimem88/goroutines/pkg/hello"
)

func main() {
	people := []string{"anne", "bob", "chris", "diana"}

	// 1 - no wait groups or channels, just use sleep
	fmt.Printf("hello using sleep...\n\n")
	helloWithSleep(people)
	// 2 - using a wait group
	fmt.Printf("hello using wait group...\n\n")
	helloWithWG(people)

	fmt.Printf("hello using channels...\n\n")
	// 3 - using channels
	helloWithChan(people)

}

func helloWithSleep(people []string) {
	hello.People(people)
	fmt.Println("sleeping for a bit to let all goroutines finish")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("all goroutines should have finished by now!")
}

func helloWithWG(people []string) {
	var wg sync.WaitGroup
	hello.PeopleWG(&wg, people)
	fmt.Println("waiting for wg to be done...")
	wg.Wait()
	fmt.Println("waiting for wg done!")
}

func helloWithChan(people []string) {
	doneCh := make(chan bool)
	// this needs to start in a go routine
	go hello.PeopleChan(people, doneCh)

	// waut for the function to finish
	<-doneCh

	fmt.Println("channel closed, we're done!")

}
