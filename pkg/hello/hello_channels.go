package hello

import "fmt"

type person struct {
	id   int
	name string
}

func helloCh(pCh chan *person) {
	defer close(pCh) // whoever is sending to the channel is in charge fo closing it!

	// this will block until we get a person in pCh
	p := <-pCh
	fmt.Printf("Hello, %s! - %d\n", p.name, p.id)
}

func peopleChan(people []string) {
	for k, p := range people {
		pCh := make(chan *person)

		// start the goroutine
		go helloCh(pCh)

		// send the person to pCh which will be received inside `helloCh`
		go func(k int, p string) {
			pCh <- &person{
				id:   k + 1,
				name: p,
			}
		}(k, p)
	}

}

// PeopleWG says hello to people in random order, a channel that will announce when it's done
func PeopleChan(people []string, doneCh chan bool) {
	peopleChan(people)
	doneCh <- true
	close(doneCh)
}
