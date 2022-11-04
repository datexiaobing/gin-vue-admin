<template>
  <div>
    <el-container>
        <el-row>
            <el-button-group class="">
              <el-space>
                <el-button icon="UploadFilled"> 上传</el-button>
                <el-button icon="Delete"> 删除</el-button>
                <el-button icon="Back" @click="getPrewFolder"> 上一级</el-button>
                <el-button icon="Refresh" @click="getTableData"> 刷新</el-button>
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

        <el-table-column align="left" label="文件名称" prop="fileName" >
          <template #default="scope">
            <el-icon v-if="scope.row.isDir"><Folder /></el-icon>
            <el-icon v-else><Files /></el-icon>
            <el-link 
             v-if="scope.row.isDir"
            :underline="false" 
            @click="nextFolder(scope.row)">{{scope.row.fileName}}</el-link>
            <span v-else >{{scope.row.fileName}}</span>
          </template>
        </el-table-column>
        <!-- if 子目录，更新datalist -->
        <el-table-column align="left" label="文件大小" width="120" >
            <template #default="scope">
                {{bytesToSize(scope.row.fileSize)}}
            </template>
        </el-table-column>
        <el-table-column align="left" label="修改时间" width="180" >
          <template #default="scope">{{ formatDate(scope.row.modTime) }}</template>
        </el-table-column>

        <el-table-column align="left" label="操作" width="330">
            <template #default="scope">             
              <el-space class="b-contain">
                <el-button type="danger"  icon="delete"  >删除</el-button>
                <el-button type="success"  icon="EditPen"   >重命名</el-button>
                <el-button v-if="scope.row.fileType === 'video'" type="primary"  icon="CopyDocument" @click="moveVideo(scope.row)" >移到视频源</el-button>
              </el-space>
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
  createFileDownload,
  deleteFileDownload,
  deleteFileDownloadByIds,
  updateFileDownload,
  findFileDownload,
  getFileDownloadList
} from '@/api/fileDownload'
import {
  getVideoListListFile,
} from '@/api/videoList'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,bytesToSize } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'


const dir=ref("/")

// 进入子目录
const nextFolder = async(row)=>{
  let foldrName = row.fileName 
  dir.value=dir.value+foldrName
  
  getTableData()

}

// 返回上一级
const getPrewFolder=async()=>{
  let newDirList = dir.value.split('/')
  // let newDirList="/女浩克.She-Hulk.Attorney.at.Law.S01E08/121221/22222".split('/')
  newDirList.pop()
  console.log(newDirList)
  if(newDirList===""){
    dir.value="/"
    getTableData()
    console.log('11')
  }else{
    let n=newDirList.join("/")
    console.log("n:",n)
    dir.value=n
    getTableData()
  }
}

// 移动视频
const moveVideo =(row)=>{
  ElMessage({
    type:'danger',
    message:"尚未开发"
  })
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        fileDegree: 0,
        fileIsDir: 0,
        fileLastDir: '',
        fileName: '',
        filePath: '',
        fileSize: 0,
        fileType: '',
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
  searchInfo.value.videoUrl=dir.value
  const table = await  getVideoListListFile({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
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
            deleteFileDownloadFunc(row)
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
      const res = await deleteFileDownloadByIds({ ids })
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
const updateFileDownloadFunc = async(row) => {
    const res = await findFileDownload({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.refileDownload
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteFileDownloadFunc = async (row) => {
    const res = await deleteFileDownload({ ID: row.ID })
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
        fileDegree: 0,
        fileIsDir: 0,
        fileLastDir: '',
        fileName: '',
        filePath: '',
        fileSize: 0,
        fileType: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createFileDownload(formData.value)
                  break
                case 'update':
                  res = await updateFileDownload(formData.value)
                  break
                default:
                  res = await createFileDownload(formData.value)
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

.b-contain{
  display: flex;
  flex-direction: row-reverse;
}
.head-class{
    font-weight: bold;
    font-size: 20px; 
}
</style>