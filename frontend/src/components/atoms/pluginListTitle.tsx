export default function PluginListTitle(props: {
	serverName: string;
}) {
	const suffix = props.serverName.length === 0 ? "" : `: ${props.serverName}`;
	return <h2 className="text-4xl mt-10 mb-5">Plugin List {suffix}</h2>;
}
