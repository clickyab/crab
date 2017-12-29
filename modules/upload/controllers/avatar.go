package controllers

import "clickyab.com/crab/modules/upload/model"

type byteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	kb byteSize = 1 << (10 * iota)
)

func init() {
	Register("avatar", int64(512*kb), model.JPGMime, model.PJPGMime, model.PNGMime)
	Register("banner", int64(512*kb), model.JPGMime, model.PJPGMime, model.PNGMime, model.GifMime)
	Register("native", int64(512*kb), model.JPGMime, model.PJPGMime, model.PNGMime)
}
