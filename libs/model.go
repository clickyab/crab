package libs

import (
	"net/http"
	"time"

	"net/url"

	"github.com/clickyab/services/gettext/t9e"
)

var epoch time.Time

func TimeToID(d time.Time) int64 {
	d = d.Truncate(time.Hour * 24)
	h := int64(d.Sub(epoch).Hours())
	return (h / 24) + 1
}

func IDToTime(d int64) time.Time {
	return epoch.AddDate(0, 0, int(d-1))
}

func init() {
	epoch, _ = time.Parse("20060102", "20180101")
}

// Redirect http redirect
func Redirect(w http.ResponseWriter, code int, url *url.URL) error {
	if code < http.StatusMultipleChoices || code > http.StatusTemporaryRedirect {
		return t9e.G("invalid redirect code")
	}
	w.Header().Set("Location", url.String())
	w.WriteHeader(code)
	return nil
}
