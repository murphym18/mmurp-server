package db
import (
   "server/model"
   "server/model/post"
   "server/model/comment"
)

// This is the Applications read-only interface to the database layer.
type DataSource interface {
   // Loads all posts from the database sorted by date (newest first)
   GetPosts() []post.Post

   // Loads all comments for a post sorted by date (newest first)
   GetComments(post.Post) []comment.Comment

   // Loads a post by ID. Returns nil if it doesn't exist.
   GetPost(string) *post.Post

   // Loads a comment by ID. Returns nil if it doesn't exist.
   GetComment(string) *comment.Comment
}

// This is the Applications write-only interface to the database layer.
type DataStore interface {
   // Creates a post in the database. The newly created post will have a valid
   // ID and zero values for any properties omitted. If the argument is nil,
   // then a new post is created with a valid ID and zero values for all
   // properties.
   CreatePost(*post.Post) *post.Post

   // Creates an empty comment in the database. The argument specifies what post
   // is associated with the comment that is created and it must not be nil.
   // The comment returned will have a valid ID and zero values for any
   // properties omitted in comment argument. if the comment is nil, then a new
   // post is created with a valid ID and zero values for all properties.
   CreateComment(*post.Post, *comment.Comment) *comment.Comment

   // Saves a post to the database. This is the same as deleting the post,
   // and creating a new post from parameter. Note the the post should have a
   // valid ID.
   SavePost(*post.Post)

   // Saves a comment to the database and behaves like SavePost. Note the the
   //comment should have a valid ID.
   SaveComment(*comment.Comment)

   // Deletes a post permanently. This also deletes its comments.
   DeletePost(*post.Post)

   // Deletes a comment permanently.
   DeleteComment(*comment.Comment)


type Database interface {
   DataSource
   DataStore
}

var providers = make(map[string]provider.DataProvider)

func RegisterProvider(name string, dao provider.DataProvider) Database {
   providers[name] = dao
}

func GetInstance(name string, config interface{}) Database {
   return providers[name].get(config)
}
