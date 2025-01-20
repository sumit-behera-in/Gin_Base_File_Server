package db

func (c *Client) initClusters() error {
	_, err := c.db.Exec(query_initClusters)
	return err
}

func (c *Client) GetClusterInfoById(clusterId int64) (*ClusterInfo, error) {

	var info ClusterInfo
	if err := c.db.Get(&info, query_getClusterInfoById, clusterId); err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *Client) AddCluster(clusterInfo ClusterInfo) error {
	_, err := c.db.Exec(query_addCluster,
		clusterInfo.ClusterName,
		clusterInfo.ClusterURL,
		clusterInfo.UserCount,
		clusterInfo.TotalStorage,
		clusterInfo.TotalUsed,
	)
	return err
}

func (c *Client) CountTotalClusters() (int, error) {
	var rowCount int
	if err := c.db.Get(&rowCount, query_countTotalClusters); err != nil {
		return 0, err
	}
	return rowCount, nil
}

func (c *Client) GetAllClusters() ([]ClusterInfo, error) {
	var clusters []ClusterInfo
	if err := c.db.Select(&clusters, query_getAllCLusters); err != nil {
		return nil, err
	}
	return clusters, nil
}
