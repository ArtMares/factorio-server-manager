package api

import (
    "fmt"
    "github.com/artmares/factorio-server-manager/model"

    "bitbucket.org/aisgteam/ai-api-client"
    "bitbucket.org/aisgteam/ai-struct"
    "github.com/afex/hystrix-go/hystrix"
)

var (
    modsName = "Factorio Mods Api Client"
    modsVersion = aistruct.Version{0, 0, 1, 0}
    modsApi *ModsApiClient
)

type ModsApiClient struct {
    client      aiapic.HystrixClient
}

func NewModsApi() (*ModsApiClient, error) {
    modsApi = &ModsApiClient{}
    modsApi.client.Name = modsName
    modsApi.client.Version = modsVersion
    modsApi.client.UserAgent = fmt.Sprintf("%s (%s)", modsName, modsVersion)
    err := modsApi.client.New("https://mods.factorio.com/api")
    if err != nil {
        return nil, err
    }

    hystrix.ConfigureCommand(modsName, hystrix.CommandConfig{
        Timeout: 50000,
        MaxConcurrentRequests: 300,
        RequestVolumeThreshold: 10,
        SleepWindow: 1000,
        ErrorPercentThreshold: 10,
    })

    return modsApi, nil
}

func (api *ModsApiClient) Mods() (*model.Mods, error) {

}