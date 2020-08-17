# SimpleCors

`simplecors` contains a single configurable Cors type and a single middlware function. It is not an exhaustive implementation. I often wind up re-writing the same 6 lines of code when implementing Cors (quickly) in dev environments and simple http server projects and grew tired of the repetition. 

```
mux := http.NewServeMux()
cors := simplecors.DefaultCors()
mux.HandleFunc("/", cors.Middleware(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("success!"))
}))
http.ListenAndServe(":8080", mux)
```
