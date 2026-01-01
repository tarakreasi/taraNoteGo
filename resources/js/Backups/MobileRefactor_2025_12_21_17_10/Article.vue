<script setup>
import { Head, Link, useForm } from '@inertiajs/vue3';
import DOMPurify from 'isomorphic-dompurify';
import { useTheme } from '@/composables/useTheme';
import { useScrollReveal } from '@/composables/useScrollReveal';
import AppNavbar from '@/Components/AppNavbar.vue';
import AppFooter from '@/Components/AppFooter.vue';

const props = defineProps({
    article: Object,
});

const formatDate = (dateString) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    });
};

// Theme management
const { isDark, toggleTheme } = useTheme();

// Scroll reveal
useScrollReveal({ threshold: 0.1, once: true });

// Noise overlay SVG (inline to avoid CSS linter warnings)
const noiseOverlayStyle = "background-image: url('data:image/svg+xml,%3Csvg viewBox=\\'0 0 200 200\\' xmlns=\\'http://www.w3.org/2000/svg\\'%3E%3Cfilter id=\\'noiseFilter\\'%3E%3CfeTurbulence type=\\'fractalNoise\\' baseFrequency=\\'0.65\\' numOctaves=\\'3\\' stitchTiles=\\'stitch\\'%3E%3C/feTurbulence%3E%3C/filter%3E%3Crect width=\\'100%25\\' height=\\'100%25\\' filter=\\'url(%23noiseFilter)\\'%3E%3C/rect%3E%3C/svg%3E');";

// Placeholder for read time calculation
const calculateReadTime = (content) => {
    if (!content) return '1 min read';
    const wordsPerMinute = 200; // Average reading speed
    const textContent = content.replace(/<[^>]*>?/gm, ''); // Remove HTML tags
    const wordCount = textContent.split(/\s+/).length;
    const minutes = Math.ceil(wordCount / wordsPerMinute);
    return `${minutes} min read`;
};

// Placeholder for Table of Contents (Ideally this would be generated from article content)
const tableOfContents = [
    { id: 'introduction', title: 'Introduction' },
    // more items would be parsed from content
];

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
    <Head :title="article.title" />

    <div class="text-[#181511] dark:text-white antialiased transition-colors duration-200 font-body relative overflow-x-hidden bg-canvas-light dark:bg-canvas-dark">
        
        <!-- Noise Overlay -->
        <div class="fixed inset-0 z-[9999] opacity-[0.15] dark:opacity-[0.07] pointer-events-none" :style="noiseOverlayStyle"></div>

        <!-- Ambient Background -->
        <!-- Adding wrapper with sticky behavior if needed or just absolute background -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none z-0" style="background-image: radial-gradient(circle at 50% 0%, rgba(68, 133, 238, 0.1) 0%, transparent 60%);">
            <div class="absolute top-[-20%] left-[-10%] w-[60%] h-[60%] bg-primary/5 dark:bg-primary/10 blur-[120px] rounded-full"></div>
            <div class="absolute bottom-[-20%] right-[-10%] w-[50%] h-[50%] bg-accent-orange/5 dark:bg-accent-orange/10 blur-[100px] rounded-full"></div>
        </div>

        <div class="relative flex h-auto min-h-screen w-full flex-col overflow-x-hidden">
        <!-- Navbar -->
        <AppNavbar />

            <main class="flex-grow flex flex-col relative w-full max-w-[1200px] mx-auto px-4 sm:px-6 lg:px-8 pt-32 pb-10 gap-10">
                <!-- Ambient Background Blob -->
                <div class="absolute top-[0%] right-[0%] w-[50%] h-[50%] bg-primary/5 dark:bg-primary/10 blur-[120px] rounded-full pointer-events-none z-[-1]"></div>
                
                <!-- Grid Layout: Sidebar Left (Social), Main Content, Sidebar Right (TOC) -->
                <div class="lg:grid lg:grid-cols-12 lg:gap-8">
                    <!-- Left Sidebar: Social Sharing (Sticky) -->
                    <aside class="hidden lg:block lg:col-span-1 lg:col-start-2">
                        <div class="sticky top-32 flex flex-col gap-6 items-center">
                            <button class="group flex flex-col items-center gap-1 text-gray-400 hover:text-primary transition-colors">
                                <div class="p-2 rounded-full group-hover:bg-primary/10 transition-colors">
                                    <span class="material-symbols-outlined">favorite</span>
                                </div>
                                <span class="text-xs font-medium">124</span>
                            </button>
                            <button class="group flex flex-col items-center gap-1 text-gray-400 hover:text-blue-400 transition-colors">
                                <div class="p-2 rounded-full group-hover:bg-blue-400/10 transition-colors">
                                    <span class="material-symbols-outlined">chat_bubble</span>
                                </div>
                                <span class="text-xs font-medium">18</span>
                            </button>
                            <button class="group flex flex-col items-center gap-1 text-gray-400 hover:text-gray-900 dark:hover:text-white transition-colors">
                                <div class="p-2 rounded-full group-hover:bg-gray-100 dark:group-hover:bg-gray-800 transition-colors">
                                    <span class="material-symbols-outlined">bookmark</span>
                                </div>
                            </button>
                            <div class="w-px h-8 bg-gray-200 dark:bg-gray-700"></div>
                            <button class="p-2 text-gray-400 hover:text-primary transition-colors">
                                <span class="material-symbols-outlined">share</span>
                            </button>
                        </div>
                    </aside>

                    <!-- Main Content Area: The Deep Dive -->
                    <article class="glass-panel relative overflow-hidden p-8 md:p-12 col-span-12 lg:col-span-8 lg:col-start-3 max-w-[760px] mx-auto w-full">
                        <!-- Ambient Glow -->
                        <div class="absolute top-0 right-0 w-[400px] h-[400px] bg-primary/20 rounded-full blur-[100px] opacity-50 pointer-events-none -z-10 translate-x-1/3 -translate-y-1/3"></div>
                        
                        <!-- Cover Image -->
                        <div v-if="article.cover_image" class="w-full h-64 md:h-96 mb-8 rounded-2xl overflow-hidden shadow-2xl relative group">
                            <img :src="article.cover_image" :alt="article.title" class="w-full h-full object-cover">
                            <!-- Overlay gradient for text readability if we overlay title later, but for now just subtle fade -->
                            <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>
                        </div>

                        <!-- Breadcrumbs -->
                        <nav class="flex flex-wrap gap-2 items-center text-sm mb-6 text-gray-500 dark:text-gray-400">
                            <Link class="hover:text-primary transition-colors" href="/">Home</Link>
                            <span class="material-symbols-outlined text-[16px] leading-none">chevron_right</span>
                            <Link class="hover:text-primary transition-colors" :href="route('articles.index')">Garden</Link>
                            <span class="material-symbols-outlined text-[16px] leading-none">chevron_right</span>
                            <span class="text-primary font-medium">{{ article.notebook?.name || 'General' }}</span>
                        </nav>

                        <!-- Article Header -->
                        <header class="mb-10">
                            <!-- Tags -->
                            <div class="flex gap-2 mb-4">
                                <span v-if="article.notebook" class="inline-flex items-center px-2.5 py-0.5 rounded-md text-xs font-medium bg-accent-orange/10 text-accent-orange border border-accent-orange/20">
                                    {{ article.notebook.name }}
                                </span>
                                <span class="inline-flex items-center px-2.5 py-0.5 rounded-md text-xs font-medium bg-blue-100 text-blue-800 dark:bg-blue-900/30 dark:text-blue-200">
                                    Article
                                </span>
                            </div>
                            <h1 class="font-display text-3xl sm:text-4xl md:text-5xl font-black tracking-tight text-gray-900 dark:text-white leading-[1.15] mb-6">
                                {{ article.title }}
                            </h1>
                            <!-- Author Meta -->
                            <div class="flex items-center justify-between border-b border-gray-100 dark:border-gray-800 pb-8">
                                <div class="flex items-center gap-4">
                                    <div class="relative">
                                        <div class="size-12 rounded-full bg-cover bg-center ring-2 ring-white dark:ring-gray-800 shadow-sm"
                                            :style="{ backgroundImage: `url(${article.user?.avatar_url || 'https://via.placeholder.com/150'})` }">
                                        </div>
                                        <div class="absolute bottom-0 right-0 size-3 bg-green-500 border-2 border-white dark:border-gray-900 rounded-full"></div>
                                    </div>
                                    <div class="flex flex-col">
                                        <Link 
                                            :href="route('portfolio.show', article.user?.username || slugify(article.user?.name))" 
                                            class="text-base font-bold text-gray-900 dark:text-white hover:text-primary transition-colors"
                                        >
                                            {{ article.user?.name }}
                                        </Link>
                                        <div class="flex items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
                                            <span>{{ formatDate(article.published_at || article.created_at) }}</span>
                                            <span>•</span>
                                            <span>{{ calculateReadTime(article.content) }}</span>
                                        </div>
                                    </div>
                                </div>
                                <div class="flex gap-2 md:hidden">
                                    <button class="p-2 rounded-full bg-gray-50 dark:bg-gray-800 text-gray-500 hover:text-primary">
                                        <span class="material-symbols-outlined text-[20px]">share</span>
                                    </button>
                                </div>
                            </div>
                        </header>

                        <!-- Article Body -->
                        <div class="prose prose-lg dark:prose-invert max-w-none text-gray-600 dark:text-gray-300" v-html="DOMPurify.sanitize(article.content)"></div>

                        <!-- Action Footer -->
                        <div class="mt-16 pt-8 border-t border-gray-100 dark:border-gray-800">
                            <div class="flex flex-col sm:flex-row justify-between items-center gap-6 p-6 bg-gray-50 dark:bg-gray-800/50 rounded-xl">
                                <div class="text-center sm:text-left">
                                    <h4 class="font-bold text-gray-900 dark:text-white text-lg">Enjoyed this article?</h4>
                                    <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">Subscribe to get the latest seeds sent to your inbox.</p>
                                </div>
                                <div class="w-full sm:w-auto flex flex-col items-center sm:items-end">
                                    <form @submit.prevent="submitNewsletter" class="flex gap-2 w-full sm:w-auto">
                                        <div class="flex-1">
                                            <input
                                                v-model="form.email"
                                                class="w-full min-w-0 sm:w-64 rounded-lg border-gray-200 dark:border-white/10 dark:bg-black/20 dark:text-white focus:ring-primary focus:border-primary"
                                                placeholder="Email address" type="email" 
                                                :disabled="form.processing"
                                            />
                                        </div>
                                        <button
                                            type="submit"
                                            class="bg-primary hover:bg-primary-hover text-white font-medium py-2 px-4 rounded-lg transition-colors shadow-sm shadow-blue-500/30 disabled:opacity-70 disabled:cursor-not-allowed"
                                            :disabled="form.processing"
                                        >
                                            <span v-if="form.processing">...</span>
                                            <span v-else>Subscribe</span>
                                        </button>
                                    </form>
                                    <div v-if="form.errors.email" class="text-red-500 text-xs mt-1">{{ form.errors.email }}</div>
                                    <div v-if="$page.props.flash?.success" class="mt-2 text-sm text-green-600 dark:text-green-400 font-bold flex items-center gap-1">
                                        <span class="material-symbols-outlined text-sm">check_circle</span>
                                        {{ $page.props.flash.success }}
                                    </div>
                                    <div v-else-if="$page.props.flash?.info" class="mt-2 text-sm text-blue-600 dark:text-blue-400 font-bold flex items-center gap-1">
                                        <span class="material-symbols-outlined text-sm">info</span>
                                        {{ $page.props.flash.info }}
                                    </div>
                                </div>
                            </div>
                        </div>
                    </article>

                    <!-- Right Sidebar: TOC (Sticky) -->
                    <aside class="hidden xl:block xl:col-span-3 xl:col-start-11">
                        <div class="sticky top-32">
                            <h5 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-4">On this page</h5>
                            <ul class="space-y-3 text-sm border-l border-gray-100 dark:border-gray-800 pl-4">
                                <li v-for="item in tableOfContents" :key="item.id">
                                    <a :class="['block hover:text-primary transition-colors', item.id === 'introduction' ? 'text-gray-900 dark:text-white font-medium -ml-[17px] pl-4 border-l-2 border-primary' : 'text-gray-500 dark:text-gray-400']" :href="`#${item.id}`">{{ item.title }}</a>
                                </li>
                            </ul>
                        </div>
                    </aside>
                </div>
            </main>

            <!-- Footer -->
            <AppFooter />
        </div>
    </div>
</template>
