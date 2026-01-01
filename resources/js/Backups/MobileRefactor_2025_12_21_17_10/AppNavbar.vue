<script setup>
import { Link, usePage } from '@inertiajs/vue3';
import { useTheme } from '@/composables/useTheme';

const { isDark, toggleTheme } = useTheme();
const page = usePage();

const slugify = (text) => {
    return text
        .toString()
        .toLowerCase()
        .trim()
        .replace(/\s+/g, '-')           // Replace spaces with -
        .replace(/[^\w\-]+/g, '')       // Remove all non-word chars
        .replace(/\-\-+/g, '-');        // Replace multiple - with single -
};
</script>

<template>
    <nav class="fixed top-8 left-0 right-0 z-50 flex justify-center px-4 w-full transition-all duration-300">
        <div class="flex flex-col max-w-[1200px] flex-1">
            <header class="glass-panel flex items-center justify-between whitespace-nowrap px-4 py-3 md:px-10">
                <!-- Branding -->
                <div class="flex items-center gap-3">
                    <div class="bg-white/90 backdrop-blur-sm rounded-xl p-2 flex items-center justify-center shadow-sm border border-white/50 hover:scale-105 transition-transform duration-300">
                        <img src="/images/logo_tarakreasi.png" alt="TaraKreasi Notes Logo" class="h-8 w-auto object-contain">
                    </div>
                </div>

                <!-- Desktop Menu -->
                <div class="hidden md:flex flex-1 justify-end gap-8 items-center">
                    <div class="flex items-center gap-6">
                        <Link 
                            href="/" 
                            class="text-sm font-medium transition-colors relative"
                            :class="$page.url === '/' ? 'text-primary font-bold after:content-[\'\'] after:absolute after:bottom-[-4px] after:left-0 after:w-full after:h-[2px] after:bg-primary' : 'text-slate-500 dark:text-slate-400 hover:text-primary'"
                        >
                            Home
                        </Link>
                        
                        <Link 
                            :href="route('about')" 
                            class="text-sm font-medium transition-colors relative"
                            :class="$page.url.startsWith('/about') ? 'text-primary font-bold after:content-[\'\'] after:absolute after:bottom-[-4px] after:left-0 after:w-full after:h-[2px] after:bg-primary' : 'text-slate-500 dark:text-slate-400 hover:text-primary'"
                        >
                            About
                        </Link>
                        <Link 
                            v-if="$page.props.auth.user" 
                            :href="route('portfolio.show', $page.props.auth.user.username || slugify($page.props.auth.user.name))" 
                            class="text-sm font-medium transition-colors relative"
                            :class="$page.url.startsWith('/portfolio') ? 'text-primary font-bold after:content-[\'\'] after:absolute after:bottom-[-4px] after:left-0 after:w-full after:h-[2px] after:bg-primary' : 'text-slate-500 dark:text-slate-400 hover:text-primary'"
                        >
                            Portfolio
                        </Link>

                        <Link 
                            :href="route('articles.index')" 
                            class="text-sm font-medium transition-colors relative"
                            :class="$page.url.startsWith('/articles') ? 'text-primary font-bold after:content-[\'\'] after:absolute after:bottom-[-4px] after:left-0 after:w-full after:h-[2px] after:bg-primary' : 'text-slate-500 dark:text-slate-400 hover:text-primary'"
                        >
                            Articles
                        </Link>

                        <Link 
                            :href="route('contact')" 
                            class="text-sm font-medium transition-colors relative"
                            :class="$page.url.startsWith('/contact') ? 'text-primary font-bold after:content-[\'\'] after:absolute after:bottom-[-4px] after:left-0 after:w-full after:h-[2px] after:bg-primary' : 'text-slate-500 dark:text-slate-400 hover:text-primary'"
                        >
                            Contact
                        </Link>
                    </div>

                    <!-- Auth Actions -->
                    <div class="flex items-center gap-4 border-l border-slate-200 dark:border-white/10 pl-6">
                        <Link 
                            v-if="$page.props.auth.user" 
                            :href="route('dashboard')" 
                            class="text-sm font-bold text-slate-700 dark:text-white hover:text-primary"
                        >
                            Dashboard
                        </Link>
                        <template v-else>
                            <Link 
                                :href="route('login')" 
                                class="text-sm font-bold text-slate-700 dark:text-white hover:text-primary"
                            >
                                Log in
                            </Link>
                        </template>
                    </div>

                    <!-- Theme Toggle -->
                    <button
                        @click="toggleTheme"
                        class="theme-toggle-btn text-primary dark:text-white"
                        aria-label="Toggle Dark Mode"
                    >
                        <span v-if="!isDark" class="material-symbols-outlined">light_mode</span>
                        <span v-else class="material-symbols-outlined">dark_mode</span>
                    </button>
                </div>

                <!-- Mobile Menu Icon (Placeholder logic) -->
                <button class="md:hidden text-slate-800 dark:text-white">
                    <span class="material-symbols-outlined">menu</span>
                </button>
            </header>
        </div>
    </nav>
</template>
