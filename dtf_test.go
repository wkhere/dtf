package dtf

import (
	"testing"
	"time"
)

func TestFmt(t *testing.T) {
	var tab = []struct {
		d time.Duration
		s string
	}{
		{0, "0ns"},
		{1, "1ns"},
		{100, "100ns"},
		{999, "999ns"},
		{time.Microsecond, "1µs"},
		{999999, "999.999µs"},
		{time.Millisecond, "1ms"},
		{999 * time.Millisecond, "999ms"},
		{999*time.Millisecond + 456*time.Microsecond, "999.456ms"},
		{999*time.Millisecond + 456*time.Microsecond + 499, "999.456ms"},
		{999*time.Millisecond + 456*time.Microsecond + 501, "999.457ms"},
		{time.Second, "1s"},
		{59 * time.Second, "59s"},
		{59*time.Second + 123*time.Millisecond, "59.123s"},
		{59*time.Second + 123*time.Millisecond + 499*time.Microsecond, "59.123s"},
		{59*time.Second + 123*time.Millisecond + 501*time.Microsecond, "59.124s"},
		{60 * time.Second, "1m"},
		{61 * time.Second, "1m1s"},
		{121 * time.Second, "2m1s"},
		{60 * time.Minute, "1h"},
		{61 * time.Minute, "1h1m"},
		{61*time.Minute + 15*time.Second, "1h1m15s"},
		{61*time.Minute + 15*time.Second + 123*time.Millisecond, "1h1m15.123s"},
		{61*time.Minute + 15*time.Second + 123*time.Millisecond + 499*time.Microsecond, "1h1m15.123s"},
		{61*time.Minute + 15*time.Second + 123*time.Millisecond + 501*time.Microsecond, "1h1m15.124s"},
		{23*time.Hour + 59*time.Minute + 59*time.Second, "23h59m59s"},
		{23*time.Hour + 59*time.Minute + 59*time.Second + 999*time.Millisecond, "23h59m59.999s"},
		{24 * time.Hour, "1d"},
		{24*time.Hour + 1*time.Minute, "1d1m"},
		{24*time.Hour + 1*time.Hour, "1d1h"},
		{24*time.Hour + 1*time.Hour + 1*time.Minute, "1d1h1m"},
		{24*time.Hour + time.Hour + time.Minute + time.Second, "1d1h1m1s"},
		{48 * time.Hour, "2d"},
		{7 * 24 * time.Hour, "7d"},
	}

	for i, tc := range tab {
		res := Fmt(tc.d)
		if res != tc.s {
			t.Errorf("tc[%d] mismatch\nhave %v\nwant %v",
				i, res, tc.s)
		}
	}
}
