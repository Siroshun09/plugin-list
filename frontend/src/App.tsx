import { useState } from "react";
import SideBar from "./components/templates/sidebar.tsx";
import DemoPluginList from "./providers/demo/demoPluginList.ts";

function App() {
	const [server, setServerName] = useState("");
	const provider = new DemoPluginList();

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
					<p className="text-1xl">
						Selected Server: <span className="font-bold">{server}</span>
					</p>
				</div>
			</div>
		</>
	);
}

export default App;
