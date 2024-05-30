import {
	type HeaderContext,
	type SortDirection,
	createColumnHelper,
	flexRender,
	getCoreRowModel,
	getFilteredRowModel,
	getSortedRowModel,
	useReactTable,
} from "@tanstack/react-table";
import { FaSort, FaSortDown, FaSortUp } from "react-icons/fa";
import isNonEmptyArray from "../../utils/utils.ts";
import PluginCount from "../atoms/pluginCount.tsx";
import {Plugin} from "../../api/backend.ts";

export default function ServerPluginTable(props: {
	plugins: readonly Plugin[];
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
			<div id="name-filter">
				<input
					placeholder="Filter plugins by name..."
					value={
						(table.getColumn("plugin_name")?.getFilterValue() as string) ?? ""
					}
					onChange={(e) => {
						table.getColumn("plugin_name")?.setFilterValue(e.target.value);
					}}
					className="flex flex-row-reverse bg-white border border-gray-400 px-1 my-1.5 right-0"
				/>
			</div>
			<table className="table-auto w-full">
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
							{row.getVisibleCells().map((cell) => (
								<td key={cell.id} className="px-4 py-2 border border-gray-300">
									{flexRender(cell.column.columnDef.cell, cell.getContext())}
								</td>
							))}
						</tr>
					))}
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
		filterFn: (row, columnId, value) => {
			const filters = ((value as string) ?? "")
				.split(" ")
				.filter((str) => str.length > 0) // Remove empty filter values
				.map((str) => str.toLowerCase()); // Make filter values lowercase to eliminate case sensitivity.
			const cell = row.getValue<string>(columnId).toLowerCase(); // Similarly, lowercase the cell values
			return (
				filters.length === 0 ||
				filters.find((filter) => cell.indexOf(filter) !== -1) !== undefined
			);
		},
	}),
	columnHelper.accessor("file_name", {
		header: "File",
		cell: (info) => info.getValue(),
	}),
	columnHelper.accessor("version", {
		header: "Version",
		cell: (info) => info.getValue(),
	}),
	columnHelper.accessor("last_updated", {
		header: (ctx) => makeSortableColumn(ctx, "Last Updated"),
		cell: (info) =>
			`${new Date(info.getValue()).toLocaleDateString()}
			${new Date(info.getValue()).toLocaleTimeString()}`,
	}),
];

function makeSortableColumn<I, O>(ctx: HeaderContext<I, O>, name: string) {
	return (
		<div className="items-center justify-center">
			<span className="mr-1">{name}</span>
			<button onClick={ctx.column.getToggleSortingHandler()} type="button">
				{getSortIcon(ctx.column.getIsSorted())}
			</button>
		</div>
	);
}

function getSortIcon(sort: false | SortDirection) {
	switch (sort) {
		case "asc":
			return <FaSortUp />;
		case "desc":
			return <FaSortDown />;
		default:
			return <FaSort />;
	}
}
