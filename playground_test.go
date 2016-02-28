package playground_test

import (
	"fmt"
	"time"

	"github.com/f2prateek/playground"
)

func ExampleTime() {
	t := playground.Time()
	fmt.Println(t.Format(time.RFC1123))
	// Output: Tue, 10 Nov 2009 23:00:00 UTC
}
