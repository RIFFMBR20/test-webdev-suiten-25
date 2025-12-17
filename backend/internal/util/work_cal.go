package util

import (
	"fmt"
	"time"
)

type WorkOvertime struct {
	WorkDays      int
	OvertimeHours int
}

func CalcWorkOvertimeByCheckout(checkOut time.Time) WorkOvertime {
	base := time.Date(
		checkOut.Year(), checkOut.Month(), checkOut.Day(),
		17, 0, 0, 0,
		checkOut.Location(),
	)

	result := WorkOvertime{WorkDays: 1, OvertimeHours: 0}

	if checkOut.Before(base) {
		// sebelum 17:00 -> (1+0) sesuai default
		return result
	}

	diff := checkOut.Sub(base)
	overtime := int(diff / time.Hour)

	result.WorkDays += overtime / 5
	result.OvertimeHours = overtime % 5

	return result
}

func (w WorkOvertime) String() string {
	return fmt.Sprintf("%d+%d", w.WorkDays, w.OvertimeHours)
}
