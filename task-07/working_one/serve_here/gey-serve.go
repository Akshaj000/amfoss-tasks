/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	filetwo, _ := os.Create("out.txt")
	defer filetwo.Close()
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.forbes.com/real-time-billionaires/#c53c1983d788", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			filetwo.WriteString(string(r.Body))
		},
	}).Start()

	filetwo.Close()

	err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
	file, _ := os.ReadFile("out.txt")
	file2, _ := os.Create("index.html")
	file2.Write(file)
	os.Remove("out.txt")

	fname := `index.html`
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()
	line, err := popLine(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("pop:", string(line))
	f.Seek(0, 0)
	f.WriteString("<!doctype html>")
	f.Close()

	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	matched, _ := filepath.Match("*.txt", fi.Name())

	if matched {
		read, err := ioutil.ReadFile("out.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println(path)

		newContents := strings.Replace(string(read), "Country/Territory", "Country", -1)

		fmt.Println(newContents)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}
func popLine(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	nw, err := io.Copy(f, buf)
	if err != nil {
		return nil, err
	}
	err = f.Truncate(nw)
	if err != nil {
		return nil, err
	}
	err = f.Sync()
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return line, nil
}
