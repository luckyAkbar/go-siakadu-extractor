package utility

func (u *Utility) GetAllProfileImage() {
	i := NewIncrementor("0000000000", "9999999999")

	for !i.IsMaxReached {
		u.GetImageFromNPM(i.GetCurrent())
		i.Next()
	}
}