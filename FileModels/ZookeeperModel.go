package FileModels

import (
	"fmt"
	"github.com/go-zookeeper/zk"
)

/*
func GetZkPath(path string) string {

	return "ok"
}
*/

func (pf *PkvFile)SyncToZk(c *zk.Conn){
	var stat *zk.Stat
	var err error
	//var acl []zk.ACL
	var res,root string

	if pf.Rootpath=="/" {
		root = ""
	} else {
		root = pf.Rootpath
	}

	fmt.Println(pf.Lines)

	for _,v := range pf.Lines {
		if ok,_,_ :=c.Exists(pf.Rootpath+v.Path);!ok {
			res,err = c.Create(root+v.Path, nil,0, zk.WorldACL(zk.PermAll))
			fmt.Println("Create:",res,err)
		}

		if v.ValType == "path" {
			stat,err = c.Set(root+v.Path, nil, 0)
		} else {
			stat,err = c.Set(root+v.Path, []byte(v.Value), 0)
		}

		if err!=nil {
			fmt.Println(root+v.Path,stat,err)
		} else {
			fmt.Printf("%s\n", root+v.Path)
		}
	}	
}
