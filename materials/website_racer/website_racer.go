package websiteracer

import (
	"errors"
	"log"
	"net/http"
	"time"
)

var (
	limitTime time.Duration = time.Second * 10

	ErrlimitTime = errors.New("limit time has got")
)

// TODO: Метод принимает 2 урла и возвращает тот, который загрузился быстрее
func WebsiteRacer(firstURL, secondURL string) (string, error) {
	ch1 := make(chan any)
	ch2 := make(chan any)

	go func() {
		http.Get(firstURL)
		close(ch1)
	}()

	go func() {
		http.Get(secondURL)
		close(ch2)
	}()

	select {
	case <-ch1:
		log.Println("AAAAA", <-ch1)
		return firstURL, nil
	case <-ch2:
		return secondURL, nil
	case <-time.After(10 * time.Second):
		return "", ErrlimitTime
	}

	// var URLToReturn string
	// var finishTimeToReturn time.Duration

	// if finishFirstURLTime > finishSecondURLTime {
	// 	URLToReturn = secondURL
	// 	finishTimeToReturn = finishSecondURLTime
	// } else {
	// 	URLToReturn = firstURL
	// 	finishTimeToReturn = finishFirstURLTime
	// }

	// if finishTimeToReturn > limitTime {
	// 	return "", ErrlimitTime
	// }

	// return URLToReturn, nil
}
