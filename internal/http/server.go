package http

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/mozillazg/go-unidecode"
    "github.com/nihildacta/line-server/pkg/file"
)

type (
    // Http server for a line reader
    Server struct {
        Reader file.Reader //Reader a line reader
    }
)

// New Given a reader returns a new Line Reader Server
func New(reader file.Reader) Server {
    return Server{
        Reader: reader,
    }
}

// hander handles the requests for this server
// Only accepts GET reqesuts and valid positive numbers
// Writes out the currespondent line, status 413 for line numbers above the existing ones
// or status 404 in case of error
func (server Server) handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }

    lineNumber, err := strconv.Atoi(r.URL.Path[len("/lines/"):])
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    if lineNumber <= 0 {
        http.Error(w, "line number must be a positive", http.StatusNotFound)
        return
    }

    line, err := server.Reader.ReadLine(lineNumber)
    if err != nil {
        http.Error(w, err.Error(), 413)
        return
    }

    fmt.Fprint(w, unidecode.Unidecode(line))
}

// ListenAndServe binds and handles users request given a port
func (server Server) ListenAndServe(port int) error {
    http.HandleFunc("/lines/", server.handler)
    return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
