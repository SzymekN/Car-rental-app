package storage

import "github.com/SzymekN/CRUD/pkg/producer"

func CloseAll() {

	if err := GetDBInstance().Close(); err != nil {
		producer.ProduceMessage("CLOSE ALL", "Postgres conn closing error: "+err.Error())
	} else {
		producer.ProduceMessage("CLOSE ALL", "Postgres conn closed")
	}
	GetCassandraInstance().Close()
	producer.ProduceMessage("CLOSE ALL", "Cassandra conn closed")
	if err := GetRDB().Close(); err != nil {
		producer.ProduceMessage("CLOSE ALL", "Redis conn closing error: "+err.Error())
	} else {
		producer.ProduceMessage("CLOSE ALL", "Redis conn closed")
	}
}
