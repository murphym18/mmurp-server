package rest
import(
   "time"
   "encoding/json"
   "strconv"
   "strings"
   "net/http"
   "server/db"
)

type AppFacade struct {
   DB: db.DataSource
}

//
type RestHandlerFunc func(http.ResponseWriter, *http.Request, *AppFacade)

type Restful struct {
   Id string `datastore:"-" json:"id,omitempty`
   Timestamp time.Time `datastore:"" json:"timestap,omitempty"`
   LastModified time.Time `datastore:"" json:"lastModified"`
}
