package storage

import "github.com/SzymekN/Car-rental-app/pkg/producer"

func Close(d DBConnector, l producer.SystemLogger) {

	switch c := d.(type) {
	case *MysqlConnect:
		if conn, err := c.GetDBInstance().DB(); err != nil {
			l.ProduceMessage("CLOSE ALL", "MYSQL conn closing error: "+err.Error())
			if err := conn.Close(); err != nil {
				l.ProduceMessage("CLOSE ALL", "MYSQL conn closing error: "+err.Error())
			}
		} else {
			l.ProduceMessage("CLOSE ALL", "MYSQL conn closed")
		}
	case *RedisConnect:
		if err := c.GetRDB().Close(); err != nil {
			l.ProduceMessage("CLOSE ALL", "Redis conn closing error: "+err.Error())
		} else {
			l.ProduceMessage("CLOSE ALL", "Redis conn closed")
		}
	}

}
