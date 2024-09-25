import{a as u,s as x,j as e,C as f,B as g,u as p}from"./CopyButton.DSfOCmN9.js";import{r as n}from"./index.k2i5OfJP.js";function j(s,r,a){let l=new Set([...r,void 0]);return s.listen((t,i,o)=>{l.has(o)&&a(t,i,o)})}/**
 * @license lucide-react v0.439.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const y=u("ChevronDown",[["path",{d:"m6 9 6 6 6-6",key:"qrunsl"}]]);/**
 * @license lucide-react v0.439.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const b=u("ChevronUp",[["path",{d:"m18 15-6-6-6 6",key:"153udz"}]]);function S(s,r={}){let a=n.useCallback(t=>r.keys?j(s,r.keys,t):s.listen(t),[r.keys,s]),l=s.get.bind(s);return n.useSyncExternalStore(a,l,l)}const N=()=>{const[s,r]=n.useState([]),[a,l]=x.useMessage(),[t,i]=n.useState(!1),o=S(p);n.useEffect(()=>{r(JSON.parse(window.localStorage.getItem("urlShorten")||"[]"))},[o]);const m=()=>{i(!t)},d=t?s:s.slice(0,3);return e.jsxs(e.Fragment,{children:[l,s.length>0&&e.jsxs("div",{className:"mt-8",children:[e.jsx("h3",{className:"text-lg font-semibold text-gray-700 mb-4",children:"Urls recientes"}),e.jsx("ul",{className:"space-y-3",children:d.map((c,h)=>e.jsx("li",{className:"bg-teal-50 p-3 rounded-md",children:e.jsxs("div",{className:"flex",children:[e.jsx("a",{href:c,target:"_blank",rel:"noopener noreferrer",className:"text-blue-500 hover:text-blue-700 font-medium break-all",children:c}),e.jsx(f,{url:c})]})},h))}),s.length>3&&e.jsx(g,{onClick:m,className:"mt-4 w-full bg-teal-100 text-teal-700 hover:bg-teal-200 font-medium py-2 px-4 rounded-md transition duration-300 ease-in-out flex items-center justify-center",children:t?e.jsxs(e.Fragment,{children:["Mostrar menos ",e.jsx(b,{className:"ml-2 h-4 w-4"})]}):e.jsxs(e.Fragment,{children:["Mostrar mas ",e.jsx(y,{className:"ml-2 h-4 w-4"})]})})]})]})};export{N as default};
