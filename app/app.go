package app

import (
   "net/http"
   ws "code.google.com/p/gorest"
   "server/services"
)

func init() {
   ws.RegisterService(new(services.PostService))
   http.Handle("/", ws.Handle())
}
