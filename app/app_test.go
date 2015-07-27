package app_test

import (
   "testing"
   "../app"
)

func TestApp(t *testing.T) {
   var a *app.App = nil
   if a == nil {
      t.Errorf("No Test Written\n")
   }

}
