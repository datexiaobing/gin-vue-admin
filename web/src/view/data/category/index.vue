<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <!-- <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item> -->
        <el-form-item :label="t('category.categoryName')">
         <el-input v-model="searchInfo.categoryName" placeholder="" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">{{t('category.search')}}</el-button>
          <el-button size="small" icon="refresh" @click="onReset">{{t('category.reset')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">{{t('category.create')}}</el-button>
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
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        header-cell-class-name="head-class"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" :label="t('category.createTime')" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="ID" prop="ID"  />
        <el-table-column align="left" :label="t('category.categoryName')" prop="categoryName"  />
        <!-- <el-table-column align="left" label="分类图片" prop="categoryPic" width="120" /> -->
        <el-table-column align="left" :label="t('videoDownload.operate')">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateVideoCategoryFunc(scope.row)">{{t('category.change')}}</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">{{t('category.delete')}}</el-button>
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
        <el-form-item :label="t('category.categoryName')"  prop="categoryName" >
          <el-input v-model="formData.categoryName" :clearable="true"  placeholder="" />
        </el-form-item>
        <!-- <el-form-item label="分类图片:"  prop="categoryPic" >
          <el-input v-model="formData.categoryPic" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">{{t('category.cancel')}}</el-button>
          <el-button size="small" type="primary" @click="enterDialog">{{t('category.confirm')}}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'VideoCategory'
}
</script>

<script setup>
import {
  createVideoCategory,
  deleteVideoCategory,
  deleteVideoCategoryByIds,
  updateVideoCategory,
  findVideoCategory,
  getVideoCategoryList
} from '@/api/videoCategory'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        categoryName: '',
        categoryPic: '',
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
  const table = await getVideoCategoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteVideoCategoryFunc(row)
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
      const res = await deleteVideoCategoryByIds({ ids })
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
const updateVideoCategoryFunc = async(row) => {
    const res = await findVideoCategory({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.revideoCategory
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteVideoCategoryFunc = async (row) => {
    const res = await deleteVideoCategory({ ID: row.ID })
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
        categoryName: '',
        categoryPic: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createVideoCategory(formData.value)
                  break
                case 'update':
                  res = await updateVideoCategory(formData.value)
                  break
                default:
                  res = await createVideoCategory(formData.value)
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
.head-class{
    font-weight: bold;
    font-size: 20px; 
}
</style>
