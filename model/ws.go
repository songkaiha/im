package models

import "github.com/gorilla/websocket"

var AllClients = make(map[string]*websocket.Conn, 0)
var Upgrader = websocket.Upgrader{}
