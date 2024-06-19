<template>
  <div>
    <!-- <el-form :model="formData" ref="form" label-width="auto" :inline="false" size="large">
      <el-form-item label="原始地址：">
        <el-input-number v-model="formData.old_address" type="number"> </el-input-number>
        <el-button type="primary" @click="test">测试</el-button> -->

    <!-- 测试按钮 -->
    <!-- </el-form-item>
      <el-form-item label="新地址：">
        <el-input-number v-model="formData.new_address" type="number"></el-input-number>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">修改地址</el-button>
      </el-form-item>
    </el-form> -->
    <el-space wrap>
      <el-tag size="large" effect="dark" v-for="i in 99" :type="addresses.indexOf(i) > -1 ? 'success' : 'info'">
        {{ padZero(i, 2) }}
      </el-tag>
    </el-space>

    <el-button type="primary" @click="test">扫描</el-button>
  </div>
</template>
<script lang="ts" setup>
import { reactive, ref } from "vue";
import { DfddfScanAddress } from "../../wailsjs/go/main/App";
import { ElMessage } from "element-plus";

const formData = reactive({
  old_address: 1,
  new_address: 0,
});

const onSubmit = () => {
  console.log(formData.old_address, formData.new_address);
};
const addresses = ref([1, 2]);

const test = () => {
  addresses.value = [];
  DfddfScanAddress().then((res) => {
    if (res.status == "error") {
      ElMessage.error(res.msg);
      return;
    }
    if (res.addresses.length == 0) {
      ElMessage.error("没有扫描到设备");
      return;
    }
    addresses.value = res.addresses;
  });
};
function padZero(number: number, length: number) {
  var str = "" + number;
  while (str.length < length) {
    str = "0" + str;
  }
  return str;
}
</script>
