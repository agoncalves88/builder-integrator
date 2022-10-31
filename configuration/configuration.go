package configuration

type Config struct {
	TESTE       string
	DATASOURCES []struct {
		NAME              string
		URL               string
		HAS_PARAM         bool
		PROPERTIES_TO_GET []struct {
			VALUE_TO_GET string
			INPUT_NAME   string
		}
	}
}
