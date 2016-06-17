package sdm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadInputFile reads the files and write to HardLocation in Memory Space
func ReadInputFile(name string, sdm *SparseDistributedMemoryImpl) {
	inFile, _ := os.Open(name)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		stringSlice := strings.Split(line, ",")

		si, _ := sliceAtoi(stringSlice)
		store(si, sdm)
	}
}

func sliceAtoi(sa []string) ([]uint8, error) {
	si := make([]uint8, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, uint8(i))
	}
	return si, nil
}

// Writetofile just enter the pic is []byte for and name like "test.ppm"
func Writetofile(pic []byte, name string) {
	filename := name

	fmt.Println("writing: " + filename)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	n, err := io.WriteString(f, "P2\n")
	n, err = io.WriteString(f, "80 80\n")
	n, err = io.WriteString(f, "250\n")

	for k := range pic {
		n, err = io.WriteString(f, fmt.Sprintf("%d ", pic[k]))
	}

	if err != nil {
		fmt.Println(n, err)
	}
	f.Close()
}
