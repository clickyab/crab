package controllers

import "clickyab.com/crab/modules/upload/model"

type byteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	kb byteSize = 1 << (10 * iota)
)

func init() {
	Register("user-avatar", int64(512*kb), 15, model.JPGMime, model.PJPGMime, model.PNGMime)

	Register("banner-image", int64(512*kb), 15, model.JPGMime, model.PJPGMime, model.PNGMime, model.GifMime)
	Register("banner-video", int64(16*512*kb), 15, model.VideoMime)

	Register("native-image", int64(512*kb), 15, model.JPGMime, model.PJPGMime, model.PNGMime)
	Register("native-video", int64(16*512*kb), 15, model.VideoMime)

	Register("vast-image", int64(512*kb), 15, model.JPGMime, model.PJPGMime, model.PNGMime)
	Register("vast-video", int64(16*512*kb), 15, model.VideoMime)
}
