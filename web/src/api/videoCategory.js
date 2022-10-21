import service from '@/utils/request'

// @Tags VideoCategory
// @Summary 创建VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoCategory true "创建VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoCategory/createVideoCategory [post]
export const createVideoCategory = (data) => {
  return service({
    url: '/videoCategory/createVideoCategory',
    method: 'post',
    data
  })
}

// @Tags VideoCategory
// @Summary 删除VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoCategory true "删除VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoCategory/deleteVideoCategory [delete]
export const deleteVideoCategory = (data) => {
  return service({
    url: '/videoCategory/deleteVideoCategory',
    method: 'delete',
    data
  })
}

// @Tags VideoCategory
// @Summary 删除VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoCategory/deleteVideoCategory [delete]
export const deleteVideoCategoryByIds = (data) => {
  return service({
    url: '/videoCategory/deleteVideoCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags VideoCategory
// @Summary 更新VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoCategory true "更新VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoCategory/updateVideoCategory [put]
export const updateVideoCategory = (data) => {
  return service({
    url: '/videoCategory/updateVideoCategory',
    method: 'put',
    data
  })
}

// @Tags VideoCategory
// @Summary 用id查询VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VideoCategory true "用id查询VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoCategory/findVideoCategory [get]
export const findVideoCategory = (params) => {
  return service({
    url: '/videoCategory/findVideoCategory',
    method: 'get',
    params
  })
}

// @Tags VideoCategory
// @Summary 分页获取VideoCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取VideoCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoCategory/getVideoCategoryList [get]
export const getVideoCategoryList = (params) => {
  return service({
    url: '/videoCategory/getVideoCategoryList',
    method: 'get',
    params
  })
}
