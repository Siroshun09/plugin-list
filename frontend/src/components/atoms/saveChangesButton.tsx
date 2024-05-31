export default function SaveChangesButton(props: { onClick: () => void }) {
	return (
		<div className="flex">
			<button
				className="text-left ml-auto hover:bg-blue-300 bg-blue-100 rounded p-2 mt-3"
				name="apply-changes-button"
				type="button"
				onClick={props.onClick}
			>
				<p className="text-1xl text-blue-700 font-bold text-left">
					Save Changes
				</p>
			</button>
		</div>
	);
}
