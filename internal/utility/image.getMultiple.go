package utility

func (u *Utility) GetImageFromNPMRange(start, to string, delay, requestOptimizer int, skipYearPlus1 bool) {
	i := NewIncrementor(start, to)
	i.SetMaxSequentialFailure(requestOptimizer)
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
