<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="0根目录1一层目录2二层目录:" prop="fileDegree">
          <el-input v-model.number="formData.fileDegree" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="1dir0文件:" prop="fileIsDir">
          <el-input v-model.number="formData.fileIsDir" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最后一层dir名:" prop="fileLastDir">
          <el-input v-model="formData.fileLastDir" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件（夹）名:" prop="fileName">
          <el-input v-model="formData.fileName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件路径:" prop="filePath">
          <el-input v-model="formData.filePath" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件大小:" prop="fileSize">
          <el-input-number v-model="formData.fileSize" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="1video2other:" prop="fileType">
          <el-input v-model="formData.fileType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FileDownload'
}
</script>

<script setup>
import {
  createFileDownload,
  updateFileDownload,
  findFileDownload
} from '@/api/fileDownload'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFileDownload({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refileDownload
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
