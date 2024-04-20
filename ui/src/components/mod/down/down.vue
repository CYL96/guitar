<template>
  <div class="common-flex-center" style="flex-direction: column;width: 100%">
    <div class="common-header">
      <el-button @click="GotoHome" style="width: 100px">
        返回
      </el-button>
    </div>
    <div class="common-flex-center">
      <el-button v-for="item in modList" style="width: 100px;height: 100px"
                 @click="OnClickBtn(item)"
      >
        {{ item.mod_name }}
      </el-button>
    </div>
  </div>

</template>

<script setup lang="ts">
import '../../common/common.css'
import {Back} from "@element-plus/icons-vue";
import {GotoHome} from "@/components/mod/home/home.ts";
import {GetGuitarModListResult} from "@/components/api/mod.ts";
import {ref} from "vue";


const OnClickBtn = (item: GetGuitarModListResult) => {
  window.open(item.mod_url, '_blank');
}
const modList = ref<GetGuitarModListResult[]>([
  {
    mod_id: 1,
    mod_name: "fox4",
    mod_url: "https://www.fox4.cn/",
  },
  {
    mod_id: 2,
    mod_name: "gtpso",
    mod_url: "https://www.gtpso.com/",
  },
  {
    mod_id: 3,
    mod_name: "cwguitar",
    mod_url: "https://www.cwguitar.cn/",
  }
])

</script>
