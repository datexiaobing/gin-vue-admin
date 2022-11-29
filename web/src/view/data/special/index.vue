<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item :label="t('special.specialName')">
         <el-input v-model="searchInfo.specialName" placeholder="" />

        </el-form-item>
        <el-form-item :label="t('special.category')">
            
             <!-- <el-input v-model.number="searchInfo.specialCategory" placeholder="搜索条件" /> -->
             <el-select v-model="searchInfo.specialCategory">
                <el-option
                v-for="(v,i) in categorys"
                :key="i"
                :label="v.label"
                :value="v.value">
                </el-option>
             </el-select>

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">{{t('category.search')}}</el-button>
          <el-button size="small" icon="refresh" @click="onReset">{{t('category.reset')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="open1">{{t('category.create')}}</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>{{t('category.deleteConfirm')}}</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">{{t('category.cancel')}}</el-button>
                <el-button size="small" type="primary" @click="onDelete">{{t('category.confirm')}}</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">{{t('category.delete')}}</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        border
        max-height="600"
        header-cell-class-name="head-class"
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" :label="t('special.specialName')" prop="specialName" min-width="300" />
        <el-table-column align="left" :label="t('special.category')"  prop="specialCategory" width="120" >
          <template #default="scope">
            {{scope.row.specialCategory ?filterDict(scope.row.specialCategory ,categorys):t('transform.noCategory')}}
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('special.videoNumber')" prop="specialVideoNum" width="120" />
        <el-table-column align="left" :label="t('category.createTime')"  width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('special.updateTime')" width="180">
            <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="专辑图片" prop="specialPic" width="120" /> -->
        <el-table-column align="left" :label="t('special.operation')" min-width="180">
            <template #default="scope">
            <el-button type="success"  icon="edit"  class="table-button" @click="updateVideoSpecialFunc(scope.row)">{{t('category.change')}}</el-button>
            <el-button type="danger"  icon="delete"  @click="deleteRow(scope.row)">{{t('category.delete')}}</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="t('category.create')">
      <el-form :model="formData" label-position="left" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item :label="t('special.specialName')"  prop="specialName" >
          <el-input v-model="formData.specialName" :clearable="true"  placeholder="" />
        </el-form-item>
        <el-form-item :label="t('special.category')" >
          <!-- <el-input v-model.number="formData.specialCategory" :clearable="true" placeholder="请输入" /> -->
                <el-select v-model="formData.specialCategory" style="width:100%">
                    <el-option
                    v-for="(v,i) in categorys"
                    :key="i"
                    :label="v.label"
                    :value="v.value">
                    </el-option>
                </el-select>
        </el-form-item>
        <!-- <el-form-item label="视频数量:"  prop="specialVideoNum" >
          <el-input v-model.number="formData.specialVideoNum" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="专辑图片:"  prop="specialPic" >
          <el-input v-model="formData.specialPic" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item :label="t('special.video')"> 
          <el-cascader
          style="width:100%"
          :show-all-levels="false"
          :options="options" 
          :props="props"
          v-model="videoLists"
          @remove-tag="removeTag"
          >
          </el-cascader>
        </el-form-item>
      </el-form>
      <el-form-item :label="t('special.searchVideo')"> 

        <el-table
        border
        max-height="600"
        style="width: 100%"
        header-cell-class-name="head-class"
        tooltip-effect="dark"
        :data="tableDataSearch"
  
        >
          <!-- <el-table-column type="selection" width="55" /> -->
          <el-table-column align="left" label="uuid" prop="transUuid" min-width="300"/>
          <el-table-column align="left" :label="t('special.videoName')" prop="transOutName" width="200" fixed />
          
          <el-table-column align="center" :label="t('special.operation')" width="180">
            <template #header>
              <el-input v-model="searchName"   >
                <template #append>
                  <el-button icon="Search" @click="getTableDataSearch"/>
                </template>
              </el-input>
            </template>
            <template #default="scope">
              <el-button type="success" size="large" icon="Select" @click="choseVideo(scope.row)">{{t('special.chose')}}</el-button>
            </template>
          </el-table-column>
        </el-table>

      </el-form-item>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">{{t('category.cancel')}}</el-button>
          <el-button size="small" type="primary" @click="enterDialog">{{t('category.confirm')}}</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="dialogFormVisible1" :before-close="closeDialog" :title="t('category.create')">
      <el-form :model="formData1" label-position="left"  :rules="rule" label-width="120px">
        <el-form-item :label="t('special.specialName')"  prop="specialName" >
          <el-input v-model="formData1.specialName" :clearable="true"  placeholder="" />
        </el-form-item>
        <el-form-item :label="t('special.category')" >
          <!-- <el-input v-model.number="formData.specialCategory" :clearable="true" placeholder="请输入" /> -->
              <el-select v-model="formData1.specialCategory" style="width:100%">
                  <el-option
                  v-for="(v,i) in categorys"
                  :key="i"
                  :label="v.label"
                  :value="v.value">
                  </el-option>
              </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog1">{{t('category.cancel')}}</el-button>
          <el-button size="small" type="primary" @click="enterDialog1">{{t('category.confirm')}}</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script>
export default {
  name: 'VideoSpecial'
}
</script>

<script setup>
import {
  createVideoSpecial,
  deleteVideoSpecial,
  deleteVideoSpecialByIds,
  updateVideoSpecial,
  findVideoSpecial,
  getVideoSpecialList
} from '@/api/videoSpecial'
import {
  createVideoSpecialFileTrans,
  deleteVideoSpecialFileTrans,
  deleteVideoSpecialFileTransByIds,
  updateVideoSpecialFileTrans,
  findVideoSpecialFileTrans,
  getVideoSpecialFileTransList
} from '@/api/videoSpecialFileTrans'
import {
  getVideoCategoryList
} from '@/api/videoCategory'

import {
  createFileTrans,
  deleteFileTrans,
  deleteFileTransByIds,
  updateFileTrans,
  findFileTrans,
  getFileTransList
} from '@/api/fileTrans'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox, ElNotification, } from 'element-plus'
import { ref, reactive } from 'vue'

import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const searchName =ref('')
const dialogFormVisible1 =ref(false)
const props = ref({ multiple: true })
const options = ref([]) //展示多级选择的视频
const videoLists =ref([]) //存放视频ID
const categorys =ref([])

// 新建专辑
const closeDialog1=()=>{
  dialogFormVisible1.value=false
}
const open1 =()=>{
  dialogFormVisible1.value=true
}
const enterDialog1 = async()=>{
  const d = await createVideoSpecial(formData1.value)
  if(d.code===0){
    ElMessage({
      type:'success',
      message:'创建专辑成功！'
    })
    dialogFormVisible1.value=false
    getTableData()
  }
}

const getVideosTypes = async()=>{
    const d = await getVideoCategoryList({page:0,pageSize:50})
    // console.log(d)
    if(d.code === 0){
      d.data.list.forEach(element => {
        categorys.value.push({
          value:element.ID,
          label:element.categoryName
        })
      })
    }
}

getVideosTypes()

// 选择视频

const choseVideo = async(row) =>{
  // console.log(row)
  options.value.push({
    value: row.ID,
    label: row.transOutName,
  })
    // 判断是否已经加入过
  let isTrue = videoLists.value.includes(row.ID)
 
  if(isTrue){
    ElNotification({
      title:'提醒',
      type:'warning',
      message:'该视频已经选过了！'
    })
    return
  }
  videoLists.value.push(row.ID)
  // 创建操作
 
  let subData ={
    videoSpecialId:formData.value.ID, //专辑ID
    fileTransId:row.ID,  //视频ID
    fileTransOutName:row.transOutName,
    fileTransUuid:row.transUuid
  }
// 发起视频插入专辑
  const d= await createVideoSpecialFileTrans(subData)
  if(d.code===0){
    ElNotification({
      title:'提示',
      type:'success',
      message:'添加专辑成功！'
    })
    // console.log(tableDataSearch.value)
  }

  // 删除已经选择过的数据
  var delete_index = (tableDataSearch.value|| []).findIndex((el) => el.ID=== row.ID)
  tableDataSearch.value.splice(delete_index, 1)
  // console.log(tableDataSearch.value)
}
// 删除专辑内的视频
const removeTag= async (e)=>{
  let deletData={
    fileTransId:e[0], //videoId
    videoSpecialId:formData.value.ID  //专辑
  }
  // 发起删除操作，对应的ID发起 ID=e[0]
  const d =await deleteVideoSpecialFileTrans(deletData)
  if(d.code===0){
    ElNotification({
      title:'删除',
      type:'success',
      message:'删除成功！'
    })
  }

}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        specialName: '',
        specialCategory: 0,
        specialVideoNum: 0,
        specialPic: '',
        })
// 创建专辑
const formData1 = ref({
        specialName: '',
        specialCategory: 1,
        specialVideoNum: 0,
        specialPic: '',
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
// ---------
const pageSearch = ref(1)
const totalSearch  = ref(0)
const pageSizeSearch  = ref(10)
const tableDataSearch  = ref([])
const searchInfoSearch  = ref({})

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


// 查询待搜索视频
const getTableDataSearch = async() => {
  searchInfoSearch.value.transStatus =2
  searchInfoSearch.value.transOutName = searchName.value
  const table = await getFileTransList({ page: pageSearch.value, pageSize: pageSizeSearch.value, ...searchInfoSearch.value })
  if (table.code === 0) {
    // console.log(table.data)
    tableDataSearch.value = table.data.list
    totalSearch.value = table.data.total
    pageSearch.value = table.data.page
    pageSizeSearch.value = table.data.pageSize
  }
}



// 查询
const getTableData = async() => {
  const table = await getVideoSpecialList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteVideoSpecialFunc(row)
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
      const res = await deleteVideoSpecialByIds({ ids })
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

// 更新行 打开时获取已经分类的视频
const updateVideoSpecialFunc = async(row) => {
    const res = await findVideoSpecial({ ID: row.ID })
    // console.log(res)
    type.value = 'update'
    if (res.code === 0) { 
      formData.value = res.data.revideoSpecial
      res.data.revideoSpecial.specialFile.forEach(row=>{
          // 已归类的视频
          options.value.push({
            value: row.fileTransId,
            label: row.fileTransOutName,
          })
          videoLists.value.push(row.fileTransId)
      })

        dialogFormVisible.value = true
    }
}


// 删除行
const deleteVideoSpecialFunc = async (row) => {
    const res = await deleteVideoSpecial({ ID: row.ID })
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
    options.value=[]
    videoLists.value=[]
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        specialName: '',
        specialCategory: 0,
        specialVideoNum: 0,
        specialPic: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
  videoLists.value.forEach(el=>{
     console.log(el)
  })
  //  发起更新操作
  dialogFormVisible.value = false

    //  elFormRef.value?.validate( async (valid) => {
    //          if (!valid) return
    //           let res
    //           switch (type.value) {
    //             case 'create':
    //               res = await createVideoSpecial(formData.value)
    //               break
    //             case 'update':
    //               res = await updateVideoSpecial(formData.value)
    //               break
    //             default:
    //               res = await createVideoSpecial(formData.value)
    //               break
    //           }
    //           if (res.code === 0) {
    //             ElMessage({
    //               type: 'success',
    //               message: '创建/更改成功'
    //             })
    //             closeDialog()
    //             getTableData()
    //           }
    //   })
}
</script>

<style>
.head-class{
    font-weight: bold;
    font-size: 20px; 
}
</style>
