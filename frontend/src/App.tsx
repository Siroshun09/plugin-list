import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import axios from "axios";
import { useState } from "react";
import AllPluginList from "./components/templates/pluginList.tsx";
import ServerPluginList from "./components/templates/serverPluginList.tsx";
import SideBar from "./components/templates/sidebar.tsx";

function App() {
	const [server, setServerName] = useState("");
	const [allPluginListMode, showAllPluginList] = useState(false);

	const apiUrl = import.meta.env.VITE_API_URL as string;

	if (apiUrl.length === 0) {
		alert("No api-url provided. Please contact an administrator.");
		console.error("No api-url provided.");
		return <></>;
	}

	const onServerSelected = (serverName: string) => {
		setServerName(serverName);
		showAllPluginList(false);
	};

	axios.defaults.baseURL = apiUrl;

	return (
		<QueryClientProvider client={new QueryClient()}>
			<div className="flex flex-wrap w-screen justify-center">
				<div id="sidebar" className="w-1/4 bg-gray-50 h-screen">
					<SideBar
						onServerSelected={(serverName) => onServerSelected(serverName)}
						allPluginListMode={allPluginListMode}
						changeAllPluginListMode={(newValue) => showAllPluginList(newValue)}
					/>
				</div>
				<div id="main" className="w-3/4">
					{allPluginListMode ? (
						<AllPluginList />
					) : (
						<ServerPluginList serverName={server} />
					)}
				</div>
			</div>
		</QueryClientProvider>
	);
}

export default App;
