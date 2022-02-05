package retry_test

import (
	"fmt"
	"time"

	"github.com/kei6u/retry"
)

func ExampleConstant() {
	r := retry.New(retry.Constant{
		Interval:    time.Millisecond,
		MaxAttempts: 5,
	})
	attempts := 0
	start := time.Now()
	for r.Next() {
		fmt.Printf("attempt %d, %s\n", attempts, time.Since(start))
		start = time.Now()
		attempts++
	}
}

func ExampleJitter() {
	r := retry.New(retry.Jitter{
		Base:        time.Millisecond,
		MaxAttempts: 30,
	})
	attempts := 0
	var ds []time.Duration
	start := time.Now()
	for r.Next() {
		d := time.Since(start)
		ds = append(ds, d)
		fmt.Printf("attempt %d, %s\n", attempts, d)
		start = time.Now()
		attempts++
	}
	fmt.Printf("durations: %v\n", ds)
}

func ExampleExponentialBackoff() {
	r := retry.New(retry.ExponentialBackoff{
		Base:        time.Millisecond,
		Max:         100 * time.Millisecond,
		MaxAttempts: 30,
	})
	attempts := 0
	var ds []time.Duration
	start := time.Now()
	for r.Next() {
		d := time.Since(start)
		ds = append(ds, d)
		fmt.Printf("attempt %d, %s\n", attempts, d)
		start = time.Now()
		attempts++
	}
	fmt.Printf("durations: %v\n", ds)
}
