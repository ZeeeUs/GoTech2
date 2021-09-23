package store

import (
	"encoding/json"
	"github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/model"
	"io/ioutil"
)

type Store struct {
	config *Config
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() ([]model.Event, error){
	file, err := ioutil.ReadFile(s.config.DatabaseURL)
	if err != nil {
		return []model.Event{}, err
	}

	var events []model.Event

	if err :=json.Unmarshal(file, &events); err != nil {
		return []model.Event{}, err
	}

	return events, nil
}

func (s *Store) Close() {
}
