package log

import (
	"fmt"
	"time"
)

type callbackFunc struct {
	now time.Time
	cb  func(duration string)
}

func (c *callbackFunc) Apply() {
	spentTime := formatTimeDuration(time.Since(c.now))
	c.cb(spentTime)
}

func CallTotalDurationLog(cb func(duration string)) func() {
	now := time.Now()
	c := &callbackFunc{now: now, cb: cb}
	return c.Apply
}

func formatTimeDuration(duration time.Duration) string {
	duration = time.Duration(duration.Seconds())
	hours := duration / 3600
	duration %= 3600
	minutes := duration / 60
	duration %= 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, duration)
}
