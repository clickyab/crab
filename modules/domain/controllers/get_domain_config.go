package controllers

import (
	"context"
	"net/http"
	"strconv"

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

// getDomainConfig get domain config by domain id
// @Rest {
// 		url = /config/:id
// 		method = get
// }
func (c *Controller) getDomainConfig(ctx context.Context, r *http.Request) (*domainConfig, error) {
	// get domain id from url params
	domainID, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	// find domain
	m := orm.NewOrmManager()
	domainObj, err := m.FindDomainByID(domainID)
	if err != nil {
		return nil, errors.DomainNotFoundError(domainID)
	}
	res := &domainConfig{
		Title:        domainObj.Title,
		Logo:         domainObj.Logo.String,
		DomainStatus: domainObj.Status,
		Theme:        domainObj.Theme,
	}
	return res, nil
}
