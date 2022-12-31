package cloud

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type CloudPostService struct {
}

// CreateCloudPost 创建CloudPost记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudPostService *CloudPostService) CreateCloudPost(cloudPost cloud.CloudPost) (err error) {

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 获取transFile 中 transPost =1的所有视频，挨个推送，推送成功的修改transPost=2
		var files []cloud.FileTrans
		// 获取所有未推送的视频
		err = tx.Where("trans_post = ?", 1).Find(&files).Error
		if err != nil {
			return err
		}
		var allNUm int
		var successNum int
		for _, v := range files {
			var body utils.PostBody
			if v.TransOssQiniuStatus == 3 {
				// qiniu已经上传
				body.OssId = 1
				allNUm += 1
			} else if v.TransOssAliStatus == 3 {
				// ali已经上传
				body.OssId = 2
				allNUm += 1
			} else {
				// 没有上传到云端，不推送
				fmt.Println("not push", v.TransUuid)
				continue
			}

			body.Title = v.TransOutName
			body.Uuid = v.TransUuid
			body.ClassifyId = v.TransType //分类ID
			body.Duration = v.TransDuration
			// var special cloud.VideoSpecialFileTrans
			// err = tx.Where("file_trans_id =?", v.ID).First(&special).Error
			// if err != nil {
			// 	return errors.New("no special record")
			// }
			if v.TransTypeNum < 1 {
				return errors.New("no special record")
			}
			body.SpecialId = v.TransTypeNum //专辑ID

			res, err := json.Marshal(body)
			if err != nil {
				return err
			}

			byte_body := []byte(string(res))

			// Post("http://xxxx","application/json;charset=utf-8",[]byte("{'aaa':'bbb'}"))
			mess, err := utils.GoPost("http://127.0.0.1:3000/updateVideo",
				"application/json;charset=utf-8", byte_body)

			if err != nil {
				return err
			}

			var message utils.Response
			err = json.Unmarshal([]byte(mess), &message)
			if err != nil {
				return err
			}
			// fmt.Println(message.Code, err)
			if message.Code == 0 {
				// 推送成功，
				fmt.Println("push success:", v.TransUuid)
				v.TransPost = 2
				tx.Save(&v)
				successNum += 1
			}
		}
		// return nil
		if allNUm > 0 {
			cloudPost.PostNum = allNUm
			cloudPost.SuccessNum = successNum
			err = tx.Create(&cloudPost).Error
			return err
		}
		return nil

	})
	return err
}

// DeleteCloudPost 删除CloudPost记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudPostService *CloudPostService) DeleteCloudPost(cloudPost cloud.CloudPost) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.CloudPost{}).Where("id = ?", cloudPost.ID).Update("deleted_by", cloudPost.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&cloudPost).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCloudPostByIds 批量删除CloudPost记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudPostService *CloudPostService) DeleteCloudPostByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.CloudPost{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.CloudPost{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCloudPost 更新CloudPost记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudPostService *CloudPostService) UpdateCloudPost(cloudPost cloud.CloudPost) (err error) {
	err = global.GVA_DB.Save(&cloudPost).Error
	return err
}

// GetCloudPost 根据id获取CloudPost记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudPostService *CloudPostService) GetCloudPost(id uint) (cloudPost cloud.CloudPost, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cloudPost).Error
	return
}

// GetCloudPostInfoList 分页获取CloudPost记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudPostService *CloudPostService) GetCloudPostInfoList(info cloudReq.CloudPostSearch) (list []cloud.CloudPost, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cloud.CloudPost{})
	var cloudPosts []cloud.CloudPost
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cloudPosts).Error
	return cloudPosts, total, err
}
