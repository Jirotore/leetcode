package leetcode

import "fmt"

const (
	hCount      = 4
	lCount      = 6
	maskMinutes = 1<<lCount - 1
	maskHours   = (1<<hCount - 1) << lCount
	maskClock   = maskHours + maskMinutes
	maxClock    = (11 << lCount) + 59
)

const dateTimeFmt = "%d:%.02d"

func readBinaryWatch(turnedOn int) []string {
	if turnedOn == 0 {
		return []string{"0:00"}
	}

	res := []string{}

	if turnedOn > 8 {
		return res
	}

	digs := []int{1 << 0, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9}

	digs2 := recs(0, turnedOn, digs)
	for x := range digs2 {
		h := (maskHours & x) >> lCount
		m := maskMinutes & x
		if !(0 <= h && h <= 11 && 0 <= m && m <= 59) {
			continue
		}
		res = append(res, fmt.Sprintf(dateTimeFmt, h, m))
	}

	return res
}

func recs(timer, depth int, digs []int) map[int]struct{} {
	res := map[int]struct{}{}

	if depth <= 0 {
		res[timer] = struct{}{}
		return res
	}

	for i, x := range digs {
		y := timer | x

		xxx := recs(y, depth-1, digs[i+1:])

		for d := range xxx {
			if _, ok := res[d]; !ok {
				res[d] = struct{}{}
			}
		}
	}

	return res
}
