package services

import (
   "github.com/go-martini/martini"
   "net/http"
   "server/platform"
   "server/db"
   . "server/util"
   . "server/model"
   _ "fmt"
   "encoding/json"
)

// var m *martini.Martini

type HttpFunc func(res http.ResponseWriter, req *http.Request)
func HelloWorld() string {
return "Hello world!"
}

func init() {
   m := martini.Classic()
   m.Use(PlatformDependencies)
   m.Get("/posts/:id", func(res http.ResponseWriter, d db.Database, params martini.Params, log Logger){
      log.Debugf("test")
      res.Header().Add("content-type", "application/json")
      p := d.GetPost(string(params["id"]))
      o := json.NewEncoder(res)
      o.Encode(p)
   })
   m.Get("/posts/", func(res http.ResponseWriter, d db.Database) {
      res.Header().Add("content-type", "application/json")
      p := d.GetPosts()
      o := json.NewEncoder(res)
      o.Encode(p)
   })

   m.Post("/posts/", func(rw http.ResponseWriter, req *http.Request, log Logger, d db.Database) {
       decoder := json.NewDecoder(req.Body)
       var t Post
       err := decoder.Decode(&t)
       if err != nil {
           log.Errorf("bad post request")
       } else {
          d.CreatePost(&t)
          log.Infof("Created a post %s\n", t.Title)
       }
   })


   //

   // m.Get("/test", func(res http.ResponseWriter) {
   //    fmt.Fprintf(res, "Hello, \n")
   // })
   http.Handle("/", m)
   // m.Group("/posts", func(r martini.Router) {
   //     r.Get("/:id", GetPost)
   //     r.Post("/new", NewBook)
   //     r.Put("/update/:id", UpdateBook)
   //     r.Delete("/delete/:id", DeleteBook)
   // }, MyMiddleware1, MyMiddleware2)
}

func PlatformDependencies(c martini.Context, req *http.Request) {
   db, log := platform.GetDependencies(req)
   c.Map(db)
   c.Map(log)

}

func GetPostfunc(res http.ResponseWriter, req *http.Request) *Post {
   // db, log := s.Dependencies()
   // log.Debugf("Getting post by id %s\n", id)
   // return *db.GetPost(id)
   return nil
}

func CreatePost(p *Post) {

}

func ListPosts() []Post {
   // db, log := s.Dependencies()

   // log.Debugf("Getting list of posts")
   // return db.GetPosts()
   return make([]Post,0)
}
