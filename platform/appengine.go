// This adapts google app engine to the app's API.
// +build appengine

package platform

import (
   "net/http"
   "strconv"
   "server/db"
   "server/util"
   "server/model"
   "appengine"
   "appengine/datastore"
)

func GetDependencies(r *http.Request) (db.Database, util.Logger) {
   c := appengine.NewContext(r)
   return &gapDatabase{c}, c
}

type gapDatabase struct {
   c appengine.Context
}

// Loads all posts from the database sorted by date (newest first)
func (g *gapDatabase) GetPosts() []model.Post {
   q := datastore.NewQuery("Post").Order("-Timestamp").Limit(50)
   buf := make([]model.Post, 50)
   keys, err := q.GetAll(g.c, &buf)
   if err != nil {
      g.c.Errorf("Error loading all posts\n%s\n", err.Error())
      return make([]model.Post, 0)
   }
   for i, key := range keys {
      if tmp := key.IntID(); tmp != 0 {
         buf[i].Id = strconv.FormatInt(tmp, 10)
      } else {
         buf[i].Id = key.StringID()
      }
   }
   if len(keys) > 0 {
      return buf
   }
   return make([]model.Post, 0)
}

// Loads all comments for a post sorted by date (newest first)
func (g *gapDatabase) GetComments(v *model.Post) []model.Comment {
   q := datastore.NewQuery("Comment").Order("-Timestamp").Limit(50)
   buf := make([]model.Post, 50)
   keys, err := q.GetAll(g.c, &buf)
   if err != nil {
      g.c.Errorf("Error loading all posts\n%s\n", err.Error())
      return make([]model.Comment, 0)
   }
   if len(keys) < 1 {
      return make([]model.Comment, 0)
   }
   for i, key := range keys {
      if tmp := key.IntID(); tmp != 0 {
         buf[i].Id = strconv.FormatInt(tmp, 10)
      } else {
         buf[i].Id = key.StringID()
      }
   }
   return nil
}

// Loads a post by ID. Returns nil if it doesn't exist.
func (g *gapDatabase) GetPost(id string) *model.Post {
   var num int64
   var kerr error
   var key *datastore.Key
   num, kerr = strconv.ParseInt(id, 10, 64)


   if kerr == nil {
      key = datastore.NewKey(g.c, "Post", "", num, nil)
   } else {
      key = datastore.NewKey(g.c, "Post", id, 0, nil)
   }

   post := &model.Post{}
   if err := datastore.Get(g.c, key, post); err != nil {
      g.c.Errorf("Error loading posts with ID '%s'\n%s\n", id, err.Error())
      return nil
   }
   return post
}

// Loads a comment by ID. Returns nil if it doesn't exist.
func (g *gapDatabase) GetComment(id string) *model.Comment {
   return nil
}

// Creates a post in the database. The newly created post will have a valid
// ID and zero values for any properties omitted. If the argument is nil,
// then a new post is created with a valid ID and zero values for all
// properties.
func (g *gapDatabase) CreatePost(v *model.Post) *model.Post {
   return nil
}

// Creates an empty comment in the database. The argument specifies what post
// is associated with the comment that is created and it must not be nil.
// The comment returned will have a valid ID and zero values for any
// properties omitted in comment argument. if the comment is nil, then a new
// post is created with a valid ID and zero values for all properties.
func (g *gapDatabase) CreateComment(p *model.Post, c *model.Comment) *model.Comment {
   return nil
}

// Saves a post to the database. This is the same as deleting the post,
// and creating a new post from parameter. Note the the post should have a
// valid ID.
func (g *gapDatabase) SavePost(v *model.Post) {

}

// Saves a comment to the database and behaves like SavePost. Note the the
//comment should have a valid ID.
func (g *gapDatabase) SaveComment(v *model.Comment) {

}

// Deletes a post permanently. This also deletes its comments.
func (g *gapDatabase) DeletePost(v *model.Post) {

}

// Deletes a comment permanently.
func (g *gapDatabase) DeleteComment(v *model.Comment) {

}
