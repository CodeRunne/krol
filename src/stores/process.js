import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'

export const useProcessStore = defineStore('processes', () => {
	const processes = ref([])
	const errors = ref([])

	const getProcesses = () => {
		axios.get('http://localhost:4000/processes')
			.then((response) => {
				processes.value = response.data
			})
	}

	const deleteProcess = (id) => {
		if (!window.confirm('Are you sure')) {
			return
		}
		axios.post(`http://localhost:4000/processes/${id}/kill`)
			.then((response) => {
				getProcesses()
				alert('process deleted successfully!')
			})
			.catch((error) => {
				alert(error)
			})
	}

	const killAllProcess = () => {
		if (!window.confirm('Are you sure')) {
			return
		}
		axios.post(`http://localhost:4000/processes/kill`)
			.then((response) => {
				alert('All process terminated successfully!')
			})
			.catch((error) => alert(error))
	}

	return { getProcesses, deleteProcess, killAllProcess, processes, errors }
})