import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		adapter: adapter({
			fallback: 'index.html',
		}),
		vite: {
			server: {
				port: 8080,
				proxy: {
					'/api': {
						target: 'http://localhost:3000',
					}
				}
			}
		}
	}
};

export default config;
