import service from '@/utils/request'

// @Tags FileInput
// @Summary 创建FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileInput true "创建FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileInput/createFileInput [post]
export const createFileInput = (data) => {
  return service({
    url: '/fileInput/createFileInput',
    method: 'post',
    data
  })
}

// @Tags FileInput
// @Summary 删除FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileInput true "删除FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileInput/deleteFileInput [delete]
export const deleteFileInput = (data) => {
  return service({
    url: '/fileInput/deleteFileInput',
    method: 'delete',
    data
  })
}

// @Tags FileInput
// @Summary 删除FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileInput/deleteFileInput [delete]
export const deleteFileInputByIds = (data) => {
  return service({
    url: '/fileInput/deleteFileInputByIds',
    method: 'delete',
    data
  })
}

// @Tags FileInput
// @Summary 更新FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileInput true "更新FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileInput/updateFileInput [put]
export const updateFileInput = (data) => {
  return service({
    url: '/fileInput/updateFileInput',
    method: 'put',
    data
  })
}

// @Tags FileInput
// @Summary 用id查询FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileInput true "用id查询FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileInput/findFileInput [get]
export const findFileInput = (params) => {
  return service({
    url: '/fileInput/findFileInput',
    method: 'get',
    params
  })
}

// @Tags FileInput
// @Summary 分页获取FileInput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FileInput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileInput/getFileInputList [get]
export const getFileInputList = (params) => {
  return service({
    url: '/fileInput/getFileInputList',
    method: 'get',
    params
  })
}
