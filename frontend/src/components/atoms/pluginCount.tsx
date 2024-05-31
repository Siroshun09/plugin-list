export default function PluginCount(props: { count: number; suffix: string }) {
	switch (props.count) {
		case 0: {
			return decorate(`No plugins ${props.suffix}`);
		}
		case 1: {
			return decorate(`1 plugin ${props.suffix}`);
		}
		default: {
			return decorate(`${props.count} plugins ${props.suffix}`);
		}
	}
}

function decorate(str: string) {
	return <p className="text-xl text-gray-700">{str}</p>;
}
