package model

type Comment struct {
   Restful
   Author
   Comment string `datastore:"" json:"comment,omitempty"`
   UserIP string `datastore:"" json:"-,omitempty"`
   UserAgent string `datastore:"" json:"-,omitempty"`
}
