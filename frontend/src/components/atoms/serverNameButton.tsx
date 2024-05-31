export default function ServerNameButton(props: {
	serverName: string;
	consumer: (serverName: string) => void;
}) {
	return (
		<div>
			<button
				className="text-left w-full hover:bg-gray-100"
				name="server-name"
				type="button"
				onClick={() => props.consumer(props.serverName)}
			>
				<p className="text-2xl w-full mx-7 my-4">{props.serverName}</p>
			</button>
		</div>
	);
}
