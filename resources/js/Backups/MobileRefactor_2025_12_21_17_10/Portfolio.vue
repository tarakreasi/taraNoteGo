<script setup>
import { Head, Link } from '@inertiajs/vue3';
import { computed } from 'vue';
import { useTheme } from '@/composables/useTheme';
import { useScrollReveal } from '@/composables/useScrollReveal';
import AppNavbar from '@/Components/AppNavbar.vue';
import AppFooter from '@/Components/AppFooter.vue';

const props = defineProps({
    user: Object,
    notes: Object,
    settings: Object
});

const heroImage = computed(() => props.user.hero_image || props.settings?.portfolio_hero_image);
const heroTagline = computed(() => props.user.hero_tagline || props.settings?.portfolio_hero_tagline);

// Theme management
const { isDark, toggleTheme } = useTheme();

// Scroll reveal
useScrollReveal({ threshold: 0.1, once: true });

// Noise overlay SVG (inline to avoid CSS linter warnings)
const noiseOverlayStyle = "background-image: url('data:image/svg+xml,%3Csvg viewBox=\\'0 0 200 200\\' xmlns=\\'http://www.w3.org/2000/svg\\'%3E%3Cfilter id=\\'noiseFilter\\'%3E%3CfeTurbulence type=\\'fractalNoise\\' baseFrequency=\\'0.65\\' numOctaves=\\'3\\' stitchTiles=\\'stitch\\'/%3E%3C/filter%3E%3Crect width=\\'100%25\\' height=\\'100%25\\' filter=\\'url(%23noiseFilter)\\'/%3E%3C/svg%3E');";

// Format date
const formatDate = (dateString) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    });
};
</script>

<template>
    <Head :title="`${user.name}'s Portfolio`" />

    <div class="min-h-screen bg-canvas-light dark:bg-canvas-dark text-slate-800 dark:text-white transition-colors duration-200 font-body relative overflow-x-hidden">
        
        <!-- Noise Overlay -->
        <div class="fixed inset-0 z-[9999] opacity-[0.15] dark:opacity-[0.07] pointer-events-none"
             :style="noiseOverlayStyle">
        </div>

        <!-- Ambient Background -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none z-0"
             style="background-image: radial-gradient(circle at 50% 0%, rgba(16, 185, 129, 0.1) 0%, transparent 60%);">
             <div class="absolute top-[-20%] left-[-10%] w-[60%] h-[60%] bg-emerald-500/5 dark:bg-emerald-500/10 blur-[120px] rounded-full"></div>
             <div class="absolute bottom-[-20%] right-[-10%] w-[50%] h-[50%] bg-lime-500/5 dark:bg-lime-500/10 blur-[100px] rounded-full"></div>
        </div>

        <!-- Navbar -->
        <AppNavbar />

        <main class="relative z-10">
            <!-- Hero Section -->
            <section class="reveal-on-scroll relative pt-28 pb-20 sm:pt-36 sm:pb-28 px-4 overflow-hidden min-h-[500px] flex items-center justify-center">
                <!-- Custom Hero Background -->
                <div v-if="heroImage" class="absolute inset-0 z-0">
                    <div class="absolute inset-0 bg-slate-900/60 z-10"></div> <!-- Dimming overlay -->
                    <img :src="heroImage" alt="Hero Background" class="w-full h-full object-cover opacity-80 blur-sm scale-110">
                </div>
                
                <!-- Ambient Spotlight (Fallback) -->
                <div v-else class="absolute top-0 left-1/2 -translate-x-1/2 w-[80%] h-[500px] bg-emerald-500/20 blur-[150px] rounded-full pointer-events-none z-0 opacity-60"></div>

                <div class="relative z-10 max-w-3xl mx-auto text-center flex flex-col gap-8 items-center">
                    <!-- User Avatar -->
                    <div class="w-32 h-32 rounded-full flex items-center justify-center text-6xl font-black text-white bg-gradient-to-br from-primary to-accent-orange shadow-xl border-4 border-white/20 backdrop-blur-sm">
                        <img v-if="user.avatar" :src="user.avatar" alt="User Avatar" class="w-full h-full rounded-full object-cover">
                        <span v-else>{{ user.name.substring(0,2).toUpperCase() }}</span>
                    </div>

                    <div class="flex flex-col gap-4">
                        <h1 class="font-display text-slate-900 dark:text-white text-5xl sm:text-6xl font-bold tracking-tight leading-tight" :class="{'!text-white': heroImage}">
                            {{ user.name }}'s <span :class="heroImage ? 'text-emerald-300' : 'text-emerald-600 dark:text-emerald-400'">{{ settings?.portfolio_title_suffix || 'Digital Garden' }}</span>
                        </h1>
                        <p class="text-slate-600 dark:text-slate-300 text-xl font-normal max-w-xl mx-auto leading-relaxed" :class="{'!text-slate-200': heroImage}">
                            {{ heroTagline || user.bio || `Explore ${user.name}'s curated collection of thoughts, tutorials, and insights.` }}
                        </p>
                    </div>
                </div>
            </section>

            <!-- Articles Grid -->
            <div class="max-w-[1200px] mx-auto px-4 sm:px-10 py-12">
                <div class="flex items-center justify-between border-b border-slate-200 dark:border-white/10 pb-4 mb-8">
                    <div>
                        <h2 class="text-slate-900 dark:text-white text-2xl font-bold tracking-tight">Published Articles</h2>
                        <div class="text-sm text-slate-500 dark:text-slate-400 mt-1">
                            {{ notes.total || notes.data.length }} {{ notes.total === 1 ? 'article' : 'articles' }}
                        </div>
                    </div>
                    
                    <Link :href="route('taranote')" class="group relative inline-flex items-center gap-2 px-5 py-2.5 bg-primary hover:bg-primary-hover text-white rounded-xl shadow-lg shadow-primary/20 transition-all duration-300 hover:scale-105 overflow-hidden">
                        <div class="absolute inset-0 bg-white/20 translate-y-full group-hover:translate-y-0 transition-transform duration-300"></div>
                        <span class="material-symbols-outlined text-[20px]">edit_note</span>
                        <span class="font-bold relative">Explore in TaraNote</span>
                    </Link>
                </div>

                <div v-if="notes.data && notes.data.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                    <Link
                        v-for="(note, index) in notes.data"
                        :key="note.id"
                        :href="route('articles.show', note.slug)"
                        class="glass-panel group flex flex-col overflow-hidden cursor-pointer hover:-translate-y-1 transition-transform duration-300"
                    >
                        <!-- Cover Image / Placeholder -->
                        <div class="h-48 w-full bg-slate-100 dark:bg-slate-800 overflow-hidden relative">
                            <template v-if="note.cover_image">
                                <img 
                                    :src="note.cover_image" 
                                    :alt="note.title" 
                                    class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                                />
                            </template>
                            <div v-else class="w-full h-full bg-gradient-to-br from-slate-200 to-slate-300 dark:from-slate-700 dark:to-slate-800 group-hover:scale-105 transition-transform duration-500 flex items-center justify-center">
                                 <span class="text-4xl filter grayscale opacity-50">{{ index % 2 === 0 ? '📝' : '💡' }}</span>
                            </div>
                            <!-- Maturity Badge -->
                            <div class="absolute top-4 right-4">
                                <span class="inline-flex items-center gap-1 rounded-md bg-white/20 backdrop-blur-md px-2 py-1 text-xs font-semibold text-white shadow-sm border border-white/20">
                                    {{ note.maturity || 'Seedling' }} 🌱
                                </span>
                            </div>
                        </div>

                        <!-- Content -->
                        <div class="p-5 flex flex-col flex-1">
                            <div class="flex items-center gap-2 text-xs font-medium text-slate-500 dark:text-slate-400 mb-3">
                                <span>{{ formatDate(note.published_at) }}</span>
                                <span>•</span>
                                <span>{{ note.notebook?.name || 'General' }}</span>
                            </div>
                            <h3 class="text-lg font-bold text-primary dark:text-white mb-2 group-hover:text-primary transition-colors">
                                {{ note.title }}
                            </h3>
                            <p class="text-slate-600 dark:text-slate-400 text-sm leading-relaxed line-clamp-2">
                                {{ note.excerpt || 'Explore this digital node...' }}
                            </p>
                        </div>
                    </Link>
                </div>

                <!-- Empty State -->
                <div v-else class="text-center py-20">
                    <div class="text-6xl mb-4">🌱</div>
                    <h3 class="text-2xl font-bold text-slate-900 dark:text-white mb-2">Garden in Progress</h3>
                    <p class="text-slate-600 dark:text-slate-400">{{ user.name }} hasn't published any articles yet.</p>
                </div>
            </div>

            <!-- Footer -->
            <AppFooter />
        </main>
    </div>
</template>
