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

func (p *createDomainPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "Name":
			res["name"] = trans.E("invalid value")

		case "Description":
			res["description"] = trans.E("invalid value")

		case "Attributes":
			res["attributes"] = trans.E("invalid value")

		case "Status":
			res["status"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (p *editDomainPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "Description":
			res["description"] = trans.E("invalid value")

		case "Attributes":
			res["attributes"] = trans.E("invalid value")

		case "Status":
			res["status"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}
