<script setup>
	import { Trash2 } from 'lucide-vue-next'
	import { Button } from '@/components/ui/button'
	import axios from 'axios'

	const props = defineProps({
		id: Number
	})

	const deleteProcess = (id) => {
		axios.post(`http://localhost:4000/processes/${id}/kill`)
			.then((response) => { })
			.catch((error) => console.log(error))
	}
</script>

<template>
	<Button variant="outline" size="icon" class="hover:bg-destructive hover:foreground-destructive" @click="deleteProcess(id)">
    	<Trash2 class="w-3.5 h-3.5" />
  	</Button>
</template>