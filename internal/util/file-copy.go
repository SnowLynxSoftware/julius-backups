package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FileCopy(source string, file os.FileInfo, destination string) {
	fmt.Println("-- COPYING >> " + destination + "/" + file.Name())
	input, err := ioutil.ReadFile(source + "/" + file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destination+"/"+file.Name(), input, 0644)
	if err != nil {
		fmt.Println("Error creating ", file.Name())
		fmt.Println(err)
		return
	}
}
