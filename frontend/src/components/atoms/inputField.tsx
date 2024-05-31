import type { ReactNode } from "react";

export default function InputField(props: {
	name: string;
	input: ReactNode;
}) {
	return (
		<>
			<h3 className="text-xl text-gray-700 mb-1">{props.name}</h3>
			{props.input}
		</>
	);
}
