package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/cloud"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
type IdsShareReq struct {
	Ids     []int  `json:"ids" form:"ids"`
	Domain  string `form:"domain" json:"domain"`
	Expires string `json:"expires" form:"expires"`
	Ip      string `jsonL:"ip" form:"ip"`
}
type GridsReq struct {
	Grids []string `json:"grids" form:"grids"`
}

// move file
type FileReq struct {
	DownloadPath string `json:"downloadPath" form:"downloadPath"`
	FileName     string `jon:"fileName" form:"fileName"`
}

// trans video
type TransVideoReq struct {
	DownloadPath string `json:"downloadPath" form:"downloadPath"`
	FileName     string `jon:"fileName" form:"fileName"`
	cloud.FileTrans
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
