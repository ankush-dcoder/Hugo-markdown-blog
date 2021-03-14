package main

import (
  "log"
  "net/http"
)

func main() {
  //simple http server serving public directory's static content
  fs := http.FileServer(http.Dir("./public"))
  http.Handle("/", fs)

  log.Println("Listening on :3000...")
  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal(err)
  }
}