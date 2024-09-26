import { useState } from "react";

export function useClipboard() {
	const [isCopied, setIsCopied] = useState(false);
	const [err, setErr] = useState<string | null>(null);

	const copyToClipboard = async (text: string) => {
		try {
			await navigator.clipboard.writeText(text);
			setIsCopied(true);
			setErr(null);
		} catch (err) {
			setErr("Failed to copy");
		}
	};

	return { isCopied, copyToClipboard, err };
}
