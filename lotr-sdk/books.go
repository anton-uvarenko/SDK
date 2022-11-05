package lotr_sdk

import (
	"encoding/json"
	"github.com/anton-uvarenko/SDK/lotr-sdk/helpers"
	"github.com/anton-uvarenko/SDK/lotr-sdk/models"
)

func (c lotrController) GetAllBooks(params ...string) ([]models.Book, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/book"+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.BookResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Docs, nil
}

func (c lotrController) GetBookById(id string, params ...string) (*models.Book, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/book/"+id+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.BookResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data.Docs[0], nil
}

func (c lotrController) GetAllBookChapters(id string, params ...string) ([]models.Chapter, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/book/"+id+"/chapter"+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.ChapterResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)

	return data.Docs, nil
}

func (c lotrController) GetAllBooksChapters(params ...string) ([]models.Chapter, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/chapter"+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.ChapterResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Docs, nil
}

func (c lotrController) GetBookChapter(id string, params ...string) (*models.Chapter, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/chapter/"+id+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.ChapterResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data.Docs[0], nil
}
