package main

import (
  "fmt";
  "net/http";
  "log";
  "os";
  "io/ioutil";
  "strings";
)

func handler(w http.ResponseWriter, r *http.Request) {
  url := strings.Join([]string{"http://blog.fefe.de", r.URL.Path, "?", r.URL.RawQuery}, "");
  response, err := http.Get(url);
  if (err != nil) {
    os.Exit(2);
  } else {
    defer response.Body.Close();
    contents, err := ioutil.ReadAll(response.Body);
    if (err != nil) {
      fmt.Printf("%s", err);
      os.Exit(1);
    }
    fmt.Fprintf(w, string(contents));
    log.Println(url);
  }
}

func main() {
  http.HandleFunc("/", handler);
  err := http.ListenAndServe(":8080", nil);
  if (err != nil) {
    log.Fatal("ListenAndServe: ", err);
  }
}