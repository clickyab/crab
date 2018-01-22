package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"clickyab.com/crab/modules/ad/add"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/xmux"
)

type bannerSize struct {
	Width  int
	Height int
}

var (
	bannerValidSizes = config.RegisterString("crab.modules.ad.banner",
		"120x600,160x600,300x250,336x280,468x60,728x90,120x240,320x50,800x440,300x600,970x90,970x250,250x250,300x1050,320x480,480x320,128x128",
		"valid banner sizes separated by x",
	)
	minNativeBannerRation = config.RegisterFloat64("crab.modules.ad.native.ratio.min", 1.55, "minimum native banner ration")
	maxNativeBannerRation = config.RegisterFloat64("crab.modules.ad.native.ratio.max", 1.65, "maximum native banner ration")
	lock                  = sync.Mutex{}

	bannerValid []bannerSize
)

// @Validate{
//}
type assignBannerPayload struct {
	Banners []struct {
		ID    int64  `json:"id,omitempty"`
		Src   string `json:"src" validate:"required"`
		Utm   string `json:"utm" validate:"required"`
		Title string `json:"title" validate:"required"`
	} `json:"banners"`
	input    []*add.Ad     `json:"-"`
	campaign *orm.Campaign `json:"-"`
	domain   *dmn.Domain   `json:"-"`
}

func (p *assignBannerPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	campaignIDInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return t9e.G("campaign id not valid")
	}
	bannerType := add.AdType(xmux.Param(ctx, "banner_type"))
	if bannerType != add.NativeAdType && bannerType != add.BannerAdType {
		return t9e.G("only native or banner is allowed")
	}
	cpManager := orm.NewOrmManager()
	d := domain.MustGetDomain(ctx)
	p.domain = d
	campaign, err := cpManager.FindCampaignByIDDomain(campaignIDInt, d.ID)
	if err != nil {
		return t9e.G("campaign not found")
	}
	p.campaign = campaign
	if string(p.campaign.Type) != string(bannerType) {
		return t9e.G("campaign is not the right type")
	}
	m := add.NewAddManager()
	if len(p.Banners) == 0 {
		return t9e.G("no banners selected")
	}
	for i := range p.Banners {
		mime, width, height, err := checkBannerImage(p.Banners[i].Src, bannerType)
		if err != nil {
			return err
		}
		if p.Banners[i].ID != 0 { //update selected
			//TODO check access for update banner
			bannerAd, err := m.FindAdByID(p.Banners[i].ID)
			if err != nil {
				return t9e.G("ad not found")
			}
			if bannerAd.CampaignID != campaign.ID {
				return t9e.G("ad not belong to campaign")
			}
			bannerAd.Mime = mime
			bannerAd.Height = height
			bannerAd.Width = width
			bannerAd.Src = p.Banners[i].Src
			bannerAd.Target = p.Banners[i].Utm
			bannerAd.Title = p.Banners[i].Title
			bannerAd.Status = add.PendingAdStatus
			bannerAd.Type = bannerType

			p.input = append(p.input, bannerAd)
		} else { //create selected
			//TODO check access for create banner
			newAd := &add.Ad{
				Target:     p.Banners[i].Utm,
				Src:        p.Banners[i].Src,
				CampaignID: campaign.ID,
				Width:      width,
				Height:     height,
				Mime:       mime,
				Title:      p.Banners[i].Title,
				Status:     add.PendingAdStatus,
				Type:       bannerType,
			}
			p.input = append(p.input, newAd)
		}
	}
	return nil
}

type adResponse []add.Ad

// assignNormalBanner assignNormalBanner module is banner type (banner/native)
// @Rest {
// 		url = /:banner_type/:id
//		protected = true
//		resource = assign_banner:self
// 		method = post
// }
func (c *Controller) assignNormalBanner(ctx context.Context, r *http.Request, p *assignBannerPayload) (adResponse, error) {
	currentUser := authz.MustGetUser(ctx)
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(p.campaign.UserID, p.domain.ID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "assign_banner", p.domain.ID)
	if !ok {
		return adResponse{}, framework.NewForbiddenError(t9e.G("don't have access for this action").Text)
	}

	res, err := add.NewAddManager().CreateUpdateCampaignNormalBanner(p.input)
	if err != nil {
		f, ok := err.(*mysql.MySQLError)
		if ok && f.Number == 1062 {
			return adResponse{}, t9e.G("duplicate src in ads")
		}
		return adResponse{}, t9e.G("cant create/update campaign")
	}
	return res, nil
}

func checkBannerDimension(width, height int, bannerType add.AdType) bool {
	switch bannerType {
	case add.BannerAdType:
		for i := range bannerValid {
			if bannerValid[i].Width == width && bannerValid[i].Height == height {
				return true
			}
		}
	case add.NativeAdType:
		rate := float64(float64(width) / float64(height))
		if rate < maxNativeBannerRation.Float64() && rate > minNativeBannerRation.Float64() {
			return true
		}

	}

	return false
}

func checkBannerImage(srcID string, bannerType add.AdType) (mime string, width int, height int, err error) {
	file, err := model.NewModelManager().FindUploadByID(srcID)

	if err != nil || (file.Attr.Banner == nil && file.Attr.Native == nil) {
		err = t9e.G("invalid uploaded file")
		return
	}
	//check banner type
	if file.Section != string(bannerType) {
		err = t9e.G("banner not selected")
		return
	}
	//TODO check access to file
	mime = file.MIME

	switch bannerType {
	case add.BannerAdType:
		width = file.Attr.Banner.Width
		height = file.Attr.Banner.Height
	case add.NativeAdType:
		width = file.Attr.Native.Width
		height = file.Attr.Native.Height
	}
	ok := checkBannerDimension(width, height, bannerType)
	if !ok {
		err = t9e.G("dimensions not valid")
		return
	}
	return
}

func init() {
	lock.Lock()
	defer lock.Unlock()
	// extract valid sizes
	mainArr := strings.Split(bannerValidSizes.String(), ",")
	for i := range mainArr {
		size := bannerSize{}
		sizeArr := strings.Split(mainArr[i], "x")
		assert.True(len(sizeArr) == 2)
		width, err := strconv.Atoi(sizeArr[0])
		assert.Nil(err)
		height, err := strconv.Atoi(sizeArr[1])
		assert.Nil(err)
		size.Width = width
		size.Height = height
		bannerValid = append(bannerValid, size)
	}
}
