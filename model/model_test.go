package model_test

import (
   "testing"
   "../model"
)

func TestCastingToModel(t *testing.T) {
   var m model.Model
   m = &model.Post{Title: "Hello"}
   m = &model.Comment{Comment: "Cool"}
   m.Touch()
}
