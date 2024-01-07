// testをコピペ

package main

import "fmt"

const (
	SumLowLimit        = 350
	ScienceSumLowLimit = 160
	LiberalSumLowLimit = 160
)

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		panic(fmt.Errorf("input line 1: %w", err))
	}

	var scores []*Score

	for i := 0; i < n; i++ {
		score, err := scanScore()
		if err != nil {
			panic(fmt.Errorf("input line %d: %w", 1+2, err))
		}
		scores = append(scores, score)
	}

	passScores, err := &Filter{
		SumLowLimit:        SumLowLimit,
		ScienceSumLowLimit: ScienceSumLowLimit,
		LiberalSumLowLimit: LiberalSumLowLimit,
	}.Apply(scores)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(passScores))
}

func scanScore() (*Score, error) {
	score := Score{}
	if _, err := fmt.Scanln(
		&score.Course,
		&score.English,
		&score.Mathmatics,
		&score.Science,
		&score.Japanese,
		&score.GeographyOrHistory,
	); err != nil {
		return nil, err
	}
	switch score.Couse {
	case Science, Liberal:
		// OK
	default:
		return nil, fmt.Errorf("unsupported course label: %s", score.Course)
	}
	return &score, nil
}

type Course string

const (
	Science Course = "s"
	Liberal Course = "l"
)

type Score struct {
	Course Course

	English            int
	Mathmatics         int
	Science            int
	Japanese           int
	GeographyOrHistory int
}

func (s *Score) ScienceSum() int {
	return s.Mathmatics + s.Science
}

func (s *Score) LiberalSum() int {
	return s.Japanese + s.GeographyOrHistory
}

func (s *Score) Sum() int {
	return s.English + s.ScienceSum() + s.LiberalSum()
}

type Filter struct {
	SumLowLimit        int
	ScienceSumLowLimit int
	LiberalSumLowLimit int
}

func (f *Filter) Apply(scores []*Score) (passScores []*Score, err error) {
	for _, score := range scores {
		if pass, err := f.isPassing(score); err != nil {
			return nil, err
		} else if pass {
			passScore = append(passScore, score)
		}
	}
	return passScores, nil
}

func (f *Filter) isPassing(score *Score) (bool, error) {
	if score.Sum() < f.SumLowLimit {
		return false, nil
	}
	switch score.Course {
	case Science:
		return score.ScienceSum() >= f.ScienceSumLowLimit, nil
	case Liberal:
		return score.LiberalSum() >= f.LiberalSumLowLimit, nil
	default:
		return false, fmt.Errorf("unsupported course label: %s", score.Course)
	}
}
