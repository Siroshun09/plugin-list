import type { ReactNode } from "react";

export default function ToggleButton(props: {
	display: ReactNode;
	onClick: () => void;
}) {
	return (
		<div>
			<button
				className="text-left w-full hover:bg-gray-100"
				name="toggle-all-plugin-list-mode"
				type="button"
				onClick={props.onClick}
			>
				<p className="text-1xl font-bold w-full mx-4 my-4 text-left">
					{props.display}
				</p>
			</button>
		</div>
	);
}
