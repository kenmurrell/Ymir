package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const fileSizeAveDefault = 100
const fileSizeDevDefault = 50
const addPuncDefault = false

var fileSizeAve int
var fileSizeDev int

var words = []string{"alpha", "beta", "omega"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createFile(wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	filename := fmt.Sprintf("%s.txt", words[rand.Intn(len(words))])
	f, err := os.Create(filename)
	defer f.Close()
	if err == nil {
		fb := bufio.NewWriter(f)
		defer fb.Flush()
		numwords := int(rand.NormFloat64()*float64(fileSizeDev)) + fileSizeAve
		if numwords < 0 {
			numwords = 0
		}

		for i := 0; i < numwords; i++ {
			k := rand.Intn(len(words))
			fb.WriteString(words[k])
			fb.WriteString(" ")
		}
	}
}

func main() {
	flag.IntVar(&fileSizeAve, "fileSizeAve", fileSizeAveDefault, "The average file size of the generated files.")
	flag.IntVar(&fileSizeDev, "fileSizeDev", fileSizeDevDefault, "The deviation in the file size of the generated files.")
	flag.Parse()

	var wg sync.WaitGroup
	numFiles := rand.Intn(3)
	for i := 0; i < numFiles; i++ {
		wg.Add(1)
		go createFile(&wg)
	}
	wg.Wait()
}
