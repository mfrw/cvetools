package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/facebookincubator/nvdtools/cpedict"
)

var in = flag.String("in", "/home/mfrw/nvd/official-cpe-dictionary_v2.3.xml.gz", "Input gz file")

func main() {
	flag.Parse()
	if *in == "" {
		log.Println("Please provide input gz file")
	}

	rd, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	zr, err := gzip.NewReader(rd)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	// its single stream & not multistream

	zr.Multistream(false)

	cpelis, err := cpedict.Decode(zr)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: make it better .. probably use search
	printFancy(cpelis)

}

func printFancy(cpelis *cpedict.CPEList) {
	for _, v := range cpelis.Items {
		fmt.Println(v.Name)
	}
}
