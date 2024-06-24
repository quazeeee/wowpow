package conf

type ServerConfig struct {
	Config
}

func NewServerConfig() *ServerConfig {
	c := newConfig()
	if c == nil {
		return nil
	}

	return &ServerConfig{
		Config: *c,
	}
}
