package summation

import "time"

const (
	ObservationSummationDayFrequency   = 1
	ObservationSummationWeekFrequency  = 2
	ObservationSummationMonthFrequency = 3
	ObservationSummationYearFrequency  = 4
)

type ObservationSummation struct {
	EntityID  uint64    `json:"entity_id"`
	Start     time.Time `json:"start"`
	Finish    time.Time `json:"finish"`
	Day       int       `json:"day"`
	Week      int       `json:"week"`
	Month     int       `json:"month"`
	Year      int       `json:"year"`
	Frequency int8      `json:"frequency"`
	Result    float64   `json:"result"`
}

type ObservationSummationFilter struct {
	EntityIDs               []uint64  `json:"entity_ids"`
	Frequency               int8      `json:"frequency"`
	StartGreaterThen        time.Time `json:"start_gt,omitempty"`
	StartGreaterThenOrEqual time.Time `json:"start_gte,omitempty"`
	FinishLessThen          time.Time `json:"finish_lt,omitempty"`
	FinishLessThenOrEqual   time.Time `json:"finish_lte,omitempty"`
}
