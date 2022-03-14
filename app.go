package main

import (
	"fmt"
	"golang-tdd/http_server"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// di.Greet(os.Stdout, "Foo")
	// log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(di.MyGreetHandler)))
	// sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	// Countdown(os.Stdout, sleeper)
	server := &http_server.PlayerServer{Store: http_server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8000", server))
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord+"\n")
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
