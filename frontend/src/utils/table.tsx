import type {
	Column,
	HeaderContext,
	Row,
	SortDirection,
} from "@tanstack/react-table";
import { FaSort, FaSortDown, FaSortUp } from "react-icons/fa";

export function makeSortableColumn<I, O>(
	ctx: HeaderContext<I, O>,
	name: string,
) {
	return (
		<div className="items-center justify-center">
			<span className="mr-1">{name}</span>
			<button onClick={ctx.column.getToggleSortingHandler()} type="button">
				{getSortIcon(ctx.column.getIsSorted())}
			</button>
		</div>
	);
}

export function getSortIcon(sort: false | SortDirection) {
	switch (sort) {
		case "asc":
			return <FaSortUp />;
		case "desc":
			return <FaSortDown />;
		default:
			return <FaSort />;
	}
}

export function createFilterInput<T>(column: Column<T> | undefined) {
	return (
		<input
			id="name-filter"
			placeholder="Filter plugins by name..."
			value={(column?.getFilterValue() as string) ?? ""}
			onChange={(e) => {
				column?.setFilterValue(e.target.value);
			}}
			className="ml-auto bg-white border border-gray-400 px-1 right-0 w-1/3 h-7"
		/>
	);
}

export function checkRowValueByFilter<T>(
	row: Row<T>,
	columnId: string,
	filter: string,
) {
	const filters = filter
		.split(" ")
		.filter((str) => str.length > 0) // Remove empty filter values
		.map((str) => str.toLowerCase()); // Make filter values lowercase to eliminate case sensitivity.
	const cell = row.getValue<string>(columnId).toLowerCase(); // Similarly, lowercase the cell values
	return (
		filters.length === 0 ||
		filters.find((filter) => cell.indexOf(filter) !== -1) !== undefined
	);
}
