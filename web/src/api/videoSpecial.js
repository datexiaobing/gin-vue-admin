import service from '@/utils/request'

// @Tags VideoSpecial
// @Summary 创建VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoSpecial true "创建VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecial/createVideoSpecial [post]
export const createVideoSpecial = (data) => {
  return service({
    url: '/videoSpecial/createVideoSpecial',
    method: 'post',
    data
  })
}

// @Tags VideoSpecial
// @Summary 删除VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoSpecial true "删除VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoSpecial/deleteVideoSpecial [delete]
export const deleteVideoSpecial = (data) => {
  return service({
    url: '/videoSpecial/deleteVideoSpecial',
    method: 'delete',
    data
  })
}

// @Tags VideoSpecial
// @Summary 删除VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoSpecial/deleteVideoSpecial [delete]
export const deleteVideoSpecialByIds = (data) => {
  return service({
    url: '/videoSpecial/deleteVideoSpecialByIds',
    method: 'delete',
    data
  })
}

// @Tags VideoSpecial
// @Summary 更新VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoSpecial true "更新VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoSpecial/updateVideoSpecial [put]
export const updateVideoSpecial = (data) => {
  return service({
    url: '/videoSpecial/updateVideoSpecial',
    method: 'put',
    data
  })
}

// @Tags VideoSpecial
// @Summary 用id查询VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VideoSpecial true "用id查询VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoSpecial/findVideoSpecial [get]
export const findVideoSpecial = (params) => {
  return service({
    url: '/videoSpecial/findVideoSpecial',
    method: 'get',
    params
  })
}

// @Tags VideoSpecial
// @Summary 分页获取VideoSpecial列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取VideoSpecial列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecial/getVideoSpecialList [get]
export const getVideoSpecialList = (params) => {
  return service({
    url: '/videoSpecial/getVideoSpecialList',
    method: 'get',
    params
  })
}
