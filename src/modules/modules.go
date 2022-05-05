package modules

type Module interface{}

func Get(moduleName string) Module {
	if moduleName == "solid" {
		return Solid{}
	}

	return Dummy{}
}
