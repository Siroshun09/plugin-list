import ServerNameButton from "../atoms/serverNameButton.tsx";

export default function ServerList(props: {
	list: readonly string[];
	consumer: (serverName: string) => void;
}) {
	return (
		<div id="server-list">
			{props.list.map((serverName) => createButton(serverName, props.consumer))}
		</div>
	);
}

function createButton(serverName: string, consumer: (serverName: string) => void) {
	return (
		<div id={`server-${serverName}`} key={serverName}>
			<ServerNameButton serverName={serverName} consumer={consumer} />
		</div>
	);
}
