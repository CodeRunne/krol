<script setup>
	import {
	    FlexRender,
	    getCoreRowModel,
	    useVueTable,
	    getFilteredRowModel,
	    getPaginationRowModel
	} from "@tanstack/vue-table"

	import {
	    Table,
	    TableBody,
	    TableCell,
	    TableHead,
	    TableHeader,
	    TableRow,
	} from "@/components/ui/table"

	import {
	  Select,
	  SelectContent,
	  SelectItem,
	  SelectTrigger,
	  SelectValue,
	} from '@/components/ui/select'
	import { Input } from '@/components/ui/input'
	import { Button } from '@/components/ui/button'
	import KillButton from '@/components/KillButton.vue'

	import { 
		ChevronsLeft,
		ChevronsRight,
		ArrowLeft,
		ArrowRight
	} from 'lucide-vue-next'

	import { ref } from 'vue'

	const props = defineProps({
		processes: {
			type: Array,
			required: true
		},
		columns: {
			type: Array,
			required: true
		}
	})

	const filter = ref('')
	const table = useVueTable({
	    get data() { return props.processes },
    	get columns() { return props.columns },
	    getCoreRowModel: getCoreRowModel(),
	    getPaginationRowModel: getPaginationRowModel(),
    	getFilteredRowModel: getFilteredRowModel(),
	    onColumnFiltersChange: updaterOrValue => {
	    	filter.value = 
	    		typeof updaterOrValue === 'function' 
		    		? (filter.value)
		    		: updaterOrValue
	    },
		state: {
	        get globalFilter() { return filter.value },
	    },
	})

</script>

<template>
	<div>
		<div class="flex items-center justify-between py-4">
            <Input class="max-w-xs" placeholder="Filter Process Names" v-model="filter" />
            <KillButton />
        </div>
		<div class="border rounded-md">
	        <Table>
	            <TableHeader>
	                <TableRow v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
	                    <TableHead v-for="header in headerGroup.headers" :key="header.id">
	                        <FlexRender v-if="!header.isPlaceholder" :render="header.column.columnDef.header"
	                            :props="header.getContext()" />
	                    </TableHead>
	                </TableRow>
	            </TableHeader>
	            <TableBody>
	                <template v-if="table.getRowModel().rows?.length">
	                    <TableRow v-for="row in table.getRowModel().rows" :key="row.id"
	                        :data-state="row.getIsSelected() ? 'selected' : undefined">
	                        <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
	                            <FlexRender :render="cell.column.columnDef.cell" :props="cell.getContext()" />
	                        </TableCell>
	                    </TableRow>
	                </template>
	                <template v-else>
	                    <TableRow>
	                        <TableCell :colSpan="columns.length" class="h-24 text-center">
	                            No results.
	                        </TableCell>
	                    </TableRow>
	                </template>
	            </TableBody>
	        </Table>
	    </div>
	    <div class="flex items-center justify-center py-4 space-x-2">
	      <div class="flex items-center justify-between px-2 gap-x-4">
		    <div class="flex-1 text-sm text-muted-foreground">
		      {{ table.getFilteredSelectedRowModel().rows.length }} of
		      {{ table.getFilteredRowModel().rows.length }} row(s) selected.
		    </div>
		    <div class="flex items-center space-x-6 lg:space-x-8">
		      <div class="flex items-center space-x-2">
		        <p class="text-sm font-medium">
		          Rows per page
		        </p>
		        <Select
		          :model-value="`${table.getState().pagination.pageSize}`"
		          @update:model-value="table.setPageSize"
		        >
		          <SelectTrigger class="h-8 w-[70px]">
		            <SelectValue :placeholder="`${table.getState().pagination.pageSize}`" />
		          </SelectTrigger>
		          <SelectContent side="top">
		            <SelectItem v-for="pageSize in [10, 20, 30, 40, 50]" :key="pageSize" :value="`${pageSize}`">
		              {{ pageSize }}
		            </SelectItem>
		          </SelectContent>
		        </Select>
		      </div>
		      <div class="flex w-[100px] items-center justify-center text-sm font-medium">
		        Page {{ table.getState().pagination.pageIndex + 1 }} of
		        {{ table.getPageCount() }}
		      </div>
		      <div class="flex items-center space-x-2">
		        <Button
		          variant="outline"
		          class="hidden w-8 h-8 p-0 lg:flex"
		          :disabled="!table.getCanPreviousPage()"
		          @click="table.setPageIndex(0)"
		        >
		          <span class="sr-only">Go to first page</span>
		          <ArrowLeft class="w-4 h-4" />
		        </Button>
		        <Button
		          variant="outline"
		          class="w-8 h-8 p-0"
		          :disabled="!table.getCanPreviousPage()"
		          @click="table.previousPage()"
		        >
		          <span class="sr-only">Go to previous page</span>
		          <ChevronsLeft class="w-4 h-4" />
		        </Button>
		        <Button
		          variant="outline"
		          class="w-8 h-8 p-0"
		          :disabled="!table.getCanNextPage()"
		          @click="table.nextPage()"
		        >
		          <span class="sr-only">Go to next page</span>
		          <ChevronsRight class="w-4 h-4" />
		        </Button>
		        <Button
		          variant="outline"
		          class="hidden w-8 h-8 p-0 lg:flex"
		          :disabled="!table.getCanNextPage()"
		          @click="table.setPageIndex(table.getPageCount() - 1)"
		        >
		          <span class="sr-only">Go to last page</span>
		          <ArrowRight class="w-4 h-4" />
		        </Button>
		      </div>
		    </div>
		  </div>
	  </div>
	</div>
</template>