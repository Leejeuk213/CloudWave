package main
import (
"io"
"net/http"
"log"
)
func HelloServer(w http.ResponseWriter, req *http.Request) {
io.WriteString(w, "Hello, Worlds!\n")
}
func main() {
http.HandleFunc("/", HelloServer)
log.Fatal(http.ListenAndServe(":80", nil))
}
