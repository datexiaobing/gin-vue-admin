<template>
   <div> 
    <el-container>
        <el-row>
            <el-button-group class="">
              <el-space>
                <el-button icon="UploadFilled" disabled> 上传</el-button>
                <el-button icon="HelpFilled"> 转码</el-button>
                <el-button icon="Delete"> 删除</el-button>
                <el-button icon="Back"  @click="getPrewFolder"> 上一级</el-button>
                <el-button icon="Refresh" @click="getTableData"> 刷新</el-button>
              </el-space>
            </el-button-group>
        </el-row>
    </el-container>
    <el-table
    border
    max-height="600"
    ref="multipleTable"
    style="width: 100%"
    tooltip-effect="dark"
    :data="tableData"
    row-key="ID"
    @selection-change="handleSelectionChange"
    >
    <el-table-column type="selection" width="55" />

    <el-table-column align="left" label="文件名称" prop="fileName"  >
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
            <el-button type="danger"  icon="delete" @click="deleteFiles(scope.row)" >删除</el-button>
            <el-button type="success"  icon="EditPen"  @click="changeFileNames(scope.row)" >重命名</el-button>
            <el-button v-if="scope.row.fileType === 'video'" 
            type="primary"  icon="HelpFilled" 
            @click="trans(scope.row)">转码</el-button>
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

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="转码">
      <el-form :model="formData" label-position="left" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="视频源:"  prop="fileName" >
          <el-input v-model="formData.fileName" disabled />
        </el-form-item>
        <el-form-item label="视频重命名:"  prop="transOutName" >
          <el-input v-model="formData.transOutName" :clearable="true"   />
        </el-form-item>
        <el-form-item label="视频分类:"  prop="transType" >
          <!-- <el-input v-model.number="formData.transType" :clearable="true" placeholder="请输入" /> -->
          <el-select v-model="formData.transType" placeholder="请选择">
            <el-option
            v-for="item in transTypeOption"
            :key="item.value"
            :label="item.label"
            :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="分辨率:"  prop="transResolution" >
          <!-- <el-input v-model="formData.transResolution" :clearable="true"  placeholder="请输入" /> -->
          <el-select v-model="formData.transResolution" placeholder="请选择">
            <el-option
            v-for="item in transResolutionOption"
            :key="item.value"
            :label="item.label"
            :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="跳过片头:"  prop="transSeektimeHeard" >
          <!-- <el-input v-model.number="formData.transSeektimeHeard" :clearable="true" placeholder="请输入" /> -->
          <el-slider v-model="formData.transSeektimeHeard" />
        </el-form-item>
        <el-form-item label="跳过片尾:"  prop="transSeektimeTail" >
          <!-- <el-input v-model.number="formData.transSeektimeTail" :clearable="true" placeholder="请输入" /> -->
          <el-slider v-model="formData.transSeektimeTail" />
        </el-form-item>
        <el-form-item label="跑马灯"   >
          <el-switch v-model="paoma" size="large" @change="changePao"/>
        </el-form-item>
        <el-form-item label="跑马灯文字:"  prop="transDrawtextString" v-show="paoma">
          <el-input v-model="formData.transDrawtextString" :clearable="true"  />
        </el-form-item>
        <el-form-item label="跑马灯位置:"  prop="transDrawtextPosition" v-show="paoma">
          <!-- <el-input v-model="formData.transDrawtextPosition" :clearable="true"  /> -->
          <el-radio-group v-model.number="formData.transDrawtextPosition" >
            <el-radio-button label="1">上</el-radio-button>
            <el-radio-button label="2">下</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="跑马灯颜色:"  prop="transDrawtextColor"  v-show="paoma">
          <!-- <el-input v-model="formData.transDrawtextColor" :clearable="true"  placeholder="请输入" /> -->
          <el-radio-group v-model.number="formData.transDrawtextColor" >
            
            <el-radio-button label="aqua" />
            <el-radio-button label="black" />
            <el-radio-button label="blue" />
            <el-radio-button label="fuchsia" />
            <el-radio-button label="gray" />
            <el-radio-button label="green" />           
            <el-radio-button label="lime" />
            <el-radio-button label="maroon" />
            <el-radio-button label="navy" />
            <el-radio-button label="olive" />
            <el-radio-button label="purple" />
            <el-radio-button label="red" />   
            <el-radio-button label="silver" />
            <el-radio-button label="teal" />
            <el-radio-button label="white" />            

          </el-radio-group>
        </el-form-item>
        <el-form-item label="跑马灯大小:"  prop="transDrawtextFontsize" v-show="paoma">
          <!-- <el-input v-model.number="formData.transDrawtextFontsize" :clearable="true" placeholder="请输入" /> -->
          <el-slider v-model="formData.transDrawtextFontsize" />
        </el-form-item>
        <el-form-item label="跑马灯时长:"  prop="transDrawtextDuration" v-show="paoma" >
          <!-- <el-input v-model="formData.transDrawtextDuration" :clearable="true"  placeholder="请输入" /> -->
          <el-slider v-model="formData.transDrawtextDuration" />
        </el-form-item>
        <el-form-item label="跑马灯间隔:"  prop="transDrawtextInterval" v-show="paoma">
          <el-slider v-model="formData.transDrawtextInterval" />
          <!-- <el-input v-model="formData.transDrawtextInterval" :clearable="true"  placeholder="请输入" /> -->
        </el-form-item>


        <!-- <el-form-item label="字幕文件:"  prop="transSubtitle" >
          <el-input v-model="formData.transSubtitle" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="字幕开关:"  prop="transSubtitleSwitch" >
          <el-input v-model.number="formData.transSubtitleSwitch" :clearable="true" placeholder="请输入" />
        </el-form-item> -->

      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="transPost">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="dialogFormVisibleFileName">
      <el-form label-position="left" label-width="120">
        <el-form-item label="新名称">
          <el-input v-model="newFileNameData.newFileName"  />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" type="primary" @click="enterChangeFileName">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  


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
  getVideoListListFileDone,
  changeFileName,
  deleteFile
} from '@/api/videoList'
import {
  createFileTrans,
  deleteFileTrans,
  deleteFileTransByIds,
  updateFileTrans,
  findFileTrans,
  getFileTransList,
  
} from '@/api/fileTrans'
import {
  getVideoCategoryList
} from '@/api/videoCategory'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,bytesToSize } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

const transResolutionOption =ref([{label:'1080P',value:3},{label:'720P',value:2},{label:'360p',value:1}])
const transTypeOption =ref([])
const dir=ref("/")
const paoma =ref(false)


const deleteFiles =async (row)=>{
  const d = await deleteFile({downloadPath:row.downloadPath})
    if(d.code===0){
     ElMessage({
      type:"success",
      message:"删除成功"
     })
     getTableData()
  }
}
// 修改文件名称
const newFileNameData=ref({
  newFileName:'',
  downloadPath:'',
  fileName:'',
})
const dialogFormVisibleFileName=ref(false)
const changeFileNames = (row)=>{
  dialogFormVisibleFileName.value=true
  
  newFileNameData.value.newFileName=row.fileName
  newFileNameData.value.downloadPath=row.downloadPath
}
const enterChangeFileName =async ()=>{
  dialogFormVisibleFileName.value=false
  const d = await changeFileName({downloadPath:newFileNameData.value.downloadPath,fileName:newFileNameData.value.newFileName})
  if(d.code===0){
     ElMessage({
      type:"success",
      message:"修改成功"
     })
     getTableData()
  }
   
}

const changePao=(val)=>{
  paoma.value = val
}

// 获取transTypeOpton
const getOptions=async ()=>{
  const d = await getVideoCategoryList({page:1,pageSize:50})
  if(d.code===0){
    // console.log(d)
    transTypeOption.value.push({
        value:0,
        label:'未分类'
      })
    d.data.list.forEach(el=> {
      transTypeOption.value.push({
        value:el.ID,
        label:el.categoryName
      })
    })

  }
}

getOptions()

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
  // console.log(newDirList)
  if(newDirList===""){
    dir.value="/"
    getTableData()
    // console.log('11')
  }else{
    let n=newDirList.join("/")
    // console.log("n:",n)
    dir.value=n
    getTableData()
  }
}


// 传递跑马灯等数据
const formData = ref(
      {
        transDrawtextSwitch:1, //是否开启跑马灯
        transInputName: '',
        transOutName: '',
        transUuid: '',
        transStatus: 0,
        transType:0,
        transTypeNum: 0,
        transResolution: 3,
        transDuration: '',
        transSeektimeHeard: 0,
        transSeektimeTail: 0,
        transDrawtextColor: 'red',
        transDrawtextPosition: 1,
        transDrawtextDuration:5,
        transDrawtextInterval: 3,
        transDrawtextString: 'this is my paoma 我的跑马',
        transDrawtextFontsize: 50,
        transSubtitle: '',
        transSubtitleSwitch: 0,
        transProgressRate: 0,
        transOssStatus: 0,
        transOssError: '',
        })
// *******trans video*****


const trans =async (row)=>{
  formData.value.downloadPath=row.downloadPath
  formData.value.fileName =row.fileName
  dialogFormVisible.value = true
}

const transPost =async()=>{
   dialogFormVisible.value = false
    const d = await createFileTrans(formData.value)
    // console.log(d)
    if(d.code===0){
      ElMessage({
        type:"success",
        message:"创建转码成功"
      })
    }
}



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
  const table = await getVideoListListFileDone({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    // formData.value = {
    //     fileDegree: 0,
    //     fileIsDir: 0,
    //     fileLastDir: '',
    //     fileName: '',
    //     filePath: '',
    //     fileSize: 0,
    //     fileType: '',
    //     }
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



/* .el-radio-button  >>> {
  background-color: aqua;
} */


.speed-cell{
    display: flex;
    flex-direction: column;
}

.b-contain{
  display: flex;
  flex-direction: row-reverse;
}

</style>