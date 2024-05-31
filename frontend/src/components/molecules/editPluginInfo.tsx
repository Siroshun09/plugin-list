import { useState } from "react";
import {
	type PluginInfoMap,
	postChangedPluginInfo,
} from "../../data/pluginInfoMap.ts";
import InputField from "../atoms/inputField.tsx";
import SaveChangesButton from "../atoms/saveChangesButton.tsx";
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
			<InputField
				name="Description"
				input={
					<textarea
						id={`${pluginName}-description`}
						className="bg-white border border-gray-400 px-1 w-full h-32"
						onChange={(event) => setDescriptionInput(event.target.value)}
						defaultValue={infoMap.get("description")?.value ?? ""}
					/>
				}
			/>
			<InputField
				name="URL"
				input={
					<input
						id={`${pluginName}-url`}
						className="bg-white border border-gray-400 px-1 h-7 w-full"
						onChange={(event) => setUrlInput(event.target.value)}
						type="url"
						defaultValue={infoMap.get("url")?.value ?? ""}
					/>
				}
			/>
			<SaveChangesButton
				onClick={() =>
					applyChanges(pluginName, infoMap, descriptionInput, urlInput, refresh)
				}
			/>
		</div>
	);
}

function applyChanges(
	pluginName: string,
	map: PluginInfoMap,
	description: string,
	url: string,
	refresh: () => void,
) {
	setValue(map, "description", description);
	setValue(map, "url", url);

	(async () => {
		await postChangedPluginInfo(pluginName, map);
	})();

	refresh();
}

function setValue(map: PluginInfoMap, key: string, value: string) {
	const holder = map.get(key);

	if (holder === undefined) {
		map.set(key, {
			original: "",
			value: value,
		});
	} else {
		holder.value = value;
	}
}
