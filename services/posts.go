package services

import (
   "code.google.com/p/gorest"
   "server/platform"
   . "server/db"
   . "server/util"
   . "server/model"
)

type PostService struct{
    gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`
    listPosts gorest.EndPoint `method:"GET" path:"/posts" output:"[]Post"`
    getPost gorest.EndPoint `method:"GET" path:"/posts/{id:string}" output:"Post"`
   //  gorest.EndPoint `method:"POST" path:"/posts"`
   //  gorest.EndPoint `method:"PUT" path:"/posts/{id:string}"`
   //  gorest.EndPoint `method:"DELETE" path:"/posts/{id:string}"`
}

func (s PostService) Dependencies() (Database, Logger) {
   return platform.GetDependencies(s.Context.Request())
}

func (s PostService) GetPost(id string) Post {
   db, log := s.Dependencies()
   log.Debugf("Getting post by id %s\n", id)
   return *db.GetPost(id)
}

func (s PostService) ListPosts() []Post {
   db, log := s.Dependencies()
   log.Debugf("Load list of posts\n")
   return db.GetPosts()
}
