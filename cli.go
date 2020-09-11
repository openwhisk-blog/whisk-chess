package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("usage: <fen> [<time>]")
		os.Exit(1)
	}

	args := make(map[string]interface{})
	if len(os.Args) > 1 {
		args["fen"] = os.Args[1]
	}
	if len(os.Args) > 2 {
		args["time"] = os.Args[2]
	}
	buf, _ := json.Marshal(args)
	fmt.Printf(">>> %s\n", buf)
	res := Main(args)
	buf, _ = json.Marshal(res)
	fmt.Printf("<<< %s\n", buf)
}
