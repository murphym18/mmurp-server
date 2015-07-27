// This adapts google app engine to the app's API.
package gap

import (
   "../../db"
   "../../util"

   "net/http"
   "../../model"

)

// "log"
// "io"
// "appengine"
// "appengine/datastore"

type GapDatabase struct {
   // c: appengine.Context
}

// Loads all posts from the database sorted by date (newest first)
func (g *GapDatabase) GetPosts() []model.Post {
   return nil
}

// Loads all comments for a post sorted by date (newest first)
func (g *GapDatabase) GetComments(v model.Post) []model.Comment {
   return nil
}

// Loads a post by ID. Returns nil if it doesn't exist.
func (g *GapDatabase) GetPost(id string) *model.Post {
   return nil
}

// Loads a comment by ID. Returns nil if it doesn't exist.
func (g *GapDatabase) GetComment(id string) *model.Comment {
   return nil
}

// Creates a post in the database. The newly created post will have a valid
// ID and zero values for any properties omitted. If the argument is nil,
// then a new post is created with a valid ID and zero values for all
// properties.
func (g *GapDatabase) CreatePost(v *model.Post) *model.Post {
   return nil
}

// Creates an empty comment in the database. The argument specifies what post
// is associated with the comment that is created and it must not be nil.
// The comment returned will have a valid ID and zero values for any
// properties omitted in comment argument. if the comment is nil, then a new
// post is created with a valid ID and zero values for all properties.
func (g *GapDatabase) CreateComment(p *model.Post, c *model.Comment) *model.Comment {
   return nil
}

// Saves a post to the database. This is the same as deleting the post,
// and creating a new post from parameter. Note the the post should have a
// valid ID.
func (g *GapDatabase) SavePost(v *model.Post) {

}

// Saves a comment to the database and behaves like SavePost. Note the the
//comment should have a valid ID.
func (g *GapDatabase) SaveComment(v *model.Comment) {

}

// Deletes a post permanently. This also deletes its comments.
func (g *GapDatabase) DeletePost(v *model.Post) {

}

// Deletes a comment permanently.
func (g *GapDatabase) DeleteComment(v *model.Comment) {

}

func GetDependencies(r *http.Request) (db.Database, util.Logger) {
   // c := appengine.NewContext(r)
   return new(GapDatabase), nil//log.New(io.Std, "", 0)
}
