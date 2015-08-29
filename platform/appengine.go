// This adapts google app engine to the app's API.
// +build appengine

package platform

import (
   "time"
   "net/http"
   "reflect"
   "strconv"
   "server/db"
   "server/util"
   "server/model"
   "appengine"
   "appengine/datastore"
)

const limit_query_posts int = 25
const limit_query_comments int = 25

func GetDependencies(r *http.Request) (db.Database, util.Logger) {
   c := appengine.NewContext(r)
   return &gapDatabase{c}, c
}

type gapDatabase struct {
   c appengine.Context
}

func (g *gapDatabase) getKeyString(key *datastore.Key) string {
   if tmp := key.IntID(); tmp != 0 {
      return strconv.FormatInt(tmp, 10)
   } else {
      return key.StringID()
   }
}

func (g* gapDatabase) handleError(err error, msg string) {
   if err != nil {
      g.c.Errorf("%s\n%s\n", msg, err.Error())
   }
}

func (g *gapDatabase) setModelIds(models interface{}, keys []*datastore.Key)  {
   arr := reflect.ValueOf(models)
   for i, key := range keys {
      arr.Index(i).FieldByName("Id").SetString(g.getKeyString(key))
   }
}

// Loads all posts from the database sorted by date (newest first)
func (g *gapDatabase) GetPosts() (res []model.Post) {
   q := datastore.NewQuery("Post").Order("-Timestamp").Limit(limit_query_posts)
   res = make([]model.Post, limit_query_posts)
   keys, err := q.GetAll(g.c, &res)
   g.handleError(err, "GetPosts returned an error.")
   g.setModelIds(res, keys)
   return res[:len(keys)]
}

// Loads all comments for a post sorted by date (newest first)
func (g *gapDatabase) GetComments(postId string) (res []model.Comment) {
   q := datastore.NewQuery("Comment").Ancestor(g.getPostId(postId)).Order("-Timestamp").Limit(limit_query_comments)
   res = make([]model.Comment, limit_query_comments)
   keys, err := q.GetAll(g.c, &res)
   g.handleError(err, "GetPosts returned an error.")
   g.setModelIds(res, keys)
   return res[:len(keys)]
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

func (g *gapDatabase) getPostId(id string) *datastore.Key {
   var num int64
   var kerr error
   var key *datastore.Key
   num, kerr = strconv.ParseInt(id, 10, 64)
   if kerr == nil {
      key = datastore.NewKey(g.c, "Post", "", num, nil)
   } else {
      key = datastore.NewKey(g.c, "Post", id, 0, nil)
   }
   return key
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
   if v == nil {
      v = &model.Post{}
   }
   v.LastModified = time.Now()
   if v.Timestamp.IsZero() {
      v.Timestamp = v.LastModified
   }
   key, err := datastore.Put(g.c, datastore.NewIncompleteKey(g.c, "Post", nil), v)
   if err != nil {
      g.c.Errorf("Error creating a post\n%s\n", err.Error())
      return nil
   }
   v.Id = g.getKeyString(key)
   return v
}

// Creates an empty comment in the database. The argument specifies what post
// is associated with the comment that is created and it must not be nil.
// The comment returned will have a valid ID and zero values for any
// properties omitted in comment argument. if the comment is nil, then a new
// post is created with a valid ID and zero values for all properties.
func (g *gapDatabase) CreateComment(p string, c *model.Comment) *model.Comment {
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
