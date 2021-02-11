package main

import (
	"fmt"
	"hermectic-random-rotation/rotor"
	"time"
)

func main() {
	count := 10
	customers := make([]rotor.Rotorable, 0)
	for i := 0; i < count; i++ {
		next := Record{
			Key: fmt.Sprint(i),
		}
		customers = append(customers, next)
	}

	seed := int64(time.Now().Minute() % 5)
	r := rotor.NewRotor(seed, customers)

	fmt.Println("List: ")
	found := true
	var next rotor.Rotorable
	for found {
		found, next = r.GetNext()
		if found {
			fmt.Print(next.GetID())
		}
	}

}

type Record struct {
	Key string
}

func (r Record) GetID() string {
	return r.Key
}
