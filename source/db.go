package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var sslmode = os.Getenv("SSLMODE")

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

// Create table in databese
func createTable() error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec(`CREATE TABLE users(ID SERIAL PRIMARY KEY, USERNAME TEXT, USER_TG_ID BIGINT UNIQUE, GROUP_ID TEXT);`); err != nil {
		return err
	}

	return nil
}

// Save default group for user
func saveGroupData(username string, userTgId int64, groupID string) error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	data := `INSERT INTO users(username, user_tg_id, group_id) 
    VALUES ($1, $2, $3)
    ON CONFLICT (user_tg_id) DO UPDATE 
    SET username = $1, group_id = $3;`

	if _, err = db.Exec(data, username, userTgId, groupID); err != nil {
		return err
	}

	return nil
}

func getGroupData(userTgId int64) (string, error) {
	var groupID string
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return "", err
	}
	defer db.Close()

	row := db.QueryRow("SELECt group_id FROM users WHERE user_tg_id = $1;", userTgId)
	err = row.Scan(&groupID)
	if err != nil {
		return "", err
	}

	return groupID, nil
}
