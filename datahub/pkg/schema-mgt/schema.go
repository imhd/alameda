package schema_mgt

import (
	"github.com/containers-ai/alameda/datahub/pkg/config"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	Log "github.com/containers-ai/alameda/pkg/utils/log"
	"sync"
)

var (
	scope = Log.RegisterScope("schema-mgt", "schema management", 0)
)

var (
	InfluxConfig  *InternalInflux.Config
	RWLock        *sync.RWMutex
	Schemas       *SchemaMap
	Schemas2Write *SchemaMap
)

func SchemaInit(config *config.Config) error {
	InfluxConfig = config.InfluxDB
	RWLock = new(sync.RWMutex)
	Schemas = NewSchemaMap()
	Schemas2Write = NewSchemaMap()
	return nil
}
