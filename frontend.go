package main
import "net/http"
import "fmt"
import "os"
const DIR = "index.html"

func handler(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open(DIR)
    if err != nil {
	fmt.Printf("ERR: Unable to open %s\n", DIR)
	return
    }

    defer file.Close()
    w.Header().Set("Content-Type", "text/html")  

    _, err = file.WriteTo(w)
    if err != nil {
	http.Error(w, "Could not write HTML file to response", http.StatusInternalServerError)
    }
}

func run_app() {
    http.HandleFunc("/", handler) // Handle requests to the root path
    fmt.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil) // Start the server on port 8080
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
