package libs

import "time"

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
