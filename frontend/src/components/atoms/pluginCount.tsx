export default function PluginCount(props: {count: number}) {
    switch (props.count) {
        case 0: {
            return decorate("No plugins installed.")
        }
        case 1: {
            return decorate("1 plugin installed.")
        }
        default: {
            return decorate(props.count + " plugins installed.")
        }
    }
}

function decorate(str: string) {
    return (<p className="text-xl text-gray-700">{str}</p>)
}
