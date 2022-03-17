package main


import (
	"fmt"
	"FileModels"
)

func test() (bool, bool){
	return true,false
}

func main(){
	var pf FileModels.PkvFile

	pf.Append("/","","","path")
	pf.Append("/config/key/k2","k2","val2","string")
	pf.Append("/config/datasource/mysql/bike","bike","jdbc://127.0.0.1:3306/asdfasdfaf","string")
	pf.Append("/config/redis/server","server","127.0.0.1:111","string")

	pf.PrintAll()

	fmt.Println("==============")
	
	pf.FillPath()

	pf.PrintAll()

	fmt.Println("==============")

	pf.PrintYaml()
}
