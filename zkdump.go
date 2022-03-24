package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"syscall"
	"time"
	"FileModels"

	"github.com/go-zookeeper/zk"
	"golang.org/x/crypto/ssh/terminal"

	"gopkg.in/alecthomas/kingpin.v2"
//	"gopkg.in/yaml.v2"
)

var (
	pkv       FileModels.PkvFile
	err       error
	c         = &zk.Conn{}
	app       = kingpin.New("zkdump", "A command-line utility to import/export Zookeeper data.").Author("Dennis Waterham <dennis.waterham@oracle.com>").Version("1.0")
	servers   = app.Flag("server", "Host name and port to connect to (host:port)").Required().Short('s').Strings()
	verbose   = app.Flag("verbose", "Print verbose.").Short('v').Bool()
	user      = app.Flag("user", "Username to use for digest authentication.").Short('u').String()
	password  = app.Flag("password", "Password to use for digest authentication (will read from TTY if not given).").Short('p').String()
	recursive = app.Flag("recursive", "Get nodes recursively.").Short('r').Bool()
	rootpath  = app.Arg("path", "Root path (default: \"/\").").Default("/").String()
	filetype  = app.Flag("type", "import/export file type, JSON or YAML ").Short('t').String()
	file      = app.Flag("file", "file name for import or export").Short('f').String()
	imp       = app.Flag("import", "import key/value pairs from file to zookeeper").Short('i').Bool()
	exp       = app.Flag("export", "export key/value pairs to file from zookeeper").Short('e').Bool()
)

type zkNode struct {
	Name     string
	Path     string
	Data     string   `json:",omitempty"`
	Children []zkNode `json:",omitempty"`
}

func (z *zkNode) getChildren() {
	items, st, err := c.Children(z.Path)

	if err != nil {
		log.Fatal(err)
	}

	if st.NumChildren == 0 {
		return
	}

	for _, child := range items {
		z.Children = append(z.Children, *getZkNode(path.Join(z.Path, child), child))
	}
}

func main() {

	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *user != "" {
		if *password == "" {
			*password = readPassword()
		}
	}

	c, _, err = zk.Connect(*servers, time.Second, zk.WithLogInfo(*verbose))
	defer c.Close()
	check(err)

	if *user != "" {
		verboseLog("Adding digest authentication for user %s", *user)
		c.AddAuth("digest", []byte(*user+":"+*password))
	}

	verboseLog("Checking if root path %s exists", *rootpath)
	exists, _, err := c.Exists(*rootpath)
	check(err)

	if !exists {
		log.Fatalf("ERROR: Path %s doesn't exist", *rootpath)
	}

	if (*imp) && (*exp) {
		*imp = false
		verboseLog("both import and export options are choosed, set option to export -e:%s", *filetype);
	}

	if *imp {
		doImport()
	} else {
		doExport()
	}
}

func doImport(){
	// TODO: check file type in arguments
	if *file=="" {
		fmt.Println("No filename set. Please use -f to set filename.")
	}

	fmt.Println("filetype:",*filetype)
    // import yaml.Unmarshal(os.Args[1:])

    switch *filetype {
    	case "PKV":
    		pkv.LoadFromPKV(*file)
    	case "JSON":
			pkv.LoadFromJSON(*file)
		default:
     		pkv.LoadFromYAML(*file)
    }

}

func doExport(){
	// Get Root node
	rootNode := getZkNode(*rootpath, path.Base(*rootpath))
	//	saveJSON(rootNode, "test.json")

	if *filetype == "" {
		//*types =  "JSON"
	} 

	if *filetype=="JSON" {
		bin, _ := json.MarshalIndent(&rootNode, "", "  ")
		fmt.Println(bin)
		pkv.PrintJSON()
	} else if *filetype=="PKV" {
		//bin, _ := yaml.Marshal(&rootNode)
		//fmt.Println(string(bin))	
		pkv.PrintAll()
	} else {
		pkv.PrintYAML(*rootpath)
	}
	
}

func verboseLog(s string, p string) {
	if *verbose {
		log.Printf(s+"\n", p)
	}
}

func getZkNode(path, name string) *zkNode {
	bin, st, err := c.Get(path)
	check(err)

	zkNode := &zkNode{Path: path, Name: name, Data: string(bin)}

	if st.NumChildren > 0 {
		pkv.Append(path,name,string(bin),"path")
		zkNode.getChildren()
	} else {
		pkv.Append(path,name,string(bin),"value")	
	}

	return zkNode
}

func readPassword() string {
	fmt.Print("Enter Password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	check(err)

	fmt.Printf("\n")
	return string(bytePassword)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printYAML(data interface{}, file string){
	
}

func saveJSON(data interface{}, file string) {
	outFile, err := os.Create(file)
	defer outFile.Close()
	if err != nil {
		log.Fatalln("Error occurred")
	}

	jsonWriter := json.NewEncoder(outFile)
	jsonWriter.SetIndent("", "   ")
	jsonWriter.Encode(&data)
}
