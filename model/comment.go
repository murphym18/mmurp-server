package model

import (
   "time"
   "encoding/json"
)

type Comment struct {
   Id string `appengine:"-" json:"id"`
   PostId string `json:"-"`
   Author string `json:"author,omitempty"`
   Comment string `json:"comment,omitempty"`
   UserIP string `json:"-,omitempty"`
   UserAgent string `json:"-,omitempty"`
   Timestamp time.Time `json:"timestap,omitempty"`
   LastModified time.Time `json:"lastModified,omitempty"`
}

func (v *Comment) Touch() {
   v.LastModified = time.Now()
   if v.Timestamp.IsZero() {
      v.Timestamp = v.LastModified
   }
}

func (v *Comment) ExportJSON() []byte {
   b, _ := json.Marshal(v)
   return b
}

func (v *Comment) ImportJSON(b []byte) {
   tmp := &Comment{}
   _ = json.Unmarshal(b, tmp)
   v = tmp
}

func (v *Comment) PatchJSON(b []byte) {
   _ = json.Unmarshal(b, v)
}
