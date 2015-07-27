package provider

type DataProvider interface {
   Get(config interface{})(Database, error)
}

type Database struct {
   DataSource
   DataStore
}
