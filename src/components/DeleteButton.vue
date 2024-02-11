<script setup>
	import axios from 'axios'
	import { Trash2 } from 'lucide-vue-next'
	import { Button } from '@/components/ui/button'
	import { useProcessStore } from '@/stores/process.js'

	const store = useProcessStore()
	const props = defineProps({
		id: Number
	})
</script>

<template>
	<Button variant="outline" size="icon" class="hover:bg-destructive hover:foreground-destructive" @click="store.deleteProcess(id)">
    	<Trash2 class="w-3.5 h-3.5" />
  	</Button>
</template>