// Wrapper for genshin-db-api by theBowja
// GitHub: https://github.com/theBowja/genshin-db-api
package genshindb_wrapper

import (
	"encoding/json"
	"errors"
	"ga/pkg/genshindb_wrapper/enums"
	"ga/pkg/genshindb_wrapper/models"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

type genshinDbApi struct {
	URL      string
	HTTP     *http.Client
	Logger   *zap.SugaredLogger
	Language enums.Language
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func Create(host string, language enums.Language, logger *zap.SugaredLogger) *genshinDbApi {
	api := genshinDbApi{
		URL:      host,
		HTTP:     httpClient(),
		Logger:   logger,
		Language: language,
	}

	return &api
}

func (api *genshinDbApi) GetCharacter(query string) (models.CharacterWeb, error) {
	var character models.CharacterWeb

	response, err := api.makeRequest(api.URL + "/characters" + "?query=" + url.QueryEscape(query) + "&resultLanguage=" + string(api.Language))
	if err != nil {
		return character, err
	}

	err = json.Unmarshal(response, &character)
	if err != nil {
		return character, err
	}

	return character, nil
}

func (api *genshinDbApi) GetAllCharacters() ([]models.CharacterWeb, error) {
	var characters []models.CharacterWeb
	charactersNames, err := api.GetAllCharactersNames()

	if err != nil {
		return characters, err
	}

	for _, characterName := range charactersNames {
		character, err := api.GetCharacter(characterName)

		if err != nil {
			return characters, err
		}

		if character.Name != characterName {
			return characters, errors.New("failed to get character " + characterName)
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func (api *genshinDbApi) GetAllCharactersNames() ([]string, error) {
	var charactersNames []string
	response, err := api.makeRequest(api.URL + "/characters" + "?query=names&matchCategories=true" + "&resultLanguage=" + string(api.Language))

	if err != nil {
		return charactersNames, err
	}

	err = json.Unmarshal(response, &charactersNames)

	if err != nil {
		return charactersNames, err
	}

	return charactersNames, nil
}

func (api *genshinDbApi) makeRequest(request string) ([]byte, error) {
	api.Logger.Debugf("[GET] %s", request)

	responce, err := api.HTTP.Get(request)

	if err != nil {
		return nil, err
	}

	defer responce.Body.Close()

	return api.decodeAPIResponse(responce)
}

func (api *genshinDbApi) decodeAPIResponse(responce *http.Response) ([]byte, error) {

	data, err := io.ReadAll(responce.Body)
	if err != nil {
		return nil, err
	}

	api.Logger.Debugw("Response",
		"status code", responce.StatusCode,
		"data", string(data))

	return data, nil
}
