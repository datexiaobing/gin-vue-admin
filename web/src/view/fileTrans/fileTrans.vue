<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="转换后的名">
         <el-input v-model="searchInfo.transOutName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="1正在2完成">
            
             <el-input v-model.number="searchInfo.transStatus" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="分类">
            
             <el-input v-model.number="searchInfo.transType" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="0等待上传1已经上传">
            
             <el-input v-model.number="searchInfo.transOssStatus" placeholder="搜索条件" />

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
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="传入视频名" prop="transInputName" width="120" />
        <el-table-column align="left" label="转换后的名" prop="transOutName" width="120" />
        <el-table-column align="left" label="uuid" prop="transUuid" width="120" />
        <el-table-column align="left" label="1正在2完成" prop="transStatus" width="120" />
        <el-table-column align="left" label="分类" prop="transType" width="120" />
        <el-table-column align="left" label="专辑视频数量" prop="transTypeNum" width="120" />
        <el-table-column align="left" label="分辨率" prop="transResolution" width="120" />
        <el-table-column align="left" label="时长" prop="transDuration" width="120" />
        <el-table-column align="left" label="跳过片头秒" prop="transSeektimeHeard" width="120" />
        <el-table-column align="left" label="跳过片尾" prop="transSeektimeTail" width="120" />
        <el-table-column align="left" label="跑马灯颜色" prop="transDrawtextColor" width="120" />
        <el-table-column align="left" label="跑马灯位置" prop="transDrawtextPosition" width="120" />
        <el-table-column align="left" label="滚动时长" prop="transDrawtextDuration" width="120" />
        <el-table-column align="left" label="滚动间隔" prop="transDrawtextInterval" width="120" />
        <el-table-column align="left" label="跑马灯文字" prop="transDrawtextString" width="120" />
        <el-table-column align="left" label="文字大小" prop="transDrawtextFontsize" width="120" />
        <el-table-column align="left" label="字幕文件" prop="transSubtitle" width="120" />
        <el-table-column align="left" label="字幕开关" prop="transSubtitleSwitch" width="120" />
        <el-table-column align="left" label="进度状态" prop="transProgressRate" width="120" />
        <el-table-column align="left" label="0等待上传1已经上传" prop="transOssStatus" width="120" />
        <el-table-column align="left" label="上传错误" prop="transOssError" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateFileTransFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="传入视频名:"  prop="transInputName" >
          <el-input v-model="formData.transInputName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="转换后的名:"  prop="transOutName" >
          <el-input v-model="formData.transOutName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="uuid:"  prop="transUuid" >
          <el-input v-model="formData.transUuid" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="1正在2完成:"  prop="transStatus" >
          <el-input v-model.number="formData.transStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类:"  prop="transType" >
          <el-input v-model.number="formData.transType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专辑视频数量:"  prop="transTypeNum" >
          <el-input v-model.number="formData.transTypeNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分辨率:"  prop="transResolution" >
          <el-input v-model="formData.transResolution" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="时长:"  prop="transDuration" >
          <el-input v-model="formData.transDuration" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跳过片头秒:"  prop="transSeektimeHeard" >
          <el-input v-model.number="formData.transSeektimeHeard" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跳过片尾:"  prop="transSeektimeTail" >
          <el-input v-model.number="formData.transSeektimeTail" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯颜色:"  prop="transDrawtextColor" >
          <el-input v-model="formData.transDrawtextColor" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯位置:"  prop="transDrawtextPosition" >
          <el-input v-model="formData.transDrawtextPosition" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="滚动时长:"  prop="transDrawtextDuration" >
          <el-input v-model="formData.transDrawtextDuration" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="滚动间隔:"  prop="transDrawtextInterval" >
          <el-input v-model="formData.transDrawtextInterval" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯文字:"  prop="transDrawtextString" >
          <el-input v-model="formData.transDrawtextString" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文字大小:"  prop="transDrawtextFontsize" >
          <el-input v-model.number="formData.transDrawtextFontsize" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="字幕文件:"  prop="transSubtitle" >
          <el-input v-model="formData.transSubtitle" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="字幕开关:"  prop="transSubtitleSwitch" >
          <el-input v-model.number="formData.transSubtitleSwitch" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="进度状态:"  prop="transProgressRate" >
          <el-input-number v-model="formData.transProgressRate"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="0等待上传1已经上传:"  prop="transOssStatus" >
          <el-input v-model.number="formData.transOssStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上传错误:"  prop="transOssError" >
          <el-input v-model="formData.transOssError" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
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
  name: 'FileTrans'
}
</script>

<script setup>
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
</script>

<style>
</style>
