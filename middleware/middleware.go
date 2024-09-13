package middleware

import (
	"github.com/deenikarim/gudu"
	"myapp/data"
)

type Middleware struct {
	App    *gudu.Gudu
	Models data.Models
}
