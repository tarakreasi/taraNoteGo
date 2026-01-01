import { ref, onMounted, watch } from 'vue';

/**
 * Composable for managing theme (dark/light mode)
 * Provides reactive theme state with localStorage persistence
 */
export function useTheme() {
    const isDark = ref(false);
    const isInitialized = ref(false);

    // Initialize theme from localStorage or system preference
    const initTheme = () => {
        if (typeof window === 'undefined') return;

        const savedTheme = localStorage.getItem('theme');
        const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;

        if (savedTheme === 'dark' || (!savedTheme && prefersDark)) {
            isDark.value = true;
            document.documentElement.classList.add('dark');
        } else {
            isDark.value = false;
            document.documentElement.classList.remove('dark');
        }

        isInitialized.value = true;
    };

    // Toggle theme
    const toggleTheme = () => {
        isDark.value = !isDark.value;

        if (isDark.value) {
            document.documentElement.classList.add('dark');
            localStorage.setItem('theme', 'dark');
        } else {
            document.documentElement.classList.remove('dark');
            localStorage.setItem('theme', 'light');
        }
    };

    // Set specific theme
    const setTheme = (theme) => {
        isDark.value = theme === 'dark';

        if (isDark.value) {
            document.documentElement.classList.add('dark');
            localStorage.setItem('theme', 'dark');
        } else {
            document.documentElement.classList.remove('dark');
            localStorage.setItem('theme', 'light');
        }
    };

    // Initialize on mount
    onMounted(() => {
        initTheme();
    });

    return {
        isDark,
        isInitialized,
        toggleTheme,
        setTheme
    };
}
