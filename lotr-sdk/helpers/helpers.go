package helpers

import (
	"anton.uvarenko/sdk/lotr-sdk/constants"
	"fmt"
	"net/http"
	"strings"
)

func CreateRequestWithRout(rout string, client *http.Client) (*http.Response, error) {
	request, err := http.NewRequest("GET", constants.Prefix+rout, nil)
	request.Header.Add("Authorization", "Bearer "+constants.Token)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ValidateParams(params ...string) string {
	var paggination []string
	var sorting []string
	var filtering []string

	for _, v := range params {
		t := strings.SplitN(v, "=", -1)[0]
		if t == "limit" || t == "offset" || t == "page" {
			paggination = append(paggination, v)
			continue
		}
		if t == "sort" {
			sorting = append(sorting, v)
			continue
		} else {
			filtering = append(filtering, v)
		}
	}

	fmt.Println(formRequestString(paggination))

	return formRequestString(paggination) + formRequestString(sorting) + formRequestString(filtering)
}

func formRequestString(strs []string) string {
	str := ""
	for _, v := range strs {
		str += "?" + v
	}

	return str
}
