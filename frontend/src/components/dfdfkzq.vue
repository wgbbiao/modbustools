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
      <el-tag size="large" style="cursor: pointer" effect="dark" v-for="i in 99" :type="addresses.indexOf(i) > -1 ? 'success' : 'info'" @click="testOne(i)">
        {{ padZero(i, 2) }}
      </el-tag>
    </el-space>

    <div>{{ scanIdex }}/99</div>

    <el-button type="primary" @click="test" v-if="scanIng == false">扫描</el-button>
    <el-button type="primary" @click="stop" v-if="scanIng">停止</el-button>
  </div>
</template>
<script lang="ts" setup>
import { reactive, ref } from "vue";
import { DfddfTestAddress } from "../../wailsjs/go/main/App";
import { ElMessage } from "element-plus";

const formData = reactive({
  old_address: 1,
  new_address: 0,
});
const scanIng = ref(false);

const onSubmit = () => {
  console.log(formData.old_address, formData.new_address);
};
const addresses = ref<Number[]>([]);
const scanIdex = ref(0);
const test = async () => {
  addresses.value = [];
  scanIng.value = true;
  for (var i = 1; i <= 99; i++) {
    if (addresses.value.indexOf(i) > -1) {
      // 删除
      addresses.value = addresses.value.filter((item) => item != i);
    }
    console.log(i);
    if (!scanIng.value) {
      console.log("stop");
      break;
    }
    scanIdex.value = i;
    var r = await DfddfTestAddress(i);
    console.log(r);
    if (r.status == "error") {
      ElMessage.error("址" + i + "失败：" + r.msg);
      continue;
    } else {
      ElMessage.success(r.msg);
    }
    addresses.value.push(i);
  }
  console.log("1sdfasdfasdf");
  scanIng.value = false;
};

const stop = () => {
  scanIng.value = false;
};
function padZero(number: number, length: number) {
  var str = "" + number;
  while (str.length < length) {
    str = "0" + str;
  }
  return str;
}

const testOne = async (i: number) => {
  if (addresses.value.indexOf(i) > -1) {
    // 删除
    addresses.value = addresses.value.filter((item) => item != i);
  }
  var r = await DfddfTestAddress(i);
  console.log(r);
  if (r.status == "error") {
    ElMessage.error("址" + i + "失败：" + r.msg);
  } else {
    ElMessage.success(r.msg);
  }
  addresses.value.push(i);
};
</script>
