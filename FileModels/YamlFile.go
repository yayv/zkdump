package FileModels

import (
	"fmt"
	"strings"
)

func (pf *PkvFile)PrintYAML(root string){
	var result []string
	
	for _,v := range pf.Lines {
		// TODO: remove root from Path
		var path string
		if strings.HasPrefix(v.Path, root) {
			path = v.Path[len(root):];
		} else {
			path = v.Path
		}

		result = strings.Split(path, "/")
		if result[0]=="" {
			result = result[1:]
		}
		
		if len(result)>0 {
			if v.ValType != "path" {
				fmt.Printf("%s%s:%s\n",strings.Repeat("\t",len(result)-1), result[len(result)-1], v.Value)
			} else {
				fmt.Printf("%s%s:\n",strings.Repeat("\t",len(result)-1), result[len(result)-1])
			}
		}	

		
	}

}

