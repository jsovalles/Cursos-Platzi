package utils

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	ContractorsTable     = "integration.contractors"
	ReferralPartnerTable = "integration.referral_partners"
	ProjectTable         = "integration.projects"
)

type Database struct {
	Client *sqlx.DB
}

var client *sqlx.DB

func NewDatabase(env Env) (database Database, err error) {
	fmt.Printf("Connecting to database")
	/*
		connInfo := fmt.Sprintf("host=%s:%s user=%s password=%s dbname=%s sslmode=disable",
			env.DbHost, env.DbPort, env.DbUsername, env.DbPassword, env.DbSchema)
	*/
	client, err = sqlx.Open("pgx", env.DatabaseUrl)
	if err != nil {
		fmt.Printf("Failed to open connection to database: %s", err.Error())
		return
	}
	if err = client.Ping(); err != nil {
		fmt.Printf("Failed to ping database: %s", err.Error())
		return
	}
	log.Println("Database ready to accept connections")

	database.Client = client

	return
}
