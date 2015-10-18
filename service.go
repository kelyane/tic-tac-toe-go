package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "fmt"
  //  "encoding/json"
)

func main(){
  // Instantiate a new router
  r := httprouter.New()

  // Get a user resource
  r.GET("/turn/:x/:y/:player", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
      // Stub an example user
      /*u := field.Turn{
          x:      p.ByName("x"),
          y:      p.ByName("y"),
          player: p.ByName("player"),
      }*/
      x :=      p.ByName("x")
      y :=      p.ByName("y")
      player := p.ByName("player")

      // Marshal provided interface into JSON structure
      //uj, _ := json.Marshal(u)

      // Write content-type, statuscode, payload
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(200)
      fmt.Fprintf(w, "%s", x)
      fmt.Fprintf(w, "%s", y)
      fmt.Fprintf(w, "%s", player)
  })

  // Fire up the server
  http.ListenAndServe("localhost:3000", r)
}
