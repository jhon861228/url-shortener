import { renderers } from './renderers.mjs';
import { s as serverEntrypointModule } from './chunks/_@astrojs-ssr-adapter_DuJeOnhO.mjs';
import { manifest } from './manifest_BCVFagt3.mjs';
import { onRequest } from './_noop-middleware.mjs';

const _page0 = () => import('./pages/_image.astro.mjs');
const _page1 = () => import('./pages/api/fetchshorturl.astro.mjs');
const _page2 = () => import('./pages/index.astro.mjs');

const pageMap = new Map([
    ["node_modules/astro/dist/assets/endpoint/generic.js", _page0],
    ["src/pages/api/fetchShortUrl.ts", _page1],
    ["src/pages/index.astro", _page2]
]);
const serverIslandMap = new Map();

const _manifest = Object.assign(manifest, {
    pageMap,
    serverIslandMap,
    renderers,
    middleware: onRequest
});
const _args = {
    "client": "file:///Users/jhonjairodiazrueda/Documents/Desarrollo/personal/url-shortener/frontend/.amplify-hosting/static/",
    "server": "file:///Users/jhonjairodiazrueda/Documents/Desarrollo/personal/url-shortener/frontend/.amplify-hosting/compute/default/",
    "host": false,
    "port": 3000,
    "assets": "_astro"
};

const _start = 'start';
if (_start in serverEntrypointModule) {
	serverEntrypointModule[_start](_manifest, _args);
}

export { pageMap };
