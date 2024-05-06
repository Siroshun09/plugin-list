import {
	createColumnHelper,
	flexRender,
	getCoreRowModel,
	useReactTable,
} from "@tanstack/react-table";
import type MCPlugin from "../../providers/mcPlugin.ts";
import isNonEmptyArray from "../../utils/utils.ts";
import PluginCount from "../atoms/pluginCount.tsx";

export default function PluginTable(props: {
	plugins: readonly MCPlugin[];
}) {
	if (props.plugins.length === 0) {
		return <PluginCount count={0} />;
	}

	return (
		<>
			<div id="plugin-count" className="mb-5">
				<PluginCount count={props.plugins.length} />
			</div>
			{createPluginTableIfInstalled(props.plugins)}
		</>
	);
}

function createPluginTableIfInstalled(plugins: readonly MCPlugin[]) {
	if (isNonEmptyArray(plugins)) {
		return createTable(plugins);
	}
	return undefined;
}

function createTable(plugins: [MCPlugin, ...MCPlugin[]]) {
	const table = useReactTable({
		data: plugins,
		columns,
		getCoreRowModel: getCoreRowModel(),
	});

	return (
		<table className="table-auto w-full">
			<thead>
				{table.getHeaderGroups().map((headerGroup) => (
					<tr key={headerGroup.id} className="text-center bg-gray-100">
						{headerGroup.headers.map((header) => (
							<th key={header.id} className="px-4 py-2 border border-gray-300">
								{header.isPlaceholder
									? null
									: flexRender(
											header.column.columnDef.header,
											header.getContext(),
										)}
							</th>
						))}
					</tr>
				))}
			</thead>
			<tbody>
				{table.getRowModel().rows.map((row) => (
					<tr key={row.id}>
						{row.getVisibleCells().map((cell) => (
							<td key={cell.id} className="px-4 py-2 border border-gray-300">
								{flexRender(cell.column.columnDef.cell, cell.getContext())}
							</td>
						))}
					</tr>
				))}
			</tbody>
		</table>
	);
}

const columnHelper = createColumnHelper<MCPlugin>();
const columns = [
	columnHelper.accessor("pluginName", {
		header: "Name",
		cell: (info) => info.getValue(),
	}),
	columnHelper.accessor("fileName", {
		header: "File",
		cell: (info) => info.getValue(),
	}),
	columnHelper.accessor("version", {
		header: "Version",
		cell: (info) => info.getValue(),
	}),
	columnHelper.accessor("lastUpdated", {
		header: "LastUpdated",
		cell: (info) =>
			`${info.getValue().toLocaleDateString()} ${info
				.getValue()
				.toLocaleTimeString()}`,
	}),
];
