import { type HeaderGroup, flexRender } from "@tanstack/react-table";

export default function TableHeader<T>(props: {
	headerGroups: HeaderGroup<T>[];
}) {
	return props.headerGroups.map((headerGroup) => (
		<tr key={headerGroup.id} className="text-center bg-gray-100">
			{headerGroup.headers.map((header) => (
				<th key={header.id} className="px-4 py-2 border border-gray-300">
					{header.isPlaceholder
						? null
						: flexRender(header.column.columnDef.header, header.getContext())}
				</th>
			))}
		</tr>
	));
}
