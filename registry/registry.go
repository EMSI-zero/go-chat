package registry

import (
	"github.com/EMSI-zero/go-chat/infra/colrepo"
	"github.com/EMSI-zero/go-chat/infra/dbrepo"
)

type ServiceRegistry interface {
	mustImplementBaseRegistry()
	GetDB() dbrepo.DBConn
	GetColDB() colrepo.ColDBConn

}

type serviceRegistry struct {
	services serviceMap
	db       dbrepo.DBConn
	colDB    colrepo.ColDBConn
}

func mustImplementBaseRegistry() {}

func (sr *serviceRegistry) mustImplementBaseRegistry() {}

type serviceMap struct {
	
}

func NewServiceRegistry(db dbrepo.DBConn, colDB colrepo.ColDBConn) *serviceRegistry {
	//create an empty service registry
	sr := new(serviceRegistry)
	sr.db = db
	sr.colDB = colDB
	return sr
}

func (sr *serviceRegistry) GetDB() dbrepo.DBConn {
	return sr.db
}

func (sr *serviceRegistry) GetColDB() colrepo.ColDBConn {
	return sr.colDB
}
