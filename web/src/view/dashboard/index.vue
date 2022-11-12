<template>
  <div>
    <el-row :gutter="8">
      <el-col :span="4" v-for="(v,i) in box_data" :key="i">
        <el-card :body-style="{ padding: '0px'}">
          <div class="content">
            <div class="left" :style="v.color">
              <!-- font-size: 36px; color: rgb(255, 255, 255); -->
              <div class="ico">
              <el-icon v-if="i<3" size="46px" color="rgb(255, 255, 255)" ><HelpFilled /></el-icon>
              <el-icon v-else size="46px" color="rgb(255, 255, 255)"><Download /></el-icon>
              </div>
            </div>
            <div class="right">
              <div class="right-content">
                  <p style="font-size:40px;text-align:center">
                    {{v.count}}
                  </p>
                  <span style="margin-top:15%;text-align:center">{{v.content}}</span>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {
getVideosStatus
} from '@/api/videoList'

const getData =async ()=>{
  const d = await getVideosStatus()
  // console.log(d)
  let dd=d.data.list
  box_data.value=[{
  count:dd.activiteNum,
  color:'background: rgb(25, 190, 107); ',
  content:'正在转码'
},{
  count:dd.waitingNum,
  color:'background: rgb(255, 153, 0); ',
  content:'转码等待'
},{
  count:dd.doneNum,
  color:'background: rgb(237, 64, 20); ',
  content:'已完成/已停止'
},
{
  count:dd.numActive,
  color:'background: rgb(25, 190, 107); ',
  content:'正在下载'
},{
  count:dd.numWaiting,
  color:'background: rgb(255, 153, 0); ',
  content:'下载等待'
},{
  count:dd.numStopped,
  color:'background: rgb(237, 64, 20); ',
  content:'已完成/已停止'
},
]
 
}

getData()

let box_data = ref(null)

</script>

<style scoped>

.content{
  display: flex;
  justify-content: space-around;
}
.left{
  width: 36%;
  text-align: center;
  height: 110px;
  display: flex;
}
.right{
  width: 64%;
  display: flex;
}
.ico{
  margin: auto;
}
.right-content{
  margin: auto;
  display: inline-flex;
  flex-direction: column;
  justify-content: space-around;
}

</style>