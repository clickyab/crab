// Code generated build with variable DO NOT EDIT.

package user

// AUTO GENERATED CODE. DO NOT EDIT!
import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/trans"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

func (pl *checkActivePayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
			res["email"] = trans.E("invalid value")

		case "Number":
			res["number"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

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

func (pl *corporation) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "Name":
			res["name"] = trans.E("invalid value")

		case "FirstName":
			res["first_name"] = trans.E("invalid value")

		case "LastName":
			res["last_name"] = trans.E("invalid value")

		case "Cellphone":
			res["cellphone"] = trans.E("invalid value")

		case "Phone":
			res["phone"] = trans.E("invalid value")

		case "Address":
			res["address"] = trans.E("invalid value")

		case "EconomicCode":
			res["economic_code"] = trans.E("invalid value")

		case "RegisterCode":
			res["register_code"] = trans.E("invalid value")

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

func (pl *loginPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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

		case "Password":
			res["password"] = trans.E("password is too short")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (pl *callBack) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "Token":
			res["token"] = trans.E("invalid value")

		case "NewPassword":
			res["new_password"] = trans.E("password is too short")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (pl *registerPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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

		case "Password":
			res["password"] = trans.E("password is too short")

		case "FirstName":
			res["first_name"] = trans.E("first name is invalid")

		case "LastName":
			res["last_name"] = trans.E("last name is invalid")

		case "UserType":
			res["user_type"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (pl *sendActivePayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
			res["email"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}
