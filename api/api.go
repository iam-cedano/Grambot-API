package api

type API struct {
	host string
}

func NewAPI(host string) *API {
	return &API{host: host}
}

func (api *API) GetHost() string {
	return api.host
}
