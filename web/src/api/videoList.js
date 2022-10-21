import service from '@/utils/request'

// @Tags VideoList
// @Summary 创建VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoList true "创建VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoList/createVideoList [post]
export const createVideoList = (data) => {
  return service({
    url: '/videoList/createVideoList',
    method: 'post',
    data
  })
}

// @Tags VideoList
// @Summary 删除VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoList true "删除VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoList/deleteVideoList [delete]
export const deleteVideoList = (data) => {
  return service({
    url: '/videoList/deleteVideoList',
    method: 'delete',
    data
  })
}

// @Tags VideoList
// @Summary 删除VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoList/deleteVideoList [delete]
export const deleteVideoListByIds = (data) => {
  return service({
    url: '/videoList/deleteVideoListByIds',
    method: 'delete',
    data
  })
}

// @Tags VideoList
// @Summary 更新VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoList true "更新VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoList/updateVideoList [put]
export const updateVideoList = (data) => {
  return service({
    url: '/videoList/updateVideoList',
    method: 'put',
    data
  })
}

// @Tags VideoList
// @Summary 用id查询VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VideoList true "用id查询VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoList/findVideoList [get]
export const findVideoList = (params) => {
  return service({
    url: '/videoList/findVideoList',
    method: 'get',
    params
  })
}

// @Tags VideoList
// @Summary 分页获取VideoList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取VideoList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoList/getVideoListList [get]
export const getVideoListList = (params) => {
  return service({
    url: '/videoList/getVideoListList',
    method: 'get',
    params
  })
}
