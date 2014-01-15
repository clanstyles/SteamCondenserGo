package main

import (
	"SteamCondenserGo/servers"
	"fmt"
)

func main() {
	fmt.Println("SteamCondenser Tests")

	//goldServer := servers.GoldServer{
	//	Address: "74.91.113.128:27015",
	//}
	//response := goldServer.GetInfo()
	//Print(response)

	minecraftServer := servers.MinecraftServer{
		Address: "mc.ecocitycraft.com:25565",
	}
	response, err := minecraftServer.GetInfo()
	fmt.Println(err)
	fmt.Println(response)
	//Print(response)

	return
}

func Print(self servers.Response) {
	fmt.Println("Hostname: ", self.Name)
	fmt.Println("Map: ", self.Map)
	fmt.Println("Players: ", self.Players, "/", self.MaxPlayers)
	fmt.Println("Server Type: ", string(self.ServerType))
	fmt.Println("Vac: ", self.Vac)
}
