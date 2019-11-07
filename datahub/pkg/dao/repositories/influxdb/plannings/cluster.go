package plannings

import (
	EntityInfluxPlanning "github.com/containers-ai/alameda/datahub/pkg/dao/entities/influxdb/plannings"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	ApiPlannings "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/plannings"
	ApiResources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	//"github.com/golang/protobuf/ptypes"
	//"github.com/golang/protobuf/ptypes/timestamp"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	//"strconv"
	"time"
)

type ClusterRepository struct {
	influxDB *InternalInflux.InfluxClient
}

func NewClusterRepository(influxDBCfg *InternalInflux.Config) *ClusterRepository {
	return &ClusterRepository{
		influxDB: &InternalInflux.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

func (c *ClusterRepository) CreatePlannings(plannings []*ApiPlannings.ClusterPlanning) error {
	points := make([]*InfluxClient.Point, 0)
	for _, planning := range plannings {
		tags := map[string]string{
			EntityInfluxPlanning.ClusterName: planning.GetObjectMeta().GetName(),
		}

		fields := map[string]interface{}{
			EntityInfluxPlanning.ClusterValue: 0,
		}

		pt, err := InfluxClient.NewPoint(string(Cluster), tags, fields, time.Unix(time.Now().UTC().Unix(), 0))
		if err != nil {
			scope.Error(err.Error())
		}

		points = append(points, pt)

	}

	err := c.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.Planning),
	})

	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (c *ClusterRepository) ListPlannings(in *ApiPlannings.ListClusterPlanningsRequest) ([]*ApiPlannings.ClusterPlanning, error) {
	influxdbStatement := InternalInflux.Statement{
		Measurement:    Cluster,
		QueryCondition: DBCommon.BuildQueryConditionV1(in.GetQueryCondition()),
	}

	for _, objMeta := range in.GetObjectMeta() {
		name := objMeta.GetName()

		if name == "" {
			influxdbStatement.WhereClause = ""
			break
		}

		keyList := []string{EntityInfluxPlanning.ClusterName}
		valueList := []string{name}

		tempCondition := influxdbStatement.GenerateCondition(keyList, valueList, "AND")
		influxdbStatement.AppendWhereClauseDirectly("OR", tempCondition)
	}

	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()

	cmd := influxdbStatement.BuildQueryCmd()

	results, err := c.influxDB.QueryDB(cmd, string(RepoInflux.Planning))
	if err != nil {
		return make([]*ApiPlannings.ClusterPlanning, 0), err
	}

	influxdbRows := InternalInflux.PackMap(results)
	plannings := c.getPlanningsFromInfluxRows(influxdbRows)

	return plannings, nil
}

func (c *ClusterRepository) getPlanningsFromInfluxRows(rows []*InternalInflux.InfluxRow) []*ApiPlannings.ClusterPlanning {
	plannings := make([]*ApiPlannings.ClusterPlanning, 0)
	for _, influxdbRow := range rows {
		for _, data := range influxdbRow.Data {
			tempPlanning := &ApiPlannings.ClusterPlanning{
				ObjectMeta: &ApiResources.ObjectMeta{
					Name: data[string(EntityInfluxPlanning.ClusterName)],
				},
			}

			plannings = append(plannings, tempPlanning)
		}
	}

	return plannings
}