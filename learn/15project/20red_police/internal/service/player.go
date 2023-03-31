package service

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"20red_police/tools"
	"errors"
	"fmt"
	"sync"
	"time"
)

type PlayerService struct {
	playerRepo       data.Player
	roomRepo         data.Room
	architectureRepo data.Architecture
}

func NewPlayerService(playerRepo data.Player, roomRepo data.Room, architectureRepo data.Architecture) *PlayerService {
	return &PlayerService{
		playerRepo:       playerRepo,
		roomRepo:         roomRepo,
		architectureRepo: architectureRepo,
	}
}

func (ps *PlayerService) Create(name string) (*model.Player, error) {
	player := model.NewPlayer(name)
	return player, nil
}

func (ps *PlayerService) BuildArchitecture(player *model.Player, roomID string, arName string) error {
	/**
	money
	sort    todo
	count
	*/

	//modey
	architecture, err := ps.architectureRepo.FetchArchitecture(arName)
	if err != nil {
		return err
	}
	needPrice := architecture.FetchArchitectureConstructionPrice()
	thredId := tools.UUID()
	if !player.SetIsBuildingAR(thredId) {
		return errors.New("There are buildings being built ")
	}
	var buildOk bool
	var tmoney int32
	var wg sync.WaitGroup

	timer := time.NewTimer(time.Second * 10)

	go func() {
		for {
			select {
			case <-timer.C:
				player.AddToCh(20000)
			default:
			}
		}
	}()

	go func() {
		defer wg.Done()
		wg.Add(1)
		for !buildOk {
			fmt.Println("***************************select**********************")
			pickMoney := player.PickFormCh()
			tmoney += pickMoney
			fmt.Println("********************pick money***************", tmoney, pickMoney, tmoney-needPrice)
			if tmoney >= needPrice {
				buildOk = true
				player.AddToCh(tmoney - needPrice)
				player.SetIsNotBuildingAR(thredId)
			}
		}
	}()
	wg.Wait()

	a := <-player.Mch
	fmt.Println("aaaaa:", a)
	//count add
	return nil
}
