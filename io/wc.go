package hello

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)
func main() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int 
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			cc += len(s)
			wc += len(strings.Fields(s))
			lc++
		}
		fmt.Printf("%7d %7d %7d %s\n", lc,wc, cc, fname)
		file.Close()

	}
}