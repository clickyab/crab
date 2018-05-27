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

func (p *changeDomainStatusPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "DomainStatus":
			res["domain_status"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

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
		case "DomainBase":
			res["domain_base"] = trans.E("invalid value")

		case "Title":
			res["title"] = trans.E("invalid value")

		case "Email":
			res["email"] = trans.E("invalid value")

		case "Password":
			res["password"] = trans.E("invalid value")

		case "Company":
			res["company"] = trans.E("invalid value")

		case "Theme":
			res["theme"] = trans.E("invalid value")

		case "Logo":
			res["logo"] = trans.E("invalid value")

		case "Description":
			res["description"] = trans.E("invalid value")

		case "Attributes":
			res["attributes"] = trans.E("invalid value")

		case "Status":
			res["status"] = trans.E("invalid value")

		case "MinTotalBudget":
			res["min_total_budget"] = trans.E("invalid value")

		case "MinDailyBudget":
			res["min_daily_budget"] = trans.E("invalid value")

		case "MinWebNativeCPC":
			res["min_web_native_cpc"] = trans.E("invalid value")

		case "MinWebBannerCPC":
			res["min_web_banner_cpc"] = trans.E("invalid value")

		case "MinWebVastCPC":
			res["min_web_vast_cpc"] = trans.E("invalid value")

		case "MinAppNativeCPC":
			res["min_app_native_cpc"] = trans.E("invalid value")

		case "MinAppBannerCPC":
			res["min_app_banner_cpc"] = trans.E("invalid value")

		case "MinAppVastCPC":
			res["min_app_vast_cpc"] = trans.E("invalid value")

		case "MinWebCPC":
			res["min_web_cpc"] = trans.E("invalid value")

		case "MinAppCPC":
			res["min_app_cpc"] = trans.E("invalid value")

		case "MinWebNativeCPM":
			res["min_web_native_cpm"] = trans.E("invalid value")

		case "MinWebBannerCPM":
			res["min_web_banner_cpm"] = trans.E("invalid value")

		case "MinWebVastCPM":
			res["min_web_vast_cpm"] = trans.E("invalid value")

		case "MinAppNativeCPM":
			res["min_app_native_cpm"] = trans.E("invalid value")

		case "MinAppBannerCPM":
			res["min_app_banner_cpm"] = trans.E("invalid value")

		case "MinAppVastCPM":
			res["min_app_vast_cpm"] = trans.E("invalid value")

		case "MinWebCPM":
			res["min_web_cpm"] = trans.E("invalid value")

		case "MinAppCPM":
			res["min_app_cpm"] = trans.E("invalid value")

		case "Advantage":
			res["advantage"] = trans.E("invalid value")

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
