// Package self
// a program that complete itself
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/posener/complete"
)

var (
	ellipsis = complete.PredictSet("./...")
	goFiles  = complete.PredictFiles("*.go")
	anyFile  = complete.PredictFiles("*")
)

func main() {

	// add a variable to the program
	var name string
	flag.StringVar(&name, "name", "", "Give your name")

	build := complete.Command{
		Flags: complete.Flags{
			"-o": anyFile,
			"-i": complete.PredictNothing,

			"-a":             complete.PredictNothing,
			"-n":             complete.PredictNothing,
			"-p":             complete.PredictAnything,
			"-race":          complete.PredictNothing,
			"-msan":          complete.PredictNothing,
			"-v":             complete.PredictNothing,
			"-work":          complete.PredictNothing,
			"-x":             complete.PredictNothing,
			"-asmflags":      complete.PredictAnything,
			"-buildmode":     complete.PredictAnything,
			"-compiler":      complete.PredictAnything,
			"-gccgoflags":    complete.PredictSet("gccgo", "gc"),
			"-gcflags":       complete.PredictSet("context", "gotypes", "netipv6zone", "printerconfig"),
			"-installsuffix": complete.PredictAnything,
			"-ldflags":       complete.PredictAnything,
			"-linkshared":    complete.PredictNothing,
			"-name":          complete.PredictSet("magnus", "gotypes", "netipv6zone", "printerconfig"),
			"-toolexec":      complete.PredictAnything,
		},
		Args: complete.PredictFiles("*.go"),
	}

	name2 := complete.Command{Flags: complete.Flags{"name": complete.PredictSet("John", "Smith")}}

	gogo := complete.Command{
		Sub: complete.Commands{
			"build":   build,
			"install": build, // install and build have the same flags
			"name":    name2,
		},
		GlobalFlags: complete.Flags{
			"-h": complete.PredictNothing,
		},
	}

	// create the complete command
	cmp := complete.New(
		"my_program",
		gogo,
	)

	// AddFlags adds the completion flags to the program flags,
	// in case of using non-default flag set, it is possible to pass
	// it as an argument.
	// it is possible to set custom flags name
	// so when one will type 'self -h', he will see '-complete' to install the
	// completion and -uncomplete to uninstall it.
	cmp.CLI.InstallName = "complete"
	cmp.CLI.UninstallName = "uncomplete"
	cmp.AddFlags(nil)

	// parse the flags - both the program's flags and the completion flags
	flag.Parse()

	// run the completion, in case that the completion was invoked
	// and ran as a completion script or handled a flag that passed
	// as argument, the Run method will return true,
	// in that case, our program have nothing to do and should return.
	if cmp.Complete() {
		return
	}

	// if the completion did not do anything, we can run our program logic here.
	if name == "" {
		fmt.Println("Your name is missing")
		os.Exit(1)
	}

	fmt.Println("Hi,", name)
}
