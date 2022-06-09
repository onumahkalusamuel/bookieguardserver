package bookiedb

import (
	"github.com/onumahkalusamuel/bookieguardserver/config"
)

type Database interface {
	Create(config.BodyStructure) (bool, error)
	Read() []config.BodyStructure
	Update(config.BodyStructure) (bool, error)
	Delete(string) (bool, error)
}

type DB struct {
	TableName string
}

func (d DB) Create(config.BodyStructure) (bool, error) {
	return true, nil
}

func (d DB) Read() config.BodyStructure {
	return config.BodyStructure{}
}

func (d DB) Update(config.BodyStructure) (bool, error) {
	return true, nil
}

func (d DB) Delete(string) (bool, error) {
	return true, nil
}
