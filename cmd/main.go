package main

import (
	"fmt"
	"youchoose/configs"
	"youchoose/internal/infra/factory"
	repository "youchoose/internal/infra/repository"
	usecase "youchoose/internal/use_case"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dsn := "host=" + configs.PostgresServer + " user=" + configs.PostgresUser + " password=" + configs.PostgresPassword + " dbname=" + configs.PostgresDb + " port=" + configs.PostgresPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(repository.Choosers{}); err != nil {
		fmt.Println("Erro durante a migração:", err)
		return
	}
	fmt.Println("Migração bem-sucedida!")

	chooserFactory := factory.NewChooserFactory(db)

	chooserFactory.CreateChooser.Execute(usecase.CreateChooserInputDTO{
		ChooserID: "b1c697a4-032f-44d4-b124-b3030ec61462",
		Name:      "Guilherme Amorim",
		Email:     "guilherme.o.a.ufal@gmail.com",
		Password:  "Abc123@",
		City:      "Aracaju",
		State:     "Sergipe",
		Country:   "Brasil",
		Day:       20,
		Month:     10,
		Year:      1986,
		ImageID:   uuid.NewString(),
	})
}
