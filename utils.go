package timex

import (
	"time"
)

const (
	CSTZoneOffset = 8 * 60 * 60
)

func Now() time.Time {
	return time.Now().In(CST)
}

func UnixTimestamp() uint64 {
	t := Now()
	return uint64(t.UnixNano() / 1e6)
}

var (
	UTC *time.Location
	CST *time.Location
)

func init() {
	UTC = time.UTC
	CST = time.FixedZone("CST", CSTZoneOffset)
}
