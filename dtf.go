package dtf

import (
	"fmt"
	"time"
)

func Fmt(d time.Duration) string {
	switch {
	case d < time.Microsecond:
		return fmt.Sprintf("%dns", d.Nanoseconds())
	case d < time.Millisecond:
		return divf(d, time.Microsecond, "µs")
	case d < time.Second:
		return divf(d, time.Millisecond, "ms")
	case d < time.Minute:
		return divf(d, time.Second, "s")
	case d < time.Hour:
		return divfrecur(d, time.Minute, "m")
	case d < 24*time.Hour:
		return divfrecur(d, time.Hour, "h")
	default:
		return divfrecur(d, 24*time.Hour, "d")
	}
}

func divf(p, q time.Duration, unit string) string {
	if p%q == 0 {
		return fmt.Sprintf("%d%s", p/q, unit)
	}
	return fmt.Sprintf("%.3f%s", float64(p)/float64(q), unit)
}

func divfrecur(p, q time.Duration, unit string) string {
	x, r := p/q, p%q
	if r == 0 {
		return fmt.Sprintf("%d%s", x, unit)
	}
	return fmt.Sprintf("%d%s%s", x, unit, Fmt(r))
}
