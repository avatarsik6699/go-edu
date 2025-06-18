package countdown

import (
	"fmt"
	"io"
	"time"
)

/**
While this is a pretty trivial program,
to test it fully we will need as always to take an iterative, test-driven approach.

It's an important skill to be able to slice up requirements as small as you can so you can have working software.
*/

const (
	finalMessage = "Go!"
	finalCount   = 3
)

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := finalCount; i > 0; i-- {
		sleeper.Sleep(time.Second)
	}

	for i := finalCount; i > 0; i-- {
		fmt.Fprintf(w, "%d\n", i)
	}

	fmt.Fprintf(w, "%s\n", finalMessage)
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}
