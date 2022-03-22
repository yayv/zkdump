package FileModels

import (
	"fmt"
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
