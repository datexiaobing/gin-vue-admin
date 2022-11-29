<template>
    <div> 
        <el-container>
            <el-row>
                <el-button-group class="">
                <el-space>
                    <!-- <el-button icon="UploadFilled"> 上传</el-button>
                    <el-button icon="HelpFilled"> 转码</el-button> -->
                    <el-button icon="Delete"> {{t('videoDownload.delete')}}</el-button>
                    <!-- <el-button icon="Back"> 上一级</el-button> -->
                    <el-button icon="Refresh" @click="getTableData"> {{ t('videoDownload.refresh')}}</el-button>
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
        <el-table-column align="left" :label="t('transform.transInputName')" prop="transInputName" width="300" fixed />
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

        <el-table-column align="left" :label="t('transform.transTypeNum')" prop="transTypeNum" width="120" />
        <el-table-column align="left" :label="t('transform.transResolution')" prop="transResolution" width="120" >
          <template #default="scope">
            {{scope.row.transResolution ===3?'1080P':scope.row.transResolution ===2?'720P':'360P'}}
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('transform.transDuration')" prop="transDuration" width="120" />
        <el-table-column align="left" :label="t('transform.transSubtitle')" prop="transSubtitle" width="180" />
        <el-table-column align="left" :label="t('videoList.SeektimeHeard')"  prop="transSeektimeHeard" width="120" >
            <template #default="scope">
                {{scope.row.transSeektimeHeard ?scope.row.transSeektimeHeard :'no'}}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('videoList.SeektimeTail')" prop="transSeektimeTail" width="120" >
             <template #default="scope">
                {{scope.row.transSeektimeTail ?scope.row.transSeektimeTail:'no'}}
            </template>
        </el-table-column>   

        <el-table-column align="left" :label="t('videoList.paoma')"  width="120" >
          <template #default="scope">
            <el-tag size="large" plain @click="seeDrawtext(scope.row)"> {{t('transform.look')}}</el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="oss" prop="transOssStatus" width="120" >
            <el-tag size="large" type="info"> {{t('transform.ossWaiting')}}</el-tag>
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
            <!-- <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template> -->
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
                :percentage="scope.row.transProgressRate.toFixed(2) || 0" />

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
            {{drawtext.transDrawtextPosition===1?t('videoList.up'):t('videoList.down')}}
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
            <el-button type="primary" @click="closeDialog"> {{t('transform.close')}}</el-button>
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
  getFileTransList
} from '@/api/fileTrans'
import {
  getVideoCategoryList
} from '@/api/videoCategory'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive,onBeforeUnmount } from 'vue'
import {start,close} from '@/utils/npgress'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

// const videosType = ref([{value:1,label:'喜剧'}])
const videosType = ref([])
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
getVideosTypes()

const drawtext = ref(null)

const seeDrawtext = (row)=>{
  dialogFormVisible.value=true
  drawtext.value=row

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
    searchInfo.value.transStatus =1
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
        }
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