<template>
<el-col :span="12">
    <el-card class="box-card">
        <template #header>
        <div class="card-header">
            <span>{{t('password.change')}}</span>
        </div>
        </template>
        <el-form class="user-account-key" ref="form" :model="form" label-width="150px">
            <!-- <el-form-item label="用户名" prop="password">
                <el-input type="text" placeholder="请输入用户名" v-model="formData.username"></el-input>
            </el-form-item> -->
            <el-form-item :label="t('password.old')" prop="password">
                <el-input type="password" placeholder="" v-model="formData.password"></el-input>
            </el-form-item>
            <el-form-item :label="t('password.new')" prop="newPassword">
                <el-input type="password" placeholder="" v-model="formData.newPassword"></el-input>
            </el-form-item>
            <!-- <el-form-item label="手机号" prop="phone">
                <el-input type="text" placeholder="请设置手机号" v-model="formData.phone"></el-input>
            </el-form-item> -->
            <el-form-item>
                <el-button type="primary" @click="onSubmit">{{t('category.modify')}}</el-button>
                <el-button @click="reset">{{t('category.reset')}}</el-button>
            </el-form-item>
        </el-form>
    </el-card>
</el-col>

  
</template>
<script setup>
import { reactive, ref } from 'vue'
import {changePassword} from '@/api/user.js'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const router = useRouter()
const route = useRoute()
// const userData = ref()
// userData.value=route.params

const userStore = useUserStore()
const form =ref(null)
const formData=reactive({
    password:'',
    newPassword:'', 
})


const reset=()=>{
    formData.password=''
    formData.newPassword=''
}
const onSubmit = async()=>{
   
    if(formData.newPassword ===formData.password){
        ElMessage({
            type:"error",
            message:"不可以和之前的密码相同！"
        })
        return
    }

    // user/changePassword
    let res= await changePassword(formData)
    // console.log(res)
    if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '修改成功'
        })
    }
    reset()

}
</script>