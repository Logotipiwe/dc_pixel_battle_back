package main

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

func (p Pixel) toDto() PixelDto {
	return PixelDto{
		p.Row, p.Column, p.Color, p.PlayerId,
	}
}

var colors = []ColorDto{
	{Color: "white"},
	{Color: "black"},
	{Color: "blue"},
	{Color: "yellow"},
	{Color: "green"},
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
