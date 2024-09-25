/// <reference path="../.astro/types.d.ts" />

interface ImportMetaEnv {
	readonly ENDPOINT_API: string;
}

interface ImportMeta {
	readonly env: ImportMetaEnv;
}
