package controllers

import (
	"errors"
	"github.com/Mmx233/beliveGO/controllers"
	"github.com/Mmx233/beliveGO/models/form"
	"github.com/Mmx233/beliveGO/modules/RateLimit"
	"github.com/Mmx233/beliveGO/modules/cache"
	"github.com/Mmx233/tool"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Avatar(c *gin.Context) {
	var f form.Avatar
	if c.ShouldBind(&f) != nil {
		controllers.CallBack.Error(c, 1, nil)
		return
	}
	url, e0 := cache.Avatar.Read(f.UID)
	if e0 == nil {
		controllers.CallBack.Success(c, url)
		return
	}

	<-RateLimit.BiApi
	if len(RateLimit.BiApi) <= 1 {
		url, e0 = cache.Avatar.Read(f.UID)
		if e0 == nil {
			controllers.CallBack.Success(c, url)
			return
		}
	}

	_, res, e := tool.HTTP.Get(&tool.GetRequest{
		Url:    "https://api.bilibili.com/x/space/acc/info",
		Header: nil,
		Query: map[string]interface{}{
			"mid": f.UID,
		},
	})
	if e != nil {
		controllers.CallBack.Error(c, 5, e)
		return
	}
	if res["code"].(float64) != 0 {
		controllers.CallBack.Error(c, 2, errors.New(res["message"].(string)))
		return
	}
	url = strings.TrimPrefix(res["data"].(map[string]interface{})["face"].(string), "http:")
	url = strings.TrimPrefix(url, "https:")
	if e0 == cache.Nil {
		_ = cache.Avatar.Cache(f.UID, url, time.Hour*24*3)
	}
	controllers.CallBack.Success(c, url)
}
