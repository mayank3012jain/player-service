package repository

import (
	"os"
	"player-service/model"

	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
)

type PlayerCsvStore struct{
	log *zap.SugaredLogger
	players []model.Player
}

func NewPlayerCsvStore (logger *zap.SugaredLogger, playersFile string) *PlayerCsvStore {
	store := &PlayerCsvStore{
		log: logger,
	}
	store.loadPlayers(playersFile)
	return store
}

func (st *PlayerCsvStore) loadPlayers(filePath string) {
	st.log.Infow("Loading players data into memory from csv")
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = gocsv.UnmarshalFile(file, &st.players)
	if err != nil {
		panic(err)
	}
}

func (st *PlayerCsvStore) GetAll() (players []model.Player) {
	players = make([]model.Player, len(st.players))
	copy(players, st.players)
	return 
}

func (st *PlayerCsvStore) GetById(id string) (player model.Player, ok bool) {
	for _, pl := range st.players {
		if pl.ID == id {
			player = pl
			ok = true
			return 
		}
	}
	return
}