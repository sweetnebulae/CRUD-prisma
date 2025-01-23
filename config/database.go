package config

import (
	"github.com/rs/zerolog/log"
	"go_prisma/prisma/db"
)

func ConnectDB() (*db.PrismaClient, error) {

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	log.Info().Msg("Connected to database!")
	return client, nil
}
