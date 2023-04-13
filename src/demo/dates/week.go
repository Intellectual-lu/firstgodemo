package dates

const Days = 7

func WeekToDay(weeks int) int {
	return weeks * Days
}

func DayToWeek(days int) float64 {
	return float64(days) / float64(Days)
}
