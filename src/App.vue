<template>
  <div class="w-[90%] lg:w-[70%] py-10 mx-auto">
    <DataTable :columns="columns" :processes="processes" />
  </div>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { columns } from '@/components/processes/columns.js'
  import DataTable from '@/components/processes/data-table.vue'
  import axios from 'axios'

  const processes = ref([])

  onMounted(() => {
    axios.get('http://localhost:4000/processes')
      .then((response) => {
        processes.value = response.data
      })
      .catch((error) => toast(error.data.error))
  })
</script>
