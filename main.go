package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"runtime/debug"

	fasthttp "github.com/valyala/fasthttp"
)

func handleError(err error) {
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
}
type MyHandler struct {
	foobar string
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

// // request handler in fasthttp style, i.e. just plain function.
// func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
// 	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
// }


func main() {
	fstree := buildFileEntityTree()


	resp, err := json.Marshal(&fstree)
	handleError(err)

	// err = ioutil.WriteFile("tree.json", resp, 0x777)
	// handleError(err)

	fmt.Printf("fstree: %+v \n", fstree)


	// request handler in fasthttp style, i.e. just plain function.
	fastHTTPHandler := func (ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Add("Content-Type", "application/json")
		ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
		fmt.Fprint(ctx, string(resp))
		// fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
		// fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
	}

	// pass bound struct method to fasthttp
	// myHandler := &MyHandler{
	// 	foobar: "foobar",
	// }
	// fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)

	// pass plain function to fasthttp
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}

type FileSystemNode struct {
	Name string 
	IsFile bool 

	Children []*FileSystemNode 
}

func populateTree(node *FileSystemNode, currentPath string, file fs.FileInfo, callstack *int) {
	node.Name = file.Name()
	node.IsFile = !file.IsDir()
	
	defer func() {
		*callstack++
		// fmt.Printf("callstack: %v \n", *callstack)
	}()
	
	if file.IsDir() {
		if file.Name()[0] == '.' {
			return
		}
		// fmt.Printf("file: %v \n", file.Name())

		thisPath := currentPath + "/" + file.Name()
		files, err := ioutil.ReadDir(thisPath)
		handleError(err)
	
		for _, f := range files {
			fmt.Printf("name: %v; 1st: %v; ch %v; eq: %v \n", f.Name(), f.Name()[0], '.', f.Name()[0] == '.')

			if f.Name()[0] == '.' {
				continue
			}

			// fmt.Printf("file: %v; c: %v \n", file, *callstack)
			
			newNode := new(FileSystemNode)
			// newNode.Parent = node
			newNode.Name = f.Name()
			newNode.IsFile = !f.IsDir()

			node.Children = append(node.Children, newNode)
			// fmt.Printf("file(len): %v; c: %v \n", len(node.Children), *callstack)

			populateTree(newNode, thisPath, f, callstack)
		}
	} else {
		newNode := new(FileSystemNode)
		// newNode.Parent = node
		newNode.Name = file.Name()
		newNode.IsFile = !file.IsDir()
		// fmt.Printf("jsut a file: %v \n", *newNode)
		node.Children = append(node.Children, newNode)
	}
}

func buildFileEntityTree() *FileSystemNode {
	fileSystem := new(FileSystemNode)

	files, err := ioutil.ReadDir("./")
	handleError(err)

	var callstack int

    for _, f := range files {
		fmt.Printf("name: %v; 1st: %v; ch %v; eq: %v \n", f.Name(), f.Name()[0], '.', f.Name()[0] == '.')
		if f.Name()[0] == '.' {
			continue
		}

		fmt.Println(f.IsDir())
        fmt.Println(f.Name())

		populateTree(fileSystem, "./", f, &callstack)
    }

	return fileSystem
}