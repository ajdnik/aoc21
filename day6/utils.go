package day6

func SimulateDay(timers []int64) []int64 {
	for idx, timer := range timers {
		if timer == 0 {
			timers = append(timers, 8)
			timers[idx] = 6
		} else {
			timers[idx]--
		}
	}
	return timers
}

func SimulateDays(timers []int64, days int) []int64 {
	for day := 0; day < days; day++ {
		timers = SimulateDay(timers)
	}
	return timers
}
