package main

import (
	"fmt"
	"mime/multipart"
	"os"
	"youchoose/configs"
	"youchoose/internal/infra/factory"
	repository "youchoose/internal/infra/repository"
	usecase "youchoose/internal/use_case"

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

	if err := db.AutoMigrate(
		repository.Choosers{},
		repository.Images{},
		repository.Actors{},
		repository.Directors{},
		repository.Genres{},
		repository.Writers{},
		repository.MovieActors{},
		repository.MovieDirectors{},
		repository.MovieGenres{},
		repository.MovieWriters{},
		repository.Movies{},
		repository.IMDBs{},
	); err != nil {
		fmt.Println("Erro durante a migração:", err)
		return
	}
	fmt.Println("Migração bem-sucedida!")

	// chooserFactory := factory.NewChooserFactory(db)

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

	// fmt.Println()

	// k, l := chooserFactory.DeactivateChooser.Execute(usecase.DeactivateChooserInputDTO{
	// 	ChooserID:             "c4ad0428-13e2-47bc-bf0f-22939694962f",
	// 	ChooserIDToDeactivate: "721b8eee-9586-4771-bf50-0543d8bfbacc",
	// })
	// if len(l.ProblemDetails) > 0 {
	// 	fmt.Println(l.ProblemDetails)
	// } else {
	// 	fmt.Println(k)
	// }

	file1, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	fileStat1, _ := file1.Stat()

	file2, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	fileStat2, _ := file2.Stat()

	file3, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	fileStat3, _ := file3.Stat()

	file4, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	fileStat4, _ := file4.Stat()

	file5, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	fileStat5, _ := file5.Stat()

	file6, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	fileStat6, _ := file6.Stat()

	file7, _ := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	fileStat7, _ := file7.Stat()

	movieFactory := factory.NewMovieFactory(db)

	m, n := movieFactory.CreateMovie.Execute(usecase.CreateMovieInputDTO{
		ChooserID:   "721b8eee-9586-4771-bf50-0543d8bfbacc",
		Title:       "Mirai",
		CountryName: "Japan",
		Flag:        "🇯🇵",
		ReleaseYear: 2018,
		ImageFile:   file1,
		ImageHandler: &multipart.FileHeader{
			Filename: file1.Name(),
			Size:     fileStat1.Size(),
		},
		Genres: []usecase.GenreDTO{
			{
				GenreID:   "",
				Name:      "Animação",
				ImageFile: file2,
				ImageHandler: &multipart.FileHeader{
					Filename: file2.Name(),
					Size:     fileStat2.Size(),
				},
			},
			{
				GenreID:   "",
				Name:      "Aventura",
				ImageFile: file3,
				ImageHandler: &multipart.FileHeader{
					Filename: file3.Name(),
					Size:     fileStat3.Size(),
				},
			},
			{
				GenreID:   "",
				Name:      "Drama",
				ImageFile: file4,
				ImageHandler: &multipart.FileHeader{
					Filename: file4.Name(),
					Size:     fileStat4.Size(),
				},
			},
		},
		Directors: []usecase.DirectorDTO{
			{
				DirectorID:  "",
				Name:        "Mamoru Hosoda",
				Day:         19,
				Month:       9,
				Year:        1967,
				CountryName: "Japan",
				Flag:        "🇯🇵",
				ImageFile:   file5,
				ImageHandler: &multipart.FileHeader{
					Filename: file5.Name(),
					Size:     fileStat5.Size(),
				},
			},
		},
		Actors: []usecase.ActorDTO{
			{
				ActorID:     "",
				Name:        "John Cho",
				Day:         16,
				Month:       6,
				Year:        1972,
				CountryName: "Korea, Republic of",
				Flag:        "🇰🇷",
				ImageFile:   file6,
				ImageHandler: &multipart.FileHeader{
					Filename: file6.Name(),
					Size:     fileStat6.Size(),
				},
			},
		},
		Writers: []usecase.WriterDTO{
			{
				WriterID:    "",
				Name:        "Mamoru Hosoda",
				Day:         19,
				Month:       9,
				Year:        1967,
				CountryName: "Japan",
				Flag:        "🇯🇵",
				ImageFile:   file7,
				ImageHandler: &multipart.FileHeader{
					Filename: file7.Name(),
					Size:     fileStat7.Size(),
				},
			},
		},
	})
	if len(n.ProblemDetails) > 0 {
		fmt.Println(n.ProblemDetails)
	} else {
		fmt.Println(m)
	}
}
