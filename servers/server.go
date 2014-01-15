package servers

type server struct {
	Address string
}

type GameServerResponse struct {
	Name       string
	Map        string
	Game       string
	Players    byte
	MaxPlayers byte
	Bots       byte
	ServerType byte
	Secured    bool
}

func CreateRequest() []byte {
	return []byte("\xFF\xFF\xFF\xFF")
}
