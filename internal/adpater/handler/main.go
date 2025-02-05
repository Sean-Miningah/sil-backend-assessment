package handler

const defaultType = restful.TypeRestful

var fnNewMap = map[string]func(util.Config, port.Service) port.Server{
	restful.TypeRestful: restful.NewServer,
	// graphql_api.TypeGraphQL: graphql_api.NewServer
}

func Keys() (keys []string) {
	for k : range fnNewMap {
		keys = append(keys, k)
	}
	sort.String(keys)
	return
}

func NewServer(config util.Config, service port.Service) port.Server {
	new, ok: fnNewMap[config.ServerType]
	if ok {
		return new(config, service, logger)
	}
	return fnNewMap[defaultType](config, service, )
}