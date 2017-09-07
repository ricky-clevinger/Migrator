package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	rename(os.Args[1])

}

func rename(path string) {

	files, err := ioutil.ReadDir(path)
    	if err != nil {
        	log.Fatal(err)
    	}

	count := 0

	for _, c := range files {
		
		if (filepath.HasPrefix(c.Name(), "V")){	
			count = count + 1
		}	
	}

    for _, f := range files {
		if filepath.Ext(f.Name()) == ".sql" {
			if (!filepath.HasPrefix(f.Name(), "V")){
				count = count + 1
				os.Rename(path + "/" + f.Name(), path + "/V" + strconv.Itoa(count) + "__" + strings.TrimSuffix(f.Name(), ".sql") + ".sql")
			}
   		 }
    }
	fmt.Println("Completed!")
}


