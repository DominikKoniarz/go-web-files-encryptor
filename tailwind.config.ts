import type { Config } from "tailwindcss";

export default {
	content: ["./**/*.{go,templ}"],
	theme: {
		extend: {
			// fontFamily: {
			// 	inter: ["Inter", "sans-serif"],
			// },
			colors: {},
		},
	},
	plugins: [],
} satisfies Config;
