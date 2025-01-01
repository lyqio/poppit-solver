package main
import "net/http"
import "fmt"
import "log"
const DIR = "index.html"

func serveStatic() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

var message string = "DefaultMessage"
func send_message(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, message)
}

func run_app() {
    message = "HELLO"

    serveStatic()
    http.HandleFunc("/", serveHTML)
    http.HandleFunc("/api/message", send_message)
    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
