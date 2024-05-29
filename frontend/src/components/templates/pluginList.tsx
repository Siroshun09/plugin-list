import { getPluginListOfServer } from "../../api/api.ts";
import PluginListTitle from "../atoms/pluginListTitle.tsx";
import PluginTable from "../molecules/pluginTable.tsx";

export default function PluginList(props: {
	serverName: string;
}) {
	return (
		<div id="sidebar" className="m-5">
			<PluginListTitle serverName={props.serverName} />
			{createSelectedPluginList(props.serverName)}
		</div>
	);
}

function createSelectedPluginList(serverName: string) {
	if (serverName.length === 0) {
		return <p className="text-2xl">⇐ Select the server from the sidebar.</p>;
	}

	const plugins = getPluginListOfServer(serverName);

	if (plugins === undefined) {
		return (
			<p className="text-xl text-red-500">The plugin list was not found.</p>
		);
	}

	return <PluginTable plugins={plugins} />;
}
