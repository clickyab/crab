// Code generated build with restful DO NOT EDIT.

package user

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// route for add/update user avatar
// @Route {
// 		url = /avatar
//		method = put
//		payload = avatarPayload
//		middleware = authz.Authenticate
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) avatarPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*avatarPayload)
	res, err := c.avatar(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// changeForgetPassword change forget password
// @Route {
// 		url = /password/change/:token
//		method = put
//		payload = callBackPayload
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) changeForgetPasswordPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*callBackPayload)
	res, err := c.changeForgetPassword(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// changePassword change password
// @Route {
// 		url = /password/change
//		method = put
//		payload = changePassword
//		middleware = authz.Authenticate
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) changePasswordPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*changePassword)
	res, err := c.changePassword(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// checkMail check mail in system
// @Route {
// 		url = /mail/check
//		method = post
//		payload = checkMailPayload
//		200 = checkMailResponse
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) checkMailPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*checkMailPayload)
	res, err := c.checkMail(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// edit route for edit personal profile
// @Route {
// 		url = /update
//		method = put
//		payload = userPayload
//		middleware = authz.Authenticate
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) editPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*userPayload)
	res, err := c.edit(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// forgetCallBack is the url coming from sent email
// @Route {
// 		url = /password/verify/
//		method = post
//		payload = forgetCodePayload
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) checkForgetCodePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*forgetCodePayload)
	res, err := c.checkForgetCode(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// forgetCallBack is the url coming from sent email
// @Route {
// 		url = /password/verify/:token
//		method = get
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) checkForgetHashGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.checkForgetHash(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// forgetPassword
// @Route {
// 		url = /password/forget
//		method = post
//		payload = forgetPayload
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) forgetPasswordPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*forgetPayload)
	res, err := c.forgetPassword(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// login userPayload in system
// @Route {
// 		url = /login
//		method = post
//		payload = loginPayload
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) loginPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*loginPayload)
	res, err := c.login(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// closeSession closes current session
// @Route {
// 		url = /logout
//		method = get
//		middleware = authz.Authenticate
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) closeSessionGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.closeSession(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// closeSession closes current session
// @Route {
// 		url = /logout/closeother
//		method = get
//		middleware = authz.Authenticate
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) closeAllOtherSessionGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.closeAllOtherSession(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// ping is for ping
// @Route {
// 		url = /ping
//		method = get
//		middleware = authz.Authenticate
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) pingGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.ping(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// register is for register user
// @Route {
// 		url = /register
//		method = post
//		payload = registerPayload
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) registerPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*registerPayload)
	res, err := c.register(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// searchByMail search user by email
// @Route {
// 		url = /search/mail
//		method = post
//		payload = searchUserPayload
//		middleware = authz.Authenticate
//		200 = userSearchResp
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) searchByMailPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*searchUserPayload)
	res, err := c.searchByMail(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// register is for register user
// @Route {
// 		url = /store
//		method = post
//		payload = storePayload
//		middleware = authz.Authenticate
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) storePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*storePayload)
	res, err := c.store(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// verifyEmail is verify code
// @Route {
// 		url = /email/verify/:token
//		method = get
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) verifyEmailGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.verifyEmail(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// verifyEmailCode is verify email
// @Route {
// 		url = /email/verify
//		method = post
//		payload = verifyEmailCodePayload
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) verifyEmailCodePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*verifyEmailCodePayload)
	res, err := c.verifyEmailCode(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// verifyResend will send an email again
// @Route {
// 		url = /email/verify/resend
//		method = post
//		payload = verifyResendPayload
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) verifyResendPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*verifyResendPayload)
	res, err := c.verifyResend(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
