import service from '@/utils/request'

// @Tags CloudPost
// @Summary 创建CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CloudPost true "创建CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudPost/createCloudPost [post]
export const createCloudPost = (data) => {
  return service({
    url: '/cloudPost/createCloudPost',
    method: 'post',
    data
  })
}

// @Tags CloudPost
// @Summary 删除CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CloudPost true "删除CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloudPost/deleteCloudPost [delete]
export const deleteCloudPost = (data) => {
  return service({
    url: '/cloudPost/deleteCloudPost',
    method: 'delete',
    data
  })
}

// @Tags CloudPost
// @Summary 删除CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloudPost/deleteCloudPost [delete]
export const deleteCloudPostByIds = (data) => {
  return service({
    url: '/cloudPost/deleteCloudPostByIds',
    method: 'delete',
    data
  })
}

// @Tags CloudPost
// @Summary 更新CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CloudPost true "更新CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cloudPost/updateCloudPost [put]
export const updateCloudPost = (data) => {
  return service({
    url: '/cloudPost/updateCloudPost',
    method: 'put',
    data
  })
}

// @Tags CloudPost
// @Summary 用id查询CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CloudPost true "用id查询CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cloudPost/findCloudPost [get]
export const findCloudPost = (params) => {
  return service({
    url: '/cloudPost/findCloudPost',
    method: 'get',
    params
  })
}

// @Tags CloudPost
// @Summary 分页获取CloudPost列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CloudPost列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudPost/getCloudPostList [get]
export const getCloudPostList = (params) => {
  return service({
    url: '/cloudPost/getCloudPostList',
    method: 'get',
    params
  })
}
