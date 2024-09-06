export async function shortUrl(url: string) {
    const endpoint = import.meta.env.ENDPOINT_API;

    const res = await fetch(`${endpoint}/short-url?id=q`, {
        headers: {
            "x-api-key": import.meta.env.ENDPOINT_API_KEY
        }
    })
    
    // these two is causing the errors.
    console.log(await res.text())
    
}

