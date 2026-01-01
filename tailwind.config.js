import defaultTheme from 'tailwindcss/defaultTheme';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
    darkMode: 'class',
    content: [
        './vendor/laravel/framework/src/Illuminate/Pagination/resources/views/*.blade.php',
        './storage/framework/views/*.php',
        './resources/views/**/*.blade.php',
        './resources/js/**/*.vue',
    ],

    theme: {
        extend: {
            colors: {
                // v6 "Electric Earth" Palette
                "primary": "#135bec",       // Tara Blue (Main Brand)
                "primary-hover": "#0e45b5",
                "accent-orange": "#FF5F1F", // Electric Orange (CTAs)
                "accent-orange-hover": "#FF8C00",

                // Backgrounds
                "canvas-light": "#F8FAFC",  // Softer than white
                "canvas-dark": "#0F172A",   // Deep Blue/Black

                // Functional
                "success": "#10B981",
                "danger": "#EF4444",
                "warning": "#F59E0B",

                // Legacy Mappings (soft deprecation)
                "brand-blue": "#4485EE",
            },
            fontFamily: {
                "display": ["Noto Sans JP", "sans-serif"],
                "body": ["Noto Sans JP", "sans-serif"],
                sans: ['Noto Sans JP', ...defaultTheme.fontFamily.sans],
            },
            backgroundImage: {
                'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
            }
        },
    },

    plugins: [forms, typography],
};