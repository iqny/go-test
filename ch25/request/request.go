package request

type ConfigFun func(c *Config)
type Config struct {
	Abc string
}

func New(cfgs ...ConfigFun) {
	c := &Config{}
	for _, f := range cfgs {
		f(c)
	}
}
