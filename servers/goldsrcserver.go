package servers

import (
	"SteamCondenserGo/helpers"
	"net"
)

type GoldServer server

type Response struct {
	Header     byte
	Protocol   byte
	Name       string
	Map        string
	Folder     string
	AppId      int64
	Game       string
	Players    byte
	MaxPlayers byte
	Bots       byte
	ServerType byte
	Enviorment byte
	Visibility byte
	Vac        byte
}

func (model GoldServer) GetInfo() (Response, error) {
	resp := Response{}

	serverAddr, err := net.ResolveUDPAddr("udp", model.Address)
	if err != nil {
		return resp, err
	}

	socket, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return resp, err
	}
	defer socket.Close()

	query := helpers.CreateNullTermByteString("TSource Engine Query")
	send := CreateRequest()
	send = append(send, query...)

	_, err = socket.Write(send)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 4096)
	_, _, err = socket.ReadFromUDP(data)
	if err != nil {
		panic(err)
	}

	resp.bufferToResponse(data)
	return resp, nil
}

func (resp *Response) bufferToResponse(b []byte) {
	position := 4

	header, position := helpers.ReadByte(b, position)
	protocol, position := helpers.ReadByte(b, position)
	name, position := helpers.ReadNullTermString(b, position)
	serverMap, position := helpers.ReadNullTermString(b, position)
	folder, position := helpers.ReadNullTermString(b, position)
	game, position := helpers.ReadNullTermString(b, position)
	appId, position := helpers.ReadShort(b, position)
	players, position := helpers.ReadByte(b, position)
	maxPlayers, position := helpers.ReadByte(b, position)
	bots, position := helpers.ReadByte(b, position)
	serverType, position := helpers.ReadByte(b, position)
	enviorment, position := helpers.ReadByte(b, position)
	visibility, position := helpers.ReadByte(b, position)
	vac, position := helpers.ReadByte(b, position)

	resp.Header = header
	resp.Protocol = protocol
	resp.Name = name
	resp.Map = serverMap
	resp.Folder = folder
	resp.Game = game
	resp.AppId = appId
	resp.Players = players
	resp.MaxPlayers = maxPlayers
	resp.Bots = bots
	resp.ServerType = serverType
	resp.Enviorment = enviorment
	resp.Visibility = visibility
	resp.Vac = vac
}
