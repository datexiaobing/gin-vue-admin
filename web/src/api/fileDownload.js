import service from '@/utils/request'

// @Tags FileDownload
// @Summary 创建FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileDownload true "创建FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileDownload/createFileDownload [post]
export const createFileDownload = (data) => {
  return service({
    url: '/fileDownload/createFileDownload',
    method: 'post',
    data
  })
}

// @Tags FileDownload
// @Summary 删除FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileDownload true "删除FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileDownload/deleteFileDownload [delete]
export const deleteFileDownload = (data) => {
  return service({
    url: '/fileDownload/deleteFileDownload',
    method: 'delete',
    data
  })
}

// @Tags FileDownload
// @Summary 删除FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileDownload/deleteFileDownload [delete]
export const deleteFileDownloadByIds = (data) => {
  return service({
    url: '/fileDownload/deleteFileDownloadByIds',
    method: 'delete',
    data
  })
}

// @Tags FileDownload
// @Summary 更新FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileDownload true "更新FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileDownload/updateFileDownload [put]
export const updateFileDownload = (data) => {
  return service({
    url: '/fileDownload/updateFileDownload',
    method: 'put',
    data
  })
}

// @Tags FileDownload
// @Summary 用id查询FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileDownload true "用id查询FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileDownload/findFileDownload [get]
export const findFileDownload = (params) => {
  return service({
    url: '/fileDownload/findFileDownload',
    method: 'get',
    params
  })
}

// @Tags FileDownload
// @Summary 分页获取FileDownload列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FileDownload列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileDownload/getFileDownloadList [get]
export const getFileDownloadList = (params) => {
  return service({
    url: '/fileDownload/getFileDownloadList',
    method: 'get',
    params
  })
}
