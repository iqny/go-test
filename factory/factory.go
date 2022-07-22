package factory

type ConfigParse interface {
	Parse(path string) string
}
type jsonConfigParse struct{}

func (json *jsonConfigParse) Parse(path string) string {
	return "解析json配置文件，路径为:" + path
}

type yamlConfigParse struct{}

func (yaml *yamlConfigParse) Parse(path string) string {
	return "解析yaml配置文件，路径为:" + path
}

type SimpleParseFactory struct {
}

func (s *SimpleParseFactory) Create(name string) ConfigParse {
	switch name {
	case "json":
		return &jsonConfigParse{}
	case "yaml":
		return &yamlConfigParse{}
	}
	return nil
}

type NormaParseFactory interface {
	createParse() ConfigParse
}

type JsonNormaParseFactory struct{}

func (json *JsonNormaParseFactory) createParse() ConfigParse {
	return &jsonConfigParse{}
}

type YamlNormaParseFactory struct{}

func (yaml *YamlNormaParseFactory) createParse() ConfigParse {
	return &yamlConfigParse{}
}

func CreateFactory(name string) NormaParseFactory {
	switch name {
	case "json":
		return &JsonNormaParseFactory{}
	case "yaml":
		return &YamlNormaParseFactory{}
	}
	return nil
}
