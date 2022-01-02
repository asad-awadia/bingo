package bingo

type (
	Decoder func([]byte, uintptr) (int, error)
	Encoder func([]byte, interface{}) (int, error)
)