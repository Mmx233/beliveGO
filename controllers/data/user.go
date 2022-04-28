package controllers

import (
	"errors"
	"github.com/Mmx233/beliveGO/controllers"
	"github.com/Mmx233/beliveGO/models/form"
	"github.com/Mmx233/tool"
	"github.com/gin-gonic/gin"
)

func Avatar(c *gin.Context) {
	var f form.Avatar
	if c.ShouldBind(&f) != nil {
		controllers.CallBack.Error(c, 1, nil)
		return
	}
	_, res, e := tool.HTTP.Get(&tool.GetRequest{
		Url:    "https://api.bilibili.com/x/space/acc/info",
		Header: nil,
		Query: map[string]interface{}{
			"mid": f.MID,
		},
		Redirect: true,
	})
	if e != nil {
		controllers.CallBack.Error(c, 5, e)
		return
	}
	if res["code"].(float64) != 0 {
		controllers.CallBack.Error(c, 2, errors.New(res["message"].(string)))
		return
	}

	controllers.CallBack.Success(c, res["data"].(map[string]interface{})["face"])
}
