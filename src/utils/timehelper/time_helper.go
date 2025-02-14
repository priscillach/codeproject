package timehelper

import (
	"fmt"
	"time"
)

const (
	DaySeconds    = 24 * 60 * 60
	HourSeconds   = 60 * 60
	MinuteSeconds = 60
)

func FormatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := int(duration.Milliseconds()) % 1000
	return fmt.Sprintf("%02d:%02d:%02d.%02d", hours, minutes, seconds, milliseconds)
}
