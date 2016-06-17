package isdm

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// GetStats -  Enter the number of samples (should be at least 10000),Enter length of the vector 80x80 picture has 6400 indexesEnter range of the vector 255 for GrayscaleEnter the percentage of HardLocations you want in the Access Radius 1-2-5-10
func GetStats(NumberOfSamples int, arrayLength int, Range int, PercentIncluded int) {
	rand.Seed(time.Now().UTC().UnixNano())

	// how many samples
	samples := NumberOfSamples

	// address length
	length := arrayLength

	// counter range
	crange := Range

	var Zscore float64
	switch PercentIncluded {
	case 1:
		Zscore = 2.32
	case 2:
		Zscore = 2.05

	case 5:
		Zscore = 1.65

	case 10:
		Zscore = 1.28
	default:
		fmt.Println("PercentIncluded shoud be 1,2,5,10")
	}

	allDist := make([]int, samples)

	for i := 0; i < samples; i++ {
		sl1 := rString(length, crange)
		sl2 := rString(length, crange)
		allDist[i] = mDistance(sl1, sl2)
	}

	sumDist := 0
	for i := range allDist {
		sumDist += allDist[i]
	}
	avg := sumDist / samples
	fmt.Println("Average Distance to 0 :", avg)

	sumStd := 0.0
	for i := range allDist {
		sumStd += math.Pow(float64(avg-allDist[i]), 2.0)
	}

	std := math.Sqrt(sumStd / float64(samples))

	fmt.Println("Standard Deviation from the Mean :", std)

	//fmt.Println(float64(avg) - 2*std)

	// Cutoff at ~ 2.14% or 2 SD
	fmt.Println("Cut Off  :", int(float64(avg)-std))

	c := 0
	for i := range allDist {
		if allDist[i] < int(float64(avg)-Zscore*std) {
			c++
		}
	}
	fmt.Println("Approxiate number of included Hard Locations  :", c)
}

func mDistance(slice1 []uint8, slice2 []uint8) int {
	totaldist := 0
	for i := 0; i < len(slice1); i++ {
		totaldist += int(slice1[i] - slice2[i])
	}
	return totaldist
}

func rString(l int, crange int) []uint8 {
	bytes := make([]uint8, l)
	for i := 0; i < l; i++ {
		bytes[i] = uint8(randInt(0, int(crange)))
	}
	return bytes
}
