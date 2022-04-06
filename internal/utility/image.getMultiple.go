package utility

import "fmt"

func (u *Utility) GetImageFromNPMRange(start, to string) {
	i := NewIncrementor(start, to)

	for !i.IsMaxReached {
		u.GetImageFromNPM(i.GetCurrent())
		i.Next()
	}

	fmt.Println(i.GeneratedCount)
}
