package keycodes

import (
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	DaoClusterType "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	"github.com/containers-ai/alameda/datahub/pkg/kubernetes/metadata"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
)

func getClustersFromDB(dao DaoClusterType.ClusterDAO) ([]*DaoClusterType.Cluster, error) {
	req := DaoClusterType.ListClustersRequest{}
	clsts, err := dao.ListClusters(&req)
	return clsts, err
}

func getClusterNodesFromDB(dao DaoClusterType.NodeDAO, clusterName string) ([]*DaoClusterType.Node, error) {
	objmeta := make([]*metadata.ObjectMeta, 0)
	objmeta = append(objmeta, &metadata.ObjectMeta{ClusterName: clusterName})
	req := DaoClusterType.ListNodesRequest{ObjectMeta: objmeta}
	ns, err := dao.ListNodes(&req)
	return ns, err
}

func GetAlamedaClusterCPUs(influxCfg *InternalInflux.Config) (int, error) {
	numCPU := 0
	clusterName := ""
	dbCluster := influxdb.NewClusterWithConfig(*influxCfg)
	dbNode := influxdb.NewNodeWithConfig(*influxCfg)

	clusters, err := getClustersFromDB(dbCluster)
	if err != nil {
		scope.Errorf("Failed to get cluster info: %s", err.Error())
		return numCPU, err
	}
	if clusters != nil && len(clusters) > 0 {
		clusterName = clusters[0].ObjectMeta.Name
		scope.Infof("check cluster [%s] CPU info", clusterName)
	}
	if clusterName != "" {
		nodes, err := getClusterNodesFromDB(dbNode, clusterName)
		if err != nil {
			scope.Errorf("Failed to get cluster [%s] node info: %s", clusterName, err.Error())
			return numCPU, err
		} else {
			scope.Infof("Number of cluster node: %d", len(nodes))
			for _, node := range nodes {
				if node != nil {
					scope.Infof("  node: %s, CPU Cores: %d", node.ObjectMeta.Name, node.Capacity.CpuCores)
					numCPU += int(node.Capacity.CpuCores)
				}
			}
		}
	}
	return numCPU, nil
}
