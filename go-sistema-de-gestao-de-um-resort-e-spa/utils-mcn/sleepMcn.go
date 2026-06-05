package utilsmcn

import "time"

func SleepMCN(s uint) {
	time.Sleep(time.Duration(s) * time.Second)
}
