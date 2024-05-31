import {
	createColumnHelper,
	flexRender,
	getCoreRowModel,
	getFilteredRowModel,
	getSortedRowModel,
	useReactTable,
} from "@tanstack/react-table";
import type { Plugin } from "../../api/backend.ts";
import {
	checkRowValueByFilter,
	createFilterInput,
	makeSortableColumn,
} from "../../utils/table.tsx";
import { isNonEmptyArray } from "../../utils/utils.tsx";

import PluginCount from "../atoms/pluginCount.tsx";
import TableBody from "../atoms/tableBody.tsx";
import TableHeader from "../atoms/tableHeader.tsx";

export default function ServerPluginTable(props: {
	plugins: readonly Plugin[];
}) {
	if (props.plugins.length === 0) {
		return <PluginCount count={0} suffix="installed" />;
	}

	return createPluginTableIfInstalled(props.plugins);
}

function createPluginTableIfInstalled(plugins: readonly Plugin[]) {
	if (isNonEmptyArray(plugins)) {
		return createTable(plugins);
	}
	return undefined;
}

function createTable(plugins: [Plugin, ...Plugin[]]) {
	const table = useReactTable({
		data: plugins,
		columns,
		getCoreRowModel: getCoreRowModel(),
		getSortedRowModel: getSortedRowModel(),
		getFilteredRowModel: getFilteredRowModel(),
		initialState: {
			sorting: [{ id: "plugin_name", desc: false }],
		},
	});

	return (
		<>
			<div id="count-and-name-filter" className="flex my-3">
				<div className="my-auto">
					<PluginCount count={plugins.length} suffix="installed" />
				</div>
				{createFilterInput(table.getColumn("plugin_name"))}
			</div>
			<table className="table-fixed w-full">
				<thead>
					<TableHeader headerGroups={table.getHeaderGroups()} />
				</thead>
				<tbody>
					<TableBody
						rowModel={table.getRowModel()}
						additionalClasses="px-4 py-2"
						cellRenderer={(cell) =>
							flexRender(cell.column.columnDef.cell, cell.getContext())
						}
					/>
				</tbody>
			</table>
		</>
	);
}

const columnHelper = createColumnHelper<Plugin>();
const columns = [
	columnHelper.accessor("plugin_name", {
		header: (ctx) => makeSortableColumn(ctx, "Name"),
		cell: (info) => info.getValue(),
		filterFn: (row, columnId, value) =>
			checkRowValueByFilter(row, columnId, (value as string) ?? ""),
	}),
	columnHelper.accessor("file_name", {
		header: "File",
		cell: (info) => info.getValue(),
	}),
	columnHelper.accessor("version", {
		header: "Version",
		cell: (info) => <div className="text-center">{info.getValue()}</div>,
	}),
	columnHelper.accessor("last_updated", {
		header: (ctx) => makeSortableColumn(ctx, "Last Updated"),
		cell: (info) =>
			`${new Date(info.getValue()).toLocaleDateString()}
			${new Date(info.getValue()).toLocaleTimeString()}`,
	}),
];
