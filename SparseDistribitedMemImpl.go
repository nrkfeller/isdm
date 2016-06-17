package isdm

// MaxIterations number of times the read can be feedbacked into the SDM
const MaxIterations uint8 = 20

// SparseDistributedMemoryImpl is the memory space that holds HardLocations
type SparseDistributedMemoryImpl struct {
	MemorySize       int // Number of Hard locations
	AddressLength    int
	ActivationRadius int
	HardLocations    []HardLocation
}

// NewSDM initializes the SDM only needs to do this once
func NewSDM(MemorySize int, AddressLength int, ActivationRadius int) *SparseDistributedMemoryImpl {

	HardLocations := make([]HardLocation, MemorySize)

	for i := 0; i < MemorySize; i++ {
		HardLocations[i] = *NewHL(randomString(AddressLength))
	}
	return &SparseDistributedMemoryImpl{MemorySize, ActivationRadius, ActivationRadius, HardLocations}
}

// store, finds neighbors, and writes to them
func store(AddressToStore []uint8, sdm *SparseDistributedMemoryImpl) {
	//count := 0
	for i := 0; i < sdm.MemorySize; i++ {
		if distance(AddressToStore, &sdm.HardLocations[i]) <= sdm.ActivationRadius {
			// write(randomString(AddressLength), myHL)
			//count++
			//fmt.Println(len(AddressToStore))
			write(AddressToStore, &sdm.HardLocations[i])
		}
	}
	//fmt.Println(count, "hits")
}

// Retrieve from isdm
func Retrieve(addressToRetrieve []uint8, sdm *SparseDistributedMemoryImpl) []uint8 {
	var neighbors []*HardLocation

	//count := 0

	for i := 0; i < sdm.MemorySize; i++ {
		if distance(addressToRetrieve, &sdm.HardLocations[i]) <= sdm.ActivationRadius {
			//count++
			neighbors = append(neighbors, &sdm.HardLocations[i])
		}
	}
	//fmt.Println(count)
	//fmt.Println(neighbors[0].counters[0][0])

	readValues := make([]uint8, AddressLength)

	for i := 0; i < AddressLength; i++ {
		counterSums := make([]int, MaxRange)
		for _, hl := range neighbors {
			for k := 0; k < int(MaxRange); k++ {
				counterSums[k] += int(hl.counters[i][k])
			}
		}
		readValues[i] = uint8(max(counterSums))
	}
	return readValues
}

// max finds the largest value in an array -- used for retrieving counters
func max(s []int) int {
	var biggest, index int
	for i, v := range s {
		if v > biggest {
			biggest = v
			index = i
		}
	}
	return index
}
