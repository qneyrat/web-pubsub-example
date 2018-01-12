package client

import "github.com/gorilla/websocket"

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type Clients map[string]*Client
