import service from '@/utils/request'

// @Tags CloudConfig
// @Summary 创建CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CloudConfig true "创建CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudConfig/createCloudConfig [post]
export const createCloudConfig = (data) => {
  return service({
    url: '/cloudConfig/createCloudConfig',
    method: 'post',
    data
  })
}

// @Tags CloudConfig
// @Summary 删除CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CloudConfig true "删除CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloudConfig/deleteCloudConfig [delete]
export const deleteCloudConfig = (data) => {
  return service({
    url: '/cloudConfig/deleteCloudConfig',
    method: 'delete',
    data
  })
}

// @Tags CloudConfig
// @Summary 删除CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloudConfig/deleteCloudConfig [delete]
export const deleteCloudConfigByIds = (data) => {
  return service({
    url: '/cloudConfig/deleteCloudConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags CloudConfig
// @Summary 更新CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CloudConfig true "更新CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cloudConfig/updateCloudConfig [put]
export const updateCloudConfig = (data) => {
  return service({
    url: '/cloudConfig/updateCloudConfig',
    method: 'put',
    data
  })
}

// @Tags CloudConfig
// @Summary 用id查询CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CloudConfig true "用id查询CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cloudConfig/findCloudConfig [get]
export const findCloudConfig = (params) => {
  return service({
    url: '/cloudConfig/findCloudConfig',
    method: 'get',
    params
  })
}

// @Tags CloudConfig
// @Summary 分页获取CloudConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CloudConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudConfig/getCloudConfigList [get]
export const getCloudConfigList = (params) => {
  return service({
    url: '/cloudConfig/getCloudConfigList',
    method: 'get',
    params
  })
}
