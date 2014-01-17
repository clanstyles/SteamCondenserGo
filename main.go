package main

import (
	"SteamCondenserGo/servers"
	"fmt"
)

func main() {
	fmt.Println("SteamCondenser Tests")

//	gameServer := servers.GoldServer{
//		Address: "94.23.120.37:27015",
//	}
//
//	response, err := gameServer.GetInfo()
//	if err != nil {
//		fmt.Println("ERROR:", err)
//		return
//	}
//
//	response.PrintDebug()

	minecraftServer := servers.MinecraftServer{
		Address: "178.32.48.244:25565",
	}

	response, err := minecraftServer.GetInfo()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Main thread response", response)
	return
}
//
//func Print(self servers.Response) {
//	fmt.Println("Hostname: ", self.Name)
//	fmt.Println("Map: ", self.Map)
//	fmt.Println("Players: ", self.Players, "/", self.MaxPlayers)
//	fmt.Println("Server Type: ", string(self.ServerType))
//	fmt.Println("Vac: ", self.Vac)
//}
