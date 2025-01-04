package main
import "net/http"
import "fmt"
import "log"
const DIR = "index.html"

func serveStatic() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

var URL string = ""
var CHANGE_URL string = ""
func serveHTML(w http.ResponseWriter, r *http.Request) {
    URL = r.URL.String()
    fmt.Println("URL UPDATE: ", URL)

    http.ServeFile(w, r, "index.html")
}

func run_app() {

    serveStatic()
    http.HandleFunc("/", serveHTML)
    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
