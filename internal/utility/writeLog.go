package utility

import (
	"fmt"
	"time"
)

func (u *Utility) WriteLog(log string) {
	now := time.Now().Local().String()

	u.logFileTarget.WriteString(fmt.Sprintf("[%s]: %s\n", now, log))
}
