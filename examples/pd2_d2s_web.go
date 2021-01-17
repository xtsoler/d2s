package main

import (
	"fmt"
	"encoding/json"
	"syscall/js"
	"github.com/xtsoler/d2s"
	//"io/ioutil"
	"bytes"
	//"log"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func printMessage(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("[golang] parser started!")
	//message := inputs[0].String()
	d2s_data := inputs[0]
	inBuf := make([]uint8, d2s_data.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, d2s_data)
	
	println("[golang] input array size is",d2s_data.Get("byteLength").Int())
	
	r := bytes.NewReader(inBuf)
	showDebugOutput := false
	char, err := d2s.Parse(r, showDebugOutput)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("[golang] parser failed to parse the data.")
	}
	
	//fmt.Println("[golang] " + string(char.Header.Name))
	
	data, err := json.Marshal(char)
    if err != nil {
        //log.Fatal(err)
		fmt.Println("[golang] error creating the json output.")
    }

    //err = ioutil.WriteFile(arg+".json", data, 0644)
	//fmt.Println(char)
	//fmt.Println(string(data))
	
	c <- true
	return string(data)
}

func main() {
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	<-c
	println("[golang] parser ended.")
}