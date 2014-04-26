/**
 * Created with IntelliJ IDEA.
 * User: harry
 * Date: 2/5/14
 * Time: 1:26 PM
 */
package main

import (
	"flag"
	"fmt"
	"github.com/hculpan/htBasic/scanner"
	//"log"
	//"os"
	//"os/user"
	//"path/filepath"
)

func init() {

}

func main() {
	fmt.Println("HtBasic Interpreter v0.1")

	flag.Parse()

	if flag.NArg() > 1 {
		fmt.Println("No source files, quitting")
		return
	} else {
		var scanner = scanner.OpenScanner(flag.Arg(0))
		var line = scanner.NextLine()
		i := 0
		for line != "" {
			i++
			fmt.Printf("%4d: %s\n", i, line)
			line = scanner.NextLine()
		}
	}

}
