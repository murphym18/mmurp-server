package provider

import "server/db"

type Provider interface {
   Open(config string) DataProvider, error
}

type DataProvider interface {
   GetDataSource() db.DataSource
   GetDataStore() db.DataStore
}
