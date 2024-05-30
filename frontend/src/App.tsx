import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import axios from "axios";
import { useState } from "react";
import ServerPluginList from "./components/templates/serverPluginList.tsx";
import SideBar from "./components/templates/sidebar.tsx";

function App() {
	const [server, setServerName] = useState("");

	const apiUrl = import.meta.env.VITE_API_URL as string;

	if (apiUrl.length === 0) {
		alert("No api-url provided. Please contact an administrator.");
		console.error("No api-url provided.");
		return <></>;
	}

	axios.defaults.baseURL = apiUrl;

	return (
		<QueryClientProvider client={new QueryClient()}>
			<div className="flex flex-wrap w-screen justify-center">
				<div id="sidebar" className="w-1/4 bg-gray-50 h-screen">
					<SideBar
						onServerSelected={(serverName) => setServerName(serverName)}
					/>
				</div>
				<div id="main" className="w-3/4">
					<ServerPluginList serverName={server} />
				</div>
			</div>
		</QueryClientProvider>
	);
}

export default App;
