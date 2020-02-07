package data

import (
	"github.com/containers-ai/alameda/datahub/pkg/config"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/influxdb"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
)

func NewDataDAO(config config.Config) types.DataDAO {
	return influxdb.NewDataWithConfig(*config.InfluxDB)
}
