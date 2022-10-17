package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Int []int

func (a Int) Len() int           { return len(a) }
func (a Int) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int) Less(i, j int) bool { return a[i] < a[j] }

type DictOrder []int

func (a DictOrder) Len() int      { return len(a) }
func (a DictOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a DictOrder) Less(i, j int) bool {
	si := strconv.Itoa(a[i])
	sj := strconv.Itoa(a[j])
	if si+sj < sj+si {
		return false
	} else {
		return true
	}
}

/**
 * @param a: the array in problem.
 * @return: represent the new number.
 */
func Combine(a []int) string {
	// write your code here
	sum := func(x []int) int {
		s := 0
		for _, v := range x {
			s += v
		}
		return s
	}
	sort.Sort(Int(a))
	var soln string
	if v := sum(a) % 3; v == 0 {
		soln = combineLargest(a)
	} else if v == 1 {
		i1 := indexOfSmallestKthWithCondition(a, 0, 1)
		a = append(a[:i1], a[i1+1:]...)
		soln = combineLargest(a)
	} else {
		i1_0 := indexOfSmallestKthWithCondition(a, 0, 1)
		i1_1 := indexOfSmallestKthWithCondition(a, 1, 1)
		i2_0 := indexOfSmallestKthWithCondition(a, 0, 2)
		if i1_0 == -1 || i1_1 == -1 {
			goto Case2
		} else if i2_0 == -1 {
			goto Case1
		}
		if a1, _ := strconv.Atoi(string(rune(a[i1_0])) + string(rune(a[i1_1]))); a[i2_0] < a1 {
			goto Case1
		} else {
			goto Case2
		}
	Case2:
		if i2_0 == len(a)-1 {
			a = a[:i2_0]
		} else {
			a = append(a[:i2_0], a[i2_0+1:]...)
		}
		goto Solve
	Case1:
		if i1_1 == len(a)-1 {
			a = a[:i1_1]
		} else {
			a = append(a[:i1_1], a[i1_1+1:]...)
		}
		if i1_0 == len(a)-1 {
			a = a[:i1_0]
		} else {
			a = append(a[:i1_0], a[i1_0+1:]...)
		}
	Solve:
		soln = combineLargest(a)
	}
	return soln
}

func combineLargest(a []int) string {
	sort.Sort(DictOrder(a))
	s := ""
	for i := 0; i < len(a); i++ {
		s += strconv.Itoa(a[i])
	}
	return s
}

func indexOfSmallestKthWithCondition(a []int, k, cond int) int {
	t := 0
	for i, v := range a {
		if v%3 == cond {
			if t == k {
				return i
			} else {
				t++
			}
		}
	}
	return -1
}

func main() {
	c := Combine([]int{557, 918, 205, 330, 110, 20, 843, 506, 748, 124, 615, 908, 476, 206, 163, 980, 909, 951, 589, 770, 536, 537, 226, 20, 777, 533, 159, 213, 165, 470, 104, 378, 28, 120, 413, 446, 521, 329, 601, 404, 872, 390, 172, 509, 63, 473, 660, 819, 576, 866, 368, 822, 590, 412, 637, 276, 630, 186, 364, 750, 609, 481, 596, 865, 87})
	fmt.Println(c)
	fmt.Println("9809519189099088787286686584382281977777075074866063763630615609601596590589576557537536533521509506481476473470446413412404390378368364330329282762262132062052020186172165163159124120110104")
}
