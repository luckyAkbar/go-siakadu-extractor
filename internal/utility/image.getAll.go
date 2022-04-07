package utility

import "time"

func (u *Utility) GetAllProfileImage(delay, maxSeqFailure int) {
	i := NewIncrementor("0000000000", "9999999999")
	i.SetMaxSequentialFailure(maxSeqFailure)

	for !i.IsMaxReached {
		if err := u.GetImageFromNPM(i.GetCurrent()); err != nil {
			i.IncrementSeqFailureCount()
		}

		i.Next()

		time.Sleep(time.Duration(delay) * time.Second)
	}
}