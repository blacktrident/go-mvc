package store

import (
	"github.com/blacktrident/go-mvc/model"
	"github.com/blacktrident/go-mvc/mongoconfig"
	"gopkg.in/mgo.v2/bson"
)
import "gopkg.in/mgo.v2"

var connectString string = mongoconfig.Config

func GetSession(collection string, pk string) *mgo.Session {
	// Dial to database and Return mgo session
	var session *mgo.Session
	var err error
	session, err = mgo.Dial(connectString)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	ensureIndex(collection, pk, session)
	return session
}

func ensureIndex(collection string, pk string, s *mgo.Session) {
	session := s.Copy()
	defer session.Close()
	c := session.DB("playstation").C(collection)
	index := mgo.Index{
		Key:        []string{pk},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func Save(game *model.Game) error {
	session := GetSession("Game", "name")
	session = session.Copy()
	defer session.Close()
	c := session.DB("playstation").C("Game")
	game.ID = bson.NewObjectId()
	err := c.Insert(&game)
	return err
}

func GetAll() ([]model.Game, error) {
	session := GetSession("Game", "name")
	session = session.Copy()
	defer session.Close()
	c := session.DB("playstation").C("Game")
	var games []model.Game
	err := c.Find(bson.M{}).All(&games)
	return games, err
}

func GetOne(Name string) (*model.Game, error) {
	session := GetSession("Game", "name")
	session = session.Copy()
	defer session.Close()
	c := session.DB("playstation").C("Game")
	var game model.Game
	var err error
	err = c.Find(bson.M{"name": Name}).One(&game)
	return &game, err
}
