package utility

import "time"

func (u *Utility) GetAllProfileImage(delay int) {
	i := NewIncrementor("0000000000", "9999999999")

	for !i.IsMaxReached {
		u.GetImageFromNPM(i.GetCurrent())
		i.Next()

		time.Sleep(time.Duration(delay) * time.Second)
	}
}