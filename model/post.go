package model

import (
   "time"
)

type Post struct {
   Id string `json:"id,omitempty`
   Timestamp time.Time `json:"timestap,omitempty"`
   LastModified time.Time `json:"lastModified,omitempty"`
   URL string `json:"url"`
   Public bool `json:"-"`
   Title string `json:"title,omitempty"`
   Description string `json:"description,omitempty"`
}
