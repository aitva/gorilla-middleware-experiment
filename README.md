# Gorilla Middleware Experiment

This repo contains a simple API built with [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
and [Gorilla Handlers](http://www.gorillatoolkit.org/pkg/handlers). The API
has only one route `/count/` with two methods: `GET` and `PUT`. Adding the CORS middleware
should allow `OPTIONS` request, but the middleware is bypass by the router
and an error `405` is returned.

The following command can be use to test CORS request: `curl -i -X OPTIONS -H 'Origin: exemple.com' -H 'Access-Control-Request-Method: PUT' localhost:3000/count/`

The code can be fixed by adding a route for all method, avoiding the router bypass:

```go
r.HandleFunc("/count/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusMethodNotAllowed)
})
```
