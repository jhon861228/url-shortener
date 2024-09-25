import{a as u,s as h,j as e,C as x,B as f,u as g}from"./CopyButton.D-JOZO8A.js";import{r as a}from"./index.k2i5OfJP.js";function j(s,r,n){let t=new Set([...r,void 0]);return s.listen((l,c,o)=>{t.has(o)&&n(l,c,o)})}/**
 * @license lucide-react v0.439.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const p=u("ChevronDown",[["path",{d:"m6 9 6 6 6-6",key:"qrunsl"}]]);/**
 * @license lucide-react v0.439.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const y=u("ChevronUp",[["path",{d:"m18 15-6-6-6 6",key:"153udz"}]]);function b(s,r={}){let n=a.useCallback(l=>r.keys?j(s,r.keys,l):s.listen(l),[r.keys,s]),t=s.get.bind(s);return a.useSyncExternalStore(n,t,t)}const w=()=>{const[s,r]=a.useState([]),[n]=h.useMessage(),[t,l]=a.useState(!1),c=b(g);a.useEffect(()=>{r(JSON.parse(window.localStorage.getItem("urlShorten")||"[]"))},[c]);const o=()=>{l(!t)},m=t?s:s.slice(0,3);return e.jsxs(e.Fragment,{children:[n,s.length>0&&e.jsxs("div",{className:"mt-8",children:[e.jsx("h3",{className:"text-lg font-semibold text-gray-700 mb-4",children:"Urls recientes"}),e.jsx("ul",{className:"space-y-3",children:m.map((i,d)=>e.jsx("li",{className:"bg-teal-50 p-3 rounded-md",children:e.jsxs("div",{className:"flex",children:[e.jsx("a",{href:i,target:"_blank",rel:"noopener noreferrer",className:"text-blue-500 hover:text-blue-700 font-medium break-all",children:i}),e.jsx(x,{url:i})]})},d))}),s.length>3&&e.jsx(f,{onClick:o,className:"mt-4 w-full bg-teal-100 text-teal-700 hover:bg-teal-200 font-medium py-2 px-4 rounded-md transition duration-300 ease-in-out flex items-center justify-center",children:t?e.jsxs(e.Fragment,{children:["Mostrar menos ",e.jsx(y,{className:"ml-2 h-4 w-4"})]}):e.jsxs(e.Fragment,{children:["Mostrar mas ",e.jsx(p,{className:"ml-2 h-4 w-4"})]})})]})]})};export{w as default};
