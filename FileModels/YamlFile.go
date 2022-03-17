package FileModels

import (
	"fmt"
	"strings"
)

func (pf *PkvFile)PrintYaml(){
	var result []string
	
	var printLine string

	for _,v := range pf.Lines {
		printLine = ""

		result = strings.Split(v.Path, "/")

		printLine = strings.Repeat("\t",len(result))
		printLine += result[len(result)-1]+":"
		if v.ValType != "path" {
			printLine += v.Value
		}

		fmt.Println(printLine)
	}

}

