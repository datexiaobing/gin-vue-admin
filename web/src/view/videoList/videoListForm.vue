<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="videoTitle">
          <el-input v-model="formData.videoTitle" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="存储地址:" prop="videoDownloadPath">
          <el-input v-model="formData.videoDownloadPath" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="下载链接:" prop="videoUrl">
          <el-input v-model="formData.videoUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="封面:" prop="videoPic">
          <el-input v-model="formData.videoPic" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="1未下载2暂停下载3下载完成:" prop="videoDownloadStatus">
          <el-input v-model.number="formData.videoDownloadStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总大小:" prop="videoSize">
          <el-input-number v-model="formData.videoSize" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="已下载大小:" prop="videoDownloadSize">
          <el-input-number v-model="formData.videoDownloadSize" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="下载速率:" prop="videoDownloadSpeed">
          <el-input-number v-model="formData.videoDownloadSpeed" :precision="2" :clearable="true"></el-input-number>
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
  name: 'VideoList'
}
</script>

<script setup>
import {
  createVideoList,
  updateVideoList,
  findVideoList
} from '@/api/videoList'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVideoList({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.revideoList
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
