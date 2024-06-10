package main

import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"
)

const p = "death & taxes"

func main() {
	// 对比全部的彩票排列
	initAllTickets()

	allCount := 17721088
	fmt.Print("输出最后一注：")
	fmt.Println(getCombinationByIndex(allCount - 1))
	for i := 0; i < allCount; i++ {
		ticket, _ := getCombinationByIndex(i)
		fmt.Print(ticket)
		redBalls, _, err := parseTicket(ticket)
		if err != nil {
			fmt.Println("Error parsing ticket:", err)
			return
		}

		if checkConsecutiveRule(redBalls) {
			fmt.Print(":violates")
		} else {
			fmt.Print(":not violate")
		}

		if checkRangeRule(redBalls) {
			fmt.Print(",range,")
		} else {
			fmt.Print(",not range,")
		}

		if checkRangeLimitRule(redBalls) {
			fmt.Println("range limit")
		} else {
			fmt.Println("not range limit")
		}
	}
}

// 计算组合数
func comb(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	if k > n-k {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res *= n - i
		res /= i + 1
	}
	return res
}

// 根据索引生成双色球组合
func getCombinationByIndex(index int) (string, error) {
	totalCombinations := comb(33, 6) * 16
	if index < 0 || index >= totalCombinations {
		return "", fmt.Errorf("index out of range")
	}

	nums := make([]int, 6)
	rem := index / 16
	blue := (index % 16) + 1

	k := 0
	for i := 1; i <= 33 && k < 6; i++ {
		c := comb(33-i, 6-k-1)
		if rem < c {
			nums[k] = i
			k++
		} else {
			rem -= c
		}
	}

	ticket := "0" + strconv.Itoa(nums[0])
	for i := 1; i < 6; i++ {
		ticket += ",0" + strconv.Itoa(nums[i])
	}
	ticket += ",0" + strconv.Itoa(blue)

	return ticket, nil
}

func checkConsecutiveRule(redBalls []int) bool {
	sort.Ints(redBalls) // 确保红球号码排序
	consecutiveCount := 1

	for i := 1; i < len(redBalls); i++ {
		if redBalls[i] == redBalls[i-1]+1 {
			consecutiveCount++
			if consecutiveCount > 4 {
				return true
			}
		} else {
			consecutiveCount = 1
		}
	}

	return false
}

func parseTicket(ticket string) ([]int, int, error) {
	parts := strings.Split(ticket, ",")
	if len(parts) != 7 {
		return nil, 0, fmt.Errorf("invalid ticket format")
	}

	var redBalls []int
	for i := 0; i < 6; i++ {
		num, err := strconv.Atoi(parts[i])
		if err != nil {
			return nil, 0, err
		}
		redBalls = append(redBalls, num)
	}

	blueBall, err := strconv.Atoi(parts[6])
	if err != nil {
		return nil, 0, err
	}

	return redBalls, blueBall, nil
}

func checkRangeRule(redBalls []int) bool {
	range1 := 0
	range2 := 0
	range3 := 0

	for _, num := range redBalls {
		switch {
		case num >= 1 && num <= 9:
			range1++
		case num >= 10 && num <= 19:
			range2++
		case num >= 20 && num <= 33:
			range3++
		}
	}

	return range1 == 2 && range2 == 2 && range3 == 2
}

func checkRangeLimitRule(redBalls []int) bool {
	range1 := 0
	range2 := 0
	range3 := 0

	for _, num := range redBalls {
		switch {
		case num >= 1 && num <= 9:
			range1++
		case num >= 10 && num <= 19:
			range2++
		case num >= 20 && num <= 33:
			range3++
		}
	}

	return range1 <= 3 && range2 <= 3 && range3 <= 3
}

func initAllTickets() {
	count := 0
	countIndex := 0
	for i := 1; i <= 33; i++ {
		for j := i + 1; j <= 33; j++ {
			for k := j + 1; k <= 33; k++ {
				for m := k + 1; m <= 33; m++ {
					for n := m + 1; n <= 33; n++ {
						for o := n + 1; o <= 33; o++ {
							for p := 1; p <= 16; p++ {
								if i != j && i != k && i != m && i != n && i != o &&
									j != k && j != m && j != n && j != o &&
									k != m && k != n && k != o &&
									m != n && m != o &&
									n != o {
									count++
									ticket := "0" + strconv.Itoa(i)
									ticket += ",0" + strconv.Itoa(j)
									ticket += ",0" + strconv.Itoa(k)
									ticket += ",0" + strconv.Itoa(m)
									ticket += ",0" + strconv.Itoa(n)
									ticket += ",0" + strconv.Itoa(o)
									ticket += ",0" + strconv.Itoa(p)
									combinationValue, _ := getCombinationByIndex(countIndex)
									if ticket == combinationValue {
										slog.Info("Ticket equal",
											slog.Int("index", count),
											slog.String("ticket", ticket),
											slog.String("combinationValue", ticket))
									} else {
										slog.Error("Ticket no equal",
											slog.Int("index", count),
											slog.String("ticket", ticket),
											slog.String("combinationValue", ticket))
									}
									countIndex++
								}
							}
						}
					}
				}
			}
		}
	}

	slog.Info("All ticket num", slog.Int("count", count))
}
