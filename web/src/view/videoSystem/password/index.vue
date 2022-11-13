<template>
<el-col :span="12">
    <el-card class="box-card">
        <template #header>
        <div class="card-header">
            <span>修改密码</span>
        </div>
        </template>
        <el-form class="user-account-key" ref="form" :model="form" label-width="100px">
            <!-- <el-form-item label="用户名" prop="password">
                <el-input type="text" placeholder="请输入用户名" v-model="formData.username"></el-input>
            </el-form-item> -->
            <el-form-item label="原密码" prop="password">
                <el-input type="password" placeholder="请输入原密码" v-model="formData.password"></el-input>
            </el-form-item>
            <el-form-item label="新密码" prop="newPassword">
                <el-input type="password" placeholder="请输入新密码" v-model="formData.newPassword"></el-input>
            </el-form-item>
            <!-- <el-form-item label="手机号" prop="phone">
                <el-input type="text" placeholder="请设置手机号" v-model="formData.phone"></el-input>
            </el-form-item> -->
            <el-form-item>
                <el-button type="primary" @click="onSubmit">修改</el-button>
                <el-button @click="reset">重置</el-button>
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