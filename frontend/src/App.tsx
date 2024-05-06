import { useEffect, useState } from "react";
import PluginList from "./components/templates/pluginList.tsx";
import SideBar from "./components/templates/sidebar.tsx";
import DemoPluginList from "./providers/demo/demoPluginList.ts";
import PluginListProvider from "./providers/pluginListProvider.ts";

function App() {
	const [server, setServerName] = useState("");
	const [provider, setProvider] = useState(PluginListProvider.createLoading);

	useEffect(() => {
		const load = async () => {
			const demo = await DemoPluginList.create();
			setProvider(demo);
		};
		load().catch((err) => {
			alert(
				"An error occurred while loading the plugin list. Please contact an administrator or see your browser's console.",
			);
			console.log(err);
		});
	}, []);

	return (
		<>
			<div className="flex flex-wrap w-screen justify-center">
				<div id="sidebar" className="w-1/4 bg-gray-50 h-screen">
					<SideBar
						provider={provider}
						onServerSelected={(serverName) => setServerName(serverName)}
					/>
				</div>
				<div id="main" className="w-3/4">
					<PluginList provider={provider} serverName={server} />
				</div>
			</div>
		</>
	);
}

export default App;
