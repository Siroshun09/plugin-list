import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import axios from "axios";
import type React from "react";
import type { ReactElement } from "react";
import {
	type PluginAllOf,
	useGetPluginsByServer,
	useGetServerNames,
} from "../../api/backend.ts";
import type MCPlugin from "../mcPlugin.ts";
import PluginListProvider from "../pluginListProvider.ts";

export default class APIPluginList extends PluginListProvider {
	public static create(apiUrl: string): APIPluginList {
		axios.defaults.baseURL = apiUrl;
		return new APIPluginList();
	}

	getServerList(): readonly string[] {
		const servers = useGetServerNames().data;
		return servers === undefined ? [] : servers.data;
	}

	getPluginList(server: string): readonly MCPlugin[] | undefined {
		const plugins = useGetPluginsByServer(server).data;

		return plugins?.data.map((plugin) => toMCPlugin(plugin));
	}

	injectQueryClient(element: { children: ReactElement }): React.JSX.Element {
		return (
			<QueryClientProvider client={new QueryClient()}>
				{element.children}
			</QueryClientProvider>
		);
	}
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
