// This adapts standard libs to the app's API.
// +build !appengine

package platform

import (
   "server/db"
   "server/util"
   "server/model"
   "log"
   "os"
   "fmt"
   "net/http"
)

type LocalLogger struct {
   out *log.Logger
}

func (g *LocalLogger) Debugf(format string, args ...interface{}) {
   g.out.Printf(fmt.Sprintf(format, args...))
}
func (g *LocalLogger) Infof(format string, args ...interface{}) {
   g.out.Printf(fmt.Sprintf(format, args...))
}
func (g *LocalLogger) Warningf(format string, args ...interface{}) {
   g.out.Printf(fmt.Sprintf(format, args...))
}
func (g *LocalLogger) Errorf(format string, args ...interface{}) {
   g.out.Printf(fmt.Sprintf(format, args...))
}
func (g *LocalLogger) Criticalf(format string, args ...interface{}) {
   g.out.Printf(fmt.Sprintf(format, args...))
}

type LocalDatabase struct {
   out *LocalLogger
}

// Loads all posts from the database sorted by date (newest first)
func (g *LocalDatabase) GetPosts() []model.Post {
   return make([]model.Post, 0)
}

// Loads all comments for a post sorted by date (newest first)
func (g *LocalDatabase) GetComments(string) []model.Comment {
   return nil
}

// Loads a post by ID. Returns nil if it doesn't exist.
func (g *LocalDatabase) GetPost(id string) *model.Post {
   return &model.Post{Id: id}
}

// Loads a comment by ID. Returns nil if it doesn't exist.
func (g *LocalDatabase) GetComment(id string) *model.Comment {
   return nil
}

// Creates a post in the database. The newly created post will have a valid
// ID and zero values for any properties omitted. If the argument is nil,
// then a new post is created with a valid ID and zero values for all
// properties.
func (g *LocalDatabase) CreatePost(v *model.Post) *model.Post {
   return nil
}

// Creates an empty comment in the database. The argument specifies what post
// is associated with the comment that is created and it must not be nil.
// The comment returned will have a valid ID and zero values for any
// properties omitted in comment argument. if the comment is nil, then a new
// post is created with a valid ID and zero values for all properties.
func (g *LocalDatabase) CreateComment(p *model.Post, c *model.Comment) *model.Comment {
   return nil
}

// Saves a post to the database. This is the same as deleting the post,
// and creating a new post from parameter. Note the the post should have a
// valid ID.
func (g *LocalDatabase) SavePost(v *model.Post) {

}

// Saves a comment to the database and behaves like SavePost. Note the the
//comment should have a valid ID.
func (g *LocalDatabase) SaveComment(v *model.Comment) {

}

// Deletes a post permanently. This also deletes its comments.
func (g *LocalDatabase) DeletePost(v *model.Post) {

}

// Deletes a comment permanently.
func (g *LocalDatabase) DeleteComment(v *model.Comment) {

}

func GetDependencies(r *http.Request) (db.Database, util.Logger) {
   // c := appengine.NewContext(r)
   l := &LocalLogger{log.New(os.Stdout, "", 0)}
   return &LocalDatabase{l}, l
}
