import { useEffect, useState } from "react";
import { getPluginNames } from "../../api/backend.ts";
import {
	type PluginInfoMap,
	fetchPluginInfo,
} from "../../data/pluginInfoMap.ts";
import PluginListTitle from "../atoms/pluginListTitle.tsx";
import AllPluginTable from "../molecules/allPluginTable.tsx";
import EditPluginInfo from "../molecules/editPluginInfo.tsx";

export default function AllPluginList() {
	const [pluginMap, setPluginMap] = useState<PluginMap>();
	const [editingPluginName, setEditingPluginName] = useState<string>();
	const [refresh, setRefresh] = useState(false);

	useEffect(() => {
		(async () => {
			setPluginMap(await getPluginMap());
		})();
	}, []);

	useEffect(() => {
		if (refresh) {
			setRefresh(false);
		}
	}, [refresh]);

	return (
		<>
			<div id="all-plugin-list" className="m-5 mr-10">
				<PluginListTitle serverName="" />
				<AllPluginTable
					plugins={toPluginInfoArray(pluginMap)}
					editorOpener={(pluginName) => setEditingPluginName(pluginName)}
				/>
				{renderPluginInfoEditor(pluginMap, editingPluginName, () =>
					setRefresh(true),
				)}
			</div>
		</>
	);
}

async function getPluginMap(): Promise<PluginMap> {
	const map = new Map<string, PluginInfoMap>();

	for (const pluginName of (await getPluginNames()).data) {
		map.set(pluginName, await fetchPluginInfo(pluginName));
	}

	return map;
}

function renderPluginInfoEditor(
	map: PluginMap | undefined,
	pluginName: string | undefined,
	refresh: () => void,
) {
	if (map === undefined) {
		return undefined;
	}

	if (pluginName === undefined) {
		return (
			<div className="flex items-center mt-3 text-xl">
				<span>TIP: Click the plugin name to edit information</span>
			</div>
		);
	}

	return (
		<EditPluginInfo
			key={pluginName}
			name={pluginName}
			map={map}
			refresh={refresh}
		/>
	);
}

export type PluginMap = Map<string, PluginInfoMap>;

function toPluginInfoArray(map: Map<string, PluginInfoMap> | undefined) {
	if (map === undefined) {
		return [];
	}

	const result = [];

	for (const entry of map.entries()) {
		result.push({
			name: entry[0],
			description: entry[1].get("description")?.value ?? "",
			url: entry[1].get("url")?.value ?? "",
		});
	}

	return result;
}
