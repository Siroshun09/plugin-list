import { getServerList } from "../../api/api.ts";
import isNonEmptyArray from "../../utils/utils.ts";
import ServerList from "../molecules/serverList.tsx";

export default function SideBar(props: {
	onServerSelected: (serverName: string) => void;
}) {
	return (
		<div id="sidebar">
			<h2 className="text-4xl m-5">Servers</h2>
			{createServerListOrErrorIfEmpty(props.onServerSelected)}
		</div>
	);
}

function createServerListOrErrorIfEmpty(
	onServerSelected: (serverName: string) => void,
) {
	const serverList = getServerList();
	if (isNonEmptyArray(serverList)) {
		return createServerList(serverList, onServerSelected);
	}

	return <p className="text-2xl my-3 mx-5 text-red-500">No servers found</p>;
}

function createServerList(
	serverList: readonly [string, ...string[]],
	onServerSelected: (serverName: string) => void,
) {
	return (
		<>
			<p className="text-xl my-3 mx-5 text-gray-700">
				Click to change the server.
			</p>
			<ServerList list={serverList} consumer={onServerSelected} />
		</>
	);
}
