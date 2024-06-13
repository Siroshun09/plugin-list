import { addPluginCustomData, getPluginCustomData } from "../api/backend.ts";

export type PluginInfoValue = {
	readonly original: string;
	value: string;
};

export type PluginInfoMap = Map<string, PluginInfoValue>;

export async function fetchPluginInfo(pluginName: string) {
	const data = await getPluginCustomData(pluginName);
	const result = new Map<string, PluginInfoValue>();

	for (const entry of Object.entries(data.data)) {
		result.set(entry[0], { original: entry[1], value: entry[1] });
	}

	return result;
}

export async function postChangedPluginInfo(
	pluginName: string,
	map: PluginInfoMap,
) {
	const changedDataMap: { [key: string]: string } = {};
	let shouldPost = false;
	for (const entry of map.entries()) {
		const key = entry[0];
		const value = entry[1];

		if (value.original !== value.value) {
			// modified
			changedDataMap[key] = value.value;
			shouldPost = true;
		}
	}

	if (shouldPost) {
		await addPluginCustomData(pluginName, changedDataMap);
	}
}
