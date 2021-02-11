package rotor

import (
	"math/rand"
)

func NewRotor(seed int64, records []string) *Rotor {
	result := &Rotor{
		records: make([]string, len(records)),
	}
	rand.Seed(seed)
	for i, next := range records {
		result.records[i] = next
	}
	rand.Shuffle(len(records), func(i, j int) {
		result.records[i], result.records[j] = result.records[j], result.records[i]
	})
	return result

}

type Rotor struct {
	records []string
}

func (r *Rotor) GetNext() (found bool, next string) {
	if len(r.records) == 0 {
		found = false
		return
	}
	next = r.records[len(r.records)-1]
	r.records = r.records[:len(r.records)-1]
	found = true
	return
}
