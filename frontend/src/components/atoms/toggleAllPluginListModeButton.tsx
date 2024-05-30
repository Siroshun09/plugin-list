export default function ToggleAllPluginListModeButton(props: {
	current: boolean;
	onClick: (newValue: boolean) => void;
}) {
	const display = props.current
		? "Click to show plugins per server"
		: "Click to show all plugins";
	return (
		<div>
			<button
				className="text-left w-full hover:bg-gray-100"
				name="toggle-all-plugin-list-mode"
				type="button"
				onClick={() => props.onClick(!props.current)}
			>
				<p className="text-1xl font-bold w-full mx-4 my-4 text-left">
					{display}
				</p>
			</button>
		</div>
	);
}
