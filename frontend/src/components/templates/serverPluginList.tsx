import type React from "react";
import { useEffect, useState } from "react";
import { getPluginsByServer } from "../../api/backend.ts";
import PluginListTitle from "../atoms/pluginListTitle.tsx";
import ServerPluginTable from "../molecules/serverPluginTable.tsx";

export default function ServerPluginList(props: {
	serverName: string;
}) {
	const [pluginList, setPluginList] = useState<React.JSX.Element>();

	useEffect(() => {
		(async () => {
			setPluginList(await createSelectedPluginList(props.serverName));
		})();
	}, [props.serverName]);

	return (
		<div id="sidebar" className="m-5 mr-10">
			<PluginListTitle serverName={props.serverName} />
			{pluginList}
		</div>
	);
}

async function createSelectedPluginList(serverName: string) {
	if (serverName.length === 0) {
		return <p className="text-2xl">⇐ Select the server from the sidebar.</p>;
	}

	const plugins = (await getPluginsByServer(serverName)).data;

	if (plugins === undefined) {
		return (
			<p className="text-xl text-red-500">The plugin list was not found.</p>
		);
	}

	return <ServerPluginTable plugins={plugins} />;
}
