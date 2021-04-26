package lesson_8

import (
	"flag"
	"log"
)

type Config struct {
	Verbose bool
	LineNumber int
	FileName string
}

func main()  {
	var (
		lineNumber int
		fileName string
	)

	verbose := flag.Bool("verbose", false, "verbosity output print")
	flag.Int64Var(&lineNumber, "n", 10, "line number")
	flag.StringVar(&fileName, "f", "", "file name")

	flag.Parse()

	config := Config{
		Verbose: *verbose,
		LineNumber: lineNumber,
		FileName: fileName,
	}

	if flag.Parsed(){
		config.checkVerbose()
	}

	if flag.NArg() != 1 {
		log.Fatal("")
	}
}

func (config *Config) checkVerbose() {
	if config.FileName != "" {
		log.Println("FileName is empty")
	}
}