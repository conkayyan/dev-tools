<template>
  <el-form :model="form" label-width="80px">
    <el-form-item label="Timestamp">
      <el-input v-model="form.timestamp" @change="setTimestamp" style="width: 240px" maxlength="10"/>
    </el-form-item>
    <el-form-item label="Timezone">
      <el-autocomplete
          v-model="form.timezone"
          placeholder="Select"
          style="width: 240px"
          :fetch-suggestions="querySearch"
          clearable
          @select="setTimezone"
      />
    </el-form-item>
    <el-form-item label="Format">
      <el-select
          v-model="form.layout"
          placeholder="Select"
          style="width: 240px"
          @change="setFormat"
      >
        <el-option
            v-for="item in formatOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="Datetime">
      <el-input v-model="form.datetime" @change="setDatetime" style="width: 240px"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="timeInit">Reset</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {GetTimeOption, SetDatetime, SetFormat, SetLocation, SetUnix, TimeInit} from '../../../wailsjs/go/main/App'
import jsonData from './timezones.json'

interface formOption {
  timestamp: number;
  timezone: string;
  layout: string;
  datetime: string;
}

const form: formOption = reactive({
  timestamp: 0,
  timezone: '',
  layout: '',
  datetime: ''
})

interface option {
  value: string;
  label: string;
}

const timezoneOptions = ref<option[]>([])

const formatOptions = ref<option[]>([])

const querySearch = (queryString: string, cb: any) => {
  const results = queryString
      ? timezoneOptions.value.filter(createFilter(queryString))
      : timezoneOptions.value
  cb(results)
}

const createFilter = (queryString: string) => {
  return (tz: option) => {
    return (
        tz.value.toLowerCase().indexOf(queryString.toLowerCase()) >= 0
    )
  }
}

const timeInit = () => {
  TimeInit().then(ret => {
    form.timestamp = ret.timestamp
    form.timezone = ret.location
    form.layout = ret.layout
    form.datetime = ret.datetime
    ret.formats.forEach((item) => {
      formatOptions.value.push({
        value: item.value,
        label: item.label,
      })
    })
  })
}

const getTimeOption = () => {
  GetTimeOption().then(ret => {
    form.timestamp = ret.timestamp
    form.timezone = ret.location
    form.layout = ret.layout
    form.datetime = ret.datetime
  })
}

const setTimestamp = () => {
  const timestamp: number = Number(form.timestamp)
  SetUnix(timestamp).then(() => {
    getTimeOption()
  })
}

const setTimezone = () => {
  SetLocation(form.timezone).then(() => {
    getTimeOption()
  })
}

const setFormat = () => {
  SetFormat(form.layout).then(() => {
    getTimeOption()
  })
}

const setDatetime = () => {
  SetDatetime(form.datetime).then(() => {
    getTimeOption()
  })
}

onMounted(() => {
  jsonData.forEach((item) => {
    item.utc.forEach((tz) => {
      timezoneOptions.value.push({
        value: tz,
        label: tz,
      })
    })
  })

  timeInit()
})

</script>

<style scoped>

</style>
