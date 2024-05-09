export default function isNonEmptyArray<T>(
	arr: T[] | readonly T[],
): arr is [T, ...T[]] {
	return arr !== null && arr.length > 0;
}
