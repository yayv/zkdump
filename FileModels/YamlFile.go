package FileModels

import (
	"fmt"
	"strings"
)

func (pf *PkvFile)PrintYAML(root string){
	var result []string
	
	var printLine string

	for _,v := range pf.Lines {
		printLine = ""

		// TODO: remove root from Path
		var path string
		if strings.HasPrefix(v.Path, root) {
			path = v.Path[len(root):];
		} else {
			path = v.Path
		}

		result = strings.Split(path, "/")

		printLine = strings.Repeat("\t",len(result))
		printLine += result[len(result)-1]+":"
		if v.ValType != "path" {
			printLine += v.Value
		}

		fmt.Println(printLine)
	}

}

