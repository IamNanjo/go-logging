package timeprefix

import (
	"time"

	"github.com/IamNanjo/go-logging/internal/common"
)

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

	return now.Format(tp.Format) + common.ColumnSeparatorString
}
