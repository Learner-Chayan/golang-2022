package practical

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("\n rect %v's length should be greater than 0", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("\n rect %v's length should be greater than 0", r)
		wg.Done()
	}

	area := r.length * r.width
	fmt.Printf("\n rect %v's area %d", r, area)
	wg.Done()
}

func Calculate() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}

	rects := []rect{r1, r2, r3}

	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("\n All go routines finished executing")

}
