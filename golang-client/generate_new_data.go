package main

import (
	"golang-client/tools"
)

func main() {
	dataconf := tools.ReadDataDescriptor()
	functionconf := tools.ReadFunctions()
	tools.GenData(dataconf, true)
	tools.GenFunction(dataconf, functionconf)

	//----------------call proto gen cmd----------------
	//batFilePath := "\\message\\generate.bat" // Adjust the relative path as needed
	//
	//// Get the current working directory
	//wd, err := os.Getwd()
	//if err != nil {
	//	log.Fatalf("Failed to get current working directory: %v", err)
	//}
	//log.Printf("Current working directory: %s", wd)
	//// Join the current working directory with the relative path
	//fullPath := filepath.Join(wd, batFilePath)
	//cmd := exec.Command("cmd", "/C", fullPath)
	//
	//// Buffers to capture standard output and standard error
	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &stderr
	//
	//// Run the command and check for errors
	//err = cmd.Run()
	//if err != nil {
	//	log.Fatalf("Failed to run .bat file: %v\nStderr: %s", err, stderr.String())
	//}
	//
	//log.Println("Successfully ran .bat file")
	//log.Printf("Output: %s", out.String())
}
