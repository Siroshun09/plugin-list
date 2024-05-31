import { useEffect, useState } from "react";
import { getServerNames } from "../../api/backend.ts";
import { isNonEmptyArray } from "../../utils/utils.tsx";
import ToggleButton from "../atoms/toggleButton.tsx";
import ServerList from "../molecules/serverList.tsx";

export default function SideBar(props: {
	onServerSelected: (serverName: string) => void;
	allPluginListMode: boolean;
	changeAllPluginListMode: (newValue: boolean) => void;
}) {
	const [serverList, setServerList] = useState<readonly string[]>([]);

	useEffect(() => {
		(async () => {
			setServerList((await getServerNames()).data);
		})();
	}, []);

	return (
		<div id="sidebar">
			<h2 className="text-4xl m-5">Servers</h2>
			{createServerListOrErrorIfEmpty(
				serverList,
				props.onServerSelected,
				props.allPluginListMode,
				props.changeAllPluginListMode,
			)}
		</div>
	);
}

function createServerListOrErrorIfEmpty(
	serverList: readonly string[],
	onServerSelected: (serverName: string) => void,
	allPluginListMode: boolean,
	changeAllPluginListMode: (newValue: boolean) => void,
) {
	if (isNonEmptyArray(serverList)) {
		return createServerList(
			serverList,
			onServerSelected,
			allPluginListMode,
			changeAllPluginListMode,
		);
	}

	return <p className="text-2xl my-3 mx-5 text-red-500">No servers found</p>;
}

function createServerList(
	serverList: readonly [string, ...string[]],
	onServerSelected: (serverName: string) => void,
	allPluginListMode: boolean,
	changeAllPluginListMode: (newValue: boolean) => void,
) {
	return (
		<>
			<p className="text-xl my-3 mx-5 text-gray-700">
				Click to change the server.
			</p>
			<ServerList list={serverList} consumer={onServerSelected} />
			<ToggleButton
				display={
					allPluginListMode
						? "Click to show plugins per server"
						: "Click to show all plugins"
				}
				onClick={() => changeAllPluginListMode(!allPluginListMode)}
			/>
		</>
	);
}
