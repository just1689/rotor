package main

import (
	"fmt"
	"hermectic-random-rotation/rotor"
	"time"
)

func main() {
	count := 10
	customers := make([]string, 0)
	for i := 0; i < count; i++ {
		next := fmt.Sprint(i)
		customers = append(customers, next)
	}

	seed := int64(time.Now().Minute() % 5)
	r := rotor.NewRotor(seed, customers)

	fmt.Println("List: ")
	found := true
	var next string
	for found {
		found, next = r.GetNext()
		if found {
			fmt.Print(next)
		}
	}

}

type Record struct {
	Key string
}

func (r Record) GetID() string {
	return r.Key
}
