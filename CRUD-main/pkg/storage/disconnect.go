package storage

import "github.com/SzymekN/Car-rental-app/pkg/producer"

func CloseAll() {

	if conn, err := MysqlConn.GetDBInstance().DB(); err != nil {
		producer.ProduceMessage("CLOSE ALL", "Postgres conn closing error: "+err.Error())
		if err := conn.Close(); err != nil {
			producer.ProduceMessage("CLOSE ALL", "Postgres conn closing error: "+err.Error())
		}
	} else {
		producer.ProduceMessage("CLOSE ALL", "Postgres conn closed")
	}
	// GetCassandraInstance().Close()
	producer.ProduceMessage("CLOSE ALL", "Cassandra conn closed")
	if err := GetRDB().Close(); err != nil {
		producer.ProduceMessage("CLOSE ALL", "Redis conn closing error: "+err.Error())
	} else {
		producer.ProduceMessage("CLOSE ALL", "Redis conn closed")
	}
}
