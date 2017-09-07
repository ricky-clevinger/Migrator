package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
	"os/exec"
	"cfg"
	"strconv"
	"strings"
)


func Rename(path string) {

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
        	fmt.Println(f.Name())
		if (!filepath.HasPrefix(f.Name(), "V")){
			count = count + 1
			os.Rename(path + "/" + f.Name(), path + "/V" + strconv.Itoa(count) + "__" + strings.TrimSuffix(f.Name(), ".sql") + ".sql")
			}
   		 }
    	}
}


func Command6(app,arg0,arg1,arg2,arg3,arg4 string) {


   	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
   	output, err := cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(output))
    		return
	} else {
    		fmt.Println(string(output))
	}

}

func Command5(app,arg0,arg1,arg2,arg3 string) {


   	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
   	output, err := cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(output))
    		return
	} else {
    		fmt.Println(string(output))
	}

}


func Command4(app,arg0,arg1,arg2 string) {


   	cmd := exec.Command(app, arg0, arg1, arg2)
   	output, err := cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(output))
    		return
	} else {
    		fmt.Println(string(output))
	}

}

func Command3(app,arg0,arg1 string) {


   	cmd := exec.Command(app, arg0, arg1)
   	output, err := cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(output))
    		return
	} else {
    		fmt.Println(string(output))
	}

}


func RunCommands(){

	mymap := make(map[string]string)
    	err := cfg.Load("config.conf", mymap)
    	if err != nil {
        	log.Fatal(err)
    	}

	path := mymap["path"]

	if (mymap["repository"] == "true"){
		Command4("git", "-C", path, "pull")
	}

	Rename(mymap["path"])

	if (mymap["repository"] == "true"){
		Command5("git", "-C", path, "add", "-A")
		Command5("git", "-C", path, "add", "*")
		Command6("git", "-C", path, "commit", "-m", "New statements")
	}

	if (mymap["flyway"] == "true"){
		Command3("flyway", "-configFile=./config.conf", "migrate")
	}

	if (mymap["repository"] == "true"){
		Command4("git", "-C", path, "push")
	}
}
