package app

import (
   "../db"
   "../util"
   "../app/gap"
   "net/http"
   //"github.com/gorilla/pat"
)
// This holds major app components like the database abstraction. The fields
// need to be resolved and injected.
type App struct {
   Req *http.Request
   Res http.ResponseWriter
   DB db.Database
   Log util.Logger
}

// This should behave like the http.HandlerFunc but it includes an App
// argument that provides access to the app's components.
type AppHandlerFunc func(*App)

// This is a place where dependency injection may happen.
// This method must create the App struct and dispatch the request to the
// appropriate AppHandlerFunc.
func (self AppHandlerFunc) ServeHTTP(res http.ResponseWriter, req *http.Request) {
   gapDB, gapLog := gap.GetDependencies(req)
   appCtx := &App{req, res, gapDB, gapLog}
   self(appCtx)
}
