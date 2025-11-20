package hello 

import ("bufio"; "fmt"; "os"; "strings")

func main() {
	for _, fname := os.Args[1:] {
		var lc, wc, cc int 

		file, err := os.OpenFile(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			cc += len(s)
			wc += len(string.Fields(s))
			lc++
		}
		fmt.Printf("%7d %7d %7d %s\n", lc,wc, cc, fname)
		file.Close()

	}
}