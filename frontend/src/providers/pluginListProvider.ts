import MCPlugin from "./mcPlugin.ts";

export default abstract class PluginListProvider {

    public static createLoading(): PluginListProvider {
        return new LoadingPluginListProvider()
    }

    public static isLoaded(provider: PluginListProvider): boolean {
        return !(provider instanceof LoadingPluginListProvider)
    }

    abstract getServerList(): readonly string[];

    abstract getPluginList(server: string): readonly MCPlugin[] | undefined
}

class LoadingPluginListProvider implements PluginListProvider {
    getServerList(): readonly string[] {
        return [];
    }

    getPluginList(_server: string): readonly MCPlugin[] | undefined {
        return undefined;
    }
}
