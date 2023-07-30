package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	. "github.com/logotipiwe/dc_go_utils/src/config"
)

type PixelResult struct {
	Row      int
	Column   int
	Color    string
	PlayerId sql.NullString
}

var db *sql.DB

func InitDb() error {
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v", GetConfig("DB_LOGIN"), GetConfig("DB_PASS"),
		GetConfig("DB_HOST"), GetConfig("DB_NAME"))
	conn, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return err
	}
	if err := conn.Ping(); err != nil {
		println(fmt.Sprintf("Error connecting database: %s", err))
		return err
	}
	db = conn
	println("Database connected!")
	return nil
}

func toNullable(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func LoadAllPixels() ([]Pixel, error) {
	var pixels []Pixel
	rows, err := db.Query("SELECT * FROM pixels")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := PixelResult{}
		err = rows.Scan(&res.Row, &res.Column, &res.Color, &res.PlayerId)
		if err != nil {
			return nil, err
		}
		pixels = append(pixels, toModel(res))
	}
	return pixels, nil
}

func toModel(p PixelResult) Pixel {
	return Pixel{
		p.Row,
		p.Column,
		p.Color,
		p.PlayerId.String,
	}
}

func (p Pixel) savePixel() error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into pixels (pixel_row, pixel_col, color, player_id) "+
		"VALUES (?, ?, ?, ?) on duplicate key update color = ?, player_id = ?",
		p.Row, p.Column, p.Color, p.PlayerId, p.Color, p.PlayerId)
	if err != nil {
		tx.Rollback()
		return err
	}
	historyId := uuid.NewString()
	_, err = tx.Exec("insert into history (id, pixel_row, pixel_col, color, user_id) "+
		"VALUES (?,?,?,?,?)",
		historyId, p.Row, p.Column, p.Color, p.PlayerId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
