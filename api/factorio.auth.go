package api

import (
    "bitbucket.org/aisgteam/ai-api-client"
    "bitbucket.org/aisgteam/ai-struct"
    "bitbucket.org/aisgteam/w-app"
)

var (
    authName    = "Factorio Auth Api Client"
    authVersion = aistruct.Version{0, 0, 1, 0}
    authApi     *AuthApi
)

type AuthApi struct {
    client  aiapic.HystrixClient
}

func NewAuthApi() (*AuthApi, error) {
    authApi = &AuthApi{}
    authApi.client.Name = authName
    authApi.client.Version = authVersion
    authApi.client.UserAgent = wapp.FullName()
}