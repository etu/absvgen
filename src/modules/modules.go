package modules

type Module interface{}

func Get(moduleName string) Module {
	return Dummy{}
}
