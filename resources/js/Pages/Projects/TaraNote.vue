<script setup>
import { Head } from '@inertiajs/vue3';
import { computed } from "vue";
import { useTheme } from "@/composables/useTheme";
import FloatingDock from "@/Components/FloatingDock.vue";

const { isDark } = useTheme();
const theme = computed(() => isDark.value ? 'dark' : 'light');
</script>

<template>
    <div :class="{ 'dark': isDark }" class="min-h-screen bg-slate-50 dark:bg-[#0B1120] text-slate-900 dark:text-slate-300 font-sans selection:bg-blue-500/30 overflow-x-hidden">
        <Head title="TaraNote - Project Detail" />
        
        <!-- Navigation -->
        <FloatingDock :vertical="true" />

        <!-- Fixed Branding -->
        <div class="fixed top-8 left-8 z-50 pointer-events-none">
            <h1 class="text-xl font-bold tracking-tighter text-slate-900 dark:text-white">TaraNote.</h1>
        </div>

        <!-- Main Content -->
        <main class="transition-all duration-500">
            
            <!-- Hero Section -->
            <section class="min-h-[90vh] flex flex-col justify-center py-20 relative">
                <!-- Background Decorative -->
                <div class="absolute top-0 right-0 w-[800px] h-[800px] bg-blue-500/5 dark:bg-blue-500/10 rounded-full blur-[120px] -z-10 translate-x-1/2 -translate-y-1/4"></div>

                <div class="max-w-6xl mx-auto pl-24 pr-6 w-full">
                    <!-- Project Tag -->
                    <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-blue-100 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400 text-sm font-bold mb-8">
                        <span class="material-symbols-outlined text-lg">edit_note</span>
                        <span>Productivity Tool</span>
                    </div>

                    <div class="grid lg:grid-cols-2 gap-16 items-center">
                        <div class="space-y-8">
                            <h2 class="text-5xl md:text-7xl font-bold text-slate-900 dark:text-white leading-[1.1] tracking-tight">
                                Structured Notes for <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-indigo-500">Builders</span>.
                            </h2>
                            <p class="text-xl text-slate-600 dark:text-slate-400 leading-relaxed max-w-lg">
                                An offline-first CMS designed to help developers organize thoughts, snippets, and documentation without distraction.
                            </p>
                            
                            <!-- Action Buttons -->
                            <div class="flex flex-wrap gap-4 pt-4">
                                <a href="/" class="px-8 py-4 rounded-full bg-slate-900 dark:bg-white text-white dark:text-slate-900 font-bold hover:scale-105 transition-transform shadow-xl flex items-center gap-2">
                                    <span class="material-symbols-outlined">arrow_back</span>
                                    Back to Portfolio
                                </a>
                                <a href="https://github.com/tarakreasi" target="_blank" class="px-8 py-4 rounded-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 font-bold hover:border-blue-500 transition-colors flex items-center gap-2">
                                    <span class="material-symbols-outlined">code</span>
                                    Source Code
                                </a>
                            </div>
                        </div>

                        <!-- Hero Visual -->
                        <div class="relative group">
                            <div class="absolute -inset-4 bg-gradient-to-tr from-blue-500 to-indigo-500 rounded-[2rem] opacity-20 group-hover:opacity-30 blur-lg transition duration-500"></div>
                            <div class="relative rounded-[2rem] overflow-hidden border border-slate-200 dark:border-slate-700 shadow-2xl bg-white dark:bg-slate-900 aspect-[4/3] flex items-center justify-center group-hover:scale-[1.02] transition duration-500">
                                <!-- Mockup Interface -->
                                <div class="w-full h-full bg-slate-100 dark:bg-slate-800 p-8 flex gap-4 opacity-50">
                                    <div class="w-1/4 h-full bg-white dark:bg-slate-900 rounded-xl shadow-sm"></div>
                                    <div class="flex-1 h-full bg-white dark:bg-slate-900 rounded-xl shadow-sm space-y-4 p-4">
                                        <div class="h-8 w-3/4 bg-slate-200 dark:bg-slate-700 rounded"></div>
                                        <div class="h-4 w-full bg-slate-200 dark:bg-slate-700 rounded"></div>
                                        <div class="h-4 w-5/6 bg-slate-200 dark:bg-slate-700 rounded"></div>
                                        <div class="h-4 w-full bg-slate-200 dark:bg-slate-700 rounded"></div>
                                        <div class="h-32 w-full bg-slate-200 dark:bg-slate-700 rounded mt-8"></div>
                                    </div>
                                </div>
                                <div class="absolute inset-0 flex items-center justify-center text-slate-400 dark:text-slate-600 font-mono text-sm">
                                    [App Screenshot Placeholder]
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Features Grid -->
            <section class="py-32 border-t border-slate-200 dark:border-slate-800/50">
                <div class="max-w-6xl mx-auto pl-24 pr-6 w-full">
                    <h3 class="text-3xl font-bold mb-16 text-slate-900 dark:text-white">Core Capabilities</h3>
                    <div class="grid md:grid-cols-3 gap-8">
                        <div class="p-8 rounded-3xl bg-white dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700/50 hover:border-blue-500/50 transition duration-300">
                            <div class="w-12 h-12 rounded-xl bg-blue-100 dark:bg-blue-900/50 text-blue-600 dark:text-blue-400 flex items-center justify-center mb-6">
                                <span class="material-symbols-outlined text-2xl">wifi_off</span>
                            </div>
                            <h4 class="text-xl font-bold mb-3 dark:text-white">Offline First</h4>
                            <p class="text-slate-500 dark:text-slate-400 leading-relaxed">
                                Work anywhere. Data is stored locally and syncs when you reconnect. Zero latency.
                            </p>
                        </div>
                        <div class="p-8 rounded-3xl bg-white dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700/50 hover:border-blue-500/50 transition duration-300">
                            <div class="w-12 h-12 rounded-xl bg-indigo-100 dark:bg-indigo-900/50 text-indigo-600 dark:text-indigo-400 flex items-center justify-center mb-6">
                                <span class="material-symbols-outlined text-2xl">markdown</span>
                            </div>
                            <h4 class="text-xl font-bold mb-3 dark:text-white">Markdown Native</h4>
                            <p class="text-slate-500 dark:text-slate-400 leading-relaxed">
                                Write in pure Markdown. Live preview, syntax highlighting, and easy export.
                            </p>
                        </div>
                        <div class="p-8 rounded-3xl bg-white dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700/50 hover:border-blue-500/50 transition duration-300">
                            <div class="w-12 h-12 rounded-xl bg-teal-100 dark:bg-teal-900/50 text-teal-600 dark:text-teal-400 flex items-center justify-center mb-6">
                                <span class="material-symbols-outlined text-2xl">encrypted</span>
                            </div>
                            <h4 class="text-xl font-bold mb-3 dark:text-white">Private by Default</h4>
                            <p class="text-slate-500 dark:text-slate-400 leading-relaxed">
                                Your thoughts belong to you. Optional end-to-end encryption for sensitive notebooks.
                            </p>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Tech Stack Footer -->
            <section class="py-20 border-t border-slate-200 dark:border-slate-800/50">
                <div class="max-w-6xl mx-auto pl-24 pr-6 flex items-center gap-8 opacity-50 grayscale hover:grayscale-0 transition duration-500">
                    <span class="text-sm font-bold uppercase tracking-widest text-slate-500">Built With</span>
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vuejs/vuejs-original.svg" class="h-8" alt="Vue">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/laravel/laravel-original.svg" class="h-8" alt="Laravel">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/tailwindcss/tailwindcss-original.svg" class="h-8" alt="Tailwind">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/sqlite/sqlite-original.svg" class="h-8" alt="SQLite">
                </div>
            </section>

        </main>
    </div>
</template>
