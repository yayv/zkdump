package FileModels

import (
	"fmt"
	"os"
	"log"
)

func (pf *PkvFile)PrintJSON(){
	for _,v := range pf.Lines {
		if v.ValType == "path" {
			fmt.Println(v.Path,"\t\t")	
		} else {
			fmt.Println(v.Path,"\t\t",v.Name,":",v.Value)
		}
	}
}

func (pf *PkvFile)LoadFromJSON(filename string){
	fmt.Println("In Load From JSON")
	fmt.Println(filename)

	// TOOD: read key/value file 
    file, err := os.Open("./README.md")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

}
