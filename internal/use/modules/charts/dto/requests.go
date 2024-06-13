package dto

import (
	"time"
)

type PeriodUserChartRequestDTO struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
}
