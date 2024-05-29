import type MCPlugin from "../data/mcPlugin.ts";
import {
	type PluginAllOf,
	useGetPluginsByServer,
	useGetServerNames,
} from "./backend.ts";

export function getServerList(): readonly string[] {
	const servers = useGetServerNames().data;
	return servers === undefined ? [] : servers.data;
}

export function getPluginListOfServer(
	server: string,
): readonly MCPlugin[] | undefined {
	const plugins = useGetPluginsByServer(server).data;
	return plugins?.data.map((plugin) => toMCPlugin(plugin));
}

function toMCPlugin(plugin: PluginAllOf): MCPlugin {
	return {
		pluginName: plugin.plugin_name,
		serverName: plugin.server_name,
		fileName: plugin.file_name,
		version: plugin.version,
		type: plugin.type,
		lastUpdated: new Date(plugin.last_updated),
	};
}
