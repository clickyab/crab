package advalidator

import (
	"clickyab.com/crab/modules/ad/controllers/advalidator/errors"
	"github.com/clickyab/services/array"
)

// Check validate ad creative
func (v *ValidatorAd) Check(i InputData) error {
	err := checkDimension(v, i)
	if err != nil {
		return err
	}

	err = checkDuration(v, i)
	if err != nil {
		return err
	}

	err = checkSize(v, i)
	if err != nil {
		return err
	}

	return checkExtension(v, i)
}

func checkDuration(v *ValidatorAd, i InputData) error {
	if v.Duration != 0 {
		if i.Duration > v.Duration {
			return errors.AdDurationErr
		}
	}
	return nil
}

func checkSize(v *ValidatorAd, i InputData) error {
	if v.Size != 0 {
		if float64(i.Size)/float64(1024) > float64(v.Size) {
			return errors.AdSizeErr
		}
	}
	return nil
}

func checkExtension(v *ValidatorAd, i InputData) error {
	if len(v.Ext) > 0 {
		if !array.StringInArray(i.Ext, v.Ext...) {
			return errors.AdExtensionDimensionErr
		}
	}
	return nil
}

func checkDimension(v *ValidatorAd, i InputData) error {
	if v.Rate != 0 {
		if i.Width/i.Height != v.Rate {
			return errors.AdRateDimensionErr
		}
	}
	if len(v.Exact) > 0 {
		found := false
		for j := range v.Exact {
			if v.Exact[j].Width == i.Width && v.Exact[j].Height == i.Height {
				found = true
				break
			}
		}
		if !found {
			return errors.AdExactDimensionErr
		}
	}
	if v.Min != nil {
		if i.Width < v.Min.Width || i.Height < v.Min.Height {
			return errors.AdMinDimensionErr
		}
	}

	if v.Max != nil {
		if i.Width > v.Max.Width || i.Height > v.Max.Height {
			return errors.AdMaxDimensionErr
		}
	}
	return nil
}

// AdTotalValidationRule AdTotalValidationRule
type AdTotalValidationRule struct {
	FAppNative *AppNative `json:"app_native,omitempty"`
	FWebNative *WebNative `json:"web_native,omitempty"`
	FAppBanner *AppBanner `json:"app_banner,omitempty"`
	FWebBanner *WebBanner `json:"web_banner,omitempty"`
}

// AppNative AppNative
type AppNative struct {
	Logo   ValidatorAd `json:"logo,omitempty"`
	Icon   ValidatorAd `json:"icon,omitempty"`
	VImage ValidatorAd `json:"v_image,omitempty"`
	HImage ValidatorAd `json:"h_image,omitempty"`
	Video  ValidatorAd `json:"video,omitempty"`
}

// AppBanner AppBanner
type AppBanner struct {
	Video      ValidatorAd `json:"video,omitempty"`
	Image      ValidatorAd `json:"image,omitempty"`
	VideoImage ValidatorAd `json:"video_image,omitempty"`
}

// WebBanner WebBanner
type WebBanner struct {
	Video      ValidatorAd `json:"video,omitempty"`
	Image      ValidatorAd `json:"image,omitempty"`
	VideoImage ValidatorAd `json:"video_image,omitempty"`
}

// WebNative WebNative
type WebNative struct {
	Logo   ValidatorAd `json:"logo,omitempty"`
	Icon   ValidatorAd `json:"icon,omitempty"`
	VImage ValidatorAd `json:"v_image,omitempty"`
	HImage ValidatorAd `json:"h_image,omitempty"`
	Video  ValidatorAd `json:"video,omitempty"`
}

// ValidatorAd ValidatorAd
type ValidatorAd struct {
	Rate     float64     `json:"rate,omitempty"`
	Min      *Dimension  `json:"min,omitempty"`
	Max      *Dimension  `json:"max,omitempty"`
	Exact    []Dimension `json:"exact,omitempty"`
	Ext      []string    `json:"ext,omitempty"`
	Duration int64       `json:"duration,omitempty"`
	Size     int64       `json:"size,omitempty"`
}

// Dimension Dimension
type Dimension struct {
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
}
