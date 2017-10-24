// Code generated build with variable DO NOT EDIT.

package controllers

// AUTO GENERATED CODE. DO NOT EDIT!
import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/trans"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

func (p *assignBannerPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := func(in interface{}) error {
		if v, ok := in.(interface {
			ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error
		}); ok {
			return v.ValidateExtra(ctx, w, r)
		}
		return nil
	}(p)
	if err != nil {
		return err
	}
	errs := validator.New().Struct(p)
	if errs == nil {
		return nil
	}
	res := middleware.GroupError{}
	for _, i := range errs.(validator.ValidationErrors) {
		switch i.Field() {
		case "Src":
			res["src"] = trans.E("invalid value")

		case "Utm":
			res["utm"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}
