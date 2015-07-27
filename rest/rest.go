package rest

import(
   "time"
   "encoding/json"
   "strconv"
   "strings"
   "net/http"
   "server/app"
   "server/db"
)

// This should behave like the http.HandlerFunc but it includes an AppFacade
// argument that provides access to the app's components.
type RestHandlerFunc func(http.ResponseWriter, *http.Request, *app.AppFacade)
