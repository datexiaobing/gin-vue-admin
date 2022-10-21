import service from '@/utils/request'

// @Tags FileTrans
// @Summary 创建FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileTrans true "创建FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileTrans/createFileTrans [post]
export const createFileTrans = (data) => {
  return service({
    url: '/fileTrans/createFileTrans',
    method: 'post',
    data
  })
}

// @Tags FileTrans
// @Summary 删除FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileTrans true "删除FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileTrans/deleteFileTrans [delete]
export const deleteFileTrans = (data) => {
  return service({
    url: '/fileTrans/deleteFileTrans',
    method: 'delete',
    data
  })
}

// @Tags FileTrans
// @Summary 删除FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileTrans/deleteFileTrans [delete]
export const deleteFileTransByIds = (data) => {
  return service({
    url: '/fileTrans/deleteFileTransByIds',
    method: 'delete',
    data
  })
}

// @Tags FileTrans
// @Summary 更新FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileTrans true "更新FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileTrans/updateFileTrans [put]
export const updateFileTrans = (data) => {
  return service({
    url: '/fileTrans/updateFileTrans',
    method: 'put',
    data
  })
}

// @Tags FileTrans
// @Summary 用id查询FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileTrans true "用id查询FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileTrans/findFileTrans [get]
export const findFileTrans = (params) => {
  return service({
    url: '/fileTrans/findFileTrans',
    method: 'get',
    params
  })
}

// @Tags FileTrans
// @Summary 分页获取FileTrans列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FileTrans列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileTrans/getFileTransList [get]
export const getFileTransList = (params) => {
  return service({
    url: '/fileTrans/getFileTransList',
    method: 'get',
    params
  })
}
