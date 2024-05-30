import {
	type Cell,
	createColumnHelper,
	flexRender,
	getCoreRowModel,
	getFilteredRowModel,
	getSortedRowModel,
	useReactTable,
} from "@tanstack/react-table";
import type React from "react";
import { FaEdit } from "react-icons/fa";
import {
	checkRowValueByFilter,
	createFilterInput,
	isNonEmptyArray,
	makeSortableColumn,
} from "../../utils/utils.tsx";
import PluginCount from "../atoms/pluginCount.tsx";

export default function AllPluginTable(props: {
	plugins: readonly PluginInfo[];
	editorOpener: (pluginName: string) => void;
}) {
	if (props.plugins.length === 0) {
		return <PluginCount count={0} suffix="found" />;
	}

	return createPluginTableIfInstalled(props.plugins, props.editorOpener);
}

function createPluginTableIfInstalled(
	plugins: readonly PluginInfo[],
	editorOpener: (pluginName: string) => void,
) {
	if (isNonEmptyArray(plugins)) {
		return createTable(plugins, editorOpener);
	}
	return undefined;
}

function createTable(
	plugins: [PluginInfo, ...PluginInfo[]],
	editorOpener: (pluginName: string) => void,
) {
	const table = useReactTable({
		data: plugins,
		columns,
		getCoreRowModel: getCoreRowModel(),
		getSortedRowModel: getSortedRowModel(),
		getFilteredRowModel: getFilteredRowModel(),
		initialState: {
			sorting: [{ id: "name", desc: false }],
		},
	});

	return (
		<>
			<div id="count-and-name-filter" className="flex my-3">
				<div className="my-auto">
					<PluginCount count={plugins.length} suffix="found" />
				</div>
				{createFilterInput(table.getColumn("name"))}
			</div>
			<table className="table-fixed w-full">
				<thead>
					{table.getHeaderGroups().map((headerGroup) => (
						<tr key={headerGroup.id} className="text-center bg-gray-100">
							{headerGroup.headers.map((header) => (
								<th
									key={header.id}
									className="px-4 py-2 border border-gray-300"
								>
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
							{row
								.getVisibleCells()
								.map((cell) => renderCell(editorOpener, cell))}
						</tr>
					))}
				</tbody>
			</table>
		</>
	);
}

const columnHelper = createColumnHelper<PluginInfo>();
const columns = [
	columnHelper.accessor("name", {
		header: (ctx) => makeSortableColumn(ctx, "Name"),
		cell: (info) => info.getValue(),
		filterFn: (row, columnId, value) =>
			checkRowValueByFilter(row, columnId, (value as string) ?? ""),
	}),
	columnHelper.accessor("description", {
		header: "Description",
		cell: (info) => info.getValue(),
	}),
	columnHelper.accessor("url", {
		header: "URL",
		cell: (info) => (
			<a
				href={info.getValue()}
				className="text-blue-400 hover:text-blue-800"
				target="_blank"
				rel="noreferrer"
			>
				{info.getValue()}
			</a>
		),
	}),
];

export type PluginInfo = {
	name: string;
	description: string;
	url: string;
};

function renderCell(
	editorOpener: (pluginName: string) => void,
	cell: Cell<PluginInfo, unknown>,
) {
	let display: React.JSX.Element;
	if (cell.column.id === "name") {
		display = (
			<button
				className="text-left w-full hover:bg-gray-100"
				name="server-name"
				type="button"
				onClick={() => editorOpener((cell.getValue() as string) ?? "")}
			>
				<div className="flex items-center px-4 py-2">
					{flexRender(cell.column.columnDef.cell, cell.getContext())}
					<FaEdit className="ml-auto" />
				</div>
			</button>
		);
	} else {
		display = (
			<div className="px-4 py-2">
				{flexRender(cell.column.columnDef.cell, cell.getContext())}
			</div>
		);
	}

	return (
		<td key={cell.id} className="border border-gray-300">
			{display}
		</td>
	);
}
