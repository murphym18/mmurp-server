package db
import (
   "server/model"
)

// This is the Applications read-only interface to the database layer.
type DataSource interface {
   // Loads all posts from the database sorted by date (newest first)
   GetPosts() []model.Post

   // Loads all comments for a post sorted by date (newest first)
   GetComments(model.Post) []model.Comment

   // Loads a post by ID. Returns nil if it doesn't exist.
   GetPost(string) *model.Post

   // Loads a comment by ID. Returns nil if it doesn't exist.
   GetComment(string) *model.Comment
}

// This is the Applications write-only interface to the database layer.
type DataStore interface {
   // Creates a post in the database. The newly created post will have a valid
   // ID and zero values for any properties omitted. If the argument is nil,
   // then a new post is created with a valid ID and zero values for all
   // properties.
   CreatePost(*model.Post) *model.Post

   // Creates an empty comment in the database. The argument specifies what post
   // is associated with the comment that is created and it must not be nil.
   // The comment returned will have a valid ID and zero values for any
   // properties omitted in comment argument. if the comment is nil, then a new
   // post is created with a valid ID and zero values for all properties.
   CreateComment(*model.Post, *model.Comment) *model.Comment

   // Saves a post to the database. This is the same as deleting the post,
   // and creating a new post from the parameter argument.
   SavePost(*mode.Post)

   // Saves a comment to the database (behavior like SavePost).
   SaveComment(*model.Comment)

   // Deletes a post permanently. This also deletes its comments.
   DeletePost(*model.Post)

   // Deletes a comment permanently.
   DeleteComment(*model.Comment)
}