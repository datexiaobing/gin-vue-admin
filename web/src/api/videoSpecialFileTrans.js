import service from '@/utils/request'

// @Tags VideoSpecialFileTrans
// @Summary 创建VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoSpecialFileTrans true "创建VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecialFileTrans/createVideoSpecialFileTrans [post]
export const createVideoSpecialFileTrans = (data) => {
  return service({
    url: '/videoSpecialFileTrans/createVideoSpecialFileTrans',
    method: 'post',
    data
  })
}

// @Tags VideoSpecialFileTrans
// @Summary 删除VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoSpecialFileTrans true "删除VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoSpecialFileTrans/deleteVideoSpecialFileTrans [delete]
export const deleteVideoSpecialFileTrans = (data) => {
  return service({
    url: '/videoSpecialFileTrans/deleteVideoSpecialFileTrans',
    method: 'delete',
    data
  })
}

// @Tags VideoSpecialFileTrans
// @Summary 删除VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoSpecialFileTrans/deleteVideoSpecialFileTrans [delete]
export const deleteVideoSpecialFileTransByIds = (data) => {
  return service({
    url: '/videoSpecialFileTrans/deleteVideoSpecialFileTransByIds',
    method: 'delete',
    data
  })
}

// @Tags VideoSpecialFileTrans
// @Summary 更新VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoSpecialFileTrans true "更新VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoSpecialFileTrans/updateVideoSpecialFileTrans [put]
export const updateVideoSpecialFileTrans = (data) => {
  return service({
    url: '/videoSpecialFileTrans/updateVideoSpecialFileTrans',
    method: 'put',
    data
  })
}

// @Tags VideoSpecialFileTrans
// @Summary 用id查询VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VideoSpecialFileTrans true "用id查询VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoSpecialFileTrans/findVideoSpecialFileTrans [get]
export const findVideoSpecialFileTrans = (params) => {
  return service({
    url: '/videoSpecialFileTrans/findVideoSpecialFileTrans',
    method: 'get',
    params
  })
}

// @Tags VideoSpecialFileTrans
// @Summary 分页获取VideoSpecialFileTrans列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取VideoSpecialFileTrans列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecialFileTrans/getVideoSpecialFileTransList [get]
export const getVideoSpecialFileTransList = (params) => {
  return service({
    url: '/videoSpecialFileTrans/getVideoSpecialFileTransList',
    method: 'get',
    params
  })
}
