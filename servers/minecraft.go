package servers

import (
	"SteamCondenserGo/helpers"
	"bytes"
	"encoding/binary"
	//"encoding/binary"
	"fmt"
	"net"
	//"strconv"
)

type MinecraftServer server

const (
	statistic = 0x00
	handshake = 0x09
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
	ChallengeToken []byte
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
	challengeCode, err := doChallengeRequest()
	if err != nil {
		return resp, err
	}
	fmt.Println("Challenge:", challengeCode)

	fmt.Println("Hummm getting status")
	serverResponse, err := getStatus(challengeCode)
	if err != nil {
		return resp, err
	}

	fmt.Println("RESPONSE", serverResponse)
	return resp, err
}

func getStatus(challengeCode string) (serverResponse, error) {
	response := serverResponse{}
	//command := createPacket(statistic, challengeCode)
	//conn.Write(command)

	fmt.Println("Waiting for read...")
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

	fmt.Println(response)
	return response, nil
}

func doChallengeRequest() (string, error) {
	request := challengeRequest{
		Magic1:    '\xFE',
		Magic2:    '\xFD',
		Type:      '\x09',
		SessionId: 1337,
	}

	encoder := new(bytes.Buffer)
	binary.Write(encoder, binary.LittleEndian, request)
	conn.Write(encoder.Bytes())

	data := make([]byte, 2048)
	_, _, err := conn.ReadFromUDP(data)
	if err != nil {
		return "", err
	}

	response := new(challengeResponse)
	err = binary.Read(encoder, binary.LittleEndian, response)
	if err != nil {
		return "", err
	}

	fmt.Println(response)
	return string(response.ChallengeToken[:]), nil
}

func (self *challengeResponse) parseChallengeResponse(data []byte) {
	position := 0

	self.Type, position = helpers.ReadByte(data, position)
	self.SessionId, position = helpers.ReadShort(data, position)
	token, position := helpers.ReadNullTermString(data, position)
	self.ChallengeToken = []byte(token)

	fmt.Println("Type", self.Type)
	fmt.Println("SessionId", self.SessionId)
	fmt.Println("Token", self.ChallengeToken)
}

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
