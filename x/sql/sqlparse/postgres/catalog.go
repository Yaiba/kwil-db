package postgres

import (
	"kwil/x/sql/catalog"
)

func NewCatalog() *catalog.Catalog {
	c := catalog.New("public")
	c.Schemas = append(c.Schemas, pgTemp())
	c.Schemas = append(c.Schemas, genPGCatalog())
	c.Schemas = append(c.Schemas, genInformationSchema())
	c.SearchPath = []string{"pg_catalog"}
	c.LoadExtension = loadExtension
	return c
}

// The generated pg_catalog is very slow to compare because it has so
// many entries. For testing, don't include it.
func newTestCatalog() *catalog.Catalog {
	c := catalog.New("public")
	c.Schemas = append(c.Schemas, pgTemp())
	c.LoadExtension = loadExtension
	return c
}