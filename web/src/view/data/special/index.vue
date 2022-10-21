<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="专辑名称">
         <el-input v-model="searchInfo.specialName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="分类">
            
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
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
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

        <el-table-column align="left" label="专辑名称" prop="specialName" min-width="300" />
        <el-table-column align="left" label="分类"  prop="specialCategory" width="120" >
          <template #default="scope">
            {{scope.row.specialCategory ?filterDict(scope.row.specialCategory ,categorys):'未分类'}}
          </template>
        </el-table-column>
        <el-table-column align="left" label="视频数量" prop="specialVideoNum" width="120" />
        <el-table-column align="left" label="创建时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="更新时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="专辑图片" prop="specialPic" width="120" /> -->
        <el-table-column align="left" label="操作" min-width="180">
            <template #default="scope">
            <el-button type="success"  icon="edit"  class="table-button" @click="updateVideoSpecialFunc(scope.row)">更新</el-button>
            <el-button type="danger"  icon="delete"  @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新建/更新">
      <el-form :model="formData" label-position="left" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="专辑名称:"  prop="specialName" >
          <el-input v-model="formData.specialName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类:" >
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
        <el-form-item label="关联视频"> 
          <el-cascader
          style="width:100%"
          :show-all-levels="false"
          :options="options" 
          :props="props"
          v-model="videoLists"
          >
          </el-cascader>
        </el-form-item>
      </el-form>
      <el-form-item label="查询视频"> 

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
          <el-table-column align="left" label="视频名称" prop="transOutName" width="200" fixed />
          
          <el-table-column align="center" label="操作" width="180">
            <template #header>
              <el-input v-model="searchName"  placeholder="Type to search" >
                <template #append>
                  <el-button icon="Search" @click="getTableDataSearch"/>
                </template>
              </el-input>
            </template>
            <template #default="scope">
              <el-button type="success" size="large" icon="Select" @click="choseVideo(scope.row)">选择</el-button>
            </template>
          </el-table-column>
        </el-table>

      </el-form-item>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
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
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

const searchName =ref('转换')

const props = ref({ multiple: true })
const options = ref([])

const videoLists =ref([])


const categorys =ref([])

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
const choseVideo = row =>{
  // console.log(row)
  options.value.push({
    value: row.ID,
    label: row.transOutName,
  })
  videoLists.value.push(row.ID)
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        specialName: '',
        specialCategory: 0,
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

// 更新行
const updateVideoSpecialFunc = async(row) => {
    const res = await findVideoSpecial({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.revideoSpecial
      res.data.revideoSpecial.specialFile.forEach(row=>{
          // 已归类的视频
          options.value.push({
            value: row.ID,
            label: row.fileTransOutName,
          })
          videoLists.value.push(row.ID)
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
