import {
	type PluginInfoMap,
	postChangedPluginInfo,
} from "../../data/pluginInfoMap.ts";

export default function ApplyPluginInfoChangeButton(props: {
	pluginName: string;
	map: PluginInfoMap;
	description: string;
	url: string;
	refresh: () => void;
}) {
	return (
		<div className="flex">
			<button
				className="text-left ml-auto hover:bg-blue-300 bg-blue-100 rounded p-2 mt-3"
				name="toggle-all-plugin-list-mode"
				type="button"
				onClick={() => applyChanges(props)}
			>
				<p className="text-1xl text-blue-700 font-bold text-left">
					Save Changes
				</p>
			</button>
		</div>
	);
}

function applyChanges(props: {
	pluginName: string;
	map: PluginInfoMap;
	description: string;
	url: string;
	refresh: () => void;
}) {
	setValue(props.map, "description", props.description);
	setValue(props.map, "url", props.url);

	(async () => {
		await postChangedPluginInfo(props.pluginName, props.map);
	})();

	props.refresh();
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
