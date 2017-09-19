package controllers

type byteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	kb byteSize = 1 << (10 * iota)
)

func init() {
	Register("avatar", int64(512*kb), "image/jpeg", "image/pjpeg", "image/png")
}
