import{u as m,j as e,c as g,B as p,C as b}from"./CopyButton.D-JOZO8A.js";import{r as a}from"./index.k2i5OfJP.js";function x(){const[s,n]=a.useState(""),[r,o]=a.useState(!1),[i,l]=a.useState(null),f=async u=>{o(!0),l(null);try{const t=await fetch("/api/fetchShortUrl",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({urlOriginal:u})});if(t.status!==200)throw new Error("Network response was not ok");const d=await t.json();n(d.urlShorted),h(d.urlShorted),m.set(d.urlShorted)}catch(t){l(t.message)}finally{o(!1)}},h=u=>{const t=JSON.parse(window.localStorage.getItem("urlShorten")||"[]");t.unshift(u),window.localStorage.setItem("urlShorten",JSON.stringify(t))};return{urlShorted:s,createUrl:f,loading:r,error:i}}const c=a.forwardRef(({className:s,type:n,...r},o)=>e.jsx("input",{type:n,className:g("flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",s),ref:o,...r}));c.displayName="Input";const w=()=>{const[s,n]=a.useState(""),{urlShorted:r,createUrl:o}=x(),i=()=>{o(s)};return e.jsxs(e.Fragment,{children:[e.jsxs("div",{className:"space-y-4",children:[e.jsx("div",{className:"space-y-2",children:e.jsx(c,{type:"url",id:"url-input",placeholder:"Enter your long URL here",onChange:l=>n(l.target.value),className:"w-full px-4 py-2 rounded-md border-2 border-teal-300 focus:border-teal-500 focus:ring focus:ring-teal-200 transition",required:!0})}),e.jsx(p,{onClick:()=>i(),className:"w-full bg-gradient-to-r from-green-500 to-blue-500 hover:from-green-600 hover:to-blue-600 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-105",children:"Acortar"})]}),r&&e.jsxs("div",{className:"mt-6 p-4 bg-green-100 rounded-md",children:[e.jsx("p",{className:"text-sm text-gray-600 mb-2",children:"Tu URL acortada:"}),e.jsxs("div",{className:"flex",children:[e.jsx("a",{href:r,target:"_blank",rel:"noopener noreferrer",className:"text-blue-500 hover:text-blue-700 font-medium break-all",children:r}),e.jsx(b,{url:r})]})]})]})};export{w as default};