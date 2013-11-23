package main

import (
	"SteamCondenser/servers"
	"fmt"
)

func main() {
	fmt.Println("SteamCondenser Tests")

	response := servers.GetInfo("74.91.113.128:27015")

	fmt.Println("Hostname: ", response.Name)
	fmt.Println("Map: ", response.Map)
	fmt.Println("Players: ", response.Players, "/", response.MaxPlayers)
	fmt.Println("Server Type: ", string(response.ServerType))
	fmt.Println("Vac: ", response.Vac)
	return
}
