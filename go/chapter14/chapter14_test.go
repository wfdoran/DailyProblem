package chapter14

import (
	"fmt"
	"testing"
)

func TestFlightItinerary(t *testing.T) {
	f := []flight{flight{src: "SFO", dst: "HKO"},
		flight{src: "YYZ", dst: "SFO"},
		flight{src: "YUL", dst: "YYZ"},
		flight{src: "HKO", dst: "ORD"},
	}

	result := FlightItinerary(f, "YUL")
	fmt.Println(result)
}

func TestAndroidCount(t *testing.T) {
	sum := 0
	for n := 1; n <= 9; n++ {
		x := AndroidCount(n)
		fmt.Println(n, x)
		sum += x
	}
	fmt.Println(sum)
}
