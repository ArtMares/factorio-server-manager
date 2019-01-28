package api

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"

    "bitbucket.org/aisgteam/ai-api-client"
    "bitbucket.org/aisgteam/ai-struct"
    "github.com/afex/hystrix-go/hystrix"
    "github.com/artmares/factorio-server-manager/model"
)

var (
    authName    = "Factorio Auth Api Client"
    authVersion = aistruct.Version{0, 0, 1, 0}
    authApi     *AuthApiClient
)

type AuthApiClient struct {
    client      aiapic.HystrixClient
}

func NewAuthApi() (*AuthApiClient, error) {
    authApi = &AuthApiClient{}
    authApi.client.Name = authName
    authApi.client.Version = authVersion
    authApi.client.UserAgent = fmt.Sprintf("%s (%s)", authName, authVersion)
    err := authApi.client.New("https://auth.factorio.com")
    if err != nil {
        return nil, err
    }

    hystrix.ConfigureCommand(authName, hystrix.CommandConfig{
        Timeout: 50000,
        MaxConcurrentRequests: 300,
        RequestVolumeThreshold: 10,
        SleepWindow: 1000,
        ErrorPercentThreshold: 10,
    })

    return authApi, nil
}

func (api *AuthApiClient) Auth(username, password string) (*model.AuthToken, error) {
    if username != "" && password != "" {
        u := api.client.PrepareUri("/api-login")
        v := url.Values{
            "username": {username},
            "password": {password},
        }
        token := model.AuthToken("")
        req, err := api.client.Prepare(http.MethodPost, u, strings.NewReader(v.Encode()))
        if err != nil {
            return nil, api.client.Err("Auth", err.Error())
        }
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
        res, err := api.client.Exec(req)
        if err != nil {
            return nil, api.client.Err("Auth", err.Error())
        }
        defer res.Body.Close()

        if res.StatusCode != 200 {
            return nil, api.client.Err("Auth", "Request error" + res.Status)
        }

        var t []string
        err = json.NewDecoder(res.Body).Decode(&t)
        if err != nil {
            return nil, api.client.Err("Auth", err.Error())
        } else {
            if len(t) > 0 {
                token = model.AuthToken(t[0])
                return &token, nil
            } else {
                return nil, api.client.Err("Auth", "Failed get token")
            }
        }
    } else {
        return nil, api.client.Err("Auth", "Username or password is empty")
    }
}