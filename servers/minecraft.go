package servers

import (
	"SteamCondenserGo/helpers"
	"fmt"
	"net"
)

type MinecraftServer server

const (
	statistic = 0x00
	handshake = 0x09
)

type challengeResponse struct {
	Type           byte
	SessionId      int64
	ChallengeToken string
}

type serverResponse struct {
	Type       byte
	SessionId  int64
	Motd       string
	GameType   string
	Map        string
	NumPlayers string
	MaxPlayers string
	HostPort   int64
	HostIp     string
}

var conn *net.UDPConn

func (self MinecraftServer) GetInfo() (GameServerResponse, error) {
	resp := GameServerResponse{}
	serverAddr, err := net.ResolveUDPAddr("udp", self.Address)
	if err != nil {
		return resp, err
	}

	conn, err = net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return resp, err
	}
	defer conn.Close()

	fmt.Println("Hummm getting challenge")
	challengeCode, err := challengeRequest()
	if err != nil {
		return resp, err
	}
	fmt.Println(challengeCode)

	fmt.Println("Hummm getting status")
	serverResponse, err := getStatus(challengeCode)
	if err != nil {
		return resp, err
	}

	fmt.Println(serverResponse)
	return resp, err
}

func getStatus(challengeCode string) (serverResponse, error) {
	response := serverResponse{}
	command := createPacket(statistic, challengeCode)
	conn.Write(command)

	data := make([]byte, 2048)
	_, _, err := conn.ReadFromUDP(data)
	if err != nil {
		return response, err
	}

	position := 0
	response.Type, position = helpers.ReadByte(data, position)
	response.SessionId, position = helpers.ReadShort(data, position)
	response.Motd, position = helpers.ReadNullTermString(data, position)
	response.GameType, position = helpers.ReadNullTermString(data, position)
	response.Map, position = helpers.ReadNullTermString(data, position)
	response.NumPlayers, position = helpers.ReadNullTermString(data, position)
	response.MaxPlayers, position = helpers.ReadNullTermString(data, position)
	response.HostPort, position = helpers.ReadShort(data, position)
	response.HostIp, position = helpers.ReadNullTermString(data, position)

	return response, nil
}

func challengeRequest() (string, error) {
	command := createPacket(handshake, "")
	conn.Write(command)

	data := make([]byte, 2048)
	_, _, err := conn.ReadFromUDP(data)
	if err != nil {
		return "", err
	}

	resp := challengeResponse{}
	resp.parseChallengeResponse(data)

	return resp.ChallengeToken, nil
}

func (self *challengeResponse) parseChallengeResponse(data []byte) {
	position := 0

	self.Type, position = helpers.ReadByte(data, position)
	self.SessionId, position = helpers.ReadShort(data, position)
	self.ChallengeToken, position = helpers.ReadNullTermString(data, position)
}

func createPacket(command byte, challenge string) []byte {
	packet := []byte("\xFE\xFD")
	packet = append(packet, command)
	packet = append(packet, []byte("\x01\x02\x03\x04")...)

	if challenge != "" {
		packet = append(packet, []byte(challenge)...)
	}

	packet = append(packet, []byte("\x00\x00\x00\x00")...)

	fmt.Println(packet)
	return packet
}
