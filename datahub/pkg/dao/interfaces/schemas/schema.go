package schemas

import (
	"github.com/containers-ai/alameda/datahub/pkg/config"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/schemas/influxdb"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/schemas/types"
)

func NewSchemaDAO(config config.Config) types.SchemaDAO {
	return influxdb.NewSchemaWithConfig(*config.InfluxDB)
}
