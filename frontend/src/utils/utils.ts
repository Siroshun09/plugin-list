export default function isNonEmptyArray<T>(
	arr: T[] | readonly T[],
): arr is [T, ...T[]] {
	return arr.length > 0;
}
