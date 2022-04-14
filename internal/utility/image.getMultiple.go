package utility

import (
	"fmt"
	"time"
)

func (u *Utility) GetImageFromNPMRange(start, to string, delay, requestOptimizer int) {
	i := NewIncrementor(start, to)
	i.SetMaxSequentialFailure(10)

	if requestOptimizer != 0 {
		i.SetMaxSequentialFailure(requestOptimizer)
	}

	for !i.IsMaxReached {
		err := u.GetImageFromNPM(i.GetCurrent())

		if err != nil && requestOptimizer != 0 {
			i.IncrementSeqFailureCount()
		}

		i.Next()

		time.Sleep(time.Duration(delay) * time.Second)
	}

	fmt.Println(i.GeneratedCount)
}
