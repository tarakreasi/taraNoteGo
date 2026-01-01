<script setup>
import { Link, usePage } from '@inertiajs/vue3';
import { ref, watch } from 'vue';
import { useTheme } from '@/composables/useTheme';

const { isDark, toggleTheme } = useTheme();
const page = usePage();
const isMobileMenuOpen = ref(false);

const toggleMobileMenu = () => {
    isMobileMenuOpen.value =!isMobileMenuOpen.value;
    if (isMobileMenuOpen.value) {
        document.body.style.overflow = 'hidden';
    } else {
        document.body.style.overflow = '';
    }
};

// Close menu on route change
watch(() => page.url, () => {
    isMobileMenuOpen.value = false;
    document.body.style.overflow = '';
});

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

                <!-- Mobile Menu Icon -->
                <button 
                    @click="toggleMobileMenu"
                    class="md:hidden text-slate-800 dark:text-white p-2 rounded-lg hover:bg-black/5 dark:hover:bg-white/10 transition-colors z-50 relative"
                    aria-label="Toggle Menu"
                >
                    <span class="material-symbols-outlined text-[28px]">
                        {{ isMobileMenuOpen ? 'close' : 'menu' }}
                    </span>
                </button>
            </header>

            <!-- Mobile Menu Overlay -->
            <transition
                enter-active-class="transition duration-300 ease-out"
                enter-from-class="opacity-0 translate-y-[-20px]"
                enter-to-class="opacity-100 translate-y-0"
                leave-active-class="transition duration-200 ease-in"
                leave-from-class="opacity-100 translate-y-0"
                leave-to-class="opacity-0 translate-y-[-20px]"
            >
                <div v-if="isMobileMenuOpen" class="absolute top-[calc(100%+8px)] left-0 w-full px-4 md:hidden">
                    <div class="glass-panel p-6 flex flex-col gap-6 shadow-2xl border border-white/20 dark:border-white/10 bg-white/95 dark:bg-slate-900/95 backdrop-blur-xl rounded-2xl">
                        <nav class="flex flex-col gap-4">
                            <Link 
                                href="/" 
                                class="flex items-center justify-between p-3 rounded-xl hover:bg-black/5 dark:hover:bg-white/5 transition-colors"
                                :class="$page.url === '/' ? 'text-primary font-bold bg-primary/5 dark:bg-primary/10' : 'text-slate-700 dark:text-slate-200'"
                            >
                                <span class="text-lg">Home</span>
                                <span v-if="$page.url === '/'" class="material-symbols-outlined text-primary">circle</span>
                            </Link>
                            
                            <Link 
                                :href="route('about')" 
                                class="flex items-center justify-between p-3 rounded-xl hover:bg-black/5 dark:hover:bg-white/5 transition-colors"
                                :class="$page.url.startsWith('/about') ? 'text-primary font-bold bg-primary/5 dark:bg-primary/10' : 'text-slate-700 dark:text-slate-200'"
                            >
                                <span class="text-lg">About</span>
                                <span v-if="$page.url.startsWith('/about')" class="material-symbols-outlined text-primary">circle</span>
                            </Link>

                            <Link 
                                v-if="$page.props.auth.user" 
                                :href="route('portfolio.show', $page.props.auth.user.username || slugify($page.props.auth.user.name))" 
                                class="flex items-center justify-between p-3 rounded-xl hover:bg-black/5 dark:hover:bg-white/5 transition-colors"
                                :class="$page.url.startsWith('/portfolio') ? 'text-primary font-bold bg-primary/5 dark:bg-primary/10' : 'text-slate-700 dark:text-slate-200'"
                            >
                                <span class="text-lg">Portfolio</span>
                                <span v-if="$page.url.startsWith('/portfolio')" class="material-symbols-outlined text-primary">circle</span>
                            </Link>

                            <Link 
                                :href="route('articles.index')" 
                                class="flex items-center justify-between p-3 rounded-xl hover:bg-black/5 dark:hover:bg-white/5 transition-colors"
                                :class="$page.url.startsWith('/articles') ? 'text-primary font-bold bg-primary/5 dark:bg-primary/10' : 'text-slate-700 dark:text-slate-200'"
                            >
                                <span class="text-lg">Articles</span>
                                <span v-if="$page.url.startsWith('/articles')" class="material-symbols-outlined text-primary">circle</span>
                            </Link>

                            <Link 
                                :href="route('contact')" 
                                class="flex items-center justify-between p-3 rounded-xl hover:bg-black/5 dark:hover:bg-white/5 transition-colors"
                                :class="$page.url.startsWith('/contact') ? 'text-primary font-bold bg-primary/5 dark:bg-primary/10' : 'text-slate-700 dark:text-slate-200'"
                            >
                                <span class="text-lg">Contact</span>
                                <span v-if="$page.url.startsWith('/contact')" class="material-symbols-outlined text-primary">circle</span>
                            </Link>
                        </nav>

                        <div class="h-px w-full bg-slate-200 dark:bg-white/10"></div>

                        <!-- Mobile Footer Actions -->
                        <div class="flex items-center justify-between gap-4">
                             <!-- Theme Toggle -->
                             <button
                                @click="toggleTheme"
                                class="flex items-center gap-2 px-4 py-2 rounded-lg bg-slate-100 dark:bg-white/5 text-slate-700 dark:text-slate-200 hover:bg-slate-200 dark:hover:bg-white/10 transition-colors"
                            >
                                <span class="material-symbols-outlined">{{ isDark ? 'light_mode' : 'dark_mode' }}</span>
                                <span class="text-sm font-medium">{{ isDark ? 'Light Mode' : 'Dark Mode' }}</span>
                            </button>

                            <!-- Auth -->
                            <Link 
                                v-if="$page.props.auth.user" 
                                :href="route('dashboard')" 
                                class="px-6 py-2 rounded-lg bg-primary text-white font-bold text-sm shadow-lg shadow-primary/30 active:scale-95 transition-all"
                            >
                                Dashboard
                            </Link>
                             <Link 
                                v-else
                                :href="route('login')" 
                                class="px-6 py-2 rounded-lg bg-slate-900 dark:bg-white text-white dark:text-slate-900 font-bold text-sm shadow-lg active:scale-95 transition-all"
                            >
                                Log in
                            </Link>
                        </div>
                    </div>
                </div>
            </transition>
        </div>
    </nav>
</template>
