package main

import (
	"github.com/google/uuid"
	"time"
)

type Pixel struct {
	Row      int
	Column   int
	Color    string
	PlayerId string
}

type PixelDto struct {
	Row      int    `json:"row"`
	Column   int    `json:"column"`
	Color    string `json:"color"`
	PlayerId string `json:"playerId"`
}

type ColorDto struct {
	Color string `json:"color"`
}

type History struct {
	Id       uuid.UUID
	Row      int
	Column   int
	Color    string
	PlayerId string
	Time     time.Time
}

type HistoryDto struct {
	Id       uuid.UUID `json:"id"`
	Row      int       `json:"row"`
	Column   int       `json:"column"`
	Color    string    `json:"color"`
	PlayerId string    `json:"playerId"`
	Time     time.Time `json:"time"`
}

func (h History) toDto() HistoryDto {
	return HistoryDto{
		Id:       h.Id,
		Row:      h.Row,
		Column:   h.Column,
		Color:    h.Color,
		PlayerId: h.PlayerId,
		Time:     h.Time,
	}
}

func (p Pixel) toDto() PixelDto {
	return PixelDto{
		p.Row, p.Column, p.Color, p.PlayerId,
	}
}

var colors = []ColorDto{
	{Color: "white"},
	{Color: "#858585"},
	{Color: "black"},
	{Color: "blue"},
	{Color: "yellow"},
	{Color: "#5F2DF2"},
	{Color: "#FF2501"},
	{Color: "#5CBF0D"},
}

func getDefaultColors() []ColorDto {
	return colors
}

func isColorExist(color string) bool {
	for _, colorDto := range colors {
		if colorDto.Color == color {
			return true
		}
	}
	return false
}
