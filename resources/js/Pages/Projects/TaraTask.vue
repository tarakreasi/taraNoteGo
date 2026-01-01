<script setup>
import { Head } from '@inertiajs/vue3';
import { computed } from "vue";
import { useTheme } from "@/composables/useTheme";
import FloatingDock from "@/Components/FloatingDock.vue";

const { isDark } = useTheme();
const theme = computed(() => isDark.value ? 'dark' : 'light');
</script>

<template>
    <div :class="{ 'dark': isDark }" class="min-h-screen bg-orange-50/50 dark:bg-[#0B1120] text-slate-900 dark:text-slate-300 font-sans selection:bg-orange-500/30 overflow-x-hidden">
        <Head title="TaraTask - Project Detail" />
        
        <FloatingDock :vertical="true" />

        <div class="fixed top-8 left-8 z-50 pointer-events-none">
            <h1 class="text-xl font-bold tracking-tighter text-slate-900 dark:text-white">TaraTask.</h1>
        </div>

        <main class="transition-all duration-500">
            
            <section class="min-h-[90vh] flex flex-col justify-center py-20 relative">
                <div class="absolute top-0 right-0 w-[800px] h-[800px] bg-orange-500/5 dark:bg-orange-500/10 rounded-full blur-[120px] -z-10 translate-x-1/2 -translate-y-1/4"></div>

                <div class="max-w-6xl mx-auto pl-24 pr-6 w-full">
                    <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-orange-100 dark:bg-orange-900/30 text-orange-600 dark:text-orange-400 text-sm font-bold mb-8">
                        <span class="material-symbols-outlined text-lg">dashboard</span>
                        <span>Project Management</span>
                    </div>

                    <div class="grid lg:grid-cols-2 gap-16 items-center">
                        <div class="space-y-8">
                            <h2 class="text-5xl md:text-7xl font-bold text-slate-900 dark:text-white leading-[1.1] tracking-tight">
                                Visual Workflows, <br>
                                <span class="text-transparent bg-clip-text bg-gradient-to-r from-orange-500 to-pink-500">Simplified</span>.
                            </h2>
                            <p class="text-xl text-slate-600 dark:text-slate-400 leading-relaxed max-w-lg">
                                A streamlined Kanban board inspired by the physical tactility of sticky notes. Drag, drop, done.
                            </p>
                            
                            <div class="flex flex-wrap gap-4 pt-4">
                                <a href="/" class="px-8 py-4 rounded-full bg-slate-900 dark:bg-white text-white dark:text-slate-900 font-bold hover:scale-105 transition-transform shadow-xl flex items-center gap-2">
                                    <span class="material-symbols-outlined">arrow_back</span>
                                    Back to Portfolio
                                </a>
                                <a href="https://github.com/tarakreasi" target="_blank" class="px-8 py-4 rounded-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 font-bold hover:border-orange-500 transition-colors flex items-center gap-2">
                                    <span class="material-symbols-outlined">code</span>
                                    Source Code
                                </a>
                            </div>
                        </div>

                        <div class="relative group">
                            <div class="absolute -inset-4 bg-gradient-to-tr from-orange-500 to-pink-500 rounded-[2rem] opacity-20 group-hover:opacity-30 blur-lg transition duration-500"></div>
                            <div class="relative rounded-[2rem] overflow-hidden border border-slate-200 dark:border-slate-700 shadow-2xl bg-white dark:bg-slate-900 aspect-[4/3] flex items-center justify-center group-hover:scale-[1.02] transition duration-500">
                                <!-- Kanban Mockup -->
                                <div class="w-full h-full bg-slate-50 dark:bg-slate-800 p-8 flex gap-4 opacity-70">
                                    <div class="w-1/3 h-full space-y-3">
                                        <div class="h-6 w-1/2 bg-slate-200 dark:bg-slate-700 rounded mb-4"></div>
                                        <div class="h-24 w-full bg-white dark:bg-slate-900 rounded-lg shadow-sm border-l-4 border-orange-400"></div>
                                        <div class="h-24 w-full bg-white dark:bg-slate-900 rounded-lg shadow-sm border-l-4 border-orange-400"></div>
                                    </div>
                                    <div class="w-1/3 h-full space-y-3">
                                        <div class="h-6 w-1/2 bg-slate-200 dark:bg-slate-700 rounded mb-4"></div>
                                        <div class="h-24 w-full bg-white dark:bg-slate-900 rounded-lg shadow-sm border-l-4 border-blue-400"></div>
                                    </div>
                                    <div class="w-1/3 h-full space-y-3">
                                        <div class="h-6 w-1/2 bg-slate-200 dark:bg-slate-700 rounded mb-4"></div>
                                        <div class="h-24 w-full bg-white dark:bg-slate-900 rounded-lg shadow-sm border-l-4 border-green-400"></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <section class="py-32 border-t border-slate-200 dark:border-slate-800/50">
                <div class="max-w-6xl mx-auto pl-24 pr-6 w-full">
                    <h3 class="text-3xl font-bold mb-16 text-slate-900 dark:text-white">Built for Flow</h3>
                    <div class="grid md:grid-cols-3 gap-8">
                        <div class="p-8 rounded-3xl bg-white dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700/50 hover:border-orange-500/50 transition duration-300">
                            <div class="w-12 h-12 rounded-xl bg-orange-100 dark:bg-orange-900/50 text-orange-600 dark:text-orange-400 flex items-center justify-center mb-6">
                                <span class="material-symbols-outlined text-2xl">drag_indicator</span>
                            </div>
                            <h4 class="text-xl font-bold mb-3 dark:text-white">Drag & Drop</h4>
                            <p class="text-slate-500 dark:text-slate-400 leading-relaxed">
                                Intuitive organization. Move tasks between stages effortlessly.
                            </p>
                        </div>
                        <div class="p-8 rounded-3xl bg-white dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700/50 hover:border-orange-500/50 transition duration-300">
                            <div class="w-12 h-12 rounded-xl bg-pink-100 dark:bg-pink-900/50 text-pink-600 dark:text-pink-400 flex items-center justify-center mb-6">
                                <span class="material-symbols-outlined text-2xl">palette</span>
                            </div>
                            <h4 class="text-xl font-bold mb-3 dark:text-white">Custom Labels</h4>
                            <p class="text-slate-500 dark:text-slate-400 leading-relaxed">
                                Categorize with color. Visual tags help you spot priorities instantly.
                            </p>
                        </div>
                        <div class="p-8 rounded-3xl bg-white dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700/50 hover:border-orange-500/50 transition duration-300">
                            <div class="w-12 h-12 rounded-xl bg-red-100 dark:bg-red-900/50 text-red-600 dark:text-red-400 flex items-center justify-center mb-6">
                                <span class="material-symbols-outlined text-2xl">history</span>
                            </div>
                            <h4 class="text-xl font-bold mb-3 dark:text-white">History Logs</h4>
                            <p class="text-slate-500 dark:text-slate-400 leading-relaxed">
                                Track every move. Know who changed what and when.
                            </p>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Tech Stack Footer -->
            <section class="py-20 border-t border-slate-200 dark:border-slate-800/50">
                <div class="max-w-6xl mx-auto pl-24 pr-6 flex items-center gap-8 opacity-50 grayscale hover:grayscale-0 transition duration-500">
                    <span class="text-sm font-bold uppercase tracking-widest text-slate-500">Built With</span>
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/laravel/laravel-original.svg" class="h-8" alt="Laravel">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/nextjs/nextjs-original.svg" class="h-8 dark:invert" alt="Next.js">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/tailwindcss/tailwindcss-original.svg" class="h-8" alt="Tailwind">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original.svg" class="h-8" alt="PostgreSQL">
                </div>
            </section>
        </main>
    </div>
</template>
