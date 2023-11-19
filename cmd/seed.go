package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/UitsHabib/ecommerce-crud/config"
	database "github.com/UitsHabib/ecommerce-crud/db"
	"github.com/UitsHabib/ecommerce-crud/repo"
	"github.com/UitsHabib/ecommerce-crud/service"
	"github.com/UitsHabib/ecommerce-crud/util"
	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seed",
	Short: "seeds database initally",
	RunE:  seed,
}

func seed(cmd *cobra.Command, args []string) error {
	dbCnf := config.GetDB()

	// connect to db
	db, err := database.Connect(dbCnf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
  if err != nil {
    panic(err)
  }

	log.Println("---------------Connected to database---------------")

	// create the brands table if it doesn't exist
	_, err = db.Exec(database.DbSchema)
	if err != nil {
		log.Fatal("can not create table: ", err)
	}

	// initialize the brand repo
	brandRepo := repo.NewBrandRepo(db)

	// create a new brand
	newBrand := &service.Brand{Name: "Lenovo", StatusID: 1, CreatedAt: util.GetCurrentTimestamp()}
	_, err = brandRepo.Add(context.Background(), newBrand)
	if err != nil {
		log.Fatal("can not create item: ", err)
	}

	fmt.Println("seed completed successfully")

	return nil
}
