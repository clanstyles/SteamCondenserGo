package servers

import (
	"SteamCondenserGo/helpers"
	"net"
	"fmt"
)

type GoldServer server

type serverResponse struct {
	Header     byte
	Protocol   byte
	Hostname   string
	Map        string
	Folder     string
	AppId      int64
	Game       string
	NumPlayers byte
	MaxPlayers byte
	Bots       byte
	ServerType byte
	Enviorment byte
	Visibility byte
	Vac        byte
}

func (model GoldServer) GetInfo() (serverResponse, error) {
	resp := serverResponse{}

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
	send := createPacket()
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

func (resp *serverResponse) bufferToResponse(b []byte) {

	reader := helpers.Init(4, b)
	resp.Header = reader.ReadByte()
	resp.Protocol = reader.ReadByte()
	resp.Hostname = reader.ReadNullTermString()
	resp.Map = reader.ReadNullTermString()
	resp.Folder = reader.ReadNullTermString()
	resp.Game = reader.ReadNullTermString()
	resp.AppId = reader.ReadShort()
	resp.NumPlayers = reader.ReadByte()
	resp.MaxPlayers = reader.ReadByte()
	resp.Bots = reader.ReadByte()
	resp.ServerType = reader.ReadByte()
	resp.Enviorment = reader.ReadByte()
	resp.Visibility = reader.ReadByte()
	resp.Vac = reader.ReadByte()
}


func createPacket() []byte {
	return []byte("\xFF\xFF\xFF\xFF")
}

func (self serverResponse) PrintDebug() {
	fmt.Println("Header: ", self.Header)
	fmt.Println("Protocol: ", self.Protocol)
	fmt.Println("Hostname: ", self.Hostname)
	fmt.Println("Map: ", self.Map)
	fmt.Println("Folder: ", self.Folder)
	fmt.Println("Game: ", self.Game)
	fmt.Println("AppId: ", self.AppId)
	fmt.Println("Players: ", self.NumPlayers, "/", self.MaxPlayers)
	fmt.Println("Bots: ", self.Bots)
	fmt.Println("Server Type: ", self.ServerType)
	fmt.Println("Vac: ", self.Vac)
}
