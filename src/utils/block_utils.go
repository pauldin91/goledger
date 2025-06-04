package utils

import "time"

func AdjustDifficulty(lastBlockDifficulty int64, lastBlockTimestamp time.Time, currentTime time.Time, mineRate int64) int64 {
	diff := lastBlockDifficulty
	var start time.Time
	if lastBlockTimestamp.IsZero() {
		start = time.Now().UTC()
	} else {
		start = lastBlockTimestamp
	}
	dur := start.UnixMilli() + int64(mineRate)

	if dur > currentTime.UnixMilli() {
		diff += 1
	} else {
		diff -= 1
		if diff <= 0 {
			diff = 1
		}
	}
	return diff
}
