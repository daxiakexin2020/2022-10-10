package service

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"20red_police/tools"
	"errors"
	"log"
	"time"
)

type PlayerService struct {
	playerRepo       data.Player
	roomRepo         data.Room
	architectureRepo data.Architecture
	armRepo          data.Arm
}

func NewPlayerService(playerRepo data.Player, roomRepo data.Room, architectureRepo data.Architecture, armRepo data.Arm) *PlayerService {
	return &PlayerService{
		playerRepo:       playerRepo,
		roomRepo:         roomRepo,
		architectureRepo: architectureRepo,
		armRepo:          armRepo,
	}
}

func (ps *PlayerService) Create(name string, initPrice int32) (*model.Player, error) {
	player := model.NewPlayer(name, initPrice)
	return player, nil
}

func (ps *PlayerService) BuildArchitecture(player *model.Player, roomID string, arName string) error {

	architecture, err := ps.architectureRepo.FetchArchitecture(arName)
	if err != nil {
		return err
	}
	needPrice := architecture.FetchArchitectureConstructionPrice()

	threaId := tools.UUID()
	if !player.SetBuildingAR(threaId) {
		return errors.New("There are buildings being built ")
	}
	var buildOk bool
	var tmoney int32

	go func() {
		outTime := time.NewTicker(time.Second * model.BuildARTimeout)
		defer outTime.Stop()
		for !buildOk {
			select {
			case pickMoney := <-player.PickReadMch():
				tmoney += pickMoney
				if tmoney >= needPrice {
					buildOk = true
					player.BuildAROver(tmoney-needPrice, threaId, arName)
					return
				}
			case <-outTime.C:
				log.Println("build ar timeout")
				return
			default:
				if player.LenGameOverch() == 1 {
					log.Println("game over, build ar out～～～～～～～～～～")
					return
				}
			}
		}
	}()
	return nil
}

func (ps *PlayerService) BuildArm(player *model.Player, roomID string, armName string) error {

	arm, err := ps.armRepo.FetchArm(armName)
	if err != nil {
		return err
	}
	needPrice := arm.FetchArmConstructionPrice()

	threadId := tools.UUID()
	if !player.SetBuildingARM(threadId) {
		return errors.New("There are buildings being built ")
	}

	var buildOk bool
	var tmoney int32

	go func() {
		outTime := time.NewTicker(time.Second * model.BuildARMTimeout)
		defer outTime.Stop()
		for !buildOk {
			select {
			case pickMoney := <-player.PickReadMch():
				tmoney += pickMoney
				if tmoney >= needPrice {
					buildOk = true
					player.BuildARMOver(tmoney-needPrice, threadId, armName)
					if ps.armRepo.IsMineCar(armName) {
						go player.MockAddPirceIntoCh()
					}
					return
				}
			case <-outTime.C:
				log.Println("build arm timeout")
				return
			default:
				if player.LenGameOverch() == 1 {
					log.Println("game over, build arm out～～～～～～～～～～")
					return
				}
			}
		}
	}()
	return nil
}
