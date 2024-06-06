// main.go
package main

import (
    "fmt"
    "net/http"
)

var webPage = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello Page</title>
    <style>
        body, html {
            margin: 0;
            padding: 0;
            height: 100%;
            display: flex;
            align-items: center;
            justify-content: center;
            font-family: Arial, sans-serif;
        }
        .center {
            text-align: center;
        }
        h1 {
            font-size: 5em;
            font-weight: bold;
        }
    </style>
</head>
<body>
    <div class="center">
        <h1>Hello Vinayak singh!</br>are we good</br></hr>docker test</h1>
    </div>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, this is the Go server!")
}
func handlerName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, webPage)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/vinayak", handlerName)
    http.ListenAndServe(":8080", nil)
}
