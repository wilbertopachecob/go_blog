package auto

import (
	"log"
	"wilbertopachecob/go_blog/cmd/web/db"
	"wilbertopachecob/go_blog/cmd/web/models"
	"wilbertopachecob/go_blog/cmd/web/utils/console"
)

func Load() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Debug().Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	err = database.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		err = database.Debug().Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
		console.PrettyPrint(user)
	}
}
