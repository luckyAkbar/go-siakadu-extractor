package utility

import "fmt"

func (u *Utility) GetImageFromNPMRange(start, to string) {
	i := NewIncrementor(start, to)
	util := New()

	for !i.IsMaxReached {
		util.GetImageFromNPM(i.GetCurrent())
		i.Next()
	}

	fmt.Println(i.GeneratedCount)
}
