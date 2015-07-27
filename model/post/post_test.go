package post_test

import (
   "time"
   "encoding/json"
   "testing"
   "server/model/post"
)

func TestTouch(t *testing.T) {
   p := &post.Post{}
   t.Errorf("ensure we start with zero value", true, p.Timestamp.IsZero())
   p.Touch()
   t.Errorf("ensure value is different", false, p.Timestamp.IsZero())
}

func TestExportJSON() {
   original := &post.Post{Title: "Test Post", Author: "Me"}
   b := original.ExportJSON()
   map := make(map[string]interface{})
   err := json.Unmarshal(b, read)
   t.Errorf("error calling json.Unmarshal", nil, err)
   t.Errorf("the exported json is wrong", map["title"], "Test Post")
   t.Errorf("the exported json is wrong", map["author"], "Me")
}
