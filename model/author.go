package model

type Author struct {
   Name  string `datastore:"" json:"name,omitempty"`
   Email  string `datastore:"" json:"-"`
   UserIP string `datastore:"" json:"-"`
   UserAgent string `datastore:"" json:"-"`
}
