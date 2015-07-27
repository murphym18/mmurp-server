package model

import (
   "time"
)

type Comment struct {
   Id string `json:"id"`
   PostId string `json:"-"`
   Author string `json:"author,omitempty"`
   Comment string `json:"comment,omitempty"`
   UserIP string `json:"-,omitempty"`
   UserAgent string `json:"-,omitempty"`
   Timestamp time.Time `json:"timestap,omitempty"`
   LastModified time.Time `json:"lastModified,omitempty"`
}

func (self *Post) Touch() {
   self.LastModified = time.TimeNow()
   if self.Timestap.IsZero() {
      self.Timestap = self.LastModified
   }
}
