// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	srcDir: "src",
	telemetry: true,

	modules: [
		// https://github.com/kevinmarrec/nuxt-pwa-module
		"@kevinmarrec/nuxt-pwa",
	],

	typescript: {
		strict: true,
	},

	pwa: {
		icon: false,
	},
})
