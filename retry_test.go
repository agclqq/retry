package retry

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func timeMatch(start time.Time, cost, diff float64) (float64, bool) {
	now := time.Now()
	s := float64(now.Sub(start)) / 1e9
	if math.Abs(s-cost) <= diff {
		return s, true
	}
	return s, false
}
func TestRetry_Run(t *testing.T) {
	r := Retry{
		InitialBackoff:    1,
		MaxBackoff:        15,
		BackoffMultiplier: 1.5,
		MaxAttempts:       15,
	}
	now := time.Now()

	r.Run(func(step uint) {
		switch step {
		case 1:
			cost := 1.0 //sleep 1.0
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}
		case 2:
			cost := 2.5 //sleep 1.5
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}
		case 3:
			cost := 4.75 //sleep 2.25
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 4:
			cost := 8.125 //sleep 3.375
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 5:
			cost := 13.1875 //sleep 5.0625
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 6:
			cost := 20.78125 //sleep 7.59375
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 7:
			cost := 32.17188 //sleep 11.39063
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 8:
			cost := 49.2578125 //sleep 17.08594  > MaxBackoff 15
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 9:
			cost := 64.2578125 //sleep 15.00
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}
		case 10:
			cost := 79.2578125 //sleep 15.00
			diff := 0.1
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 11:
			cost := 94.2578125 //sleep 15.00
			diff := 0.2
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 12:
			cost := 109.2578125 //sleep 15.00
			diff := 0.2
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}

		case 13:
			cost := 124.2578125 //sleep 15.00
			diff := 0.2
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}
		case 14:
			cost := 139.2578125 //sleep 15.00
			diff := 0.2
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}
		case 15:
			cost := 154.2578125 //sleep 15.00
			diff := 0.2
			if got, ok := timeMatch(now, cost, diff); !ok {
				t.Errorf("want %f±%f,got %f", cost, diff, got)
			}
		default:
			t.Error("It's not supposed to be here")
		}
	})
}

func TestRetry_Reset(t *testing.T) {
	r := Retry{
		InitialBackoff:    1,
		MaxBackoff:        15,
		BackoffMultiplier: 1.5,
		MaxAttempts:       5,
	}
	go func() {
		time.Sleep(5 * time.Second)
		r.Reset()
	}()
	oneNum := 0
	r.Run(func(step uint) {
		if step == 1 {
			oneNum++
		}
	})
	fmt.Println(oneNum)
	if oneNum <= 1 {
		t.Errorf("Reset failed")
	}
}

func TestRetry_Cancel(t *testing.T) {
	r := Retry{
		InitialBackoff:    1,
		MaxBackoff:        15,
		BackoffMultiplier: 1.5,
		MaxAttempts:       5,
	}
	go func() {
		time.Sleep(5 * time.Second) //5 seconds is between 3 and 4 steps
		r.Cancel()
	}()
	maxNum := uint(0)
	r.Run(func(step uint) {
		if step >= maxNum {
			maxNum = step
		}
	})
	if maxNum >= 5 {
		t.Errorf("Cancel failed")
	}
}
