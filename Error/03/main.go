package main 

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	
)
func main(){
	defer foo()
	f,err:=os.Open("names.txt")
	if err != nil { 
		//fmt.Println(err)
		//log.Println(err)
		log.Panicln(err)
		//log.Fatalln(err)
		return
	}
	defer f.Close()
	bs,err := ioutil.ReadAll(f);
	if err != nil { 
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}

func foo(){ 
	fmt.Println("when os.Exit() is called , the deffered function doesnt run")
}