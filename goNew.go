package main

import (
	"log"
	"os"
	"time"
	
	"github.com/brunetto/goutils/debug"
)


func main () () {
	defer debug.TimeMe(time.Now())
	
	if len(os.Args) < 2 {
		log.Fatal("Provide a name for the new file without the extension.")
	}
	
	var (
		newFileTemplate string 
		outFile *os.File
		err error
	)
		
	newFileTemplate = `package main

import (
	"log"
	"time"
	
	"github.com/brunetto/goutils/debug"
)

func main () () {
	defer debug.TimeMe(time.Now())


}

`
	if outFile, err = os.Create(os.Args[1]+".go"); err != nil {
		log.Fatal(err)
	}
	
	if _, err = outFile.WriteString(newFileTemplate); err != nil {
		log.Fatal(err)
	}
	
	
}


