package lotr_sdk

import (
	"anton.uvarenko/sdk/lotr-sdk/helpers"
	"anton.uvarenko/sdk/lotr-sdk/models"
	"encoding/json"
)

func (c lotrController) GetAllMovies(params ...string) ([]models.Movie, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/movie"+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.MovieResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Docs, nil
}

func (c lotrController) GetMovieById(id string, params ...string) (*models.Movie, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/movie/"+id+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.MovieResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data.Docs[0], nil
}

func (c lotrController) GetAllMovieQuotes(id string, params ...string) ([]models.Quote, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/movie/"+id+"/quote"+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.QuoteResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Docs, nil
}

func (c lotrController) GetAllQoutes(params ...string) ([]models.Quote, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/quote"+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.QuoteResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Docs, nil
}

func (c lotrController) GetMovieQoute(id string, params ...string) (*models.Quote, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/quote/"+id+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.QuoteResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data.Docs[0], nil
}
