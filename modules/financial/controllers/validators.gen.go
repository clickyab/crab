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

func (pl *initPaymentPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "ChargeAmount":
			res["charge_amount"] = trans.E("invalid value")

		case "GateWay":
			res["gate_way"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (pl *registerBankSnapPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "Amount":
			res["amount"] = trans.E("invalid value")

		case "TraceNumber":
			res["trace_number"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}
