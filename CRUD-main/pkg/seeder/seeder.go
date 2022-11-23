package seeder

import (
	"fmt"
	"log"
	"reflect"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
	"github.com/gocql/gocql"

	"gorm.io/gorm"
)

type SeedPG struct {
	Name  string
	Model model.DataModel
	Run   func(*gorm.DB, string) error
}
type SeedCASS struct {
	Name string
	Run  func(*gocql.Session, string) error
}

func AllPG() []SeedPG {
	return []SeedPG{
		{
			Name:  "users",
			Model: &model.User{},
			Run: func(db *gorm.DB, tableName string) error {
				var err error

				for _, user := range Users {
					err = db.Table(tableName).Create(&user).Error
					if err != nil {
						return err
					}
				}

				return nil
			},
		},
		{
			Name:  "operators",
			Model: &model.Operator{},
			Run: func(db *gorm.DB, tableName string) error {
				var err error

				for _, operator := range Operators {
					err = db.Table(tableName).Create(&operator).Error
					if err != nil {
						return err
					}
				}

				return nil
			},
		},
	}
}

// func AllCASS() []SeedCASS {
// 	return []SeedCASS{
// 		{
// 			Name: "Create users",
// 			Run: func(cas *gocql.Session) error {
// 				var err error

// 				for _, u := range Users {
// 					err = cas.Query(`Insert into userapi.users(id, firstname,lastname, age) values (?,?,?,?)`, u.Id, u.Firstname, u.Lastname, u.Age).Exec()
// 					if err != nil {
// 						return err
// 					}
// 				}

// 				return nil
// 			},
// 		},
// 		{
// 			Name: "Create operators",
// 			Run: func(cas *gocql.Session) error {
// 				var err error

// 				for _, o := range Operators {
// 					pwd, _ := auth.GeneratehashPassword(o.Password)
// 					err = cas.Query(`Insert into userapi.operators(username, email,password, role) values (?,?,?,?)`, o.Username, o.Email, pwd, o.Role).Exec()
// 					if err != nil {
// 						return err
// 					}
// 				}

// 				return nil
// 			},
// 		},
// 	}
// }

func CreateAndSeed() {
	CreateAndSeedPG(storage.MysqlConn.GetDBInstance())
	// CreateAndSeedCASS(storage.GetCassandraInstance())
}

func dropTable(db *gorm.DB, tableName string) {

	tableExists := db.Migrator().HasTable(tableName)
	if tableExists {

		err := db.Migrator().DropTable(tableName).Error

		if err != nil {
			log.Fatalf("PG: Dropping table users failed with error: %s", err)
		} else {
			fmt.Printf("PG: SUCCESFULLY dropped table '%s'!", tableName)
		}

	}
}

func createTable(db *gorm.DB, tableName string, dataModel model.DataModel) {

	var err error
	// err = db.Table(tableName).AutoMigrate(&model.User).Error

	fmt.Println(reflect.TypeOf(dataModel))

	switch dataModel.(type) {
	case *model.User:
		err = db.Table(tableName).AutoMigrate(&model.User{})
	case *model.Operator:
		err = db.Table(tableName).AutoMigrate(&model.Operator{})
	default:
		log.Fatal("ERROR COULD NOT CREATE")
	}

	if err != nil {
		log.Fatalf("PG: Creating table users failed with error: %s", err)
	} else {
		fmt.Printf("PG: SUCCESFULLY created table '%s'!", tableName)
	}

}

func CreateAndSeedPG(db *gorm.DB) {

	for _, seed := range AllPG() {
		dropTable(db, seed.Name)
		createTable(db, seed.Name, seed.Model)
		if err := seed.Run(db, seed.Name); err != nil {
			log.Fatalf("PG: Running seed '%s', failed with error: %s", seed.Name, err)
		}
		fmt.Printf("PG: SUCCESFULLY seeded table database '%s'!", seed.Name)
	}

}

// func CreateAndSeedCASS(cas *gocql.Session) {

// 	err := cas.Query("CREATE KEYSPACE IF NOT EXISTS userapi WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };").Exec()
// 	if err != nil {
// 		log.Fatalf("CASS: Creating table users failed with error: %s", err)
// 	} else {
// 		fmt.Println("CASS: SUCCESFULLY dropped table users!")
// 	}

// 	err = cas.Query(`CREATE TABLE IF NOT EXISTS userapi.users (
// 		id int PRIMARY KEY,
// 		firstname text ,
// 		lastname text,
// 		age int
// 		);`).Exec()

// 	if err != nil {
// 		log.Fatalf("CAS: Creating table users failed with error: %s", err)
// 	} else {
// 		fmt.Println("CAS: SUCCESFULLY created table users!")
// 	}

// 	err = cas.Query(`CREATE TABLE IF NOT EXISTS userapi.operators (
// 		username text PRIMARY KEY,
// 		email text,
// 		password text,
// 		role text
// 		);`).Exec()

// 	if err != nil {
// 		log.Fatalf("CAS: Creating table operators failed with error: %s", err)
// 	} else {
// 		fmt.Println("CAS: SUCCESFULLY created table operators!")
// 	}

// 	for _, seed := range AllCASS() {
// 		if err := seed.Run(cas); err != nil {
// 			log.Fatalf("CAS: Running seed '%s', failed with error: %s", seed.Name, err)
// 		}
// 	}
// 	fmt.Println("CAS: SUCCESFULLY seeded table database!")

// }
