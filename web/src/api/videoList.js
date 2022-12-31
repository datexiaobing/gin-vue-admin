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

// active videos list
export const getVideoListListActive = (params) => {
  return service({
    url: '/videoList/getVideoListListActive',
    method: 'get',
    params
  })
}
// waiting videos list
export const getVideoListListWaiting = (params) => {
  return service({
    url: '/videoList/getVideoListListWaiting',
    method: 'get',
    params
  })
}
// stop videos list
export const getVideoListListStop = (params) => {
  return service({
    url: '/videoList/getVideoListListStop',
    method: 'get',
    params
  })
}
// 暂停下载
export const getVideoListListPause = (data) => {
  return service({
    url: '/videoList/getVideoListListPause',
    method: 'POST',
    data
  })
}
// 恢复下载
export const getVideoListListUnpause = (data) => {
  return service({
    url: '/videoList/getVideoListListUnpause',
    method: 'POST',
    data
  })
}
// 恢复下载
export const getVideoListListRemove = (data) => {
  return service({
    url: '/videoList/getVideoListListRemove',
    method: 'POST',
    data
  })
}

// file list
export const getVideoListListFile = (params) => {
  return service({
    url: '/videoList/getVideoListListFile',
    method: 'get',
    params
  })
}

// remove file
export const renameFile = (data) => {
  return service({
    url: '/videoList/renameFile',
    method: 'put',
    data
  })
}

// file list removed
export const getVideoListListFileDone = (params) => {
  return service({
    url: '/videoList/getVideoListListFileDone',
    method: 'get',
    params
  })
}
// 文件重名名
export const changeFileName = (data) => {
  return service({
    url: '/videoList/changeFileName',
    method: 'post',
    data
  })
}
// 删除文件
export const deleteFile = (data) => {
  return service({
    url: '/videoList/deleteFile',
    method: 'post',
    data
  })
}


// videos staus
export const getVideosStatus = (params) => {
  return service({
    url: '/videoList/getVideosStatus',
    method: 'get',
    params
  })
}


// down m3u8
export const downM3u8 = (data) => {
  return service({
    url: '/videoList/downM3u8',
    method: 'post',
    data
  })
}