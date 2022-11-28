package seeder

// import (
// 	"fmt"
// 	"log"
// 	"reflect"

// 	"github.com/SzymekN/Car-rental-app/pkg/model"
// 	"github.com/SzymekN/Car-rental-app/pkg/storage"

// 	"gorm.io/gorm"
// )

// type seed struct {
// 	// name of the seed
// 	Name string
// 	// datatype
// 	Model model.DataModel
// 	// function to run
// 	Run func(*gorm.DB, string) error
// }

// func allSeeds() []seed {

// 	return []seed{
// 		{
// 			Name:  "users",
// 			Model: &model.User{},
// 			Run: func(db *gorm.DB, tableName string) error {
// 				var err error

// 				for _, user := range Users {
// 					err = db.Table(tableName).Create(&user).Error
// 					if err != nil {
// 						return err
// 					}
// 				}

// 				return nil
// 			},
// 		},
// 		{
// 			Name:  "operators",
// 			Model: &model.Operator{},
// 			Run: func(db *gorm.DB, tableName string) error {
// 				var err error

// 				for _, operator := range Operators {
// 					err = db.Table(tableName).Create(&operator).Error
// 					if err != nil {
// 						return err
// 					}
// 				}

// 				return nil
// 			},
// 		},
// 	}
// }

// func dropTable(db *gorm.DB, tableName string) {

// 	tableExists := db.Migrator().HasTable(tableName)
// 	if tableExists {

// 		err := db.Migrator().DropTable(tableName).Error

// 		if err != nil {
// 			log.Fatalf("PG: Dropping table users failed with error: %s", err)
// 		} else {
// 			fmt.Printf("PG: SUCCESFULLY dropped table '%s'!", tableName)
// 		}

// 	}
// }

// func createTable(db *gorm.DB, tableName string, dataModel model.DataModel) {

// 	var err error
// 	// err = db.Table(tableName).AutoMigrate(&model.User).Error

// 	fmt.Println(reflect.TypeOf(dataModel))

// 	switch dataModel.(type) {
// 	case *model.User:
// 		err = db.Table(tableName).AutoMigrate(&model.User{})
// 	case *model.Operator:
// 		err = db.Table(tableName).AutoMigrate(&model.Operator{})
// 	default:
// 		log.Fatal("ERROR COULD NOT CREATE")
// 	}

// 	if err != nil {
// 		log.Fatalf("PG: Creating table users failed with error: %s", err)
// 	} else {
// 		fmt.Printf("PG: SUCCESFULLY created table '%s'!", tableName)
// 	}

// }

// // deletes, creates and populates tables
// func CreateAndSeed() {
// 	db := storage.MysqlConn.GetDBInstance()

// 	// run all defined seeds
// 	for _, seed := range allSeeds() {
// 		dropTable(db, seed.Name)
// 		createTable(db, seed.Name, seed.Model)
// 		if err := seed.Run(db, seed.Name); err != nil {
// 			log.Fatalf("PG: Running seed '%s', failed with error: %s", seed.Name, err)
// 		}
// 		fmt.Printf("PG: SUCCESFULLY seeded table database '%s'!", seed.Name)
// 	}

// }
