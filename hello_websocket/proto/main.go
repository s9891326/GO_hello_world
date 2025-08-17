package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

func main() {
	// os.RemoveAll("./platform/go")
	//os.MkdirAll("./proto/go", os.ModePerm)

	absPath, _ := filepath.Abs("./hello_websocket/proto")
	pb := absPath + "\\pb"
	goFile := absPath + "\\go"
	fmt.Println("pb path:", pb)
	fmt.Println("Go file:", goFile)

	cmd := exec.Command("protoc",
		"--proto_path="+pb,
		"--go_out="+goFile,
		"--go_opt=paths=source_relative",
		"--go-grpc_out="+goFile,
		"--go-grpc_opt=paths=source_relative",
		pb+"\\*",
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Output:", stdout.String())
	fmt.Println("Error Output:\n", stderr.String())

}
