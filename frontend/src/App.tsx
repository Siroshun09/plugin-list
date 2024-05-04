import { useState } from "react";
import viteLogo from "/vite.svg";
import reactLogo from "./assets/react.svg";
import "./App.css";
import SideBar from "./components/templates/SideBar.tsx";

function App() {
	const [count, setCount] = useState(0);
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
					<div>
						<a href="https://vitejs.dev" target="_blank" rel="noreferrer">
							<img src={viteLogo} className="logo" alt="Vite logo" />
						</a>
						<a href="https://react.dev" target="_blank" rel="noreferrer">
							<img src={reactLogo} className="logo react" alt="React logo" />
						</a>
					</div>
					<h1>Vite + React</h1>
					<div className="card">
						<button onClick={() => setCount((count) => count + 1)}>
							count is {count}
						</button>
						<p>
							Edit <code>src/App.tsx</code> and save to test HMR
						</p>
					</div>
					<p className="read-the-docs">
						Click on the Vite and React logos to learn more
					</p>
					<p className="text-1xl ">
						Selected Server: <span className="font-bold">{server}</span>
					</p>
				</div>
			</div>
		</>
	);
}

export default App;
