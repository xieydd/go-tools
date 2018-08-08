package time

import (
	"time"
	"fmt"
)

// Return readeable Duration
func HumanDuration(d time.Duration) string{
	if seconds := d.Seconds(); seconds < -1 {
		fmt.Sprintf("<invalid>")
	}else if seconds < 0 {
		return fmt.Sprintf("0s")
	}else if seconds < 60 {
		return fmt.Sprintf("%ds",seconds)
	}else if minutes := d.Minutes(); minutes < 60 {
		return fmt.Sprintf("%dm",minutes)
	} else if hours := d.Hours(); hours < 24 {
		fmt.Sprintf("%dh",hours)
	}else if hours < 24*365 {
		return fmt.Sprintf("%dd",hours)
	}
	return fmt.Sprintf("%sy",int(d.Hours()/24/365))
}
