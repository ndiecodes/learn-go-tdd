package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type ConfigurablaSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurablaSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)

	// for i := countdownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// }

	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }

	// sleeper.Sleep()
	// fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurablaSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
