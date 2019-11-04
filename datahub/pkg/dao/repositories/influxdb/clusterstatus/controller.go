package clusterstatus

import (
	"fmt"
	EntityInfluxClusterStatus "github.com/containers-ai/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	ApiResources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"strconv"
	"strings"
	"time"
)

type ControllerRepository struct {
	influxDB *InternalInflux.InfluxClient
}

func NewControllerRepository(influxDBCfg *InternalInflux.Config) *ControllerRepository {
	return &ControllerRepository{
		influxDB: &InternalInflux.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

func (c *ControllerRepository) CreateControllers(controllers []*ApiResources.Controller) error {
	points := make([]*InfluxClient.Point, 0)
	for _, controller := range controllers {
		controllerNamespace := controller.GetObjectMeta().GetNamespace()
		controllerName := controller.GetObjectMeta().GetName()
		controllerKind := controller.GetKind().String()
		controllerExecution := controller.GetAlamedaControllerSpec().GetEnableRecommendationExecution()
		controllerPolicy := controller.GetAlamedaControllerSpec().GetPolicy().String()

		ownerNamespace := ""
		ownerName := ""
		ownerKind := ""

		if len(controller.GetOwnerReferences()) > 0 {
			ownerNamespace = controller.GetOwnerReferences()[0].GetObjectMeta().GetNamespace()
			ownerName = controller.GetOwnerReferences()[0].GetObjectMeta().GetName()
			ownerKind = controller.GetOwnerReferences()[0].GetKind().String()
		}

		tags := map[string]string{
			string(EntityInfluxClusterStatus.ControllerNamespace):      controllerNamespace,
			string(EntityInfluxClusterStatus.ControllerName):           controllerName,
			string(EntityInfluxClusterStatus.ControllerOwnerNamespace): ownerNamespace,
			string(EntityInfluxClusterStatus.ControllerOwnerName):      ownerName,
		}

		fields := map[string]interface{}{
			string(EntityInfluxClusterStatus.ControllerKind):            controllerKind,
			string(EntityInfluxClusterStatus.ControllerOwnerKind):       ownerKind,
			string(EntityInfluxClusterStatus.ControllerReplicas):        controller.GetReplicas(),
			string(EntityInfluxClusterStatus.ControllerEnableExecution): strconv.FormatBool(controllerExecution),
			string(EntityInfluxClusterStatus.ControllerPolicy):          controllerPolicy,
			string(EntityInfluxClusterStatus.ControllerSpecReplicas):    controller.GetSpecReplicas(),
		}

		pt, err := InfluxClient.NewPoint(string(Controller), tags, fields, time.Unix(0, 0))
		if err != nil {
			scope.Error(err.Error())
		}
		points = append(points, pt)
	}

	err := c.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.ClusterStatus),
	})

	if err != nil {
		scope.Error(err.Error())
	}

	return nil
}

func (c *ControllerRepository) ListControllers(in *ApiResources.ListControllersRequest) ([]*ApiResources.Controller, error) {
	namespace := in.GetObjectMeta()[0].GetNamespace()
	name := in.GetObjectMeta()[0].GetName()

	whereStr := c.convertQueryCondition(namespace, name)

	influxdbStatement := InternalInflux.Statement{
		Measurement: Controller,
		WhereClause: whereStr,
		GroupByTags: []string{EntityInfluxClusterStatus.ControllerNamespace, EntityInfluxClusterStatus.ControllerName},
	}

	cmd := influxdbStatement.BuildQueryCmd()

	results, err := c.influxDB.QueryDB(cmd, string(RepoInflux.ClusterStatus))
	if err != nil {
		return make([]*ApiResources.Controller, 0), err
	}

	influxdbRows := InternalInflux.PackMap(results)

	controllerList := c.getControllersFromInfluxRows(influxdbRows)
	return controllerList, nil
}

func (c *ControllerRepository) DeleteControllers(in *ApiResources.DeleteControllersRequest) error {
	controllers := in.GetControllers()
	whereStr := ""

	for _, controller := range controllers {
		namespace := controller.GetObjectMeta().GetNamespace()
		name := controller.GetObjectMeta().GetName()
		whereStr += fmt.Sprintf(" (\"name\"='%s' AND \"namespace\"='%s') OR", name, namespace)
	}

	whereStr = strings.TrimSuffix(whereStr, "OR")

	if whereStr != "" {
		whereStr = "WHERE" + whereStr
	}
	cmd := fmt.Sprintf("DROP SERIES FROM %s %s", string(Controller), whereStr)

	_, err := c.influxDB.QueryDB(cmd, string(RepoInflux.ClusterStatus))
	if err != nil {
		return err
	}

	return nil
}

func (c *ControllerRepository) convertQueryCondition(namespace string, name string) string {
	ret := ""

	if namespace != "" {
		ret += fmt.Sprintf("\"namespace\"='%s' ", namespace)
	}

	if name != "" {
		ret += fmt.Sprintf("AND \"name\"='%s' ", name)
	}

	ret = strings.TrimPrefix(ret, "AND")
	if ret != "" {
		ret = "WHERE " + ret
	}
	return ret
}

func (c *ControllerRepository) getControllersFromInfluxRows(rows []*InternalInflux.InfluxRow) []*ApiResources.Controller {
	controllerList := make([]*ApiResources.Controller, 0)
	for _, row := range rows {
		namespace := row.Tags[EntityInfluxClusterStatus.ControllerNamespace]
		name := row.Tags[EntityInfluxClusterStatus.ControllerName]

		tempController := &ApiResources.Controller{
			ObjectMeta: &ApiResources.ObjectMeta{
				Namespace: namespace,
				Name:      name,
			},
			OwnerReferences:       make([]*ApiResources.OwnerReference, 0),
			AlamedaControllerSpec: &ApiResources.AlamedaControllerSpec{},
		}

		ownerReferences := make([]*ApiResources.OwnerReference, 0)
		for _, data := range row.Data {
			ownerNamespace := data[EntityInfluxClusterStatus.ControllerOwnerNamespace]
			ownerName := data[EntityInfluxClusterStatus.ControllerOwnerName]
			tempOwnerKind := data[EntityInfluxClusterStatus.ControllerOwnerKind]
			var ownerKind ApiResources.Kind

			if val, found := ApiResources.Kind_value[tempOwnerKind]; found {
				ownerKind = ApiResources.Kind(val)
			}

			tempOwner := &ApiResources.OwnerReference{
				ObjectMeta: &ApiResources.ObjectMeta{
					Namespace: ownerNamespace,
					Name:      ownerName,
				},
				Kind: ownerKind,
			}

			ownerReferences = append(ownerReferences, tempOwner)

			//------
			tempKind := data[EntityInfluxClusterStatus.ControllerKind]
			var kind ApiResources.Kind
			if val, found := ApiResources.Kind_value[tempKind]; found {
				kind = ApiResources.Kind(val)
			}
			tempController.Kind = kind

			tempReplicas, _ := strconv.ParseInt(data[string(EntityInfluxClusterStatus.ControllerReplicas)], 10, 32)
			tempController.Replicas = int32(tempReplicas)

			enableExecution, _ := strconv.ParseBool(data[EntityInfluxClusterStatus.ControllerEnableExecution])
			tempController.AlamedaControllerSpec.EnableRecommendationExecution = enableExecution

			tempPolicy := data[EntityInfluxClusterStatus.ControllerPolicy]
			var policy ApiResources.RecommendationPolicy
			if val, found := ApiResources.RecommendationPolicy_value[tempPolicy]; found {
				policy = ApiResources.RecommendationPolicy(val)
			}
			tempController.AlamedaControllerSpec.Policy = policy
		}

		tempController.OwnerReferences = ownerReferences
		controllerList = append(controllerList, tempController)
	}

	return controllerList
}
