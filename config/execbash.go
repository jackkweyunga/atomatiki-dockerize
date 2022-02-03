package config

import (
	"os"

	"os/exec"
)

var tempFilePattern string = "cmdtemp-*.sh"

func BashCmdExec(bashCmd string) (*exec.Cmd, error, *os.File) {

	// create a temporary file for commands.
	f, err := os.CreateTemp(".", tempFilePattern)
	if err != nil {
		return nil, err, nil
	}


	// write some commands to the temporay file.
	data := []byte(bashCmd)
	if _, err := f.Write(data); err != nil {
		return nil, err, f
	}
	

	cmd := exec.Command("sh", f.Name())


	return cmd, nil, f

}


func ClearTempFile(f *os.File) {
	// close the temp command file at the end of the program
	f.Close()
	os.Remove(f.Name())
}

