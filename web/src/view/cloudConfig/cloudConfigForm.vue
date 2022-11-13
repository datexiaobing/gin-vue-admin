<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="跑马灯颜色:" prop="transDrawtextColor">
          <el-input v-model="formData.transDrawtextColor" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="滚动时长:" prop="transDrawtextDuration">
          <el-input v-model.number="formData.transDrawtextDuration" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯大小:" prop="transDrawtextFontsize">
          <el-input v-model.number="formData.transDrawtextFontsize" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="滚动间隔:" prop="transDrawtextInterval">
          <el-input v-model.number="formData.transDrawtextInterval" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯位置:" prop="transDrawtextPosition">
          <el-input v-model.number="formData.transDrawtextPosition" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="滚动文字:" prop="transDrawtextString">
          <el-input v-model="formData.transDrawtextString" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分辨率:" prop="transResolution">
          <el-input v-model.number="formData.transResolution" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跳过片头:" prop="transSeektimeHeard">
          <el-input v-model.number="formData.transSeektimeHeard" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跳过片尾:" prop="transSeektimeTail">
          <el-input v-model.number="formData.transSeektimeTail" :clearable="true" placeholder="请输入" />
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
  name: 'CloudConfig'
}
</script>

<script setup>
import {
  createCloudConfig,
  updateCloudConfig,
  findCloudConfig
} from '@/api/cloudConfig'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            transDrawtextColor: '',
            transDrawtextDuration: 0,
            transDrawtextFontsize: 0,
            transDrawtextInterval: 0,
            transDrawtextPosition: 0,
            transDrawtextString: '',
            transResolution: 0,
            transSeektimeHeard: 0,
            transSeektimeTail: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCloudConfig({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recloudConfig
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
               res = await createCloudConfig(formData.value)
               break
             case 'update':
               res = await updateCloudConfig(formData.value)
               break
             default:
               res = await createCloudConfig(formData.value)
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
