package sqrt_test

import (
	"log"
	"math"
	"testing"

	sqrt "github.com/moledoc/letsgo/Sqrt"
)

func TestSqrt(t *testing.T) {
	// init variables for convenience
	var err error
	var sqrtVal float64

	// positive nr sqrt

	sqrtVal, err = sqrt.Sqrt(1)
	if err != nil {
		log.Fatal(err)
	}
	if math.Abs(sqrtVal-1) > 0.001 {
		t.Fatalf("Sqrt(%v) returns incorrect result: expected %v, got %v", 1, 1, sqrtVal)
	}

	sqrtVal, err = sqrt.Sqrt(2)
	if err != nil {
		log.Fatal(err)
	}
	if math.Abs(math.Round(sqrtVal*1000)/1000-1.414) > 0.001 {
		t.Fatalf("Sqrt(%v) returns incorrect result: expected %v, got %v", 2, 1.414, sqrtVal)
	}

	sqrtVal, err = sqrt.Sqrt(64)
	if err != nil {
		log.Fatal(err)
	}
	if math.Abs(sqrtVal-8) > 0.001 {
		t.Fatalf("Sqrt(%v) returns incorrect result: expected %v, got %v", 64, 8, sqrtVal)
	}

	// negative nr sqrt
	sqrtVal, err = sqrt.Sqrt(-1)
	if err == nil {
		t.Fatalf("Sqrt(%v) returns incorrect result: expected %v, got %v", -1, nil, sqrtVal)
	}

	// sqrt(0)
	sqrtVal, err = sqrt.Sqrt(0)
	if err != nil {
		log.Fatal(err)
	}
	if sqrtVal != 1 {
		t.Fatalf("Sqrt(%v) returns incorrect result: expected %v, got %v", 0, 1, sqrtVal)
	}

}
