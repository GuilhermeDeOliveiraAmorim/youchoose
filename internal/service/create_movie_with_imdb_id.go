package service

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"youchoose/configs"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"

	"github.com/gocolly/colly"
	"github.com/oklog/ulid/v2"
)

type GenreDTO struct {
	GenreID string `json:"genre_id"`
	Name    string `json:"name"`
}

type DirectorDTO struct {
	DirectorID  string `json:"director_id"`
	Name        string `json:"name"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
	CountryName string `json:"country_name"`
	Flag        string `json:"flag"`
	ImageName   string `json:"image_name"`
}

type ActorDTO struct {
	ActorID     string `json:"actor_id"`
	Name        string `json:"name"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
	CountryName string `json:"country_name"`
	Flag        string `json:"flag"`
	ImageName   string `json:"image_name"`
}

type WriterDTO struct {
	WriterID    string `json:"writer_id"`
	Name        string `json:"name"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
	CountryName string `json:"country_name"`
	Flag        string `json:"flag"`
	ImageName   string `json:"image_name"`
}

type CreateMovieWithIMDBIdServiceInputDTO struct {
	IMDBID string `json:"imdb_id"`
}

type CreateMovieWithIMDBIdServiceOutputDTO struct {
	ChooserID   string        `json:"chooser_id"`
	Title       string        `json:"title"`
	CountryName string        `json:"country_name"`
	Flag        string        `json:"flag"`
	ReleaseYear int           `json:"release_year"`
	ImageName   string        `json:"image_name"`
	Genres      []GenreDTO    `json:"genres"`
	Directors   []DirectorDTO `json:"directors"`
	Actors      []ActorDTO    `json:"actors"`
	Writers     []WriterDTO   `json:"writers"`
}

type CreateMovieWithIMDBIdService struct {
	IMDBRepository repositoryinterface.IMDBRepositoryInterface
}

func NewCreateMovieWithIMDBIdService(
	IMDBRepository repositoryinterface.IMDBRepositoryInterface,
) *CreateMovieWithIMDBIdService {
	return &CreateMovieWithIMDBIdService{
		IMDBRepository: IMDBRepository,
	}
}

func scrapingIMDB(IMDBID string) (string, string, string, string, int, []string, []string, []string, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}
	imdbLink := "https://www.imdb.com"

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var title, poster, country string
	var directors, writers, actors, releaseYears []string

	c.OnHTML("span.hero__primary-text[data-testid='hero__primary-text']", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML("div[data-testid='hero-media__poster'] img", func(e *colly.HTMLElement) {
		url := e.Attr("srcset")
		url = strings.Replace(strings.Split(url, ", ")[2], " 380w", "", -1)
		poster = url
	})

	c.OnHTML("li.ipc-metadata-list__item[data-testid='title-details-origin'] a.ipc-metadata-list-item__list-content-item--link", func(e *colly.HTMLElement) {
		country = e.Text
	})

	var firstListProcessed bool

	var count = 0

	c.OnHTML("ul.ipc-metadata-list.ipc-metadata-list--dividers-all.title-pc-list.ipc-metadata-list--baseAlt", func(e *colly.HTMLElement) {
		if count > 0 {
			return
		}

		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			if strings.Contains(el.Text, "Director") {
				el.ForEach("a", func(_ int, ell *colly.HTMLElement) {
					directors = append(directors, strings.Split(ell.Attr("href"), "/")[2])
				})
			}

			if strings.Contains(el.Text, "Writer") {
				el.ForEach("a", func(_ int, ell *colly.HTMLElement) {
					if !strings.Contains(ell.Attr("href"), "writer?") {
						writers = append(writers, strings.Split(ell.Attr("href"), "/")[2])
					}
				})
			}

			if strings.Contains(el.Text, "Star") {
				el.ForEach("a", func(_ int, ell *colly.HTMLElement) {
					if !strings.Contains(ell.Attr("href"), "cast?") {
						actors = append(actors, strings.Split(ell.Attr("href"), "/")[2])
					}
				})
			}
		})

		count++
	})

	c.OnHTML("ul.ipc-inline-list li.ipc-inline-list__item:first-of-type a.ipc-link--baseAlt", func(e *colly.HTMLElement) {
		releaseYear := e.Text
		releaseYears = append(releaseYears, releaseYear)
	})

	c.OnRequest(func(r *colly.Request) {
		if !firstListProcessed {
			firstListProcessed = true
			return
		}
		r.Abort()
	})

	c.Visit(imdbLink + "/title/" + IMDBID)
	c.Wait()

	flag, err := getFlagFromCountry(country)
	if err != nil {
		var nothing []string
		return "", "", "", "", 0, nothing, nothing, nothing, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	imageName, err := loadImageFromURL(poster)
	if err != nil {
		var nothing []string
		return "", "", "", "", 0, nothing, nothing, nothing, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	releaseYear, err := strconv.Atoi(releaseYears[1])
	if err != nil {
		var nothing []string
		return "", "", "", "", 0, nothing, nothing, nothing, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	return title, country, flag, imageName, releaseYear, directors, writers, actors, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}

func personsTreatment(directors, actors, writers []string) ([]DirectorDTO, []ActorDTO, []WriterDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	var directorsDTO []DirectorDTO
	var actorsDTO []ActorDTO
	var writersDTO []WriterDTO

	for _, director := range directors {
		newDirector, err := getDirector(director)
		if err != nil {
			return []DirectorDTO{}, []ActorDTO{}, []WriterDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}
		directorsDTO = append(directorsDTO, newDirector)
	}

	for _, actor := range actors {
		newActor, err := getActor(actor)
		if err != nil {
			return []DirectorDTO{}, []ActorDTO{}, []WriterDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}
		actorsDTO = append(actorsDTO, newActor)
	}

	for _, writer := range writers {
		newWriter, err := getWriter(writer)
		if err != nil {
			return []DirectorDTO{}, []ActorDTO{}, []WriterDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}
		writersDTO = append(writersDTO, newWriter)
	}

	return directorsDTO, actorsDTO, writersDTO, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}

func (cm *CreateMovieWithIMDBIdService) Execute(input CreateMovieWithIMDBIdServiceInputDTO) (CreateMovieWithIMDBIdServiceOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	title, country, flag, imageName, releaseYear, directors, writers, actors, problemsDetailsOutput := scrapingIMDB(input.IMDBID)
	if len(problemsDetailsOutput.ProblemDetails) > 0 {
		return CreateMovieWithIMDBIdServiceOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetailsOutput.ProblemDetails,
		}
	}

	var imdbIDs []string
	var imdbs []entity.IMDB

	imdbIDs = append(imdbIDs, directors...)
	imdbIDs = append(imdbIDs, actors...)
	imdbIDs = append(imdbIDs, writers...)
	imdbIDs = append(imdbIDs, input.IMDBID)

	directorsDTO, actorsDTO, writersDTO, problemsDetailsOutput := personsTreatment(directors, actors, writers)
	if len(problemsDetailsOutput.ProblemDetails) > 0 {
		return CreateMovieWithIMDBIdServiceOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetailsOutput.ProblemDetails,
		}
	}

	for _, imdbID := range imdbIDs {
		newIMDB, err := entity.NewIMDB(imdbID)
		if err != nil {
			return CreateMovieWithIMDBIdServiceOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		imdbs = append(imdbs, *newIMDB)
	}

	err := cm.IMDBRepository.CreateMany(&imdbs)
	if err != nil {
		return CreateMovieWithIMDBIdServiceOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newMovie := CreateMovieWithIMDBIdServiceOutputDTO{
		Title:       title,
		CountryName: country,
		Flag:        flag,
		ReleaseYear: releaseYear,
		ImageName:   imageName,
		Directors:   directorsDTO,
		Actors:      actorsDTO,
		Writers:     writersDTO,
	}

	return newMovie, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}

func getFlagFromCountry(country string) (string, error) {
	countries := valueobject.NewCountries()

	for _, c := range countries {
		if c.Name == country {
			return c.Flag, nil
		}
	}

	return "", nil
}

func loadImageFromURL(poster string) (string, error) {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	path := configs.LocalImagePath

	if len(poster) == 0 {
		return "no-image", err
	}

	response, err := http.Get(poster)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	imageName := ulid.Make().String() + ".jpg"
	imagePath := path + imageName

	file, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return imageName, err
}

func getDirector(imdbID string) (DirectorDTO, error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var name, countryName, flag, poster string
	var birthDate, birthPlace, dayStr, yearStr string
	var day, month, year int

	c.OnHTML("span.hero__primary-text[data-testid='hero__primary-text']", func(e *colly.HTMLElement) {
		name = e.Text
	})

	c.OnHTML("li[data-testid='nm_pd_bl']", func(e *colly.HTMLElement) {
		birthDate = e.ChildText("ul > li:nth-child(1) > a")
		birthPlace = e.ChildText("ul > li:nth-child(2) > a")
	})

	c.OnHTML(".ipc-poster__poster-image img", func(e *colly.HTMLElement) {
		srcSet := e.Attr("srcset")
		parts := strings.Split(srcSet, ", ")
		for _, part := range parts {
			if strings.Contains(part, "280w") {
				poster = strings.Fields(part)[0]
				break
			}
		}
	})

	c.Visit("https://www.imdb.com/name/" + imdbID)
	c.Wait()

	if len(birthPlace) > 0 {
		birthPlaceParts := strings.Split(birthPlace, ", ")

		if len(birthPlaceParts) == 1 {
			countryName = birthPlaceParts[0]
		} else if len(birthPlaceParts) == 2 {
			countryName = birthPlaceParts[1]
		} else if len(birthPlaceParts) == 3 {
			countryName = birthPlaceParts[2]
		} else if len(birthPlaceParts) == 4 {
			countryName = birthPlaceParts[2]
		}

		var err error

		flag, err = getFlagFromCountry(countryName)
		if err != nil {
			return DirectorDTO{}, err
		}
	} else {
		countryName = "no-name"
		flag = "no-flag"
	}

	if len(birthDate) > 0 {
		birthDateParts := strings.Split(birthDate, " ")

		monthStr := birthDateParts[0]
		birthDatePart := birthDateParts[1]

		if len(birthDatePart) == 0 {
			dayStr = "1"
			yearStr = "1900"
		} else if len(birthDatePart) == 1 {
			dayStr = birthDatePart[:1]
			yearStr = "1900"
		} else if len(birthDatePart) == 2 {
			dayStr = birthDatePart[:2]
			yearStr = "1900"
		} else if len(birthDatePart) == 5 {
			dayStr = birthDatePart[:1]
			yearStr = birthDatePart[1:]
		} else if len(birthDatePart) == 6 {
			dayStr = birthDatePart[:2]
			yearStr = birthDatePart[2:]
		}

		var err error

		day, err = strconv.Atoi(dayStr)
		if err != nil {
			return DirectorDTO{}, err
		}

		month, err = getMonthNumber(monthStr)
		if err != nil {
			return DirectorDTO{}, err
		}

		year, err = strconv.Atoi(yearStr)
		if err != nil {
			return DirectorDTO{}, err
		}
	} else if len(birthDate) == 4 {
		day = 1
		month = 1

		var err error

		year, err = strconv.Atoi(birthDate)
		if err != nil {
			return DirectorDTO{}, err
		}
	} else {
		day = 1
		month = 1
		year = 1900
	}

	imageName, err := loadImageFromURL(poster)
	if err != nil {
		return DirectorDTO{}, err
	}

	newDirector := DirectorDTO{
		DirectorID:  "",
		Name:        name,
		Day:         day,
		Month:       month,
		Year:        year,
		CountryName: countryName,
		Flag:        flag,
		ImageName:   imageName,
	}

	return newDirector, nil
}

func getActor(imdbID string) (ActorDTO, error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var name, countryName, flag, poster string
	var birthDate, birthPlace, dayStr, yearStr string
	var day, month, year int

	c.OnHTML("span.hero__primary-text[data-testid='hero__primary-text']", func(e *colly.HTMLElement) {
		name = e.Text
	})

	c.OnHTML("li[data-testid='nm_pd_bl']", func(e *colly.HTMLElement) {
		birthDate = e.ChildText("ul > li:nth-child(1) > a")
		birthPlace = e.ChildText("ul > li:nth-child(2) > a")
	})

	c.OnHTML(".ipc-poster__poster-image img", func(e *colly.HTMLElement) {
		srcSet := e.Attr("srcset")
		parts := strings.Split(srcSet, ", ")
		for _, part := range parts {
			if strings.Contains(part, "280w") {
				poster = strings.Fields(part)[0]
				break
			}
		}
	})

	c.Visit("https://www.imdb.com/name/" + imdbID)
	c.Wait()

	if len(birthPlace) > 0 {
		birthPlaceParts := strings.Split(birthPlace, ", ")

		if len(birthPlaceParts) == 1 {
			countryName = birthPlaceParts[0]
		} else if len(birthPlaceParts) == 2 {
			countryName = birthPlaceParts[1]
		} else if len(birthPlaceParts) == 3 {
			countryName = birthPlaceParts[2]
		} else if len(birthPlaceParts) == 4 {
			countryName = birthPlaceParts[2]
		}

		var err error

		flag, err = getFlagFromCountry(countryName)
		if err != nil {
			return ActorDTO{}, err
		}
	} else {
		countryName = "no-name"
		flag = "no-flag"
	}

	if len(birthDate) > 0 {
		birthDateParts := strings.Split(birthDate, " ")

		monthStr := birthDateParts[0]
		birthDatePart := birthDateParts[1]

		if len(birthDatePart) == 0 {
			dayStr = "1"
			yearStr = "1900"
		} else if len(birthDatePart) == 1 {
			dayStr = birthDatePart[:1]
			yearStr = "1900"
		} else if len(birthDatePart) == 2 {
			dayStr = birthDatePart[:2]
			yearStr = "1900"
		} else if len(birthDatePart) == 5 {
			dayStr = birthDatePart[:1]
			yearStr = birthDatePart[1:]
		} else if len(birthDatePart) == 6 {
			dayStr = birthDatePart[:2]
			yearStr = birthDatePart[2:]
		}

		var err error

		day, err = strconv.Atoi(dayStr)
		if err != nil {
			return ActorDTO{}, err
		}

		month, err = getMonthNumber(monthStr)
		if err != nil {
			return ActorDTO{}, err
		}

		year, err = strconv.Atoi(yearStr)
		if err != nil {
			return ActorDTO{}, err
		}
	} else if len(birthDate) == 4 {
		day = 1
		month = 1

		var err error

		year, err = strconv.Atoi(birthDate)
		if err != nil {
			return ActorDTO{}, err
		}
	} else {
		day = 1
		month = 1
		year = 1900
	}

	imageName, err := loadImageFromURL(poster)
	if err != nil {
		return ActorDTO{}, err
	}

	newActor := ActorDTO{
		ActorID:     "",
		Name:        name,
		Day:         day,
		Month:       month,
		Year:        year,
		CountryName: countryName,
		Flag:        flag,
		ImageName:   imageName,
	}

	return newActor, nil
}

func getWriter(imdbID string) (WriterDTO, error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var name, countryName, flag, poster string
	var birthDate, birthPlace, dayStr, yearStr string
	var day, month, year int

	c.OnHTML("span.hero__primary-text[data-testid='hero__primary-text']", func(e *colly.HTMLElement) {
		name = e.Text
	})

	c.OnHTML("li[data-testid='nm_pd_bl']", func(e *colly.HTMLElement) {
		birthDate = e.ChildText("ul > li:nth-child(1) > a")
		birthPlace = e.ChildText("ul > li:nth-child(2) > a")
	})

	c.OnHTML(".ipc-poster__poster-image img", func(e *colly.HTMLElement) {
		srcSet := e.Attr("srcset")
		parts := strings.Split(srcSet, ", ")
		for _, part := range parts {
			if strings.Contains(part, "280w") {
				poster = strings.Fields(part)[0]
				break
			}
		}
	})

	c.Visit("https://www.imdb.com/name/" + imdbID)
	c.Wait()

	if len(birthPlace) > 0 {
		birthPlaceParts := strings.Split(birthPlace, ", ")

		if len(birthPlaceParts) == 1 {
			countryName = birthPlaceParts[0]
		} else if len(birthPlaceParts) == 2 {
			countryName = birthPlaceParts[1]
		} else if len(birthPlaceParts) == 3 {
			countryName = birthPlaceParts[2]
		} else if len(birthPlaceParts) == 4 {
			countryName = birthPlaceParts[2]
		}

		var err error

		flag, err = getFlagFromCountry(countryName)
		if err != nil {
			return WriterDTO{}, err
		}
	} else {
		countryName = "no-name"
		flag = "no-flag"
	}

	if len(birthDate) > 4 {
		birthDateParts := strings.Split(birthDate, " ")

		monthStr := birthDateParts[0]
		birthDatePart := birthDateParts[1]

		if len(birthDatePart) == 0 {
			dayStr = "1"
			yearStr = "1900"
		} else if len(birthDatePart) == 1 {
			dayStr = birthDatePart[:1]
			yearStr = "1900"
		} else if len(birthDatePart) == 2 {
			dayStr = birthDatePart[:2]
			yearStr = "1900"
		} else if len(birthDatePart) == 5 {
			dayStr = birthDatePart[:1]
			yearStr = birthDatePart[1:]
		} else if len(birthDatePart) == 6 {
			dayStr = birthDatePart[:2]
			yearStr = birthDatePart[2:]
		}

		var err error

		day, err = strconv.Atoi(dayStr)
		if err != nil {
			return WriterDTO{}, err
		}

		month, err = getMonthNumber(monthStr)
		if err != nil {
			return WriterDTO{}, err
		}

		year, err = strconv.Atoi(yearStr)
		if err != nil {
			return WriterDTO{}, err
		}
	} else if len(birthDate) == 4 {
		day = 1
		month = 1

		var err error

		year, err = strconv.Atoi(birthDate)
		if err != nil {
			return WriterDTO{}, err
		}
	} else {
		day = 1
		month = 1
		year = 1900
	}

	imageName, err := loadImageFromURL(poster)
	if err != nil {
		return WriterDTO{}, err
	}

	newWriter := WriterDTO{
		WriterID:    "",
		Name:        name,
		Day:         day,
		Month:       month,
		Year:        year,
		CountryName: countryName,
		Flag:        flag,
		ImageName:   imageName,
	}

	return newWriter, nil
}

func getMonthNumber(monthName string) (int, error) {
	var month int

	months := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	for key, value := range months {
		if key == monthName {
			month = value
		}
	}

	if month == 0 {
		return 0, errors.New("month not found")
	}

	return month, nil
}
