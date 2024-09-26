import { defineConfig } from 'astro/config';
import tailwind from "@astrojs/tailwind";
import awsAmplify from 'astro-aws-amplify';

import react from '@astrojs/react';

// https://astro.build/config
export default defineConfig({
  build: {
    minimize: true,
  },
  output: "server",
  adapter: awsAmplify(),
  integrations: [tailwind(), react()],
  plugins: ['prettier-plugin-astro'],
  vite: {
    resolve: {
      alias: {
        '@': '/src'
      }
    }
  }
});