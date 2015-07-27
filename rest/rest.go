package rest

import(
   "time"
   "encoding/json"
   "strconv"
   "strings"
   "net/http"
   "../app"
   "../db"
)

type RestHandler interface {
   InjectDB(db.Database)
   InjectLogger(app.Log)
   ServeHTTP(ResponseWriter, *Request)
}

func init() {
      http.HandleFunc("/api/", )
      http.HandleFunc("/posts/sign", sign)
}

func DispatchApiRequest(ResponseWriter, *Request){

}
