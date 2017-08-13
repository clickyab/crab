// Code generated build with variable DO NOT EDIT.

package user

// AUTO GENERATED CODE. DO NOT EDIT!
import (
	"context"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/trans"
	validator "gopkg.in/go-playground/validator.v9"
)

func (pl *checkMailPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := func(in interface{}) error {
		if v, ok := in.(interface {
			ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error
		}); ok {
			return v.ValidateExtra(ctx, w, r)
		}
		return nil
	}(pl)
	if err != nil {
		return err
	}
	errs := validator.New().Struct(pl)
	if errs == nil {
		return nil
	}
	res := middleware.GroupError{}
	for _, i := range errs.(validator.ValidationErrors) {
		switch i.Field() {
		case "Email":
			res["email"] = trans.E("email is invalid")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (pl *personalPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := func(in interface{}) error {
		if v, ok := in.(interface {
			ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error
		}); ok {
			return v.ValidateExtra(ctx, w, r)
		}
		return nil
	}(pl)
	if err != nil {
		return err
	}
	errs := validator.New().Struct(pl)
	if errs == nil {
		return nil
	}
	res := middleware.GroupError{}
	for _, i := range errs.(validator.ValidationErrors) {
		switch i.Field() {
		case "FirstName":
			res["first_name"] = trans.E("invalid value")

		case "LastName":
			res["last_name"] = trans.E("invalid value")

		case "Gender":
			res["gender"] = trans.E("invalid value")

		case "CellPhone":
			res["cellphone"] = trans.E("invalid value")

		case "Phone":
			res["phone"] = trans.E("invalid value")

		case "Address":
			res["address"] = trans.E("invalid value")

		case "CityID":
			res["city_id"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}