package cmd

import (
	"fmt"
	"log"

	"github.com/UitsHabib/ecommerce-crud/config"
	"github.com/UitsHabib/ecommerce-crud/db"
	"github.com/UitsHabib/ecommerce-crud/repo"
	"github.com/UitsHabib/ecommerce-crud/rest"
	"github.com/UitsHabib/ecommerce-crud/service"
	"github.com/spf13/cobra"
)

var serveRestCmd = &cobra.Command{
	Use:   "serve-rest",
	Short: "start a rest server",
	RunE:  serveRest,
}

func serveRest(cmd *cobra.Command, args []string) error {
	appCnf := config.GetApp()
	dbCnf := config.GetDB()

	fmt.Printf("App config: %+v, db config: %+v\n", appCnf, dbCnf)

	// connect to db
	db, err := db.Connect(dbCnf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
  if err != nil {
    panic(err)
  }

	// create all the repos 
	brandRepo := repo.NewBrandRepo(db)
	ctgryRepo := repo.NewCategoryRepo(db)
	spplrRepo := repo.NewSupplierRepo(db)
	productRepo := repo.NewProductRepo(db)

	svc := service.NewService(brandRepo, ctgryRepo, spplrRepo, productRepo)

	server, err := rest.NewServer(svc, appCnf)
	if err != nil {
		log.Fatal("cannot create the server: ", err)
	}

	err = server.Start(appCnf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}

	return nil
}
