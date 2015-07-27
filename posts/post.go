package post

import (
   "html/template"
   "net/http"
   "time"
   "appengine"
   "appengine/datastore"
   "appengine/user"
   "server/models"
)

func init() {
      http.HandleFunc("/posts/", root)
      http.HandleFunc("/posts/sign", sign)
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Topic", "post_comments", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("Comment").Ancestor(guestbookKey(c)).Order("-Timestamp").Limit(10)
        entitySlice := make([]Comment, 0, 10)


}





func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        c.Infof("I am here %s", "sds")

        g := Comment{
             Comment: r.FormValue("content"),
             Restful: rest.Restful{Timestamp: time.Now().UTC()},
        }
        if u := user.Current(c); u != nil {
                //g.Author = u.String()
        }

        // We set the same parent key on every Greeting entity to ensure each Greeting
        // is in the same entity group. Queries across the single entity group
        // will be consistent. However, the write rate to a single entity group
        // should be limited to ~1/second.
        key := datastore.NewIncompleteKey(c, "Comment", guestbookKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/posts", http.StatusFound)
}
