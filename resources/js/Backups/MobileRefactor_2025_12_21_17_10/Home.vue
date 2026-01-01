<script setup>
import { Head, Link, useForm, usePage } from '@inertiajs/vue3';
import { ref, computed } from 'vue';
import { useScrollReveal } from '@/composables/useScrollReveal';
import AppNavbar from '@/Components/AppNavbar.vue';
import AppFooter from '@/Components/AppFooter.vue';

const props = defineProps({
    canLogin: {
        type: Boolean,
    },
    canRegister: {
        type: Boolean,
    },
    laravelVersion: {
        type: String,
        required: true,
    },
    phpVersion: {
        type: String,
        required: true,
    },
    latestArticles: {
        type: Array,
        default: () => [],
    },
    settings: {
        type: Object,
        default: () => ({})
    }
});

// Scroll reveal animations
useScrollReveal({ threshold: 0.1, once: true });

// Filter state for Digital Garden section (Client-side filtering for simplicity on home)
const activeFilter = ref('all');

const setFilter = (filter) => {
    activeFilter.value = filter;
};

const filteredArticles = computed(() => {
    if (activeFilter.value === 'all') {
        return props.latestArticles;
    }
    // Assuming 'maturity' might not be a field yet, default to showing all or check if field exists
    return props.latestArticles.filter(article => article.maturity === activeFilter.value);
});

// Format Date
const formatDate = (dateString) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    });
};

// Noise overlay SVG
const noiseOverlayStyle = "background-image: url('data:image/svg+xml,%3Csvg viewBox=\\'0 0 200 200\\' xmlns=\\'http://www.w3.org/2000/svg\\'%3E%3Cfilter id=\\'noiseFilter\\'%3E%3CfeTurbulence type=\\'fractalNoise\\' baseFrequency=\\'0.65\\' numOctaves=\\'3\\' stitchTiles=\\'stitch\\'/%3E%3C/filter%3E%3Crect width=\\'100%25\\' height=\\'100%25\\' filter=\\'url(%23noiseFilter)\\'/%3E%3C/svg%3E');";

const slugify = (text) => {
    return text
        .toString()
        .toLowerCase()
        .trim()
        .replace(/\s+/g, '-')           // Replace spaces with -
        .replace(/[^\w\-]+/g, '')       // Remove all non-word chars
        .replace(/\-\-+/g, '-');        // Replace multiple - with single -
};
const form = useForm({
    email: '',
});

const submitNewsletter = () => {
    form.post(route('newsletter.subscribe'), {
        preserveScroll: true,
        onSuccess: () => form.reset(),
    });
};
</script>

<template>
    <Head title="Welcome to TaraKreasi" />

    <div class="min-h-screen bg-canvas-light dark:bg-canvas-dark text-slate-800 dark:text-white transition-colors duration-200 font-body relative overflow-x-hidden">
        
        <!-- Noise Overlay -->
        <div class="fixed inset-0 z-[9999] opacity-[0.15] dark:opacity-[0.07] pointer-events-none"
             :style="noiseOverlayStyle">
        </div>

        <!-- Ambient Background -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none z-0"
             style="background-image: radial-gradient(circle at 50% 0%, rgba(68, 133, 238, 0.1) 0%, transparent 60%);">
             <div class="absolute top-[-20%] left-[-10%] w-[60%] h-[60%] bg-primary/5 dark:bg-primary/10 blur-[120px] rounded-full"></div>
             <!-- Note: home.html doesn't explicitly have the bottom-right blob in the same way, but it does have the noise and radial gradient. 
                  However, looking closely at home.html:
                  <div class="absolute top-[-20%] left-[-10%] ..."></div> is present.
                  The background-image radial-gradient is on the main wrapper div in home.html line 285.
                  I will add the radial gradient here and keep the blobs as they add nice depth consistent with V6 vibe if they match.
             -->
        </div>

        <!-- Navbar Component -->
        <AppNavbar />

        <main class="relative z-10 flex-grow">
            <!-- Hero Section -->
            <section class="reveal-on-scroll relative pt-24 pb-20 md:pt-36 md:pb-32 px-4 sm:px-6 lg:px-8 max-w-[1200px] mx-auto">
                <div class="grid lg:grid-cols-2 gap-12 items-center">
                    <!-- Text Content -->
                    <div class="flex flex-col gap-8 order-2 lg:order-1">
                        <h1 class="font-display text-slate-900 dark:text-white text-5xl md:text-7xl font-bold leading-[1.05] tracking-tight">
                             <span v-if="settings?.welcome_hero_title">{{ settings.welcome_hero_title }}</span>
                             <span v-else>Deep dives into <span class="bg-clip-text text-transparent bg-gradient-to-r from-accent-orange to-warning">Code</span> & <span class="text-primary">Creation</span>.</span>
                        </h1>
                        <p class="text-slate-600 dark:text-slate-300 text-lg md:text-xl font-medium leading-relaxed max-w-xl">
                            {{ settings?.welcome_hero_subtitle || 'A living collection of thoughts, tutorials, and experiments growing in my digital garden.' }}
                        </p>
                        <div class="flex flex-wrap gap-4 mt-2">
                            <Link :href="route('articles.index')" class="flex items-center justify-center rounded-xl h-14 px-8 bg-primary text-white text-lg font-bold shadow-lg shadow-primary/25 hover:shadow-xl hover:shadow-primary/40 hover:bg-primary-hover transition-all transform hover:-translate-y-1">
                                Explore Notes
                            </Link>
                            <Link v-if="$page.props.auth.user" :href="route('portfolio.show', $page.props.auth.user.username || slugify($page.props.auth.user.name))" class="flex items-center justify-center rounded-xl h-14 px-8 bg-white dark:bg-white/5 backdrop-blur-sm text-slate-900 dark:text-white border border-slate-200 dark:border-white/10 hover:bg-slate-50 dark:hover:bg-white/10 text-lg font-bold transition-all">
                                View Portfolio
                            </Link>
                            <Link v-else :href="route('about')" class="flex items-center justify-center rounded-xl h-14 px-8 bg-white dark:bg-white/5 backdrop-blur-sm text-slate-900 dark:text-white border border-slate-200 dark:border-white/10 hover:bg-slate-50 dark:hover:bg-white/10 text-lg font-bold transition-all">
                                View About
                            </Link>
                        </div>

                         <!-- Heatmap -->
                         <div class="flex flex-col gap-2 mt-6">
                            <div class="flex items-end gap-2 text-xs font-bold text-slate-500 uppercase tracking-wider">
                                <span class="material-symbols-outlined text-success text-sm animate-pulse">check_circle</span>
                                <div>System Activity</div>
                            </div>
                            <div class="flex gap-1 p-2 bg-white/50 dark:bg-black/20 rounded-lg border border-white/20 w-fit backdrop-blur-sm">
                                <div class="grid grid-rows-3 grid-flow-col gap-1">
                                    <div v-for="n in 21" :key="n" class="w-2.5 h-2.5 rounded-[1px] bg-success" :class="`opacity-${Math.floor(Math.random() * 8 + 2)}0`"></div>
                                </div>
                                <div class="flex flex-col justify-end ml-2 text-[10px] text-slate-500 dark:text-slate-400 font-mono leading-tight">
                                    <div class="font-bold text-slate-900 dark:text-white">842</div>
                                    <div>commits</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Hero Visual with Portal Glow -->
                    <div class="order-1 lg:order-2 relative group perspective-1000">
                        <div class="portal-glow transition-transform duration-500 group-hover:scale-[1.02] group-hover:rotate-1">
                             <div class="relative w-full aspect-[4/3] bg-slate-100 dark:bg-slate-800 rounded-2xl overflow-hidden shadow-2xl border border-white/20 dark:border-white/5 z-10">
                                <img 
                                    alt="Coding monitor screen"
                                    class="w-full h-full object-cover opacity-90 hover:scale-105 transition-transform duration-700 ease-out"
                                    src="https://images.unsplash.com/photo-1555066931-4365d14bab8c?q=80&w=2070&auto=format&fit=crop" 
                                />
                                <!-- Glass Overlay -->
                                <div class="absolute bottom-4 left-4 right-4 p-4 glass-panel flex items-center gap-4 bg-white/10 backdrop-blur-md border-white/20">
                                    <div class="size-10 rounded-full bg-success/20 flex items-center justify-center">
                                        <span class="material-symbols-outlined text-success">terminal</span>
                                    </div>
                                    <div>
                                        <div class="text-xs text-white/70 uppercase tracking-wider font-bold">Current Status</div>
                                        <div class="text-sm text-white font-bold">System Operational</div>
                                    </div>
                                </div>
                             </div>
                        </div>
                    </div>
                </div>
            </section>

             <!-- Stats Section -->
             <section class="reveal-on-scroll px-4 sm:px-6 lg:px-8 max-w-[1200px] mx-auto mb-20">
                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                    <div class="glass-panel p-6 group hover:-translate-y-1 transition-transform stagger-delay-1 reveal-on-scroll">
                        <div class="flex items-center justify-between mb-4">
                            <p class="text-slate-500 dark:text-slate-400 font-medium text-sm uppercase tracking-wider">{{ settings?.welcome_feature_1_title || 'Latest Note' }}</p>
                            <span class="material-symbols-outlined text-accent-orange">article</span>
                        </div>
                        <h3 class="text-primary dark:text-white text-2xl font-bold leading-tight mb-2">Refactoring UI Logic</h3>
                        <a href="#" class="text-sm font-semibold text-primary flex items-center gap-1 mt-4">
                            Read Note <span class="material-symbols-outlined text-base">arrow_forward</span>
                        </a>
                    </div>
                     <div class="glass-panel p-6 group hover:-translate-y-1 transition-transform stagger-delay-2 reveal-on-scroll">
                        <div class="flex items-center justify-between mb-4">
                            <p class="text-slate-500 dark:text-slate-400 font-medium text-sm uppercase tracking-wider">{{ settings?.welcome_feature_2_title || 'Garden Size' }}</p>
                            <span class="material-symbols-outlined text-accent-orange">forest</span>
                        </div>
                        <h3 class="text-primary dark:text-white text-2xl font-bold leading-tight mb-2">42 Notes Planted</h3>
                        <div class="flex gap-2 mt-4">
                            <span class="px-2 py-1 rounded-full bg-success/10 text-success text-xs font-bold border border-success/20">12 Evergreen</span>
                            <span class="px-2 py-1 rounded-full bg-warning/10 text-warning text-xs font-bold border border-warning/20">30 Seedlings</span>
                        </div>
                    </div>
                     <div class="glass-panel p-6 group hover:-translate-y-1 transition-transform stagger-delay-3 reveal-on-scroll">
                        <div class="flex items-center justify-between mb-4">
                            <p class="text-slate-500 dark:text-slate-400 font-medium text-sm uppercase tracking-wider">{{ settings?.welcome_feature_3_title || 'Availability' }}</p>
                            <span class="material-symbols-outlined text-accent-orange">person_check</span>
                        </div>
                        <h3 class="text-primary dark:text-white text-2xl font-bold leading-tight mb-2">{{ settings?.welcome_availability_text || 'Open to Colab' }}</h3>
                        <p class="text-slate-600 dark:text-slate-400 text-sm mt-2">Ready to work on interesting projects.</p>
                    </div>
                </div>
            </section>

            <!-- Digital Garden (Notes) Section -->
            <section class="py-16 border-t border-slate-200 dark:border-white/5">
                <div class="px-4 sm:px-6 lg:px-8 max-w-[1200px] mx-auto">
                    <!-- Section Header -->
                    <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-10 reveal-on-scroll">
                        <div>
                            <h2 class="text-primary dark:text-white text-3xl font-bold tracking-tight mb-2">{{ settings?.welcome_notes_title || 'Digital Notes' }}</h2>
                            <p class="text-slate-600 dark:text-slate-400">{{ settings?.welcome_notes_subtitle || 'Notes ranging from raw ideas to polished tutorials.' }}</p>
                        </div>
                        <!-- Filter Chips -->
                        <div class="flex flex-wrap gap-2">
                            <button
                                @click="setFilter('all')"
                                class="filter-chip px-4 py-2 rounded-full text-sm font-semibold shadow-md"
                                :class="activeFilter === 'all' ? 'active' : 'bg-white dark:bg-slate-800 border border-slate-200 dark:border-white/10 text-slate-600 dark:text-slate-300'"
                            >All</button>
                            <button
                                @click="setFilter('evergreen')"
                                class="filter-chip px-4 py-2 rounded-full text-sm font-medium"
                                :class="activeFilter === 'evergreen' ? 'active' : 'bg-white dark:bg-slate-800 border border-slate-200 dark:border-white/10 text-slate-600 dark:text-slate-300'"
                            >Evergreen 🌲</button>
                            <button
                                @click="setFilter('budding')"
                                class="filter-chip px-4 py-2 rounded-full text-sm font-medium"
                                :class="activeFilter === 'budding' ? 'active' : 'bg-white dark:bg-slate-800 border border-slate-200 dark:border-white/10 text-slate-600 dark:text-slate-300'"
                            >Budding 🌿</button>
                            <button
                                @click="setFilter('seedling')"
                                class="filter-chip px-4 py-2 rounded-full text-sm font-medium"
                                :class="activeFilter === 'seedling' ? 'active' : 'bg-white dark:bg-slate-800 border border-slate-200 dark:border-white/10 text-slate-600 dark:text-slate-300'"
                            >Seedling 🌱</button>
                        </div>
                    </div>
                    <!-- Article Grid -->
                    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                        <Link 
                            v-for="(article, index) in filteredArticles" 
                            :key="article.id"
                            :href="route('articles.show', article.slug)"
                            class="glass-panel flex flex-col overflow-hidden group reveal-on-scroll hover:-translate-y-1 transition-transform duration-300"
                            :class="`stagger-delay-${index + 1}`"
                        >
                            <div class="h-48 w-full bg-slate-100 dark:bg-slate-800 overflow-hidden relative">
                                <template v-if="article.cover_image">
                                    <img 
                                        :src="article.cover_image" 
                                        :alt="article.title"
                                        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                                    />
                                </template>
                                <!-- Using placeholder if no image -->
                                <div v-else class="w-full h-full bg-gradient-to-br from-slate-200 to-slate-300 dark:from-slate-700 dark:to-slate-800 group-hover:scale-105 transition-transform duration-500 flex items-center justify-center">
                                     <span class="text-4xl filter grayscale opacity-50">{{ index % 2 === 0 ? '📝' : '💡' }}</span>
                                </div>
                                <div class="absolute top-4 right-4">
                                    <span class="inline-flex items-center gap-1 rounded-md bg-white/20 backdrop-blur-md px-2 py-1 text-xs font-semibold text-white shadow-sm border border-white/20">
                                        Note
                                    </span>
                                </div>
                            </div>
                            <div class="p-6 flex flex-col flex-1">
                                <div class="flex items-center gap-2 text-xs font-medium text-slate-500 mb-3">
                                    <span>{{ formatDate(article.published_at) }}</span>
                                    <span>•</span>
                                    <span>{{ article.notebook?.name || 'General' }}</span>
                                </div>
                                <h3 class="text-xl font-bold text-primary dark:text-white mb-3 group-hover:text-primary transition-colors leading-tight line-clamp-2">
                                    {{ article.title }}
                                </h3>
                                <p class="text-slate-600 dark:text-slate-400 text-sm leading-relaxed mb-6 flex-1 line-clamp-3">
                                    {{ article.excerpt || 'Read more about this topic...' }}
                                </p>
                                <span class="inline-flex items-center text-sm font-bold text-primary dark:text-white hover:text-primary transition-colors group/link">
                                    Read Note
                                    <span class="material-symbols-outlined ml-1 text-lg group-hover/link:translate-x-1 transition-transform">arrow_forward</span>
                                </span>
                            </div>
                        </Link>
                    </div>
                </div>
            </section>

             <!-- Newsletter Footer -->
            <footer class="mt-20 reveal-on-scroll">
                <!-- Newsletter Section -->
                <div class="px-4 sm:px-6 lg:px-8 max-w-[1200px] mx-auto mb-12">
                    <div class="glass-panel newsletter-section p-8 md:p-12 relative overflow-hidden border-2 border-primary/10">
                        <!-- Glow effect -->
                        <div class="newsletter-glow"></div>

                        <div class="relative z-10 flex flex-col md:flex-row gap-8 items-center justify-between">
                            <div class="text-center md:text-left max-w-lg">
                                <h2 class="font-display text-3xl md:text-4xl font-bold text-slate-900 dark:text-white mb-4">
                                    Join the <span class="bg-clip-text text-transparent bg-gradient-to-r from-accent-orange to-warning">Digital Notes</span>
                                </h2>
                                <p class="text-slate-600 dark:text-slate-400 text-lg">
                                    Receive weekly seeds of thought on coding, design, and building living software. No spam, just growth.
                                </p>
                            </div>
                            <div class="w-full max-w-md">
                                <form class="flex flex-col sm:flex-row gap-3" @submit.prevent="submitNewsletter">
                                    <div class="flex-1">
                                        <input 
                                            v-model="form.email"
                                            type="email" 
                                            placeholder="your@email.com"
                                            class="w-full bg-white/50 dark:bg-black/20 backdrop-blur-sm border border-slate-200 dark:border-white/10 rounded-xl px-4 py-3 outline-none focus:border-primary/50 text-slate-900 dark:text-white placeholder-slate-500 transition-colors" 
                                            :disabled="form.processing"
                                        />
                                        <div v-if="form.errors.email" class="text-red-500 text-xs mt-1 ml-1">{{ form.errors.email }}</div>
                                    </div>
                                    <button 
                                        type="submit"
                                        class="bg-primary hover:bg-primary-hover text-white font-bold py-3 px-6 rounded-xl transition-all hover:shadow-lg hover:shadow-primary/25 active:scale-95 disabled:opacity-70 disabled:cursor-not-allowed"
                                        :disabled="form.processing"
                                    >
                                        <span v-if="form.processing">Sending...</span>
                                        <span v-else>Subscribe</span>
                                    </button>
                                </form>
                                <div v-if="$page.props.flash?.success" class="mt-3 text-sm text-green-600 dark:text-green-400 font-bold flex items-center gap-2">
                                    <span class="material-symbols-outlined text-lg">check_circle</span>
                                    {{ $page.props.flash.success }}
                                </div>
                                <div v-else-if="$page.props.flash?.info" class="mt-3 text-sm text-blue-600 dark:text-blue-400 font-bold flex items-center gap-2">
                                    <span class="material-symbols-outlined text-lg">info</span>
                                    {{ $page.props.flash.info }}
                                </div>
                                <p v-else class="text-xs text-slate-500 mt-3 text-center md:text-left">
                                    Join 2,000+ other gardeners. Unsubscribe anytime.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- App Footer -->
                <AppFooter :settings="settings" />
            </footer>

        </main>
    </div>
</template>
