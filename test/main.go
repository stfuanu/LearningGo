package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {

	var depth int
	flag.IntVar(&depth, "d", 4, " recursion depth")

	flag.Parse()

	domainname := flag.Arg(0)
	wordListFile := flag.Arg(1)

	if domainname == "" {
		fmt.Fprintln(os.Stderr, "usage: brute [-d=<int>] <domain> [<wordfile>|-]")
		return
	}

	var f io.Reader
	var err error

	// default to stdin for the wordlist
	f = os.Stdin

	if wordListFile != "" && wordListFile != "-" {
		f, err = os.Open(wordListFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file : %s\n", err)
			return
		}
	}

	sc := bufio.NewScanner(f)

	words := make([]string, 0)
	for sc.Scan() {
		words = append(words, sc.Text())
	}

	brute(domainname, words, 1, depth)
}

func brute(domainname string, words []string, depth, maxDepth int) {
	if depth > maxDepth {
		return
	}

	var waaaaait sync.WaitGroup

	for _, w := range words {
		candidate := fmt.Sprintf("%s.%s", w, domainname)

		waaaaait.Add(1)
		// go routine
		go func() {
			brute(candidate, words, depth+1, maxDepth)
			waaaaait.Done()
		}()
		// prints the result
		fmt.Println(candidate)
	}

	waaaaait.Wait()
}
