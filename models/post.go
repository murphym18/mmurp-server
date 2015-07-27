
type Post struct {
   rest.Restful
   rest.Keywords
   URL string `json:"url"`
   Public bool `json:"public,omitempty"`
   Title string `json:"title,omitempty"`
   Description string `json:"description,omitempty"`
}

type Keywords struct {
   JsonForm []string `datastore:"-" json:"keywords,omitempty`
   Keywords string `datastore:"keywords" json:"-"`
}

func (words *Keywords) Load(c <-chan datastore.Property) error {
   if err := datastore.LoadStruct(words, c); err != nil {
      return err
   }
   words.JsonForm = strings.Fields(words.Keywords)
   return nil
}

func (words *Keywords) Save(c chan<- datastore.Property) error {
    defer close(c)
    c <- datastore.Property{
        Name:  "DataForm",
        Value: strings.Join(words.JsonForm, " "),
    }
    return nil
}