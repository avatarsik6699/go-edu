package countdown

import (
	"bytes"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("should just print 3", func(t *testing.T) {
		var sleeper SpySleeper
		var buff bytes.Buffer

		CountDown(&buff, &sleeper)

		got := buff.String()
		want := "3\n2\n1\nGo!\n"

		if sleeper.counts != time.Duration(time.Second*3) {
			t.Errorf("got -> %s, %s <- want", sleeper.counts, time.Duration(time.Second*3))
		}

		if got != want {
			t.Errorf("got -> %q, %q <- want", got, want)
		}

	})
}

type SpySleeper struct {
	counts time.Duration
}

func (s *SpySleeper) Sleep(duration time.Duration) {
	s.counts += duration
}
