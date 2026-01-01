<script setup>
import { Head, Link } from '@inertiajs/vue3';
import { useTheme } from '@/composables/useTheme';
import { useScrollReveal } from '@/composables/useScrollReveal';
import AppNavbar from '@/Components/AppNavbar.vue';
import AppFooter from '@/Components/AppFooter.vue';

const props = defineProps({
    settings: {
        type: Object,
        default: () => ({})
    }
});

// Theme management
const { isDark, toggleTheme } = useTheme();

// Noise overlay SVG (inline to avoid CSS linter warnings)
const noiseOverlayStyle = "background-image: url('data:image/svg+xml,%3Csvg viewBox=\\'0 0 200 200\\' xmlns=\\'http://www.w3.org/2000/svg\\'%3E%3Cfilter id=\\'noiseFilter\\'%3E%3CfeTurbulence type=\\'fractalNoise\\' baseFrequency=\\'0.65\\' numOctaves=\\'3\\' stitchTiles=\\'stitch\\'/%3E%3C/filter%3E%3Crect width=\\'100%25\\' height=\\'100%25\\' filter=\\'url(%23noiseFilter)\\'/%3E%3C/svg%3E');";
</script>

<template>
    <Head title="About - Living Portfolio" />

    <div class="min-h-screen bg-canvas-light dark:bg-canvas-dark text-slate-800 dark:text-white transition-colors duration-200 font-body relative overflow-x-hidden">
        
        <!-- Noise Overlay -->
        <div class="fixed inset-0 z-[9999] opacity-[0.15] dark:opacity-[0.07] pointer-events-none"
             :style="noiseOverlayStyle">
        </div>

        <!-- Ambient Background -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none z-0"
             style="background-image: radial-gradient(circle at 50% 0%, rgba(68, 133, 238, 0.1) 0%, transparent 60%);">
             <div class="absolute top-[-20%] left-[-10%] w-[60%] h-[60%] bg-primary/5 dark:bg-primary/10 blur-[120px] rounded-full"></div>
             <!-- Keeping bottom-right blob for depth consistency -->
             <div class="absolute bottom-[-20%] right-[-10%] w-[50%] h-[50%] bg-accent-orange/5 dark:bg-accent-orange/10 blur-[100px] rounded-full"></div>
        </div>

        <!-- Navbar -->
        <AppNavbar />

        <!-- Main Content -->
        <main class="relative z-10">
            <!-- Hero Section -->
            <section class="pt-24 pb-16 md:pt-32 px-4 sm:px-6 lg:px-8 max-w-[1200px] mx-auto">
                <div class="flex flex-col-reverse md:flex-row gap-12 items-center">
                    <!-- Text Content -->
                    <div class="flex-1 flex flex-col gap-6">
                        <div class="inline-flex items-center gap-2 px-3 py-1.5 bg-accent-orange/10 border border-accent-orange/20 rounded-full w-fit">
                            <span class="material-symbols-outlined text-accent-orange text-[18px]">waving_hand</span>
                            <span class="text-accent-orange text-xs font-bold uppercase tracking-wide">About Me</span>
                        </div>
                        <h1 class="font-display text-4xl md:text-5xl lg:text-6xl font-black leading-tight tracking-tight">
                            {{ settings?.about_greeting || "Hi, I'm Tara." }}<br>
                            <span class="text-primary">{{ settings?.about_tagline || 'A Digital Craftsman.' }}</span>
                        </h1>
                        <p class="text-slate-600 dark:text-slate-400 text-lg leading-relaxed max-w-[500px]" v-html="settings?.about_description || 'Building TaraKreasi Notes...'"></p>
                        <div class="flex gap-4">
                            <Link :href="route('articles.index')" class="flex cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-6 bg-primary text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary-hover transition-all hover:scale-105 shadow-md">
                                <span class="truncate">{{ settings?.about_cta_text || 'Read Notes' }}</span>
                            </Link>
                        </div>
                    </div>

                    <!-- Image -->
                    <div class="w-full md:w-1/2 flex justify-center">
                        <div class="w-full max-w-[400px] aspect-square bg-gradient-to-br from-primary/20 to-accent-orange/20 rounded-2xl glass-panel flex items-center justify-center text-6xl">
                            👨‍💻
                        </div>
                    </div>
                </div>
            </section>

            <!-- Philosophy Section -->
            <section class="py-16 px-4 sm:px-6 lg:px-8 max-w-[1200px] mx-auto border-y border-slate-200 dark:border-white/5">
                <div class="flex flex-col gap-10">
                    <div class="flex flex-col gap-4 text-center items-center">
                        <h2 class="text-primary text-4xl font-bold leading-tight tracking-tight">
                            {{ settings?.about_philosophy_title || 'The Living Portfolio Philosophy' }}
                        </h2>
                        <p class="text-slate-600 dark:text-slate-400 text-lg leading-normal max-w-[720px]" v-html="settings?.about_philosophy_subtitle || 'This portfolio is a garden, not a museum...'"></p>
                    </div>
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                        <!-- Card 1 -->
                        <div class="glass-panel flex flex-col gap-4 p-6">
                            <div class="text-primary bg-blue-50 dark:bg-primary/10 w-12 h-12 rounded-lg flex items-center justify-center">
                                <span class="material-symbols-outlined text-[28px]">update</span>
                            </div>
                            <div class="flex flex-col gap-2">
                                <h3 class="text-slate-900 dark:text-white text-lg font-bold leading-tight">{{ settings?.about_principle1_title || 'Always Evolving' }}</h3>
                                <p class="text-slate-600 dark:text-slate-400 text-sm leading-relaxed">{{ settings?.about_principle1_desc || 'Work is never truly finished...' }}</p>
                            </div>
                        </div>
                        <!-- Card 2 -->
                        <div class="glass-panel flex flex-col gap-4 p-6">
                            <div class="text-primary bg-blue-50 dark:bg-primary/10 w-12 h-12 rounded-lg flex items-center justify-center">
                                <span class="material-symbols-outlined text-[28px]">construction</span>
                            </div>
                            <div class="flex flex-col gap-2">
                                <h3 class="text-slate-900 dark:text-white text-lg font-bold leading-tight">{{ settings?.about_principle2_title || 'Process First' }}</h3>
                                <p class="text-slate-600 dark:text-slate-400 text-sm leading-relaxed">{{ settings?.about_principle2_desc || "Documenting the 'how' is as important..." }}</p>
                            </div>
                        </div>
                        <!-- Card 3 -->
                        <div class="glass-panel flex flex-col gap-4 p-6">
                            <div class="text-primary bg-blue-50 dark:bg-primary/10 w-12 h-12 rounded-lg flex items-center justify-center">
                                <span class="material-symbols-outlined text-[28px]">menu_book</span>
                            </div>
                            <div class="flex flex-col gap-2">
                                <h3 class="text-slate-900 dark:text-white text-lg font-bold leading-tight">{{ settings?.about_principle3_title || 'Open Learning' }}</h3>
                                <p class="text-slate-600 dark:text-slate-400 text-sm leading-relaxed">{{ settings?.about_principle3_desc || 'Building in public...' }}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Footer -->
            <AppFooter :settings="settings" />
        </main>
    </div>
</template>
