<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="传入视频名:" prop="transInputName">
          <el-input v-model="formData.transInputName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="转换后的名:" prop="transOutName">
          <el-input v-model="formData.transOutName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="uuid:" prop="transUuid">
          <el-input v-model="formData.transUuid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="1正在2完成:" prop="transStatus">
          <el-input v-model.number="formData.transStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类:" prop="transType">
          <el-input v-model.number="formData.transType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专辑视频数量:" prop="transTypeNum">
          <el-input v-model.number="formData.transTypeNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分辨率:" prop="transResolution">
          <el-input v-model="formData.transResolution" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="时长:" prop="transDuration">
          <el-input v-model="formData.transDuration" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跳过片头秒:" prop="transSeektimeHeard">
          <el-input v-model.number="formData.transSeektimeHeard" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跳过片尾:" prop="transSeektimeTail">
          <el-input v-model.number="formData.transSeektimeTail" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯颜色:" prop="transDrawtextColor">
          <el-input v-model="formData.transDrawtextColor" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯位置:" prop="transDrawtextPosition">
          <el-input v-model="formData.transDrawtextPosition" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="滚动时长:" prop="transDrawtextDuration">
          <el-input v-model="formData.transDrawtextDuration" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="滚动间隔:" prop="transDrawtextInterval">
          <el-input v-model="formData.transDrawtextInterval" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跑马灯文字:" prop="transDrawtextString">
          <el-input v-model="formData.transDrawtextString" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文字大小:" prop="transDrawtextFontsize">
          <el-input v-model.number="formData.transDrawtextFontsize" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="字幕文件:" prop="transSubtitle">
          <el-input v-model="formData.transSubtitle" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="字幕开关:" prop="transSubtitleSwitch">
          <el-input v-model.number="formData.transSubtitleSwitch" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="进度状态:" prop="transProgressRate">
          <el-input-number v-model="formData.transProgressRate" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="0等待上传1已经上传:" prop="transOssStatus">
          <el-input v-model.number="formData.transOssStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上传错误:" prop="transOssError">
          <el-input v-model="formData.transOssError" :clearable="true" placeholder="请输入" />
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
  name: 'FileTrans'
}
</script>

<script setup>
import {
  createFileTrans,
  updateFileTrans,
  findFileTrans
} from '@/api/fileTrans'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFileTrans({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refileTrans
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
