package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dir = flag.String("dir", "./", "where is the files you wanna rename?")

func main() {
	flag.Parse()
	filelist, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Println(err)
	}
	for _, f := range filelist {
		if strings.HasSuffix(f.Name(), ".exe") {
			continue
		}
		os.Rename(filepath.Join(*dir, f.Name()),
			filepath.Join(*dir, match(f.Name())))
	}
}

func match(filename string) string {
	pieces := strings.Split(filename, "-UDP-")
	prefix := pieces[0]
	secondPiece := strings.Join(pieces[1:], "")
	atPieces := strings.Split(secondPiece, "@")
	subfix := atPieces[len(atPieces)-1]
	return fmt.Sprintf("[%s]UDP@%s", prefix, subfix)
}
