/// <reference path="../.astro/types.d.ts" />

interface ImportMetaEnv {
	readonly ENDPOINT_API: string;
	readonly ENDPOINT_API_KEY: string;
}

interface ImportMeta {
	readonly env: ImportMetaEnv;
}
