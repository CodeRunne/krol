<template>
  <div class="w-[90%] lg:w-[70%] py-10 mx-auto">
    <DataTable :columns="columns" :processes="store.processes" />
  </div>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { columns } from '@/components/processes/columns.js'
  import DataTable from '@/components/processes/data-table.vue'
  import { useProcessStore } from '@/stores/process'

  const store = useProcessStore()
  onMounted(() => {
    store.getProcesses()
  })
</script>
