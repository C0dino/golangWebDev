package main

import (
	"log"
	"io"
	"strings"
)

func main()  {
    name := "Stephen Pendino"
    str := fmt.Sprint(`
       <!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>Hello World</title>
    </head>
    <body>
    <h1>` + name + `</h1>
    </body>
    </html>`)

    nf, err := os.Create("index.html")
    if, err != nil {
        log.Fatal("error creating file", err)
    }
    defer nf.Close()

    io.Copy(nf, strings.NewReader(str))
}