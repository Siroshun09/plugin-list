import type PluginListProvider from "../pluginListProvider.ts";

export default class DemoPluginList implements PluginListProvider {
	getServerList(): string[] {
		return ["test-1", "test-2", "test-3"];
	}
}
