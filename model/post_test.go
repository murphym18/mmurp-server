package model_test

import (
   "encoding/json"
   "fmt"
   "testing"
   "server/model"
)

func TestTouch(t *testing.T) {
   p := &model.Post{}
   if !p.Timestamp.IsZero() {
      t.Errorf("Didn't start with zero value\n")
      t.Fail()
   }
   p.Touch()
   if p.Timestamp.IsZero() {
      t.Errorf("Timestamp didn't change\n")
      t.Fail()
   }
}

func TestExportJSON(t *testing.T) {
   original := &model.Post{Title: "Test Post", Author: "Me"}
   b := original.ExportJSON()
   m := make(map[string]interface{})

   if err := json.Unmarshal(b, &m); err != nil {
      t.Errorf("Timestamp didn't change\n")
      t.Fail()
   }

   if m["title"] != "Test Post" || m["author"] != "Me" {
      t.Errorf("Expected '%s' and found '%s'\n", "Test Post", m["title"])
      t.Errorf("Expected '%s' and found '%s'\n", "Me", m["author"])
      t.Fail()
   }
}

func TestImportJSON(t *testing.T) {
   p1 := &model.Post{Title: "Test Post", Author: "Me"}
   p2 := &model.Post{Title: "something", Author: "Not Me", URL: "/something"}
   b := p1.ExportJSON()
   if fmt.Sprintf("%s", p1) == fmt.Sprintf("%s", p2) {
      t.Errorf("'%s' should equal '%s'\n", p1, p2)
      t.Fail()
   }
   p2.ImportJSON(b)
   if fmt.Sprintf("%s", p1) == fmt.Sprintf("%s", p2) {
      t.Errorf("'%s' should not equal '%s'\n", p1, p2)
      t.Fail()
   }
}

func TestPatchJSON(t *testing.T) {
   original := &model.Post{Title: "Test Post", Author: "Me"}
   result := &model.Post{Title: "Test Post", Author: "Me", URL: "/something"}
   patch := &model.Post{URL: "/something"}
   b := patch.ExportJSON()
   original.PatchJSON(b)
   if fmt.Sprintf("%s", original) != fmt.Sprintf("%s", result) {
      t.Errorf("'%s' should equal '%s'\n", original, result)
      t.Fail()
   }
}
