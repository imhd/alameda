package clusterstatus

import (
	"github.com/containers-ai/alameda/datahub/pkg/config"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewContainerDAO(config config.Config) types.ContainerDAO {
	return influxdb.NewContainerWithConfig(*config.InfluxDB)
}
