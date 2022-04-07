package utility

import (
	"fmt"
	"strconv"
	"time"
)

type NPMIncrementor struct {
	Start              string
	To                 string
	Current            string
	CurrentYear        string
	CurrentDepartement string
	CurrentDivision    string
	CurrentAbsent      string
	GeneratedCount     int
	SequentialFailureCount int
	MaxSequentialFailure int
	IsMaxReached       bool
}

func NewIncrementor(start, to string) *NPMIncrementor {
	return &NPMIncrementor{
		Start:              start,
		To:                 to,
		Current:            start,
		CurrentYear:        start[0:2],
		CurrentDepartement: start[2:4],
		CurrentDivision:    start[4:7],
		CurrentAbsent:      start[7:10],
		GeneratedCount:     0,
		SequentialFailureCount: 0, // this and max seq set to this val
		MaxSequentialFailure: -1,	// so will not take effect if not needed
		IsMaxReached:       false,
	}
}

func (n *NPMIncrementor) GenerateNPMRangeList() []string {
	var NPMLists []string = []string{}
	var t int64 = int64(0)

	for !n.IsMaxReached {
		NPMLists = append(NPMLists, n.GetCurrent())

		n.incrementAbsent()
		t++

		if n.GetCurrent() == n.To {
			n.IsMaxReached = true
		}
	}

	NPMLists = append(NPMLists, n.To)

	return NPMLists
}

func (n *NPMIncrementor) Next() {
	n.incrementAbsent()
	n.GeneratedCount++

	if n.Current == n.To {
		n.IsMaxReached = true
	}
}

func (n *NPMIncrementor) incrementYear() {
	tmp, _ := strconv.Atoi(n.CurrentYear)
	tmp++

	last2YearDigitsStr := strconv.Itoa(time.Now().Year())
	last2YearDigits, _ := strconv.Atoi(last2YearDigitsStr[0:2])

	if tmp > last2YearDigits {
		n.IsMaxReached = true

		return
	}

	if tmp == 100 {
		n.IsMaxReached = true
		return
	}

	if tmp > 9 {
		n.CurrentYear = strconv.Itoa(tmp)
		return
	}

	n.CurrentYear = fmt.Sprintf("0%s", strconv.Itoa(tmp))
}

func (n *NPMIncrementor) incrementDepartement() {
	tmp, _ := strconv.Atoi(n.CurrentDepartement)
	tmp++

	if tmp > 99 {
		n.CurrentDepartement = "00"
		n.incrementYear()
		return
	}

	if tmp > 9 {
		n.CurrentDepartement = strconv.Itoa(tmp)
		return
	}

	n.CurrentDepartement = fmt.Sprintf("0%s", strconv.Itoa(tmp))
}

func (n *NPMIncrementor) incrementDivision() {
	tmp, _ := strconv.Atoi(n.CurrentDivision)
	tmp++

	if tmp == 1000 {
		n.CurrentDivision = "000"
		n.incrementDepartement()

		return
	}

	if tmp > 99 {
		n.CurrentDivision = strconv.Itoa(tmp)
		return
	}

	if tmp > 9 {
		n.CurrentDivision = fmt.Sprintf("0%s", strconv.Itoa(tmp))
		return
	}

	n.CurrentDivision = fmt.Sprintf("00%s", strconv.Itoa(tmp))
}

func (n *NPMIncrementor) incrementAbsent() {
	tmp, _ := strconv.Atoi(n.CurrentAbsent)
	tmp++

	if n.SequentialFailureCount > n.MaxSequentialFailure {
		n.CurrentAbsent = "000"

		n.ResetSeqFailureCount()
		n.incrementDivision()

		return
	}

	if tmp == 1000 {
		n.CurrentAbsent = "000"
		n.incrementDivision()

		return
	}

	if tmp > 99 {
		n.CurrentAbsent = strconv.Itoa(tmp)
		return
	}

	if tmp > 9 {
		n.CurrentAbsent = fmt.Sprintf("0%s", strconv.Itoa(tmp))
		return
	}

	n.CurrentAbsent = fmt.Sprintf("00%s", strconv.Itoa(tmp))
}

func (n *NPMIncrementor) GetCurrent() string {
	n.Current = fmt.Sprintf("%s%s%s%s", n.CurrentYear, n.CurrentDepartement, n.CurrentDivision, n.CurrentAbsent)

	return n.Current
}

func (n *NPMIncrementor) SetMaxSequentialFailure(value int) {
	n.MaxSequentialFailure = value
}

func (n *NPMIncrementor) IncrementSeqFailureCount() {
	n.SequentialFailureCount++
}

func (n *NPMIncrementor) ResetSeqFailureCount() {
	n.SequentialFailureCount = 0
}