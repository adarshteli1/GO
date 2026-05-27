package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	Enforcer = e
}
