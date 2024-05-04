import ServerNameButton from "../atoms/ServerNameButton.tsx";

export default function ServerList(props: {
	list: string[];
	consumer: (serverName: string) => void;
}) {
	return (
		<div id="server-list">
			{props.list.map((serverName) => toHTML(serverName, props.consumer))}
		</div>
	);
}

function toHTML(serverName: string, consumer: (serverName: string) => void) {
	return (
		<div id={`server-list-${serverName}`}>
			<ServerNameButton serverName={serverName} consumer={consumer} />
		</div>
	);
}
