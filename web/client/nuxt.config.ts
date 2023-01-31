// API backend
const apiPrefix = "/api"
const apiBackendURL = process.env?.BACKEND || "http://127.0.0.1:8080";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	srcDir: "src",
	telemetry: true,

	modules: [
		"@nuxt-alt/proxy",
	],

	typescript: {
		strict: true,
	},

	proxy: {
		proxies: {
			[apiPrefix]: apiBackendURL,
		},
	},
});
