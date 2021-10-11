package main

// #include <Python.h>
// int PyArg_ParseTuple_s(PyObject* args, char** a)
import "C"
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//export command
func command(self *C.PyObject, args *C.PyObject) *C.PyObject {

	var cmd *C.char

	if C.PyArg_ParseTuple_s(args, &cmd) == 0 {
		fmt.Println("Could not parse args")
		os.Exit(1)
	}

	var cmdStr string = C.GoString(cmd)

	cmdParts := strings.Split(cmdStr, " ")

	runCmd(cmdParts[0], cmdParts[1:]...)

	C.Py_IncRef(C.Py_None)
	return C.Py_None
}

func runCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		errMsg := stderr.String()
		if len(errMsg) > 0 {
			log.Fatal(errMsg)
		}
	}

	info := stdout.String()
	if len(info) > 0 {
		log.Fatalf("%s\n", info)
	}

	log.Printf("Finished '%s'\n", cmd)
}

func main() {}
