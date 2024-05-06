import PluginListProvider from "../../providers/pluginListProvider.ts";
import PluginListTitle from "../atoms/pluginListTitle.tsx";
import PluginTable from "../molecules/pluginTable.tsx";

export default function PluginList(props: {
    provider: PluginListProvider;
    serverName: string;
}) {
    return (
        <div id="sidebar" className="m-5">
            <PluginListTitle serverName={props.serverName}/>
            {createSelectedPluginList(props.provider, props.serverName)}
        </div>
    );
}

function createSelectedPluginList(provider: PluginListProvider, serverName: string) {
    if (serverName.length == 0) {
        return (<p className="text-2xl">⇐ Select the server from the sidebar to show plugin list.</p>)
    }

    const plugins = provider.getPluginList(serverName)

    if (plugins == undefined) {
        return (<p className="text-xl text-red-500">The plugin list was not found.</p>)
    }

    return (<PluginTable plugins={plugins} />)
}
