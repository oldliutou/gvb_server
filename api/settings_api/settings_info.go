package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (SettingsApi) SettingInfoView(c *gin.Context) {
	res.OK(map[string]string{},"xxxx",c)
}