package db_test

import (
   "testing"
   "server/db"
)

func TestDatabase(t *testing.T) {
   var testDb db.Database = nil
   if testDb == nil {
      t.Errorf("No Test Written\n")
   }

}
