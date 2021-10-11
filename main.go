package main

// #include <Python.h>
// int PyArg_ParseTuple_ss(PyObject* args, char** a, char** b);
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
	var cwd *C.char

	if C.PyArg_ParseTuple_ss(args, &cmd, &cwd) == 0 {
		fmt.Println("Could not parse args")
		os.Exit(1)
	}

	var cmdStr string = C.GoString(cmd)
	var cwdStr string = C.GoString(cwd)

	cmdParts := strings.Split(cmdStr, " ")

	if len(cmdParts) == 0 {
		log.Fatal("No valid command provided")
	}

	runCmd(cwdStr, cmdParts[0], cmdParts[1:]...)

	C.Py_IncRef(C.Py_None)
	return C.Py_None
}

func runCmd(cwd, command string, args ...string) {
	cmd := exec.Command(command, args...)

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = cwd

	err := cmd.Run()

	if err != nil {
		errMsg := stderr.String()

		if len(errMsg) > 0 {
			log.Fatal(errMsg)
		}
	}

	info := stdout.String()
	if len(info) > 0 {
		fmt.Printf("%s\n", info)
	}
}

func main() {}
