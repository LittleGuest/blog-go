package handler

import (
	"blog/consts"
	"blog/model"
	"blog/model/param"
	"blog/model/result"
	"blog/repo"
	"blog/repo/db"
	"blog/response"
	"blog/tool"
	"crypto/sha1"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

// IsInstall check installation status
// @title check installation status
// @router /api/v1/admin/is_installed [get]
func IsInstall(ctx *gin.Context) {
	option := repo.GetPropertyByKey("is_installed")
	var isInstall bool
	if option != nil && option.OptionValue == "true" {
		isInstall = true
	}
	response.Return(ctx, http.StatusOK, http.StatusOK, "", isInstall)
}

// SendResetPasswordCode send reset password verify code
// @title send reset password verify code
// @router /api/v1/admin/password/code [post]
func SendResetPasswordCode(ctx *gin.Context) {
	param := &param.RestPasswordParam{}
	_ = ctx.Bind(param)
	valid := validate.Struct(param)
	if !valid.Validate() {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "", valid.Errors)
		return
	}

	code, err := db.GetRedisClient().Get("code").Result()
	if err != nil {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "获取验证码失败", nil)
		return
	}
	if tool.IsNotBlank(code) {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "已经获取过验证码，不能重复获取", nil)
		return
	}
	if !repo.VerifyUser(param.Account, param.Email) {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "用户名或者邮箱验证错误", nil)
		return
	}

	code = tool.RandomNumbersToString(6)
	db.GetRedisClient().Set("code", code, time.Minute*5)

	option := repo.GetPropertyByKey("email_enabled")
	if option == nil || option.OptionValue == "false" {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError,
			"未启用 SMTP 服务，无法发送邮件，但是你可以通过系统日志找到验证码", nil)
		return
	}

	content := fmt.Sprintf("您正在进行密码重置操作，如不是本人操作，请尽快做好相应措施。密码重置验证码如下（五分钟有效）：\n%v", code)
	if err := tool.SendTextEmail("2190975784@qq.com", "找回密码验证码", content); err != nil {
		response.Return(ctx, http.StatusOK, consts.StatusEmailSendFail, err.Error(), nil)
		return
	}
}

// ResetPasswordByCode reset password by verify code
// @title reset password by verify code
// @router /api/v1/admin/password/reset [put]
func ResetPasswordByCode(ctx *gin.Context) {
	param := &param.RestPasswordParam{}
	_ = ctx.Bind(param)
	valid := validate.Struct(param)
	if !valid.Validate() {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "", valid.Errors)
		return
	}

	if tool.IsBlank(param.Code) {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "验证码不能为空", nil)
		return
	}
	if tool.IsBlank(param.Password) {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "密码不能为空", nil)
		return
	}
	if !repo.VerifyUser(param.Account, param.Email) {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "用户名或者邮箱验证错误", nil)
		return
	}

	code, err := db.GetRedisClient().Get("code").Result()
	if err != nil {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "获取验证码失败", nil)
		return
	}
	if tool.IsBlank(code) {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "未获取过验证码", nil)
	}

	if code != param.Code {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "验证码不正确", nil)
	}

	user := repo.GetCurrentUser()
	if user == nil {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "未查询到博主信息", nil)
	}

	s1 := sha1.New()
	_, _ = s1.Write([]byte(param.Password))
	user.Password = fmt.Sprintf("%x", s1.Sum(nil))
	db.GetMDB().Model(user).Update("password", user.Password)

	db.GetRedisClient().Del("code")
}

// GetCount get count info
// @title get count info
// @router /api/v1/admin/counts [get]
func GetCount(ctx *gin.Context) {
	count := result.StatisticResult{}
	count.PostCount = repo.CountPostByStatus(consts.PostStatusPublished) // TODO
	count.AttachmentCount = repo.CountAttachment()

	comments := repo.CountCommentByStatus(consts.CommentStatusPublished)
	// TODO
	journals := repo.CountJournalByStatus(consts.JournalStatusPublished)
	count.CommentCount = comments + journals

	birthday := repo.GetBirthday()
	count.Birthday = birthday
	count.EstablishDays = (time.Now().Unix() - birthday) / (60 * 60 * 24)

	count.LinkCount = repo.CountLink()

	count.VisitCount = repo.CountPostVisit()
	// TODO
	count.LikeCount = repo.CountPostLike()
	// TODO

	response.Return(ctx, http.StatusOK, http.StatusOK, "", count)
}

// GetEnvironments get environment
// @title get count info
// @router /api/v1/admin/environments [get]
func GetEnvironments(ctx *gin.Context) {
	env := result.EnvironmentResult{}
	env.StartTime = 0 // TODO
	env.Database = "mysql"
	env.Version = consts.BlogVersion
	response.Return(ctx, http.StatusOK, http.StatusOK, "", env)
}

// InstallBlog install blog
// @title install blog
// @router /api/v1/admin/installations [post]
func InstallBlog(ctx *gin.Context) {
	param := &param.InstallParam{}
	_ = ctx.Bind(param)
	valid := validate.Struct(param)
	if !valid.Validate() {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "", valid.Errors)
		return
	}
	if tool.IsBlank(param.Password) {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "密码为空", nil)
		return
	}

	isInstalled := repo.GetPropertyByKey("is_installed")
	var isInstall bool
	if isInstalled != nil && isInstalled.OptionValue == "true" {
		isInstall = true
	}
	if isInstall {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "该博客已初始化，不能再次安装！", nil)
		return
	}

	properties := map[string]string{}
	properties["is_installed"] = "true"
	properties["blog_locale"] = param.Locale
	properties["blog_title"] = param.Title
	var blogURL string
	if tool.IsBlank(param.Url) {
		blogURL = repo.GetBlogBaseURL()
	} else {
		blogURL = param.Url
	}
	properties["blog_url"] = blogURL
	properties["birthday"] = strconv.FormatInt(repo.GetBirthday(), 10)
	gape := repo.GetPropertyByKey("global_absolute_path_enabled")
	if gape == nil {
		properties["global_absolute_path_enabled"] = "false"
	}

	tx := db.GetMDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for key, value := range properties {
		o := repo.GetPropertyByKey(key)
		if o == nil || tool.IsBlank(o.OptionKey) {
			o = &model.Option{
				OptionKey:   key,
				OptionValue: value,
			}
			if err := tx.Create(o).Error; err != nil {
				tx.Rollback()
				response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "安装失败", err)
				return
			}
		} else {
			if value != o.OptionValue {
				if err := tx.Model(&model.Option{}).Where("option_key = ?", key).Update("option_value", value); err != nil {
					tx.Rollback()
					response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "安装失败", err)
					return
				}
			}
		}
	}
	if err := tx.Commit().Error; err != nil {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "安装失败", nil)
		return
	}
}
