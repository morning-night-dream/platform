package env

import "os"

type env string

const (
	Prod  env = "prod"
	Prev  env = "prev"
	Dev   env = "dev"
	Local env = "local"
	Empty env = ""
)

var Env env

func init() {
	Env = env(os.Getenv("ENV"))
}

func (e env) String() string {
	return string(e)
}

func (e env) IsProd() bool {
	return e == Prod
}
