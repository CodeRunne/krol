<script setup>
	import { onUpdated } from 'vue'
	import { Trash2 } from 'lucide-vue-next'
	import { Button } from '@/components/ui/button'
	import { useProcessStore } from '@/stores/process'

	const store = useProcessStore()
	const props = defineProps({
		id: Number
	})
</script>

<template>
	<Button variant="outline" size="" class="bg-destructive foreground-destructive border-destructive" @click="store.killAllProcess()">
    	Kill All Processes
  	</Button>
</template>