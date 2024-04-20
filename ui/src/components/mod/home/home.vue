<template>
  <div class="home_div_container">
    <div
        class="common-header">
      <div>
        <el-button @click="GotoDown">下载中心</el-button>
        <el-button @click="GetPicList">重新加载</el-button>
        <el-button @click="IsEdit=!IsEdit">{{ IsEdit ? '退出编辑模式' : '编辑模式' }}</el-button>
      </div>
      <div style="display: flex;flex-direction: row;">
        <el-input v-model="NowSelectPara.search"></el-input>
        <el-button @click="GetPicList" style="width: 100px;margin-left: 10px;margin-right: 20px">搜索</el-button>
      </div>
    </div>

    <div class="pic_list_container">
      <el-scrollbar view-class="pic_list_container_2">
        <div v-for="item in NowPicList"
             class="common-flex-center common-border"
             style="margin: 2px;flex-direction: column">
          <div @click="SelectPic(item)" style="width: 200px;height: 250px;">
            <div class="common-flex-center"
                 style="height: 50px;width: 100%;background: rgba(0,255,53,0.66)">
              <el-text truncated style="width: 95%;text-align: center" >
                {{ item.class_name }}
              </el-text>
            </div>
            <div style="width: 100%;height: 200px;display: flex;align-items: center;justify-content: center">
              <el-text v-if="item.picList==null || item.picList.length==0">
                无图片
              </el-text>
              <el-image v-if="item.picList!=null && item.picList.length>0" fit="cover" style="width: 95%;height:95%"
                        :src="GetPicUrl(item.picList[0])"
              />
            </div>

          </div>
          <div v-if="IsEdit" style="width: 80%;height: 40px;display: flex;">
            <el-button style="width: 50%;background: rgba(204,139,52,0.79)"
                       @click="ClickEditPicBtn(item)">编辑
            </el-button>
            <el-button style="width: 50%;background: rgba(220,76,76,0.79)"
                       @click="ClickDelPicBtn(item)">删除
            </el-button>
          </div>
        </div>


      </el-scrollbar>
    </div>
  </div>
  <!-- 查看弹框 -->
  <el-dialog
      v-model="IsSelectVisible"
      fullscreen
  >
    <div style="position: fixed;left: 10px;top: 10px;z-index: 9999">
      <el-button
          @click="IsSelectVisible=false"
          style="width: 100px;height: 50px;background: red;">
        关闭
      </el-button>

    </div>
    <div style="display: flex;flex-wrap: wrap;align-items: flex-start;justify-content: center;z-index: 1000;">
      <el-image v-for="item in NowPicSelect.picList" :src="GetPicUrl(item)"
                style="width: 675px;height: 900px;border: antiquewhite solid 2px ;margin: 2px"/>
    </div>
  </el-dialog>
  <!-- 删除确认弹框 -->
  <el-dialog
      v-model="IsDelVisible"
      title="删除确认"
      width="30%"
      center>
    <span style="width: 30%;text-align: center">
      确认删除 {{ NowPicSelect.class_name }} ?
    </span>
    <template #footer>
      <el-button @click="IsDelVisible = false">取消</el-button>
      <el-button type="primary" @click="ClickDelPicSubBtn">
        确认
      </el-button>
    </template>
  </el-dialog>

  <!-- 编辑弹框 -->
  <el-dialog
      v-model="IsEditVisible"
      title="修改名称"
      width="30%"
      center>
    <div>
      <el-form>
        <el-form-item label="名称">
          <el-input v-model="EditPicReq.name"></el-input>
        </el-form-item>

      </el-form>
    </div>
    <template #footer>
      <el-button @click="IsEditVisible = false">取消</el-button>
      <el-button type="primary" @click="ClickEditPicSubBtn">
        确认
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import './home.css'
import {GotoDown} from "@/components/mod/down/down.ts";
import {onMounted, ref} from "vue";
import {
  ApiDeleteGuitarClassServer,
  ApiGetGuitarPicListServer, ApiRenameGuitarClassServer, ApiRenameGuitarClassServerReq, CopyGuitarPicListResult,
  GetGuitarPicListReq,
  GetGuitarPicListResult, NewApiRenameGuitarClassServerReq,
  NewGetGuitarPicListResult
} from "@/components/api/pic.ts";
import {runConfig} from "@/components/common/config.ts";

const NowPicList = ref<GetGuitarPicListResult[]>([])
const NowPicSelect = ref<GetGuitarPicListResult>(NewGetGuitarPicListResult())
const NowSelectPara = ref<GetGuitarPicListReq>({} as GetGuitarPicListReq)
const IsSelectVisible = ref(false)
const IsEdit = ref(false)
const IsDelVisible = ref(false)
const IsEditVisible = ref(false)
const EditPicReq = ref(NewApiRenameGuitarClassServerReq())
const GetPicList = () => {
  ApiGetGuitarPicListServer(NowSelectPara.value).then(res => {
    if (res.state === 0) {
      NowPicList.value = res.data.list
      NowPicSelect.value = NewGetGuitarPicListResult()
    }
  })
}

const SelectPic = (item: GetGuitarPicListResult) => {
  NowPicSelect.value = item
  IsSelectVisible.value = true
}

const GetPicUrl = (path: string): string => {
  return runConfig.server + '/' + path
}
onMounted(() => {
  GetPicList()
})

const ClickDelPicBtn = (item: GetGuitarPicListResult) => {
  NowPicSelect.value = CopyGuitarPicListResult(item)
  IsDelVisible.value = true
}

const ClickDelPicSubBtn = () => {
  ApiDeleteGuitarClassServer(NowPicSelect.value.class_name).then(res => {
    if (res.state === 0) {
      GetPicList()
    }
    IsDelVisible.value = false
  })
}

const ClickEditPicBtn = (item: GetGuitarPicListResult) => {
  EditPicReq.value.oldName = item.class_name
  EditPicReq.value.name = item.class_name
  IsEditVisible.value = true
}

const ClickEditPicSubBtn = () => {
  ApiRenameGuitarClassServer(EditPicReq.value).then(res => {
    if (res.state === 0) {
      GetPicList()
    }
    IsEditVisible.value = false
  })
}

</script>