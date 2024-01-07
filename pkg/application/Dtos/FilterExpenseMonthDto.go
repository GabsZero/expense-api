package dtos

import "time"

type FilterExpenseMonthDto struct {
	Date time.Time `time_format:"2006-01" name:"date" json:"date" form:"date"`
}
