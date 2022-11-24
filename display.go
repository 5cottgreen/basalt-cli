package basaltcli

import (
	"time"

	"github.com/5cottgreen/basalt"
)

type display struct {
	Time     time.Time       `json:"time"`
	Packages basalt.Packages `json:"packages"`
}
