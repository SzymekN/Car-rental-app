package storage

import "github.com/SzymekN/Car-rental-app/pkg/producer"

func Close(d DBConnector) {

	switch c := d.(type) {
	case *MysqlConnect:
		if conn, err := c.GetDBInstance().DB(); err != nil {
			producer.ProduceMessage("CLOSE ALL", "MYSQL conn closing error: "+err.Error())
			if err := conn.Close(); err != nil {
				producer.ProduceMessage("CLOSE ALL", "MYSQL conn closing error: "+err.Error())
			}
		} else {
			producer.ProduceMessage("CLOSE ALL", "MYSQL conn closed")
		}
	case *RedisConnect:
		if err := c.GetRDB().Close(); err != nil {
			producer.ProduceMessage("CLOSE ALL", "Redis conn closing error: "+err.Error())
		} else {
			producer.ProduceMessage("CLOSE ALL", "Redis conn closed")
		}
	}

}
