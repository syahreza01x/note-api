package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Note struct
type Note struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

var notes []Note

func getNotes(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if notes == nil {
        notes = []Note{}
    }

    jsonData, err := json.MarshalIndent(notes, "", "  ")
    if err != nil {
        http.Error(w, "Failed to encode notes", http.StatusInternalServerError)
        return
    }

    w.Write(jsonData)
}

func viewNotes(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    html := `<html><head><title>Daftar Notes</title>
    <style>
      body { font-family: Arial, sans-serif; margin: 20px; }
      h1 { color: #333; }
      ul { list-style-type: none; padding: 0; }
      li { background: #f4f4f4; margin-bottom: 10px; padding: 10px; border-radius: 5px; }
      .title { font-weight: bold; font-size: 1.2em; }
      .content { margin-top: 5px; }
    </style>
    </head><body><h1>Daftar Notes</h1><ul>`

    for _, note := range notes {
        html += "<li><div class='title'>" + note.Title + "</div>"
        html += "<div class='content'>" + note.Content + "</div></li>"
    }

    html += "</ul></body></html>"

    w.Write([]byte(html))
}

func getNote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, note := range notes {
        if note.ID == params["id"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(note)
            return
        }
    }
    http.Error(w, "Note not found", http.StatusNotFound)
}

func createNote(w http.ResponseWriter, r *http.Request) {
    var note Note
    err := json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    note.ID = strconv.Itoa(rand.Intn(1000000))
    notes = append(notes, note)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(note)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, note := range notes {
        if note.ID == params["id"] {
            notes = append(notes[:i], notes[i+1:]...)
            var updated Note
            err := json.NewDecoder(r.Body).Decode(&updated)
            if err != nil {
                http.Error(w, "Invalid input", http.StatusBadRequest)
                return
            }
            updated.ID = note.ID
            notes = append(notes, updated)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updated)
            return
        }
    }
    http.Error(w, "Note not found", http.StatusNotFound)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, note := range notes {
        if note.ID == params["id"] {
            notes = append(notes[:i], notes[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "Note not found", http.StatusNotFound)
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/notes", getNotes).Methods("GET")
    r.HandleFunc("/notes/view", viewNotes).Methods("GET")
    r.HandleFunc("/notes/{id}", getNote).Methods("GET")
    r.HandleFunc("/notes", createNote).Methods("POST")
    r.HandleFunc("/notes/{id}", updateNote).Methods("PUT")
    r.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")

    log.Println("Server running on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
