package "app"

import (
   "../db"
   "../util"
)

// This holds major app components like the database abstraction. The fields
// need to be resolved and injected.
type App struct {
   DB: db.Database
   Log: util.Logger
   Res: http.ResponseWriter
   Req: *http.Request
}

// This should behave like the http.HandlerFunc but it includes an App
// argument that provides access to the app's components.
type AppHandlerFunc func(*App)

// This is a place where dependency injection may happen.
// This method must create the App struct and dispatch the request to the
// appropriate AppHandlerFunc.
func (app *AppHandlerFunc) ServeHTTP(ResponseWriter, *Request) {

}
