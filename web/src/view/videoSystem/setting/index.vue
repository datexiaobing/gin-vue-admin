<template>
    <div>
        <el-card>
            <el-form :model="formData" label-position="left" ref="elFormRef" :rules="rule" label-width="120px">

                <el-form-item :label="t('transform.transResolution')"  prop="transResolution" >
                <!-- <el-input v-model="formData.transResolution" :clearable="true"  placeholder="请输入" /> -->
                <el-select v-model="formData.transResolution" placeholder="">
                    <el-option
                    v-for="item in transResolutionOption"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                    </el-option>
                </el-select>
                </el-form-item>
                <el-form-item :label="t('videoList.oss')">
                    <el-radio-group v-model="formData.ossStatus" @change="changeOss">
                        <el-radio label="1"  border>Qiniu</el-radio>
                        <el-radio label="2"  border>Ali</el-radio>
                    </el-radio-group>
                </el-form-item>

                <el-form-item :label="t('videoList.SeektimeHeard')"  prop="transSeektimeHeard" >
                <!-- <el-input v-model.number="formData.transSeektimeHeard" :clearable="true" placeholder="请输入" /> -->
                <el-slider v-model="formData.transSeektimeHeard" />
                </el-form-item>
                <el-form-item :label="t('videoList.SeektimeTail')"  prop="transSeektimeTail" >
                <!-- <el-input v-model.number="formData.transSeektimeTail" :clearable="true" placeholder="请输入" /> -->
                <el-slider v-model="formData.transSeektimeTail" />
                </el-form-item>
                <el-form-item :label="t('videoList.paoma')"   >
                <el-switch v-model="paoma" size="large" @change="changePao"/>
                </el-form-item>
                <el-form-item :label="t('videoList.drawtextString')"  prop="transDrawtextString" v-show="paoma">
                <el-input v-model="formData.transDrawtextString" :clearable="true"  />
                </el-form-item>
                <el-form-item :label="t('videoList.drawtextPosition')"  prop="transDrawtextPosition" v-show="paoma">
                <!-- <el-input v-model="formData.transDrawtextPosition" :clearable="true"  /> -->
                <el-radio-group v-model.number="formData.transDrawtextPosition" >
                    <el-radio-button label="1">{{t('videoList.up')}}</el-radio-button>
                    <el-radio-button label="2">{{t('videoList.down')}}</el-radio-button>
                </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('videoList.drawtextColor')"  prop="transDrawtextColor"  v-show="paoma">
                <!-- <el-input v-model="formData.transDrawtextColor" :clearable="true"  placeholder="请输入" /> -->
                <el-radio-group v-model.number="formData.transDrawtextColor" >
                    
                    <el-radio-button label="aqua" />
                    <el-radio-button label="black" />
                    <el-radio-button label="blue" />
                    <el-radio-button label="fuchsia" />
                    <el-radio-button label="gray" />
                    <el-radio-button label="green" />           
                    <el-radio-button label="lime" />
                    <el-radio-button label="maroon" />
                    <el-radio-button label="navy" />
                    <el-radio-button label="olive" />
                    <el-radio-button label="purple" />
                    <el-radio-button label="red" />   
                    <el-radio-button label="silver" />
                    <el-radio-button label="teal" />
                    <el-radio-button label="white" />            

                </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('videoList.drawtextFontsize')"  prop="transDrawtextFontsize" v-show="paoma">
                <!-- <el-input v-model.number="formData.transDrawtextFontsize" :clearable="true" placeholder="请输入" /> -->
                <el-slider v-model="formData.transDrawtextFontsize" />
                </el-form-item>
                <el-form-item :label="t('videoList.drawtextDuration')"  prop="transDrawtextDuration" v-show="paoma" >
                <!-- <el-input v-model="formData.transDrawtextDuration" :clearable="true"  placeholder="请输入" /> -->
                <el-slider v-model="formData.transDrawtextDuration" />
                </el-form-item>
                <el-form-item :label="t('videoList.drawtextInterval')"  prop="transDrawtextInterval" v-show="paoma">
                    <el-slider v-model="formData.transDrawtextInterval" />
                    <!-- <el-input v-model="formData.transDrawtextInterval" :clearable="true"  placeholder="请输入" /> -->
                </el-form-item>
                <!-- ali oss keys -->
                <el-form-item :label="t('videoList.ali')"    >
                    <el-switch v-model="ali" size="large" @change="changeAli"/>
                </el-form-item>
                <el-form-item :label="t('ali.endPoint')" v-show="ali">
                    <el-input v-model="formData.aliEndPoint" :clearable="true"  placeholder="" />
                </el-form-item>
                <el-form-item :label="t('ali.accessKeyId')" v-show="ali">
                    <el-input v-model="formData.aliAccessKeyId" :clearable="true"  placeholder="" />
                </el-form-item>
                <el-form-item :label="t('ali.accessKeySecret')" v-show="ali">
                    <el-input v-model="formData.aliAccessKeySecret" :clearable="true"  placeholder="" />
                </el-form-item>
                <el-form-item :label="t('ali.bucketName')" v-show="ali">
                    <el-input v-model="formData.aliBucketName" :clearable="true"  placeholder="" />
                </el-form-item> 
                <el-form-item >
                    <el-button  size="large" @click="onReset">{{t('category.reset')}}</el-button>
                    <el-button type="primary" size="large" @click="onSubmit">{{t('category.confirm')}}</el-button>
                </el-form-item>
            </el-form>
            
            
        </el-card>
    </div>
</template>

<script setup>
import {
  createCloudConfig,
  deleteCloudConfig,
  deleteCloudConfigByIds,
  updateCloudConfig,
  findCloudConfig,
  getCloudConfigList
} from '@/api/cloudConfig'
import { ElMessage, ElMessageBox } from 'element-plus'
import {ref} from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const paoma =ref(false)
const ali =ref(false)
const radio1 = ref('0')

const changeOss =  (e)=>{
    // if(e==='2'){
    //   formData.value.transOssAliStatus=2
    //   formData.value.transOssQiniuStatus=1
    // }else if(e==='1'){
    //   formData.value.transOssQiniuStatus=2
    //   formData.value.transOssAliStatus=1
    // }
    console.log(e)
}


const transResolutionOption =ref([{label:'1080P',value:3},{label:'720P',value:2},{label:'360p',value:1}])

const getConfig =async ()=>{
  const d= await findCloudConfig({ID:1})
  if(d.code===0){
    formData.value=d.data.recloudConfig
  }
}

getConfig()

const changePao=(val)=>{
  paoma.value = val
}
const changeAli=(val)=>{
  ali.value = val
}

const formData = ref(
      {
        transDrawtextSwitch:1, //是否开启跑马灯
        transDrawtextColor: '',
        transDrawtextDuration: 0,
        transDrawtextFontsize: 0,
        transDrawtextInterval: 0,
        transDrawtextPosition: 0,
        transDrawtextString: '',
        transResolution: 1,
        transSeektimeHeard: 0,
        transSeektimeTail: 0,
        aliEndPoint:'',
        aliAccessKeyId:'',
        aliAccessKeySecret:'',
        aliBucketName:''
        })

const onSubmit = async()=>{
    const d=  await createCloudConfig(formData.value)
    if(d.code===0){
        ElMessage({
            type:'success',
            message:'修改成功'
        })
    }
}

const onReset=()=>{
    formData.value= {
        transDrawtextSwitch:1, //是否开启跑马灯
        transDrawtextColor: '',
        transDrawtextDuration: 0,
        transDrawtextFontsize: 0,
        transDrawtextInterval: 0,
        transDrawtextPosition: 0,
        transDrawtextString: '',
        transResolution: 1,
        transSeektimeHeard: 0,
        transSeektimeTail: 0,
    }
}   

</script>
