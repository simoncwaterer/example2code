package ex

import "fmt"

type T struct{}

func (t *T) Bar() {
	fmt.Println("Bar")
}
