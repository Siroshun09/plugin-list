import { useEffect, useState } from "react";
import PluginList from "./components/templates/pluginList.tsx";
import SideBar from "./components/templates/sidebar.tsx";
import APIPluginList from "./providers/api/apiPluginList.tsx";
import DemoPluginList from "./providers/demo/demoPluginList.ts";
import PluginListProvider from "./providers/pluginListProvider.ts";

function App() {
	const [server, setServerName] = useState("");
	const [provider, setProvider] = useState(PluginListProvider.createLoading);

	const apiUrl = import.meta.env.VITE_API_URL as string;

	useEffect(() => {
		const load = async () => {
			if (apiUrl.length === 0) {
				const demo = await DemoPluginList.create();
				setProvider(demo);
			} else {
				setProvider(APIPluginList.create(apiUrl));
			}
		};
		load().catch((err) => {
			alert(
				"An error occurred while loading the plugin list. Please contact an administrator or see your browser's console.",
			);
			console.log(err);
		});
	}, []);

	return (
		<provider.injectQueryClient>
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
		</provider.injectQueryClient>
	);
}

export default App;
