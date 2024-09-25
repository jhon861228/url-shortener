import React, { useState } from "react";
import { urlCreated } from "../store/urlStore";

export function useCreateUrl() {
	const [urlShorted, setUrlShorted] = useState<string>("");
	const [loading, setLoading] = useState<boolean>(false);
	const [error, setError] = useState<string | null>(null);

	const createUrl = async (urlOriginal: string) => {
		setLoading(true);
		setError(null);
		try {
			const res = await fetch("/api/fetchShortUrl", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ urlOriginal }),
			});

			if (res.status !== 200) {
				throw new Error("Network response was not ok");
			}

			const result = await res.json();
			setUrlShorted(result.urlShorted);
			saveOnStorage(result.urlShorted);
			urlCreated.set(result.urlShorted);
		} catch (error: any) {
			setError(error.message);
		} finally {
			setLoading(false);
		}
	};

	const saveOnStorage = (url: string) => {
		const urls: string[] = JSON.parse(
			window.localStorage.getItem("urlShorten") || "[]",
		);
		urls.unshift(url);
		window.localStorage.setItem("urlShorten", JSON.stringify(urls));
	};

	return { urlShorted, createUrl, loading, error };
}
