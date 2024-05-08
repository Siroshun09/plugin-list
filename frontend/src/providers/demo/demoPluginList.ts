import axios from "axios";
import type MCPlugin from "../mcPlugin.ts";
import PluginListProvider from "../pluginListProvider.ts";

export default class DemoPluginList extends PluginListProvider {
	public static async create(): Promise<DemoPluginList> {
		return new DemoPluginList(await fetchDemoPluginLists());
	}

	private pluginListByServer;

	constructor(map: Map<string, readonly MCPlugin[]>) {
		super();
		this.pluginListByServer = map;
	}

	getServerList(): readonly string[] {
		return Array.from(this.pluginListByServer.keys());
	}

	getPluginList(server: string): readonly MCPlugin[] | undefined {
		return this.pluginListByServer.get(server);
	}
}

async function fetchDemoPluginLists(): Promise<
	Map<string, readonly MCPlugin[]>
> {
	return axios.get("./demo.json").then((res) => readDemoPluginLists(res.data));
}

// biome-ignore lint/suspicious/noExplicitAny: This method reads a json file.
function readDemoPluginLists(data: any): Map<string, readonly MCPlugin[]> {
	const result = new Map<string, readonly MCPlugin[]>();
	for (const key of Object.keys(data)) {
		result.set(key, createPluginList(data[key], key));
	}
	return result;
}

// biome-ignore lint/suspicious/noExplicitAny: This method reads a json file.
function createPluginList(data: any, serverName: string): readonly MCPlugin[] {
	if (Array.isArray(data)) {
		return data.map((e) => createMCPlugin(e, serverName));
	}

	throw Error(`Expected array, but got ${JSON.stringify(data)}`);
}

// biome-ignore lint/suspicious/noExplicitAny: This method reads a json file.
function createMCPlugin(data: any, serverName: string): MCPlugin {
	return {
		pluginName: data.plugin_name,
		serverName: serverName,
		fileName: data.file_name,
		version: data.version,
		type: data.type,
		lastUpdated: new Date(data.last_updated * 1000), // unix time -> epoch milliseconds
	};
}
