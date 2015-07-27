package db_test

import (
   "testing"
   "server/db"
)

func TestDatabase(t *testing.T) {
   t.Errorf("This should end the test", true, false)
   DB := db.GetInstance("")
}
