package main

import (
	"flag"
	"fmt"
	"github.com/sourcefellows/mongo-query/internal"
	"log"
	"os"
	"strings"
)

func main() {

	inFile := flag.String("in", "", "path to file with Golang structs")
	outDirectory := flag.String("outDir", "", "path to output directory - a subdirectory \"filter\" will be generated automatically")
	explicitStructs := flag.String("only", "", "list of struct names - only given struct names will be used for code generation")

	flag.Parse()

	if inFile == nil || *inFile == "" {
		log.Println("no input file given")
		flag.PrintDefaults()
		return
	}

	if outDirectory == nil || *outDirectory == "" {
		workingDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		outDirectory = &workingDir
	}

	outDir := fmt.Sprintf("%s/filter", *outDirectory)
	_, err := os.Stat(outDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(outDir, 0774)
		if err != nil {
			panic(err)
		}
	}

	mongoDbStructs, err := internal.ParseFile(*inFile, *explicitStructs)
	if err != nil {
		panic(err)
	}

	writerType := internal.StructWriter
	for _, mongoDbStruct := range mongoDbStructs {
		//out := os.Stdout
		outFile := fmt.Sprintf("%s/%sFilter.go", outDir, strings.ToLower(mongoDbStruct.Name))
		out, err := os.OpenFile(outFile, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}

		err = internal.Write(mongoDbStruct, writerType, out)
		if err != nil {
			log.Fatalln(err)
		}

		out.Close()
	}
}
