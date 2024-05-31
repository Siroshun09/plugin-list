import type { Cell, RowModel } from "@tanstack/react-table";
import type { ReactNode } from "react";

export default function TableBody<T>(props: {
	rowModel: RowModel<T>;
	additionalClasses: string;
	cellRenderer: (cell: Cell<T, unknown>) => ReactNode;
}) {
	return props.rowModel.rows.map((row) => (
		<tr key={row.id}>
			{row.getVisibleCells().map((cell) => (
				<td
					key={cell.id}
					className={`border border-gray-300 ${props.additionalClasses}`}
				>
					{props.cellRenderer(cell)}
				</td>
			))}
		</tr>
	));
}
