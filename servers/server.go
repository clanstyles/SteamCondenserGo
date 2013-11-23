package servers

type Server struct {
	Address string
	Port    int
}

func CreateRequest() []byte {
	return []byte("\xFF\xFF\xFF\xFF")
}
