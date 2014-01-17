package servers

type server struct {
	Address string
}

type GameServerResponse struct {
	Name       string
	Map        string
	Game       string
	Players    int
	MaxPlayers int
	Bots       int
	ServerType int
	Secured    bool
}
