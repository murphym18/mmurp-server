package post

import (
   "time"
   "encoding/json"
)

type Post struct {
   Id string `json:"id,omitempty"`
   Author string `json:"author,omitempty"`
   URL string `json:"url,omitempty"`
   Public bool `json:"-"`
   Title string `json:"title,omitempty"`
   Description string `json:"description,omitempty"`
   Timestamp time.Time `json:"timestap,omitempty"`
   LastModified time.Time `json:"lastModified,omitempty"`
}

func (post *Post) Touch() {
   post.LastModified = time.Now()
   if post.Timestamp.IsZero() {
      post.Timestamp = post.LastModified
   }
}

func (v *Post) ExportJSON() []byte {
   b, _ := json.Marshal(v)
   return b
}

func (v *Post) ImportJSON(b []byte) {
   tmp := &Post{}
   _ = json.Unmarshal(b, tmp)
   v = tmp
}

func (v *Post) PatchJSON(b []byte) {
   _ = json.Unmarshal(b, v)
}
