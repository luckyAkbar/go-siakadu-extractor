package utility

func (u *Utility) GetAllProfileImage(delay, maxSeqFailure int, start string, to string, skipYearPlus1 bool) {
	i := NewIncrementor(start, to)
	i.SetMaxSequentialFailure(maxSeqFailure)
	i.SetDelaySec(delay)

	if skipYearPlus1 {
		i.EnableYearPlus1Skipping()
	}

	for !i.IsMaxReached {
		if err := u.GetImageFromNPM(i.GetCurrent()); err != nil {
			i.IncrementSeqFailureCount()
		}

		i.Next()
	}
}
