package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"github.com/xtsoler/d2s"
	"io/ioutil"
)

func main() {


	
	arg := os.Args[1]
	//fmt.Println(arg)

	// path := "papadela.d2s"
	path := arg
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening .d2s file", err)
	}

	defer file.Close()

	char, err := d2s.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// Prints character name and class.
	//fmt.Println(char.Header.Name)
	//fmt.Println(char.Header.Class)
	//fmt.Println(char.Header)
	
	data, err := json.Marshal(char)
    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile(arg+".json", data, 0644)
	//fmt.Println(char)
	//fmt.Println(string(data))
	fmt.Println("parse completed.")
}