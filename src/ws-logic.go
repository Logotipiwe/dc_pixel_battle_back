package main

import (
	"fmt"
	"strconv"
)

const (
	Join     = "JOIN"
	Leave    = "LEAVE"
	SetPixel = "SET_PIXEL"
)

func HandleMessage(client *Client, event map[string]string) {
	switch event["type"] {
	case SetPixel:
		row, err := strconv.Atoi(event["row"])
		column, err := strconv.Atoi(event["col"])
		color := event["color"]
		if err != nil {
			fmt.Println(err)
			return
		}
		pixel := Pixel{
			Row:      row,
			Column:   column,
			Color:    color,
			PlayerId: client.User.Id,
		}
		err = pixel.savePixel()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Pixel " + strconv.Itoa(row) + ":" + strconv.Itoa(column) + " is set to color " + color + " by user " + client.User.Id)

		msg := MessageWithClient{
			Message: Message{
				Type: SetPixel,
				Body: event,
			},
			Client: client,
		}
		client.Pool.Broadcast <- msg
		break
	default:
		fmt.Println("no type on message")
	}
}
func (pool *Pool) HandleRegister(client *Client) {
	body := map[string]string{
		"count": strconv.Itoa(len(client.Pool.Clients)),
	}
	msg := Message{
		Type: Join,
		Body: body,
	}
	client.Pool.BroadcastAll <- msg
}
func (pool *Pool) HandleUnregister(client *Client) {
	body := map[string]string{
		"count": strconv.Itoa(len(client.Pool.Clients)),
	}
	msg := MessageWithClient{
		Message: Message{
			Type: Leave,
			Body: body,
		},
		Client: client,
	}
	client.Pool.Broadcast <- msg
}
