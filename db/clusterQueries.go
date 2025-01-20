package db

var (
	query_initClusters = `CREATE TABLE IF NOT EXISTS clusters ( 
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		url VARCHAR(100) UNIQUE NOT NULL,
		user_count SMALLINT,
		total_storage DOUBLE PRECISION,
		total_used DOUBLE PRECISION,
		created_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)`

	query_getClusterInfoById = `SELECT name,url,user_count,total_storage,total_used
								From clusters where id = $1`

	query_addCluster = `INSERT INTO clusters (name,url,user_count,total_storage,total_used) VALUES ($1,$2,$3,$4,$5)`

	query_countTotalClusters = `SELECT COUNT(*) FROM clusters`

	query_getAllCLusters = ` SELECT name,url,user_count,total_storage,total_used FROM clusters`
)
