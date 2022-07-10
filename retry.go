package retry

import (
	"time"
)

type Retry struct {
	InitialBackoff    float32 //The initial retry interval, in seconds, must be greater than 0
	MaxBackoff        float32 //The maximum retry interval, in seconds, must be greater than or equal to initialBackoff
	BackoffMultiplier float32 //The retry interval growth factor needs to be greater than 1
	MaxAttempts       uint    //Maximum number of retries. If the value is 0, there is no limit
	step              uint
	cancel            bool
}

func (rt *Retry) Run(f func(step uint)) {
	sleepTime := rt.InitialBackoff
	for {
		rt.step++
		if rt.MaxAttempts > 0 && rt.step > rt.MaxAttempts || rt.cancel {
			break
		}
		time.Sleep(time.Duration(sleepTime*1e6) * time.Microsecond)
		f(rt.step)

		if sleepTime >= rt.MaxBackoff {
			sleepTime = rt.MaxBackoff
		} else if sleepTime < rt.MaxBackoff {
			sleepTime = sleepTime * rt.BackoffMultiplier
		}
	}
}
func (rt *Retry) Reset() {
	rt.step = 0
}
func (rt *Retry) Cancel() {
	rt.cancel = true
}
