package model_test

import (
   "encoding/json"
   "fmt"
   "testing"
   "../model"
)

func TestCommentTouch(t *testing.T) {
   p := &model.Comment{}
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

func TestCommentExportJSON(t *testing.T) {
   original := &model.Comment{Comment: "Test Post", Author: "Me"}
   b := original.ExportJSON()
   m := make(map[string]interface{})

   if err := json.Unmarshal(b, &m); err != nil {
      t.Errorf("Timestamp didn't change\n")
      t.Fail()
   }

   if m["comment"] != "Test Post" || m["author"] != "Me" {
      t.Errorf("Expected '%s' and found '%s'\n", "Test Post", m["comment"])
      t.Errorf("Expected '%s' and found '%s'\n", "Me", m["author"])
      t.Fail()
   }
}

func TestCommentImportJSON(t *testing.T) {
   p1 := &model.Comment{Comment: "Test Post", Author: "Me"}
   p2 := &model.Comment{Comment: "something", Author: "Not Me", UserAgent: "/something"}
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

func TestCommentPatchJSON(t *testing.T) {
   original := &model.Comment{Comment: "Test Post", Author: "Me"}
   result := &model.Comment{Comment: "Test Post", Author: "Me", Id: "something"}
   patch := &model.Comment{Id: "something"}
   b := patch.ExportJSON()
   original.PatchJSON(b)
   if fmt.Sprintf("%s", original) != fmt.Sprintf("%s", result) {
      t.Errorf("'%s' should equal '%s'\n", original, result)
      t.Fail()
   }
}
