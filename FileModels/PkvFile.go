package FileModels

import (
	"fmt"
	"sort"
	"strings"
)

/*
This File is a Path-Key-Value Format File Model Object
*/

type PkvLine struct {
	Path     string
	Name     string
	Value    string
	ValType	 string
}

func (pl *PkvLine)String() string{
	return pl.Path
}


type PkvLines []PkvLine
func (a PkvLines) Len() int { return len(a) }
func (a PkvLines) Swap(i, j int) { a[i],a[j]=a[j],a[i] }
func (a PkvLines) Less(i, j int) bool { return a[i].Path < a[j].Path }

type PkvFile struct {
	Lines PkvLines
	Rootpath string
}

func (pf *PkvFile)Append(path, key, value,valueType string){
	pf.Lines = append(pf.Lines, PkvLine{path,key,value,valueType})
}


func (pf *PkvFile)FillPath(){
	var result []string
	max 	:= len(pf.Lines)

	//Paths   := map[string] string {"/":"path"}
	initstack := map[string] int{"/":1}
	for i:=0;i<max;i++ {
		initstack[pf.Lines[i].Path] = 1;
	}

	stack := initstack
	for i:=0;i<max;i++ {
		// split path string
		result = strings.Split(pf.Lines[i].Path, "/")
		
		// push into array
		p := "/"
		for _,v := range result {
			if v=="" {
				continue 
			}
			p = p + v

			if _,ok := stack[p]; ok {
				stack[p] += 1
			} else {
				pf.Append(p,"","","path")
				stack[p] = 1
			}
			p = p + "/"
		}
	}

	sort.Sort(pf.Lines)
}

func (pf *PkvFile)PrintAll(){
	for _,v := range pf.Lines {
		if v.ValType == "path" {
			fmt.Println(v.Path,"\t\t")	
		} else {
			fmt.Println(v.Path,"\t\t",v.Name,":",v.Value)
		}
	}
}

func (pf *PkvFile)LoadFromPKV(filename string){
	fmt.Println(filename)
}
