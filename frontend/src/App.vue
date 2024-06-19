<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <div style="padding-top: 40px">USB连接状态</div>
        <div style="padding-top: 40px">
          <el-tag type="info" size="large" v-if="usbStatus == false">未连接</el-tag>
          <el-tag type="success" size="large" v-else>已连接</el-tag>
        </div>
        <div style="padding-top: 40px">
          <el-button type="primary" size="large" @click="open">连接</el-button>
        </div>
      </el-aside>
      <el-main>
        <el-tabs type="border-card" v-model="activeName" class="demo-tabs" @tab-click="handleClick">
          <el-tab-pane label="东阀蝶阀控制器" name="first">
            <dfdfkzq></dfdfkzq>
          </el-tab-pane>
        </el-tabs>
      </el-main>
    </el-container>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { ElMessage, type TabsPaneContext } from "element-plus";
import Dfdfkzq from "./components/dfdfkzq.vue";
import { GetUsbStatus, OpenUsb } from "../wailsjs/go/main/App";

const activeName = ref("first");

const usbStatus = ref(false);

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};

var interval = null;
onMounted(() => {
  interval = setInterval(() => {
    GetUsbStatus().then((res) => {
      usbStatus.value = res.UsbStatus;
    });
  }, 1000);
});

const open = () => {
  OpenUsb().then((res) => {
    console.log(res);
    if (res.status == "success") {
      ElMessage.success("连接成功");
    } else {
      ElMessage.error(res.msg);
    }
  });
};
</script>
<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
