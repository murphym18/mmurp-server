package model

import (
   "time"
)

type Post interface {
   Model
   Data() *PostData
}

type JsonData struct {
   Id string `datastore:"-" json:"id,omitempty`
   Timestamp time.Time `json:"timestap,omitempty"`
   LastModified time.Time `json:"lastModified,omitempty"`
   URL string `json:"url"`
   Public bool `json:"-"`
   Title string `json:"title,omitempty"`
   Description string `json:"description,omitempty"`
}

func NewPost() *Post {
   post := new(postData)
   post.Timestap = time.TimeNow()
   post.LastModified = post.Timestamp
}
