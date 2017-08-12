package controllers

type ByteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	kb ByteSize = 1 << (10 * iota)
	mb
	gb
)

func init() {
	Register("avatar", int64(512*kb), "image/jpeg", "image/pjpeg", "image/png")
}
