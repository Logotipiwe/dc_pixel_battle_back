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

func (p Pixel) toDto() PixelDto {
	return PixelDto{
		p.Row, p.Column, p.Color, p.PlayerId,
	}
}
