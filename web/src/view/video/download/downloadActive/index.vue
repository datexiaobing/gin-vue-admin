<template>
  <div>
    <el-container>
        <el-row>
            <el-button-group class="">
              <el-space>
                <el-button icon="UploadFilled" @click="createDownLoad">{{ t('videoDownload.createDownLoad')}}</el-button>
                <el-button icon="VideoPlay" disabled> {{ t('videoDownload.begin')}}</el-button>
                <el-button icon="VideoPause" @click="pause"> {{ t('videoDownload.pause')}}</el-button>
                <el-button icon="Delete" @click="remove"> {{ t('videoDownload.delete')}}</el-button>
              </el-space>
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
            {{scope.row.path}}
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
                    :color="(scope.row.videoDownloadSize/scope.row.videoSize).toFixed(2)*100<100?'':'#67c23a'"
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
        
     <el-dialog v-model="createDown" :title="t('videoDownload.create')">
        <el-tabs v-model="activeName" class="demo-tabs" >
          <el-tab-pane :label="t('videoDownload.link')" name="first">
              <el-input
                v-model="textarea"
                :rows="10"
                type="textarea"
                :placeholder="t('videoDownload.linkDetail')"
              />
              <el-divider content-position="left">add headers</el-divider>
              <el-input
                v-model="textareaHeader"
                :rows="5"
                type="textarea"
                placeholder="referer: https://tv11.avsee.in/"
              />

          </el-tab-pane>
          <el-tab-pane :label="t('videoDownload.seedFile')" name="second">
            <el-upload
            :action="`${path}/api/videoList/getVideoListListUpload`"
            :headers="{ 'x-token': userStore.token }"
            :show-file-list="false"
            class="upload-btn"
            drag
          >
              <el-icon class="el-icon--upload"><upload-filled /></el-icon>
              <div class="el-upload__text">
                {{t('videoDownload.dragSeed')}} <em>{{t('videoDownload.uploadSeed')}}</em>
              </div>
            </el-upload>
          </el-tab-pane>

          <el-tab-pane label="m3u8" name="there">
              <el-input
                v-model="textarea3"
                :rows="10"
                type="textarea"
                placeholder="https://xxx.m3u8  , newName"
              />
          <!-- <el-divider content-position="left">add headers</el-divider>
          <el-input
            v-model="textareaHeader"
            :rows="5"
            type="textarea"
            placeholder="referer: https://tv11.avsee.in/"
          /> -->

          </el-tab-pane>
        </el-tabs>
      <template #footer>
        <el-button size="small" type="success" @click="downNow"> {{t('videoDownload.downloadNow')}}
        </el-button>
      </template>
     </el-dialog>   
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
  getVideoListListActive,
  getVideoListListPause,
   getVideoListListRemove,
   downM3u8,
} from '@/api/videoList'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,bytesToSize } from '@/utils/format'
import { ElDescriptions, ElMessage, ElMessageBox, ElNotification } from 'element-plus'
import { ref, reactive ,onBeforeUnmount} from 'vue'
import { useUserStore } from '@/pinia/modules/user'
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
  },8000)

}
setSearch()

onBeforeUnmount(()=>{
  clearInterval(timer)
})


const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_PATH)
const createDown =ref(false)
const activeName = ref('first')
const textarea = ref('')
const textareaHeader =ref('')
const createDownLoad =()=>{
  createDown.value=true

}
const textarea3 = ref('')

// 立即下载
const downNow= async ()=>{
   createDown.value=false
   if(activeName.value ==="first"){
    // videoDownloadPath
    const res = await createVideoList({videoDownloadPath:textarea.value,videoUrl:textareaHeader.value})
    // console.log("1",textarea.value)
    getTableData()
   }else if(activeName.value ==="there"){
    // 下载m3u8
    let urls =textarea3.value.split('\n')
    urls.forEach(e=>{
      let path =''
      let name =''
      let paths = e.split(',')
      console.log(paths)
      if(paths.length===1){
        path =paths[0]
      }else{
        path =paths[0]
        name=paths[1]
      }
      console.log(path,name)
      downM3u8({downloadPath:path,fileName:name})
      ElNotification({
        title:'success',
        type:'success',
        message:'creating ...'
      })
    })

   }else{
    getTableData()
   }
}


// 暂停
const pause= async()=>{
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要暂停下载的视频'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.gid)
        })
      // console.log(ids)
      const res = await getVideoListListPause({ grids:ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '暂停成功'
        })
        getTableData()
      }
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
// 查询
const getTableData = async() => {
  tableData.value=[]
  const table = await getVideoListListActive({ page: page.value, pageSize: pageSize.value})
  if (table.code === 0) {
    let d =table.data.list

      d.forEach(de=>{
        de.files.forEach(fe=>{
          let td =Object.assign({}, de)
          td.path=fe.path
          td.fileLength=fe.length
          td.completedLength=fe.completedLength
          tableData.value.push(td)
        })

      })
    

  console.log(tableData.value)
    // tableData.value = table.data.list
    // tableData.value.forEach(e=>{
    //   // 文件名
    //   let dir = e.dir
    //   let path = e.files[0].path ||0 
    //   e.fileLength = e.files[0].length 
    //   e.completedLength =e.files[0].completedLength || 0
    //   let temp = path.replace(dir,"").split('/')
    //   // console.log(temp,e.fileLength)
    //   e.fileName=path

    // })

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


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

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