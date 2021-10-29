package main

import (
	. "gvm/vm"
)

func main() {
	cmd := ParseCmd()

	if cmd.VersionFlag {
		println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		PrintUsage()
	} else {
		NewJVM(cmd).Start()
	}
}
