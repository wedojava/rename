package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
		if newname := match(f.Name()); newname != "" {
			os.Rename(filepath.Join(*dir, f.Name()), filepath.Join(*dir, newname))
		}
	}
}

func match(filename string) string {
	re := regexp.MustCompile(`(.*?.dat)-(UDP-\d{6}).*@(.*)`)
	pieces := re.FindAllStringSubmatch(filename, -1)
	if pieces == nil || len(pieces) == 0 {
		return ""
	}
	return fmt.Sprintf("[%s][%s]@%s", pieces[0][1], pieces[0][2], pieces[0][3])
}
