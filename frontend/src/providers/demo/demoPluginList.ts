import type PluginListProvider from "../pluginListProvider.ts";
import MCPlugin from "../mcPlugin.ts";
import axios from "axios";

export default class DemoPluginList implements PluginListProvider {

    public static async create(): Promise<DemoPluginList> {
        return new DemoPluginList(await fetchDemoPluginLists())
    }

    private pluginListByServer;

    constructor(map: Map<string, readonly MCPlugin[]>) {
        this.pluginListByServer = map
    }

    getServerList(): readonly string[] {
        return Array.from(this.pluginListByServer.keys());
    }

    getPluginList(server: string): readonly MCPlugin[] | undefined {
        return this.pluginListByServer.get(server)
    }
}

async function fetchDemoPluginLists(): Promise<Map<string, readonly MCPlugin[]>> {
    return axios.get("./demo.json")
        .then((res) => readDemoPluginLists(res.data))
}

function readDemoPluginLists(data: any): Map<string, readonly MCPlugin[]> {
    const result = new Map<string, readonly MCPlugin[]>
    Object.keys(data).forEach(key => result.set(key, createPluginList(data[key], key)))
    return result
}

function createPluginList(data: any, serverName: string): readonly MCPlugin[] {
    if (Array.isArray(data)) {
        return data.map(e => createMCPlugin(e, serverName))
    }

    throw Error("Expected array, but got " + JSON.stringify(data))
}

function createMCPlugin(data: any, serverName: string): MCPlugin {
    return {
        pluginName: data["plugin_name"],
        serverName: serverName,
        fileName: data["file_name"],
        version: data["version"],
        type: data["type"],
        lastUpdated: new Date(data["last_updated"] * 1000) // unix time -> epoch milliseconds
    }
}
