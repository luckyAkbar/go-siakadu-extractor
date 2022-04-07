package utility

import (
	"fmt"
	"time"
)

func (u *Utility) GetImageFromNPMRange(start, to string, delay int) {
	i := NewIncrementor(start, to)

	for !i.IsMaxReached {
		u.GetImageFromNPM(i.GetCurrent())
		i.Next()

		time.Sleep(time.Duration(delay) * time.Second)
	}

	fmt.Println(i.GeneratedCount)
}
