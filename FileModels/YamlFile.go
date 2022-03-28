package FileModels

import (
	"os"
	"fmt"
	"log"
	"regexp"
	"strings"
	"bufio"
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
				fmt.Printf("%s%s:%s\n",strings.Repeat("  ",len(result)-1), result[len(result)-1], v.Value)
			} else {
				fmt.Printf("%s%s:\n",strings.Repeat("  ",len(result)-1), result[len(result)-1])
			}
		}	

		
	}

}

func printPath(path []string) string{
	var fullpath string
	for _,v := range(path) {
		fullpath += "/"+v
	}

	return fullpath
}

func (pf *PkvFile)LoadFromYAML(filename string){
	// TOOD: read key/value file 
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //indent  := 0
    scanner := bufio.NewScanner(file)
    var pkv,key,value string
    var path []string

    rule,_  := regexp.Compile(`([ ]*)([a-zA-Z0-9-_]*):(.*)`)
    rule.Longest()
    // optionally, resize scanner's capacity for lines over 64K, see next example
    i:=0
    for scanner.Scan() {
    	i+=1
    	pkv = scanner.Text()

    	all := rule.FindStringSubmatch(pkv)
    	if len(all)>0 {
	    	//indent 	= len(all[1])/2
	    	key 	= all[2]
	    	value 	= all[3]

    		path = append(path, key)

	        fmt.Println(printPath(path), key, ":",value)
    	} else {
    		fmt.Printf("!%s!",pkv)
    	}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }	

}
