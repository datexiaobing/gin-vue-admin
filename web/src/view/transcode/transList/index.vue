<template>
    <div> 
        <el-container>
            <el-row>
                <el-button-group class="">
                <el-space>
                    <el-button icon="Share" @click="getShare"> {{t('videoDownload.share')}} </el-button>
                    <!-- <el-button icon="HelpFilled"> 转码</el-button> -->
                    <el-button icon="Delete"> {{t('videoDownload.delete')}}</el-button>
                    <el-button icon="Pointer" @click="updateType"> {{t('videoDownload.update')}}</el-button>
                    <el-button icon="Refresh" @click="getTableData"> {{t('videoDownload.refresh')}}</el-button>
                </el-space>
                </el-button-group>
            </el-row>
        </el-container>

        <el-table
        border
        max-height="600"
        ref="multipleTable"
        style="width: 100%"
        header-cell-class-name="head-class"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" :label="t('transform.transInputName')" prop="transInputName" min-width="300" fixed />
        <el-table-column align="left" label="uuid" prop="transUuid" min-width="300"/>
        <el-table-column align="left" :label="t('transform.transOutName')" prop="transOutName" min-width="200" />
        
        <!-- <el-table-column align="left" label="1正在2完成" prop="transStatus" width="120" /> -->
        <el-table-column align="left" :label="t('transform.category')"  width="120" >
            <template #default="scope">
            <el-tag size="large"  type="success"> 
                {{scope.row.transType ? filterDict(scope.row.transType,videosType):t('transform.noCategory')}}
            </el-tag>
            </template>
        </el-table-column>

        <el-table-column align="left" :label="t('transform.transTypeNum')"  width="120" >
          <template #default="scope">
            <el-tag size="large" >
              {{scope.row.transType ? filterDict(scope.row.transTypeNum,videosTypeNum):0}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('transform.transResolution')" prop="transResolution" width="120" >
          <template #default="scope">
            {{scope.row.transResolution ===3?'1080P':scope.row.transResolution ===2?'720P':'360P'}}
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('transform.transDuration')" prop="transDuration" width="120" />
        <el-table-column align="left" :label="t('transform.transSubtitle')" prop="transSubtitle" width="180" />
        <el-table-column align="left" :label="t('videoList.SeektimeHeard')" prop="transSeektimeHeard" width="120" >
            <template #default="scope">
                {{scope.row.transSeektimeHeard ?scope.row.transSeektimeHeard :'否'}}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('videoList.SeektimeTail')" prop="transSeektimeTail" width="120" >
             <template #default="scope">
                {{scope.row.transSeektimeTail ?scope.row.transSeektimeTail:'否'}}
            </template>
        </el-table-column>   

        <el-table-column align="left" :label="t('videoList.paoma')"  width="120" >
          <template #default="scope">
            <el-tag size="large" plain @click="seeDrawtext(scope.row)">{{t('transform.look')}}</el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="QiNiuOss"  width="120" >
          <template #default="scope">
            <el-button 
            v-if="scope.row.transOssQiniuStatus ===1"
            type="primary" @click="upQiniu(scope.row)" size="small">{{t('transform.ossUp')}}</el-button>
            <el-tag 
            v-else-if="scope.row.transOssQiniuStatus ===2"
            size="large" type="info">{{t('transform.ossWaiting')}}</el-tag>
            <el-tag 
            v-else
            size="large" type="success">{{t('transform.ossComplete')}}</el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="AliOss"  width="120" >
          <template #default="scope">
            <el-button 
            v-if="scope.row.transOssAliStatus === 1"
            type="primary" @click="upAli(scope.row)" size="small">{{t('transform.ossUp')}}</el-button>
            <el-tag 
            v-else-if="scope.row.transOssAliStatus ===2"
            size="large" type="info">{{t('transform.ossWaiting')}} </el-tag>
            <el-tag 
            v-else
            size="large" type="success">{{t('transform.ossComplete')}}</el-tag>
          </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="跑马灯颜色" prop="transDrawtextColor" width="120" />
        <el-table-column align="left" label="跑马灯位置" prop="transDrawtextPosition" width="120" />
        <el-table-column align="left" label="滚动时长" prop="transDrawtextDuration" width="120" />
        <el-table-column align="left" label="滚动间隔" prop="transDrawtextInterval" width="120" />
        <el-table-column align="left" label="跑马灯文字" prop="transDrawtextString" width="120" />
        <el-table-column align="left" label="文字大小" prop="transDrawtextFontsize" width="120" />
        <el-table-column align="left" label="字幕开关" prop="transSubtitleSwitch" width="120" /> -->
        <el-table-column align="left" :label="t('transform.transStartTime')" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('transform.transEndTime')" width="180">
            <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('transform.createTime')" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('transform.updateTime')" width="180">
            <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>        
        
        <el-table-column align="left" :label="t('transform.transProgressRate')" prop="transProgressRate" min-width="280" fixed="right">
           <template #default="scope">
            <div class="pro-cell">
                <el-progress 
                :status="scope.row.transProgressRate === 100?'success':''"
                :stroke-width="8" 
                :percentage="scope.row.transProgressRate" />
            </div>
           </template>
        </el-table-column>
        
        <!-- <el-table-column align="left" label="上传错误" prop="transOssError" width="120" /> -->
        <el-table-column align="left" :label="t('videoDownload.operate')" fixed="right" width="120">
            <template #default="scope">
            <el-button type="danger"  icon="delete"  @click="deleteRow(scope.row)">{{ t('videoDownload.delete')}}</el-button>
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
      <!-- 跑马灯 -->
      <el-dialog v-model="dialogFormVisible"  :title="t('videoList.paoma')">
        <el-descriptions
          :column="1"
          border
          >
          <el-descriptions-item>
            <template #label>
             {{t('transform.subtitleSwitch')}}
            </template>
            {{drawtext.transSubtitleSwitch? t('transform.open'):t('transform.close')}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>
               {{t('videoList.drawtextPosition')}}
            </template>
            {{drawtext.transDrawtextPosition===1? t('videoList.up'):t('videoList.down')}}
          </el-descriptions-item>
           <el-descriptions-item>
            <template #label>
               {{t('videoList.drawtextColor')}}
            </template>
            {{drawtext.transDrawtextColor?drawtext.transDrawtextColor:''}}
          </el-descriptions-item>   
          <el-descriptions-item>
            <template #label>
              {{t('videoList.drawtextFontsize')}}
            </template>
            {{drawtext.transDrawtextFontsize?drawtext.transDrawtextFontsize:''}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>
              {{t('videoList.drawtextDuration')}}
            </template>
            {{drawtext.transDrawtextDuration?drawtext.transDrawtextDuration:''}}
          </el-descriptions-item>    
          <el-descriptions-item>
            <template #label>
               {{t('videoList.drawtextInterval')}}
            </template>
            {{drawtext.transDrawtextInterval?drawtext.transDrawtextInterval:''}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>
              {{t('videoList.drawtextString')}}
            </template>
            {{drawtext.transDrawtextString?drawtext.transDrawtextString:''}}
          </el-descriptions-item>           

        </el-descriptions>

        <template #footer>
          <div class="dialog-footer">
            <el-button type="primary" @click="closeDialog">{{t('transform.close')}}</el-button>
          </div>
        </template>
      </el-dialog>

      <!--分享  -->
      <el-dialog v-model="dialogFormVisibleShare"  :title="t('videoDownload.share')">
        <el-form :model="formDataShare" label-position="left" label-width="90px">
          <el-form-item :label="t('transform.videoSource')"  >
            <el-input v-for="(k,index) in videoNames"
            :key="index"
            disabled
            :placeholder="k.value" />
          </el-form-item>

          <el-form-item label="ip" prop="ip" >
            <el-input v-model="formDataShare.ip"   />
          </el-form-item>
          <el-form-item :label="t('transform.expires')" prop="expires" >
            <el-input v-model="formDataShare.expires"   />
          </el-form-item>
          <el-form-item :label="t('transform.domain')" prop="domain" >
            <el-input v-model="formDataShare.domain"   />
          </el-form-item>
          <el-form-item :label="t('transform.format')" >
            <el-radio-group v-model="formatDetail"  >
              <el-radio-button label="1">M3U8</el-radio-button>
              <el-radio-button label="2">{{t('transform.player')}}</el-radio-button>
              <el-radio-button label="3">{{t('transform.pic')}}</el-radio-button>
              <el-radio-button label="4">{{t('transform.qn')}}</el-radio-button>
              <el-radio-button label="5">{{t('transform.al')}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item :label="t('transform.content')">
            <el-input
              v-model="textarea"
              :rows="10"
              type="textarea"
              
            />
          </el-form-item>
        </el-form>

        <template #footer>
          <div class="dialog-footer">
            <!-- <el-button type="primary" @click="closeDialogShare">关闭</el-button> -->
            <el-button type="primary" @click="configShare">
              <el-icon><Share /></el-icon>
              {{t('transform.fetch')}}
              </el-button>
          </div>
        </template>
      </el-dialog>
      <!-- 修改分类 -->
      <el-dialog v-model="dialogChange"  :title="t('transform.change')">
        <el-form :model="formDataChange" label-position="left" label-width="90px">
        <el-form-item :label="t('transform.transOutName')"  >
            <el-input v-model="formDataChange.transOutName"   />
        </el-form-item>

        <el-form-item :label="t('transform.transTypeNum')"  >
          <el-select v-model="formDataChange.transTypeNum" placeholder="">
            <el-option
            v-for="item in videosTypeNum"
            :key="item.value"
            :label="item.label"
            :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="t('transform.category')"  >
          <el-select v-model="formDataChange.transType" placeholder="">
            <el-option
            v-for="item in videosType"
            :key="item.value"
            :label="item.label"
            :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>

        </el-form>

        <template #footer>
          <div class="dialog-footer">
            <!-- <el-button type="primary" @click="closeDialogShare">关闭</el-button> -->
            <el-button type="primary" @click="configChange">
              {{t('transform.change')}}
              </el-button>
          </div>
        </template>
      </el-dialog>

    </div>
</template>
    
<script setup>
import {
  createFileTrans,
  deleteFileTrans,
  deleteFileTransByIds,
  updateFileTrans,
  findFileTrans,
  getFileTransList,
  getShareList,
  uploadQiniu,
  uploadAli
} from '@/api/fileTrans'
import {
  getVideoCategoryList
} from '@/api/videoCategory'
import{  
  getVideoSpecialList
} from '@/api/videoSpecial'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox, ElNotification } from 'element-plus'
import { ref, reactive,watch } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

// 多选数据
const multipleSelection = ref([])
// const base_host="http://127.0.0.1"
const base_host= ref(import.meta.env.VITE_BASE_PATH)
// 弹窗控制标记
const dialogFormVisible = ref(false)
const dialogFormVisibleShare =ref(false)
// 分享链接的formData
const formDataShare =ref({
  ids:[],
  ip:'0.0.0.0',
  expires:'100',
  domain:'*'
})
const videoNames =ref([])
const formatDetail =ref('1')
const textarea =ref('')

// 上传到七牛
const upQiniu=async(row)=>{
  // console.log(row)
  const d = await uploadQiniu(row)
  if(d.code===0){
    ElNotification({
      title:'success',
      type:'success',
      message:'Upload To Qiniu....'
    })
  }
}

// 上传到ali
const upAli=async(row)=>{
  // console.log(row)
  const d = await uploadAli(row)
  if(d.code===0){
    ElNotification({
      title:'success',
      type:'success',
      message:'Upload To Ali....'
    })
  }
}

// 获取 分类 专辑
const videosType = ref([])
const videosTypeNum = ref([])
const getVideosTypes = async()=>{
    const d = await getVideoCategoryList({page:0,pageSize:50})
    // console.log(d)
    if(d.code === 0){
      d.data.list.forEach(element => {
        videosType.value.push({
          value:element.ID,
          label:element.categoryName
        })
      })
    }
}
const getVideosTypesNum = async()=>{
    const d = await getVideoSpecialList({page:0,pageSize:50})
    // console.log(d)
    if(d.code === 0){
      d.data.list.forEach(element => {
        videosTypeNum.value.push({
          value:element.ID,
          label:element.specialName
        })
      })
    }
}
getVideosTypes()
getVideosTypesNum()

// 更新分类 专辑
const dialogChange =ref(false)
const updateType = ()=>{
  let d =multipleSelection.value 
  if(d.length ===0){
    ElNotification({
      title:"no select",
      type:"error",
      message:"please chose a video"
    })
  }else if(d.length >1){
    ElNotification({
      title:"only one",
      type:"warning",
      message:"only one video can be chose "
    })
  }else{
    // console.log(d)
    
    formDataChange.value.transOutName = d[0].transOutName
    formDataChange.value.transType = d[0].transType
    formDataChange.value.transTypeNum = d[0].transTypeNum
    dialogChange.value=true
  }
}
const configChange =async ()=>{
  let d =multipleSelection.value[0] 
  d.transType =formDataChange.value.transType
  d.transTypeNum =  formDataChange.value.transTypeNum
  d.transOutName = formDataChange.value.transOutName
  // console.log(d)
  d.transPost =1 //重置推送状态
  const res = await updateFileTrans(d)
  if(res.code===0){
      ElMessage({
      type: 'success',
      message: 'success !'
    })
  }
   dialogChange.value=false
}

const formDataChange=ref({
  transType:0,
  transTypeNum:0,
  transOutName:''
})

const drawtext = ref(null)

const seeDrawtext = (row)=>{
  dialogFormVisible.value=true
  drawtext.value=row

}

watch(formatDetail,(newVal,oldVal)=>{
  // console.log(newVal)
  configShare()
})

// 获取分享链接
const getShare = ()=>{
  textarea.value=[]
  dialogFormVisibleShare.value=true
  videoNames.value=[]
  formDataShare.value.ids=[]
  multipleSelection.value.map(item => {
      formDataShare.value.ids.push(item.ID)
      videoNames.value.push({
        key:item.ID,
        value:item.transOutName
      })
    })
 
}
const configShare=async()=>{
    textarea.value=[]
    const d = await getShareList(formDataShare.value)
    if(d.code===0){
      d.data.list.forEach(el=>{
          if(formatDetail.value ==='3'){
            // 缩略图
            textarea.value += base_host.value + el.picUrl+'\n'
          }else if(formatDetail.value ==='4'){
            // 七牛
            textarea.value +=el.qiniuUrl+'\n'
          }else if(formatDetail.value ==='5'){
            // 七牛
            textarea.value +=el.aliUrl+'\n'
          }else{
            // 视频链接
            textarea.value +=base_host.value +el.M3Url+'\n'
          }
      })
    }
  
}
// 关闭分享链接
const closeDialogShare=()=>{
  dialogFormVisibleShare.value=false
  textarea.value=[]
}


// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        transInputName: '',
        transOutName: '',
        transUuid: '',
        transStatus: 0,
        transType: 0,
        transTypeNum: 0,
        transResolution: '',
        transDuration: '',
        transSeektimeHeard: 0,
        transSeektimeTail: 0,
        transDrawtextColor: '',
        transDrawtextPosition: '',
        transDrawtextDuration: '',
        transDrawtextInterval: '',
        transDrawtextString: '',
        transDrawtextFontsize: 0,
        transSubtitle: '',
        transSubtitleSwitch: 0,
        transProgressRate: 0,
        transOssStatus: 0,
        transOssError: '',
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
  searchInfo.value.transStatus =2
  const table = await getFileTransList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteFileTransFunc(row)
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
      const res = await deleteFileTransByIds({ ids })
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
const updateFileTransFunc = async(row) => {
    const res = await findFileTrans({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.refileTrans
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteFileTransFunc = async (row) => {
    const res = await deleteFileTrans({ ID: row.ID })
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



// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false

}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createFileTrans(formData.value)
                  break
                case 'update':
                  res = await updateFileTrans(formData.value)
                  break
                default:
                  res = await createFileTrans(formData.value)
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

<style>
.pro-cell{
    padding-top: 5%;
}
.head-class{
    font-weight: bold;
    font-size: 20px; 
}
</style>