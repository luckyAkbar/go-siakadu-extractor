package utility

import "time"

func (u *Utility) GetAllProfileImage(delay, maxSeqFailure int, start string, to string) {
	i := NewIncrementor(start, to)
	i.SetMaxSequentialFailure(maxSeqFailure)

	for !i.IsMaxReached {
		if err := u.GetImageFromNPM(i.GetCurrent()); err != nil {
			i.IncrementSeqFailureCount()
		}

		i.Next()

		time.Sleep(time.Duration(delay) * time.Second)
	}
}
