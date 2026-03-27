package timeprefix

import "time"

type TimePrefix struct {
	UTC    bool
	Format string
}

func (tp *TimePrefix) Get() string {
	if tp == nil {
		return ""
	}

	now := time.Now()
	if tp.UTC {
		now = now.UTC()
	}

	return now.Format(tp.Format) + " "
}
