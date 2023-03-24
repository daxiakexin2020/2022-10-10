package model

import (
	"fmt"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	tm := NewMJBarracks()
	log.Println(tm.barracks.armList, tm.armList)

	productionEngineer := ProductionEngineer()
	productionEngineer2 := ProductionEngineer()
	fmt.Println(productionEngineer == productionEngineer2)
}
