package "app"

import (
   "server/db"
)

type AppLogger interface {
   Debugf(format string, args ...interface{})
   Infof(format string, args ...interface{})
   Warningf(format string, args ...interface{})
   Errorf(format string, args ...interface{})
}

type AppFacade struct {
   DB: db.Database
   Log: AppLogger
}
