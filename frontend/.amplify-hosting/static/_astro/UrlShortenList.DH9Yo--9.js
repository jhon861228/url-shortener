import{a as u,j as e,C as d,B as h,u as x}from"./CopyButton.CIJoukfj.js";import{r as a}from"./index.D-Q10PB_.js";function f(s,t,r){let l=new Set([...t,void 0]);return s.listen((n,i,o)=>{l.has(o)&&r(n,i,o)})}/**
 * @license lucide-react v0.439.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const j=u("ChevronDown",[["path",{d:"m6 9 6 6 6-6",key:"qrunsl"}]]);/**
 * @license lucide-react v0.439.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const p=u("ChevronUp",[["path",{d:"m18 15-6-6-6 6",key:"153udz"}]]);function y(s,t={}){let r=a.useCallback(n=>t.keys?f(s,t.keys,n):s.listen(n),[t.keys,s]),l=s.get.bind(s);return a.useSyncExternalStore(r,l,l)}const S=()=>{const[s,t]=a.useState([]),[r,l]=a.useState(!1),n=y(x);a.useEffect(()=>{t(JSON.parse(window.localStorage.getItem("urlShorten")||"[]"))},[n]);const i=()=>{l(!r)},o=r?s:s.slice(0,3);return e.jsx(e.Fragment,{children:s.length>0&&e.jsxs("div",{className:"mt-8",children:[e.jsx("h3",{className:"text-lg font-semibold text-gray-700 mb-4",children:"Urls recientes"}),e.jsx("ul",{className:"space-y-3",children:o.map((c,m)=>e.jsx("li",{className:"bg-teal-50 p-3 rounded-md",children:e.jsxs("div",{className:"flex",children:[e.jsx("a",{href:c,target:"_blank",rel:"noopener noreferrer",className:"text-blue-500 hover:text-blue-700 font-medium break-all",children:c}),e.jsx(d,{url:c})]})},m))}),s.length>3&&e.jsx(h,{onClick:i,className:"mt-4 w-full bg-teal-100 text-teal-700 hover:bg-teal-200 font-medium py-2 px-4 rounded-md transition duration-300 ease-in-out flex items-center justify-center",children:r?e.jsxs(e.Fragment,{children:["Mostrar menos ",e.jsx(p,{className:"ml-2 h-4 w-4"})]}):e.jsxs(e.Fragment,{children:["Mostrar mas ",e.jsx(j,{className:"ml-2 h-4 w-4"})]})})]})})};export{S as default};
