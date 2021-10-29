package vm

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag         bool
	VersionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	cpOption         string
	XjreOption       string
	Class            string
	args             []string
}

// 解析jvm启动参数
func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	// 参数列表的第一个赋值给class名 后续的参数重新构成参数列表
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
