package repository

import (
	"fmt"

	modelsdb "github.com/wrtgvr/microblog/internal/models/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func MustOpenDb(dsn string) *Database {
	db, err := OpenDb(dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func OpenDb(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database open error: %w", err)
	}

	err = db.AutoMigrate(&modelsdb.User{}, &modelsdb.Post{})
	if err != nil {
		return nil, err
	}

	return &Database{
		Conn: db,
	}, nil
}

func (DB *Database) CloseConn() {
	sqlDB, _ := DB.Conn.DB()
	sqlDB.Close()
}
