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
	{Color: "#5F2DF2"},
	{Color: "#4D2D9B"},
	{Color: "#858585"},
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
