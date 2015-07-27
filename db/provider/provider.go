package provider

import "server/db"

type DataProvider interface {
   Get(config interface{})(db.Database, error)
}
