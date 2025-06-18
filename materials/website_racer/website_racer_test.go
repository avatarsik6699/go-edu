package websiteracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	createServerWithDelay := func(delay time.Duration) *httptest.Server {
		return httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if delay > 0 {
					time.Sleep(delay)
				}

				w.WriteHeader(http.StatusOK)
			}))
	}

	t.Run("should return first url as more fast", func(t *testing.T) {
		fastServer := createServerWithDelay(0)
		slowServer := createServerWithDelay(time.Millisecond * 20)
		defer func() {
			slowServer.Close()
			fastServer.Close()
		}()

		got, err := WebsiteRacer(fastServer.URL, slowServer.URL)
		want := fastServer.URL

		if err != nil {
			t.Errorf("got unexpectad error")
		}

		if got != want {
			t.Errorf("want -> %q, %q <- got", want, got)
		}
	})

	t.Run("should return err if a server doesn't respond within 10s", func(t *testing.T) {
		fastServer := createServerWithDelay(time.Second * 11)
		slowServer := createServerWithDelay(time.Second * 12)

		defer func() {
			slowServer.Close()
			fastServer.Close()
		}()

		_, err := WebsiteRacer(fastServer.URL, slowServer.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})
}
