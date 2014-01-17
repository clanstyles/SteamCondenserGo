package servers

import (
	"SteamCondenserGo/helpers"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
//	"time"
	"strconv"
)

type MinecraftServer server

const (
	magic1 = 0xFE
	magic2 = 0xFD
	statistic = 0x00
	handshake = 0x09
	sessionId = 13371
)

type challengeRequest struct {
	Magic1    byte
	Magic2    byte
	Type      byte
	SessionId int32
}

type challengeResponse struct {
	Type           byte
	SessionId      int64
	ChallengeToken string
}

type mcQueryRequest struct {
	Magic1    byte
	Magic2    byte
	Type      byte
	SessionId int32
	Challenge int32
}

type mcServerResponse struct {
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

	code, err := requestChallengeCode()
	if err != nil {
		return resp, err
	}

	basicInfo, err := getStatus(code)
	if err != nil {
		return resp, err
	}

	fmt.Println("Server resposne",basicInfo)
	return resp, err
}

func getStatus(challengeCode string) (mcServerResponse, error) {
	response := mcServerResponse{}

	fmt.Println("Trying to convert challenge code", challengeCode)
	code, err := strconv.Atoi(challengeCode)
	if err != nil {
		return response, err
	}

	request := mcQueryRequest{
		// stupid magic, fuck you mojang
		Magic1:    magic1,
		Magic2:    magic2,
		Type:      statistic,
		SessionId: sessionId,
		Challenge: int32(code),
	}

	ec := new(bytes.Buffer)
	binary.Write(ec, binary.LittleEndian, request)
	fmt.Println("Dump request", hex.Dump(ec.Bytes()))
	conn.Write(ec.Bytes())

	data := make([]byte, 4096)
	_, _, err = conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("Dying")
		return response, err
	}

//	reader := helpers.Init(4, data)
//	response.Type = reader.ReadByte()
//	response.SessionId = reader.ReadShort()
//	response.Motd = reader.ReadNullTermString()
//	response.GameType = reader.ReadNullTermString()
//	response.Map = reader.ReadNullTermString()
//	response.NumPlayers = reader.ReadNullTermString()
//	response.MaxPlayers = reader.ReadNullTermString()
//	response.HostPort = reader.ReadShort()
//	response.HostIp = reader.ReadNullTermString()
//
//	fmt.Println(response)
	return response, nil
}

func requestChallengeCode() (string, error) {
	request := challengeRequest{
		Magic1:    magic1,
		Magic2:    magic2,
		Type:      handshake,
		SessionId: sessionId,
	}

	encoder := new(bytes.Buffer)
	binary.Write(encoder, binary.LittleEndian, request)
	conn.Write(encoder.Bytes())

	data := make([]byte, 2048)
	_, _, err := conn.ReadFromUDP(data)
	if err != nil {
		return "", err
	}

	reader := helpers.Init(4, data)
	response := challengeResponse{}

	response.Type = reader.ReadByte()
	response.SessionId = reader.ReadShort()
	response.ChallengeToken = reader.ReadNullTermString()

	return response.ChallengeToken, nil
}

//func (self *challengeResponse) parseChallengeResponse(data []byte) {
//	reader := helpers.Init(4, data)
//	self.Type = reader.ReadByte()
//	self.SessionId = reader.ReadShort()
//
//	token, position := helpers.ReadNullTermString(data, position)
//	self.ChallengeToken = []byte(token)
//
//	fmt.Println("Type", self.Type)
//	fmt.Println("SessionId", self.SessionId)
//	fmt.Println("Token", self.ChallengeToken)
//}

//func createPacket(command byte, challenge string) []byte {
//	packet := []byte("\xFE\xFD")
//	packet = append(packet, command)

//	buf := new(bytes.Buffer)
//	var codebyte int32 = 1337
//	err := binary.Write(buf, binary.BigEndian, codebyte)
//	if err != nil {
//		fmt.Println("binary.Write failed:", err)
//	}
//	packet = append(packet, buf.Bytes()...)
//	//packet = append(packet, []byte("\x31\x32\x33\x34")...)
//	//packet = append(packet, []byte("\x01\x02\x03\x04")...)

//	if challenge != "" {
//		fmt.Println("Challenge BEFORE:", challenge)
//		code, _ := strconv.Atoi(challenge)
//		fmt.Println("Challenge AFTER:", code)

//		buf := new(bytes.Buffer)
//		var codebyte int32 = int32(code)
//		err := binary.Write(buf, binary.LittleEndian, codebyte)
//		if err != nil {
//			fmt.Println("binary.Write failed:", err)
//		}

//		fmt.Println("Challenge BYTE:", buf.Bytes())
//		//var i int = code
//		//b := make([]byte, 4)
//		//binary.BigEndian.PutUint32(b, uint32(i))

//		packet = append(packet, buf.Bytes()...)
//	}

//	//packet = append(packet, []byte("\x00\x00\x00\x00")...)

//	fmt.Println("Packet", packet)
//	return packet
//}
