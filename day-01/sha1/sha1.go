package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("sha1.go.gz")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(sig)
}

/*
if file naemes end with .gs
	cat http.log.gz| gunzip |sha1sum
else
	cat http.log | sha1sum
*/
//
func sha1Sum(filename string) (string, error) {
	// idiom: acquire a resource, check for error, defer release
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("can't open: %s\n", err)
		return "", err
	}
	defer file.Close() // deferred are called in LIFO order

	var r io.Reader = file

	if strings.HasSuffix(filename, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return "", err
		}
		defer gz.Close()
		r = gz
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		fmt.Printf("error: %s", err)
		return "", err
	}

	sig := w.Sum(nil)
	// fmt.Sprintf("%x\n", sig)
	return fmt.Sprintf("%x", sig), nil
}
