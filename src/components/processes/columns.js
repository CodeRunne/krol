import { h } from 'vue'
import DeleteButton from '@/components/DeleteButton.vue'

export const columns = [
	{
		accessorKey: 'id',
		header: 'S/N'
	}, 
	{
		accessorKey: 'name',
		header: 'Name'
	},
	{
		accessorKey: 'pid',
		header: 'PID'
	},
	{
		accessorKey: 'ppid',
		header: 'PPID',
		cell: info => {
			if (info.getValue() == 1) {
				return 'root'
			} else if(info.getValue() == 2) {
				return 'user'
			} else {
				return info.getValue()
			}
		}
	},
	{
		accessorKey: 'Action',
		header: ' ',
		cell: ({ row }) => h(DeleteButton, { id: row.original.id } )
	}
]