package testhw

import "math/rand"

// MaxRange is maxumum range for counters and values
const MaxRange uint8 = 255

// HardLocation is the active memory in the SDM
type HardLocation struct {
	address    []uint8
	counters   [][]uint8 //[][]uint8
	wordLength int
	writeCount int
}

// NewHL makes a new HardLocation
func NewHL(NewAddress []uint8) *HardLocation {

	l := len(NewAddress)

	picture := make([][]uint8, AddressLength)

	for i := range picture {
		picture[i] = make([]uint8, MaxRange)
	}

	return &HardLocation{NewAddress, picture, l, 0}
}

func (hl *HardLocation) getAddress() []uint8 {
	return hl.address
}

func (hl *HardLocation) getWriteCount() int {
	return hl.writeCount
}

func (hl *HardLocation) getWordLength() int {
	return hl.wordLength
}

// func (hl *HardLocation) getCounters() [][]byte {
// 	return hl.counters
// }

func write(writeAddress []uint8, hl *HardLocation) {
	hl.writeCount++
	for k, v := range writeAddress {
		if v == 255 {
			v = 254
		}
		//fmt.Println(k, v)
		hl.counters[k][v] = (hl.counters[k][v] + uint8(1)) % MaxRange
	}
}

func distance(addressToWrite []uint8, hl *HardLocation) int {

	totalDist := 0

	for i := range addressToWrite {
		totalDist += int(addressToWrite[i] - hl.address[i])
	}

	return totalDist

}

func randomString(l int) []uint8 {
	bytes := make([]uint8, l)
	for i := 0; i < l; i++ {
		bytes[i] = uint8(randInt(0, int(MaxRange)))
	}
	return bytes
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
