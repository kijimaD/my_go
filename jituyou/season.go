package main

type Season int

const (
	Peak   Season = iota + 1 // 繁忙期
	Normal                   // 通常期
	Off                      // 閑散期
)

func (s Season) Price(price float64) bool {
	if s == Peak {
		return s + 200 // 繁忙期のみ割増
	}
	return s
}
