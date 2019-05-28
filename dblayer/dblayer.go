package dblayer

import (
	"errors"
	"log"
)

const (
	mongo = "mongodb"
	mysql = "mysql"
)

var ErrDBTypeNotFound = errors.New("database type not found")

type DBLayer interface {
	AddMember(cm *CrewMember) error
	FindMember(id int) (CrewMember, error)
	AllMembers() (crew, error)
}

type CrewMember struct {
	ID           int    `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	SecClearance int    `json:"clearance" bson:"security clearance"`
	Position     string `json:"position" bson:"position"`
}

type crew []CrewMember

func ConnectDatabase(o string, cstring string) (DBLayer, error) {
	switch o {
	case mongo:
		return NewMongoStore(cstring)
	case mysql:
		return NewMySQLDataStore(cstring)
	}
	log.Println("Could not find ", o)
	return nil, ErrDBTypeNotFound
}
