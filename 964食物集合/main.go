package main

import (
	"math"
	"sort"
)

/**
 * @param lunch: an array that contains each lunch food's calories and value
 * @param dinner: an array that contains each dinner food's calories and value
 * @param t: the minest limit value
 * @return: return the min calories
 */
type Foods [][]int

func (ct Foods) Len() int           { return len(ct) }
func (ct Foods) Swap(i, j int)      { ct[i], ct[j] = ct[j], ct[i] }
func (ct Foods) Less(i, j int) bool { return ct[i][1] < ct[j][1] }

func GetMinCalories(lunch [][]int, dinner [][]int, t int) int {
	// write your code here
	hot := math.MaxInt
	badlunch := lunch
	baddinner := dinner
	sort.Sort(Foods(badlunch))
	sort.Sort(Foods(baddinner))
	for _, vi := range badlunch {
		if vi[1] >= t {
			if vi[0] < hot {
				hot = vi[0]
			}
			continue
		}
		for j := len(baddinner) - 1; j >= 0; j-- {
			if baddinner[j][1] >= t {
				if baddinner[j][0] < hot {
					hot = baddinner[j][0]
				}
				continue
			}
			if vi[1]+baddinner[j][1] >= t {
				if vi[0]+baddinner[j][0] < hot {
					hot = vi[0] + baddinner[j][0]
				}
			} else {
				goto Outer
			}
		}
	Outer:
	}
	if hot == math.MaxInt {
		return -1
	}
	return hot
}
