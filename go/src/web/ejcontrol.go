package main
  import (
                  "html/template"
                  "net/http"
  )
  type Page struct {
                  Title string
                  Body  []byte
  }
  func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
                  t, _ := template.ParseFiles(tmpl + ".html")
                  t.Execute(w, p)
  }
  func editHandler(w http.ResponseWriter, r *http.Request) {
                  renderTemplate(w, "edit", &Page{})
  }
  func main() {
                  http.HandleFunc("/sync/", editHandler)
                  http.ListenAndServe(":8080", nil)
  }