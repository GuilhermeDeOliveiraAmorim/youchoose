package main

import (
	"fmt"
	"youchoose/configs"
	repository "youchoose/internal/infra/repository"

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

	if err := db.AutoMigrate(repository.Choosers{}, repository.Images{}); err != nil {
		fmt.Println("Erro durante a migração:", err)
		return
	}
	fmt.Println("Migração bem-sucedida!")
  
	//chooserFactory := factory.NewChooserFactory(db)

	// file, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	// fileStat, _ := file.Stat()

	// input := usecase.CreateChooserInputDTO{
	// 	ChooserID: "721b8eee-9586-4771-bf50-0543d8bfbacc",
	// 	Name:      "Guilherme Amorim",
	// 	Email:     "guilherme.a.ufal@bol.com.br",
	// 	Password:  "Abc123@",
	// 	City:      "Aracaju",
	// 	State:     "Sergipe",
	// 	Country:   "Brasil",
	// 	Day:       20,
	// 	Month:     10,
	// 	Year:      1986,
	// 	ImageFile: file,
	// 	ImageHandler: &multipart.FileHeader{
	// 		Filename: file.Name(),
	// 		Size:     fileStat.Size(),
	// 	},
	// }

	// a, b := chooserFactory.CreateChooser.Execute(input)
	// if len(b.ProblemDetails) > 0 {
	// 	fmt.Println(b.ProblemDetails)
	// } else {
	// 	fmt.Println(a)
	// }

	// fmt.Println()

	// c, d := chooserFactory.FindChooserByID.Execute(usecase.GetChooserInputDTO{
	// 	ChooserID:       "721b8eee-9586-4771-bf50-0543d8bfbacc",
	// 	ChooserIDToFind: "c4ad0428-13e2-47bc-bf0f-22939694962f",
	// })
	// if len(d.ProblemDetails) > 0 {
	// 	fmt.Println(d.ProblemDetails)
	// } else {
	// 	fmt.Println(c)
	// }

	// fmt.Println()

	// e, f := chooserFactory.GetChoosers.Execute(usecase.GetChoosersInputDTO{
	// 	ChooserID: "721b8eee-9586-4771-bf50-0543d8bfbacc",
	// })
	// if len(f.ProblemDetails) > 0 {
	// 	fmt.Println(f.ProblemDetails)
	// } else {
	// 	fmt.Println(e)
	// }

	// fmt.Println()

	// g, h := chooserFactory.UpdateChooser.Execute(usecase.UpdateChooserInputDTO{
	// 	ChooserID: "c9f6c34a-a2fa-43df-bc33-7f0b9fb5ecf8",
	// 	Name:      "Novo Nome",
	// 	City:      "Maceió",
	// 	State:     "Alagoas",
	// 	Country:   "Brasil",
	// 	Day:       11,
	// 	Month:     12,
	// 	Year:      1986,
	// 	ImageID:   "",
	// 	ImageFile: file,
	// 	ImageHandler: &multipart.FileHeader{
	// 		Filename: file.Name(),
	// 		Size:     fileStat.Size(),
	// 	},
	// })
	// if len(h.ProblemDetails) > 0 {
	// 	fmt.Println(h.ProblemDetails)
	// } else {
	// 	fmt.Println(g)
	// }

	// fmt.Println()

	// i, j := chooserFactory.UpdateChooser.Execute(usecase.UpdateChooserInputDTO{
	// 	ChooserID:    "c9f6c34a-a2fa-43df-bc33-7f0b9fb5ecf8",
	// 	Name:         "Novo Novo Nome Do Guilherme",
	// 	City:         "Recife",
	// 	State:        "Pernambuco",
	// 	Country:      "Brasil",
	// 	Day:          20,
	// 	Month:        10,
	// 	Year:         1986,
	// 	ImageID:      "3b50916b-bb23-438e-8896-1cbc44273cc4",
	// 	ImageFile:    nil,
	// 	ImageHandler: nil,
	// })
	// if len(j.ProblemDetails) > 0 {
	// 	fmt.Println(j.ProblemDetails)
	// } else {
	// 	fmt.Println(i)
	// }
}
