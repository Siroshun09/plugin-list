import { useState } from "react";
import SideBar from "./components/templates/sidebar.tsx";

function App() {
	const [server, setServerName] = useState("");

	return (
		<>
			<div className="flex flex-wrap w-screen justify-center">
				<div id="sidebar" className="w-1/4 bg-gray-50 h-screen">
					<SideBar
						serverList={["test-1", "test-2", "test-3"]}
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
