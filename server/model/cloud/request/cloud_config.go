package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CloudConfigSearch struct{
    cloud.CloudConfig
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
