<script setup>
	import { Trash2 } from 'lucide-vue-next'
	import { Button } from '@/components/ui/button'
	import { toast } from 'vue3-toastify';
	import axios from 'axios'

	const props = defineProps({
		id: Number
	})

	const killAllProcess = () => {
		axios.post(`http://localhost:4000/processes/kill`)
			.then((response) => toast(response.data.message))
			.catch((error) => toast(error.data.error))
	}
</script>

<template>
	<Button variant="outline" size="" class="bg-destructive foreground-destructive" @click="killAllProcess">
    	Kill All Processes
  	</Button>
</template>