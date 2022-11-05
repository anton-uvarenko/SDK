package lotr_sdk

import (
	"github.com/anton-uvarenko/SDK/lotr-sdk/models"
	"net/http"
)

type lotrController struct {
	client *http.Client
}

type Lotr interface {
	GetAllCharacters(params ...string) ([]models.Character, error)
	GetCharacterById(id string, params ...string) (*models.Character, error)
	GetAllChatacterQuotes(id string, params ...string) ([]models.Quote, error)
	GetAllBooks(params ...string) ([]models.Book, error)
	GetBookById(id string, params ...string) (*models.Book, error)
	GetAllBookChapters(id string, params ...string) ([]models.Chapter, error)
	GetAllBooksChapters(params ...string) ([]models.Chapter, error)
	GetBookChapter(id string, params ...string) (*models.Chapter, error)
	GetAllMovies(params ...string) ([]models.Movie, error)
	GetMovieById(id string, params ...string) (*models.Movie, error)
	GetAllMovieQuotes(id string, params ...string) ([]models.Quote, error)
	GetAllQoutes(params ...string) ([]models.Quote, error)
	GetMovieQoute(id string, params ...string) (*models.Quote, error)
}

func NewSdk() Lotr {
	client := &http.Client{}
	return &lotrController{client}
}
