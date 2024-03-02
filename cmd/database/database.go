package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/alcb1310/kanban/cmd/types"
)

var dbdir = os.Getenv("XDG_DATA_HOME") + "/kanban"
var file = dbdir + "/kanban.db"

type Database struct {
	*sql.DB
}

func Connect() Database {
	if dbdir == "/kanban" {
		dbdir = os.Getenv("HOME") + "/.local/state/kanban"
		file = dbdir + "/kanban.db"
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal("could not open database", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		os.MkdirAll(dbdir, 0755)
		os.Create(file)
	}

	x := Database{db}
	x.CreateSchema()

	return x
}

func (d *Database) CreateSchema() {
	_, err := d.Exec("CREATE TABLE IF NOT EXISTS lists (id INTEGER PRIMARY KEY, title TEXT NOT NULL, status INTEGER NOT NULL);")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (db *Database) GetByStatus(status uint8) []types.List {
	sql := "SELECT id, title, status FROM lists WHERE status = ?"
	rows, err := db.Query(sql, status)
	if err != nil {
		log.Fatal(err)
	}

	var lists []types.List
	for rows.Next() {
		var l types.List
		err := rows.Scan(&l.ID, &l.Title, &l.Status)
		if err != nil {
			log.Fatal(err)
		}
		lists = append(lists, l)
	}
	return lists
}

func (db *Database) UpdateList(l types.List) {
	sql := "UPDATE lists set status = status +1 WHERE id = ?"
	db.Exec(sql, l.ID)
}

func (db *Database) AddList(l types.List) {
	sql := "INSERT INTO lists (title, status) VALUES (?, ?)"
	db.Exec(sql, l.Title, types.TODO)
}
