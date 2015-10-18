package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "fmt"
    "strconv"
    . "./src/"
  //  "encoding/json"
)

func main(){
  field := NewField()
  r := httprouter.New()

  r.GET("/turn/:x/:y/:player", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
      x, y, player := ExtractGameParams(p)

      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(200)

      if(field.IsGameFinished(x,y,player)){
        fmt.Fprintf(w, "%s", "Is Winner!!!\n")
      }else{
        field.Turn(x,y,player)
        fmt.Fprintf(w, "%s", field.FieldToString())
      }
  })

  http.ListenAndServe("localhost:3000", r)
}

func ExtractGameParams(p httprouter.Params) (x, y int, player Player) {
  x, _ = strconv.Atoi(p.ByName("x"))
  y, _ = strconv.Atoi(p.ByName("y"))
  m := map[string] Player {
    "X": X,
    "O": O,
  }
  pString := p.ByName("player")
  player, _ = m[pString]

  return x, y, player
}
