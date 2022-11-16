package config

type Config struct {
	Port    string `default:"8080"`
	Twitter struct {
		Token            string
		Secret           string
		RequestURI       string
		AuthorizationURI string
		TokenRequestURI  string
		CallbackURI      string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.Twitter.Token = ""
	c.Twitter.Secret = ""
	c.Twitter.RequestURI = ""
	c.Twitter.AuthorizationURI = ""
	c.Twitter.TokenRequestURI = ""
	c.Twitter.CallbackURI = ""

	return c
}
