package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	Host     = "178.154.254.105"
	Port     = 5432
	User     = "inotech"
	Password = "platex"
	Dbname   = "postgres"
)

type DbConfig struct {
	Host string
	Port int
	User string
	Password string
	Dbname string
}
func (dc DbConfig) Connect() (*sqlx.DB,error){
	db, err := sqlx.Connect(
		"postgres",
		"host="+Host+" user="+User+" dbname="+Dbname+" password="+Password+ " sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}
	fmt.Println("Successfully connected!")
	return db, err
}

func Insert(db *sqlx.DB, r string, table interface{}) error{
	query := r
	fmt.Println(query)
	_, err := db.NamedExec(query, table)
	return err
}
