package domain

import (
	"player-service/model"

	"go.uber.org/zap"
)

type PlayerStore interface{
	GetAll() (players []model.Player)
	GetById(id string) (player model.Player, ok bool)
}

type PlayerService struct {
	log *zap.SugaredLogger
	store PlayerStore
}

func NewPlayerService(logger *zap.SugaredLogger, store PlayerStore) *PlayerService{
	service := &PlayerService{
		log: logger,
		store: store,
	}
	return service
}

func (svc *PlayerService) GetAll() ([]model.Player){
	return svc.store.GetAll()
}

func (svc *PlayerService) GetById(id string) (player model.Player, ok bool){
	return svc.store.GetById(id)
}