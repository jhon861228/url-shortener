import React, { useState } from "react";
import { urlCreated } from "../store/urlStore";
import type { UrlShortenModel } from "@/models/UrlShorten.model";
import { parseDate } from "@/utils/DateUtil";

export function useCreateUrl() {
	const [urlShorted, setUrlShorted] = useState<UrlShortenModel>();
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
			setUrlShorted({
				url: result.urlShorted,
				createdAt: parseDate(result.createdAt).toLocaleDateString(),
			});
			saveOnStorage(urlShorted);
			urlCreated.set(result.urlShorted);
		} catch (error: any) {
			setError(error.message);
		} finally {
			setLoading(false);
		}
	};

	const saveOnStorage = (url: UrlShortenModel | undefined) => {
		if(url) {
			const urls: UrlShortenModel[] = JSON.parse(
				window.localStorage.getItem("urlShorten") || "[]",
			);
			urls.unshift(url);
			window.localStorage.setItem("urlShorten", JSON.stringify(urls));
		}
	};

	return { urlShorted, createUrl, loading, error };
}
