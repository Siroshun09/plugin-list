import { useState } from "react";
import type { PluginInfoMap } from "../../data/pluginInfoMap.ts";
import ApplyPluginInfoChangeButton from "../atoms/applyPluginInfoChangeButton.tsx";
import type { PluginMap } from "../templates/pluginList.tsx";

export default function EditPluginInfo(props: {
	name: string;
	map: PluginMap;
	refresh: () => void;
}) {
	const infoMap = props.map.get(props.name);

	if (infoMap === undefined) {
		return undefined;
	}

	return NewEditor(props.name, infoMap, props.refresh);
}

function NewEditor(
	pluginName: string,
	infoMap: PluginInfoMap,
	refresh: () => void,
) {
	const [descriptionInput, setDescriptionInput] = useState(
		infoMap.get("description")?.value ?? "",
	);
	const [urlInput, setUrlInput] = useState(infoMap.get("url")?.value ?? "");

	return (
		<div
			id="edit-plugin-info"
			className="w-3/4 border-2 border-black p-3 rounded-xl mt-5"
		>
			<h2 className="text-2xl text-gray-700 mb-2">
				Edit information of {pluginName}
			</h2>
			<h3 className="text-xl text-gray-700 mb-1">Description</h3>
			<textarea
				id={`${pluginName}-description`}
				className="bg-white border border-gray-400 px-1 w-full h-32"
				onChange={(event) => setDescriptionInput(event.target.value)}
				defaultValue={infoMap.get("description")?.value ?? ""}
			/>
			<h3 className="text-xl text-gray-700 mb-1">URL</h3>
			<input
				id={`${pluginName}-url`}
				className="bg-white border border-gray-400 px-1 h-7 w-full"
				onChange={(event) => setUrlInput(event.target.value)}
				type="url"
				defaultValue={infoMap.get("url")?.value ?? ""}
			/>
			<ApplyPluginInfoChangeButton
				pluginName={pluginName}
				map={infoMap}
				description={descriptionInput}
				url={urlInput}
				refresh={refresh}
			/>
		</div>
	);
}
