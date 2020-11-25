package config

import (
	"log"

	"github.com/gofiber/session/v2"
	"github.com/gofiber/session/v2/provider/sqlite3"
)

func Sessions() (sessions *session.Session) {

	provider, err := sqlite3.New(sqlite3.Config{
		DBPath:    "./test.db",
		TableName: "session",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	sessions = session.New(session.Config{
		Provider: provider,
	})

	return sessions
}
