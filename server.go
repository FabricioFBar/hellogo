package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    // Mensagem principal
    fmt.Fprintf(w, "<h1>Olá, teste de aplicação GO no OpenShift</h1>")

    // Data e hora de execução
    currentTime := time.Now().Format("02/01/2006 15:04:05")
    fmt.Fprintf(w, "<p><strong>Data de execução:</strong> %s</p>", currentTime)

    // Nome do servidor (hostname)
    hostname, err := os.Hostname()
    if err != nil {
        hostname = "Servidor desconhecido"
    }
    fmt.Fprintf(w, "<p><strong>Servidor:</strong> %s</p>", hostname)
}

func main() {
    http.HandleFunc("/", helloWorld)
    fmt.Println("Servidor iniciado em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
