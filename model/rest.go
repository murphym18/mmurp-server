package model

import(
   "time"
   "encoding/json"
   "strconv"
   "strings"
   "appengine"
   "appengine/datastore"
)

type RestfulEntity interface {
   SetId(string)
}

type Restful struct {
   Id string `datastore:"-" json:"id,omitempty`
   Timestamp time.Time `datastore:"" json:"timestap,omitempty"`
   LastModified time.Time `datastore:"" json:"lastModified"`
}

func (self *Restful) SetId(id string) {
   self.Id = id
}

func GetJson(c appengine.Context, q datastore.Query, slice []interface{}) {
   keys, err := q.GetAll(c, slice);
   if err != nil {
       c.Warningf("%s\n", err.Error())
   }

   for i, key := range keys {
      slice[i].(RestfulEntity).SetId(strconv.FormatInt(key.IntID(), 10))
   }

   data, err := json.Marshal(slice); if err != nil {
      c.Warningf("%s\n", err.Error())
   }
   return data
   w.Write(data)
}
