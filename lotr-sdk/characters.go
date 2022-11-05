package lotr_sdk

import (
	"encoding/json"
	"github.com/anton-uvarenko/SDK/lotr-sdk/helpers"
	"github.com/anton-uvarenko/SDK/lotr-sdk/models"
)

func (c lotrController) GetAllCharacters(params ...string) ([]models.Character, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/character"+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.CharaterResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Docs, nil
}

func (c lotrController) GetCharacterById(id string, params ...string) (*models.Character, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/character/"+id+str, c.client)
	if err != nil {
		return nil, err
	}

	data := models.CharaterResponse{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data.Docs[0], nil
}

func (c lotrController) GetAllChatacterQuotes(id string, params ...string) ([]models.Quote, error) {
	str := helpers.ValidateParams(params...)

	resp, err := helpers.CreateRequestWithRout("/character/"+id+str, c.client)
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
