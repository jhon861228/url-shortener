import type { APIRoute } from "astro";

export const POST: APIRoute = async ({ request }) => {
	try {
		const body = await request.json();

		const apiUrl =
			"https://4o41uoibs7.execute-api.us-east-1.amazonaws.com/prod/short-url";
		const response = await fetch(apiUrl, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(body),
		});

		if (response.status !== 201) {
			throw new Error("Network response was not ok");
		}

		const result = await response.json();
		return new Response(JSON.stringify(result), {
			headers: { "Content-Type": "application/json" },
		});
	} catch (error) {
		console.error("Error:", error);
		return new Response(
			JSON.stringify({ error: "Failed to process request" }),
			{ status: 500 },
		);
	}
};
