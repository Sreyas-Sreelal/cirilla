package types

//Config structure to store configuration datas
//nolint:lll
type Config struct {
	TelegramToken string `envconfig:"CIRILLA_TOKEN"`
	Debug         bool   `envconfig:"CIRILLA_DEBUG" default:"false"`
	Timeout       int    `ennconfig:"CIRILLA_TIMEOUT"`
}
