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
	"os"
	//"os/user"
	//"path/filepath"
)

func init() {

}

func main() {
	fmt.Println("HtBasic Interpreter v0.1.1")

	flag.Parse()

	if flag.NArg() == 1 {
		if _, err := os.Stat(flag.Arg(0)); os.IsNotExist(err) {
			fmt.Printf("%s: no such file or directory\n", flag.Arg(0))
			return
		} else {
			var scanner = scanner.OpenScanner(flag.Arg(0))
			var value = scanner.NextTokenValue()
			i := 0
			for value != "" {
				fmt.Printf("%4d: %s\n", value)
				i++
				value = scanner.NextTokenValue()
			}
			/*			var line = scanner.NextLine()
						i := 0
						for line != "" {
							i++
							fmt.Printf("%4d: %s\n", i, line)
							line = scanner.NextLine()
						}*/
		}
	} else if flag.NArg() == 0 {
		fmt.Println("No source files, quitting")
		return
	} else {
		fmt.Println("No support for multiple files...yet")
		return
	}

}
