package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	env "github.com/logotipiwe/dc_go_env_lib"
)

type PixelResult struct {
	Row      int
	Column   int
	Color    string
	PlayerId sql.NullString
}

func ConnectDb() (*sql.DB, error) {
	println("Database connected!")
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v", env.GetDbLogin(), env.GetDbPassword(),
		env.GetDbHost(), env.GetDbName())
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
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
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}
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
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	exec, err := db.Exec("insert into pixels (pixel_row, pixel_col, color, player_id) "+
		"VALUES (?, ?, ?, ?) on duplicate key update color = ?, player_id = ?",
		p.Row, p.Column, p.Color, p.PlayerId, p.Color, p.PlayerId)
	if err != nil {
		return err
	}
	println(exec)
	return nil
}
