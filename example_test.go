package retry_test

import (
	"fmt"
	"time"

	"github.com/keisku/retry"
)

func ExampleConstant() {
	r := retry.New(retry.Constant{
		Interval:    time.Second,
		MaxAttempts: 5,
	})
	retries := 0
	start := time.Now()
	for r.Next() {
		fmt.Printf("retry %d, %s\n", retries, time.Since(start))
		start = time.Now()
		retries++
	}
}

func ExampleJitter() {
	r := retry.New(retry.Jitter{
		Base:        time.Millisecond,
		MaxAttempts: 30,
	})
	retries := 0
	var ds []time.Duration
	start := time.Now()
	for r.Next() {
		d := time.Since(start)
		ds = append(ds, d)
		fmt.Printf("retry %d, %s\n", retries, d)
		start = time.Now()
		retries++
	}
	fmt.Printf("durations: %v\n", ds)
}

func ExampleExponentialBackoff() {
	r := retry.New(retry.ExponentialBackoff{
		Base:        time.Millisecond,
		Max:         10 * time.Second,
		MaxAttempts: 20,
	})
	retries := 0
	var ds []time.Duration
	start := time.Now()
	for r.Next() {
		d := time.Since(start)
		ds = append(ds, d)
		fmt.Printf("retry %d, %s\n", retries, d)
		start = time.Now()
		retries++
	}
	fmt.Printf("durations: %v\n", ds)
}
