<template>
  <div> 
    <el-container>
        <el-row>
            <el-button-group class="">
                <el-button icon="UploadFilled" disabled>{{ t('videoDownload.createDownLoad')}}</el-button>
                <el-button icon="VideoPlay" @click="reDown"> {{ t('videoDownload.begin')}}</el-button>
                <el-button icon="VideoPause" disabled> {{ t('videoDownload.pause')}}</el-button>
                <el-button icon="Delete" @click="remove"> {{ t('videoDownload.delete')}}</el-button>
            </el-button-group>
        </el-row>
    </el-container>
        <el-table
        border
        max-height="600"
        ref="multipleTable"
        header-cell-class-name="head-class"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" :label="t('videoDownload.fileName')"  >
          <template #default="scope">
            {{scope.row.fileName}}
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('videoDownload.fileSize')" width="120" >
            <template #default="scope">
                {{bytesToSize(scope.row.fileLength)}}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('videoDownload.progress')" width="200" >
            <template #default="scope">
                <div class="pro-cell">
                    <el-progress 
                    color="#67c23a"
                    :format="barformat"
                    :stroke-width="10" 
                    :percentage="(scope.row.completedLength/scope.row.fileLength).toFixed(2)*100 || 0" />
                </div>
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('videoDownload.speed')" prop="videoTitle" width="120" >
            <template #default="scope">
                <div class="speed-cell">
                    <div class="in-cell">
                        <el-icon><Top /></el-icon>
                        <span>{{scope.row.downloadSpeed}} b/s</span>
                    </div>
                     <div>
                        <el-icon><Bottom /></el-icon>
                        <span>{{scope.row.uploadSpeed}} b/s</span>
                    </div>
                </div>
            </template>
        </el-table-column>

        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
  </div>
    
</template>

<script setup>
import {
  createVideoList,
  deleteVideoList,
  deleteVideoListByIds,
  updateVideoList,
  findVideoList,
  getVideoListList,
  getVideoListListWaiting,
  getVideoListListUnpause,
  getVideoListListRemove,
} from '@/api/videoList'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,bytesToSize } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive,onBeforeUnmount } from 'vue'
import {start,close} from '@/utils/npgress'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

let timer =ref(null)
// 设置查询定时器
const setSearch =()=>{
  timer = setInterval(() => {
      //需要定时执行的代码
      start()
      getTableData()
      close()
  },5000)

}
setSearch()

onBeforeUnmount(()=>{
  clearInterval(timer)
})


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}


// 删除任务
const remove=async()=>{
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的任务'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.gid)
    })
  // console.log(ids)
  const res = await  getVideoListListRemove({ grids:ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    getTableData()
  }
}

// 恢复任务
const reDown= async()=>{
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要恢复下载的视频'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.gid)
        })
      // console.log(ids)
      const res = await getVideoListListUnpause({ grids:ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '恢复成功'
        })
        getTableData()
      }
}

const barformat =  (percentage) => (percentage === 100 ? '100%' : `${percentage}%`)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        videoTitle: '',
        videoDownloadPath: '',
        videoUrl: '',
        videoPic: '',
        videoDownloadStatus: 0,
        videoSize: 0,
        videoDownloadSize: 0,
        videoDownloadSpeed: 0,
        })

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  searchInfo.value.videoDownloadStatus = 2
  const table = await getVideoListListWaiting({ page: page.value, pageSize: pageSize.value})
  if (table.code === 0) {
    tableData.value = table.data.list
    tableData.value.forEach(e=>{
      // 文件名
      let dir = e.dir
      let path = e.files[0].path
      e.fileLength = e.files[0].length
      e.completedLength =e.files[0].completedLength
      let temp = path.replace(dir,"").split('/')
      // console.log(temp)
      e.fileName=path

    })

    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()




// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteVideoListFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteVideoListByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateVideoListFunc = async(row) => {
    const res = await findVideoList({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.revideoList
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteVideoListFunc = async (row) => {
    const res = await deleteVideoList({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        videoTitle: '',
        videoDownloadPath: '',
        videoUrl: '',
        videoPic: '',
        videoDownloadStatus: 0,
        videoSize: 0,
        videoDownloadSize: 0,
        videoDownloadSpeed: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createVideoList(formData.value)
                  break
                case 'update':
                  res = await updateVideoList(formData.value)
                  break
                default:
                  res = await createVideoList(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}



</script>

<style scoped>

.pro-cell{
    padding-top: 5%;
}

.speed-cell{
    display: flex;
    flex-direction: column;
}

.head-class{
    font-weight: bold;
    font-size: 20px; 
}
</style>