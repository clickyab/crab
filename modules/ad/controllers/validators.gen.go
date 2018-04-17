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

func (pl *NativeAssetPayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "Title":
			res["title"] = trans.E("invalid value")

		case "Description":
			res["description"] = trans.E("invalid value")

		case "CTA":
			res["cta"] = trans.E("invalid value")

		case "Icon":
			res["icon"] = trans.E("invalid value")

		case "Images":
			res["images"] = trans.E("invalid value")

		case "Video":
			res["video"] = trans.E("invalid value")

		case "Logo":
			res["logo"] = trans.E("invalid value")

		case "Rating":
			res["rating"] = trans.E("invalid value")

		case "Price":
			res["price"] = trans.E("invalid value")

		case "SalePrice":
			res["sale_price"] = trans.E("invalid value")

		case "Downloads":
			res["downloads"] = trans.E("invalid value")

		case "Phone":
			res["phone"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (p *createNativePayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "CampaignID":
			res["campaign_id"] = trans.E("invalid value")

		case "URL":
			res["url"] = trans.E("invalid value")

		case "MaxBid":
			res["max_bid"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}

func (p *editNativePayload) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		case "URL":
			res["url"] = trans.E("invalid value")

		case "MaxBid":
			res["max_bid"] = trans.E("invalid value")

		default:
			logrus.Panicf("the field %s is not translated", i)
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}
