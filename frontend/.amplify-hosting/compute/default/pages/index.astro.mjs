/* empty css                                 */
import { c as createComponent, r as renderTemplate, a as addAttribute, d as renderHead, e as renderSlot, b as createAstro, f as renderComponent } from '../chunks/astro/server_Cv7uYwqQ.mjs';
import 'kleur/colors';
import { jsx } from 'react/jsx-runtime';
import * as React from 'react';
import { clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
export { renderers } from '../renderers.mjs';

function cn(...inputs) {
  return twMerge(clsx(inputs));
}

const Card = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  "div",
  {
    ref,
    className: cn(
      "rounded-lg border bg-card text-card-foreground shadow-sm",
      className
    ),
    ...props
  }
));
Card.displayName = "Card";
const CardHeader = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  "div",
  {
    ref,
    className: cn("flex flex-col space-y-1.5 p-6", className),
    ...props
  }
));
CardHeader.displayName = "CardHeader";
const CardTitle = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  "h3",
  {
    ref,
    className: cn(
      "text-2xl font-semibold leading-none tracking-tight",
      className
    ),
    ...props
  }
));
CardTitle.displayName = "CardTitle";
const CardDescription = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  "p",
  {
    ref,
    className: cn("text-sm text-muted-foreground", className),
    ...props
  }
));
CardDescription.displayName = "CardDescription";
const CardContent = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx("div", { ref, className: cn("p-6 pt-0", className), ...props }));
CardContent.displayName = "CardContent";
const CardFooter = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  "div",
  {
    ref,
    className: cn("flex items-center p-6 pt-0", className),
    ...props
  }
));
CardFooter.displayName = "CardFooter";

const $$Astro = createAstro();
const $$IndexLayout = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$IndexLayout;
  return renderTemplate`<html lang="en"> <head><meta charset="utf-8"><link rel="icon" type="image/svg+xml" href="/favicon.svg"><meta name="viewport" content="width=device-width"><meta name="generator"${addAttribute(Astro2.generator, "content")}><title>Astro</title>${renderHead()}</head> <body> <main class="min-h-screen bg-gradient-to-br from-green-400 via-teal-500 to-blue-500 flex items-center justify-center p-4"> ${renderSlot($$result, $$slots["default"])} </main> </body></html>`;
}, "/Users/jhonjairodiazrueda/Documents/Desarrollo/personal/url-shortener/frontend/src/layouts/IndexLayout.astro", void 0);

const $$Index = createComponent(($$result, $$props, $$slots) => {
  return renderTemplate`${renderComponent($$result, "IndexLayout", $$IndexLayout, { "title": "Acortador de Urls" }, { "default": ($$result2) => renderTemplate` ${renderComponent($$result2, "Card", Card, { "className": "w-full max-w-md bg-white bg-opacity-90 backdrop-blur-md shadow-xl" }, { "default": ($$result3) => renderTemplate` ${renderComponent($$result3, "CardHeader", CardHeader, { "className": "text-center" }, { "default": ($$result4) => renderTemplate` ${renderComponent($$result4, "CardTitle", CardTitle, { "className": "text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-green-500 to-blue-500" }, { "default": ($$result5) => renderTemplate`
Acortador de Urls
` })} ` })} ${renderComponent($$result3, "CardContent", CardContent, {}, { "default": ($$result4) => renderTemplate` ${renderComponent($$result4, "UrlShortenCreation", null, { "client:only": "react", "client:component-hydration": "only", "client:component-path": "/Users/jhonjairodiazrueda/Documents/Desarrollo/personal/url-shortener/frontend/src/components/UrlShortenCreation", "client:component-export": "default" })} ${renderComponent($$result4, "UrlShortenList", null, { "client:only": "react", "client:component-hydration": "only", "client:component-path": "/Users/jhonjairodiazrueda/Documents/Desarrollo/personal/url-shortener/frontend/src/components/UrlShortenList", "client:component-export": "default" })} ` })} ` })} ` })}`;
}, "/Users/jhonjairodiazrueda/Documents/Desarrollo/personal/url-shortener/frontend/src/pages/index.astro", void 0);

const $$file = "/Users/jhonjairodiazrueda/Documents/Desarrollo/personal/url-shortener/frontend/src/pages/index.astro";
const $$url = "";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
	__proto__: null,
	default: $$Index,
	file: $$file,
	url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
