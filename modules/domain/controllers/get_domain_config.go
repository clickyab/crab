package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/orm"
	"github.com/rs/xmux"
)

type domainConfig struct {
	DomainStatus orm.DomainStatus `json:"domain_status"`
	Theme        string           `json:"theme"`
	Logo         string           `json:"logo"`
	Title        string           `json:"title"`
}

// getDomainConfig get domain config by domain name
// @Rest {
// 		url = /config/:name
// 		method = get
// }
func (c *Controller) getDomainConfig(ctx context.Context, r *http.Request) (*domainConfig, error) {
	// get domain id from url params
	domainName := xmux.Param(ctx, "name")
	// find domain
	m := orm.NewOrmManager()
	domainObj, err := m.FindDomainByDomainBase(domainName)
	if err != nil {
		return nil, errors.DomainNotFoundErrorByName(domainName)
	}
	res := &domainConfig{
		Title:        domainObj.Title,
		Logo:         domainObj.Logo.String,
		DomainStatus: domainObj.Status,
		Theme:        domainObj.Theme,
	}
	return res, nil
}
