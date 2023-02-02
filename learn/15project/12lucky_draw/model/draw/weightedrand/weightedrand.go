package weightedrand

import (
	weightedrand "github.com/mroth/weightedrand/v2"
	"math/rand"
	"time"
)

type Weightedrand struct {
	choices map[any]int32
}

func NewWeightedrand(choices map[any]int32) *Weightedrand {
	return &Weightedrand{choices: choices}
}

func (wd *Weightedrand) Pick() interface{} {

	rand.Seed(time.Now().UTC().UnixNano())
	choices := []weightedrand.Choice[any, int32]{}
	for k, v := range wd.choices {
		choices = append(choices, weightedrand.NewChoice(k, v))
	}
	chooser, _ := weightedrand.NewChooser(choices...)
	return chooser.Pick()
}
