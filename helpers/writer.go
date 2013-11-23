package helpers

type Request []byte

func CreateNullTermByteString(data string) []byte {
	arr := make([]byte, 0)

	arr = append(arr, []byte(data)...)
	return append(arr, byte('\x00'))
}
