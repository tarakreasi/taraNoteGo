<script setup>
import { Head, Link, useForm } from '@inertiajs/vue3';
import DOMPurify from 'isomorphic-dompurify';
import { useTheme } from '@/composables/useTheme';
import { useScrollReveal } from '@/composables/useScrollReveal';
import ArticleLayout from '@/Layouts/ArticleLayout.vue';

defineOptions({ layout: ArticleLayout });

const props = defineProps({
    article: Object,
    settings: {
        type: Object,
        default: () => ({})
    }
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

const slugify = (text) => {
    return text.toString().toLowerCase()
        .replace(/\s+/g, '-')
        .replace(/[^\w\-]+/g, '')
        .replace(/\-\-+/g, '-')
        .replace(/^-+/, '') 
        .replace(/-+$/, '');
};

// Share Logic
import { ref } from 'vue';
const copyStatus = ref('');

const shareToLinkedin = () => {
    const url = window.location.href;
    const title = props.article.title;
    window.open(`https://www.linkedin.com/sharing/share-offsite/?url=${encodeURIComponent(url)}`, '_blank');
};

const copyUrl = async () => {
    try {
        await navigator.clipboard.writeText(window.location.href);
        copyStatus.value = 'Copied!';
        setTimeout(() => copyStatus.value = '', 2000);
    } catch (err) {
        console.error('Failed to copy', err);
    }
};
</script>

<template>
    <Head :title="article.title" />

    <div class="flex-1 max-w-[1200px] w-full mx-auto px-4 sm:px-6 py-8">
        <div class="flex flex-col lg:flex-row gap-12">

            <!-- Left Sidebar: Author Profile (Sticky) -->
            <aside class="w-full lg:w-[260px] shrink-0 hidden lg:block">
                <div class="sticky top-24 flex flex-col gap-8">
                    
                    <!-- Back Button -->
                    <Link :href="route('articles.index')" class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors flex items-center gap-1">
                        <span class="material-symbols-outlined">first_page</span>
                        <span class="text-xs font-bold uppercase tracking-wider">Back</span>
                    </Link>

                    <!-- Author Card -->
                    <div class="flex flex-col">
                        <Link :href="route('portfolio.show', article.user?.username || slugify(article.user?.name))" class="group flex items-center gap-4 mb-4">
                            <div class="size-14 rounded-full overflow-hidden bg-slate-100 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 group-hover:opacity-80 transition-opacity shrink-0">
                                <img v-if="article.user?.avatar" :src="article.user.avatar" class="w-full h-full object-cover">
                                <div v-else class="w-full h-full flex items-center justify-center text-slate-400 font-bold text-lg">
                                    {{ article.user?.name?.charAt(0) || 'U' }}
                                </div>
                            </div>
                            <div class="flex flex-col">
                                <h3 class="font-bold text-slate-900 dark:text-white text-base leading-tight group-hover:underline decoration-2 underline-offset-4">
                                    {{ article.user?.name }}
                                </h3>
                                <span class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">{{ article.user?.hero_tagline || 'Note Taker' }}</span>
                            </div>
                        </Link>
                        
                        <p class="text-sm text-slate-600 dark:text-slate-300 leading-relaxed mb-4">
                            {{ article.user?.bio || 'Documenting the journey of learning and creating. Digital gardener.' }}
                        </p>
                        
                        <button class="w-full py-2 px-4 rounded-md bg-slate-900 dark:bg-white text-white dark:text-slate-900 text-sm font-bold hover:opacity-90 transition-opacity flex items-center justify-center gap-2">
                            <span class="material-symbols-outlined text-[18px]">add</span> Follow
                        </button>

                        <!-- Social Links -->
                        <div class="flex gap-4 mt-6 border-t border-slate-100 dark:border-slate-800 pt-6">
                            <a v-if="settings.footer_social_twitter" :href="settings.footer_social_twitter" target="_blank" class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors">
                                <span class="material-symbols-outlined text-[20px]">link</span>
                            </a>
                            <a v-if="settings.footer_social_github" :href="settings.footer_social_github" target="_blank" class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors">
                                <svg class="w-5 h-5 fill-current" viewBox="0 0 24 24"><path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/></svg>
                            </a>
                        </div>
                    </div>

                </div>
            </aside>

            <!-- Main Content -->
            <main class="flex-1 min-w-0">
                <!-- Cover Image -->
                <div v-if="article.cover_image" class="aspect-[2/1] rounded-2xl overflow-hidden mb-8 border border-slate-100 dark:border-slate-800">
                    <img :src="article.cover_image" class="w-full h-full object-cover">
                </div>
                
                <!-- Article Header -->
                <header class="mb-10">
                    <div class="flex items-center gap-2 mb-4 text-sm">
                        <Link :href="route('articles.index', { topic: article.notebook?.slug })" class="px-2 py-1 bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-300 rounded font-medium hover:bg-emerald-500 hover:text-white transition-colors">
                            {{ article.notebook?.name || 'Uncategorized' }}
                        </Link>
                        <span class="text-slate-300">•</span>
                        <span class="text-slate-400">{{ formatDate(article.published_at) }}</span>
                        <span class="text-slate-300">•</span>
                        <span class="text-slate-400">{{ calculateReadTime(article.content) }}</span>
                    </div>

                    <h1 class="text-3xl md:text-5xl font-extrabold text-slate-900 dark:text-white leading-tight mb-6">
                        {{ article.title }}
                    </h1>
                </header>

                <!-- Content Body -->
                <article class="prose prose-lg dark:prose-invert max-w-none prose-slate
                    prose-headings:font-bold prose-headings:tracking-tight prose-headings:text-slate-900 dark:prose-headings:text-white
                    prose-p:text-slate-600 dark:prose-p:text-slate-300 prose-p:leading-relaxed
                    prose-a:text-emerald-500 prose-a:no-underline hover:prose-a:underline
                    prose-img:rounded-xl prose-img:shadow-sm"
                >
                    <div v-html="DOMPurify.sanitize(article.content)"></div>
                </article>

                <!-- MOBILE AUTHOR CARD (Visible on small screens) -->
                <div class="lg:hidden mt-12 pt-8 border-t border-slate-100 dark:border-slate-800">
                    <div class="bg-slate-50 dark:bg-slate-800/50 rounded-2xl p-6">
                        <div class="flex items-center gap-4 mb-4">
                            <Link :href="route('portfolio.show', article.user?.username || slugify(article.user?.name))" class="shrink-0">
                                <div class="size-12 rounded-full overflow-hidden bg-slate-200 dark:bg-slate-700">
                                    <img v-if="article.user?.avatar" :src="article.user.avatar" class="w-full h-full object-cover">
                                    <div v-else class="w-full h-full flex items-center justify-center text-slate-400 font-bold">
                                        {{ article.user?.name?.charAt(0) || 'U' }}
                                    </div>
                                </div>
                            </Link>
                            <div>
                                <h3 class="font-bold text-slate-900 dark:text-white">{{ article.user?.name }}</h3>
                                <p class="text-xs text-slate-500 dark:text-slate-400">{{ article.user?.hero_tagline || 'Note Taker' }}</p>
                            </div>
                            <button class="ml-auto px-4 py-1.5 bg-slate-900 dark:bg-white text-white dark:text-slate-900 text-xs font-bold rounded-full hover:opacity-90 transition-opacity">
                                Follow
                            </button>
                        </div>
                        <p class="text-sm text-slate-600 dark:text-slate-300 mb-4 leading-relaxed">
                            {{ article.user?.bio || 'Documenting the journey of learning and creating. Digital gardener.' }}
                        </p>
                        <!-- Mobile Socials -->
                         <div class="flex gap-4">
                            <a v-if="settings.footer_social_twitter" :href="settings.footer_social_twitter" target="_blank" class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors">
                                <span class="material-symbols-outlined text-[20px]">link</span>
                            </a>
                            <a v-if="settings.footer_social_github" :href="settings.footer_social_github" target="_blank" class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors">
                                <svg class="w-5 h-5 fill-current" viewBox="0 0 24 24"><path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/></svg>
                            </a>
                        </div>
                    </div>
                </div>

                <!-- Tags / Footer of Article -->
                <div class="mt-12 pt-8 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between">
                        <div class="flex gap-2">
                        <!-- Tags could go here -->
                        </div>
                        
                        <div class="flex items-center gap-4 relative">
                            <!-- Share Dropdown Wrapper -->
                            <div class="relative group">
                                <button class="flex items-center gap-2 text-slate-400 hover:text-sky-500 transition-colors">
                                    <span class="material-symbols-outlined text-[20px]">ios_share</span>
                                    <span class="font-bold text-sm">Share</span>
                                </button>
                                
                                <!-- Dropdown Menu -->
                                <div class="absolute right-0 bottom-full mb-2 w-48 bg-white dark:bg-slate-800 rounded-lg shadow-xl border border-slate-100 dark:border-slate-700 py-1 invisible opacity-0 group-hover:visible group-hover:opacity-100 transition-all z-20">
                                    <button 
                                        @click="shareToLinkedin"
                                        class="w-full text-left px-4 py-2 text-sm text-slate-700 dark:text-slate-200 hover:bg-slate-50 dark:hover:bg-slate-700 flex items-center gap-2"
                                    >
                                        <svg class="w-4 h-4 fill-[#0077b5]" viewBox="0 0 24 24"><path d="M19 0h-14c-2.761 0-5 2.239-5 5v14c0 2.761 2.239 5 5 5h14c2.762 0 5-2.239 5-5v-14c0-2.761-2.238-5-5-5zm-11 19h-3v-11h3v11zm-1.5-12.268c-.966 0-1.75-.79-1.75-1.764s.784-1.764 1.75-1.764 1.75.79 1.75 1.764-.783 1.764-1.75 1.764zm13.5 12.268h-3v-5.604c0-3.368-4-3.113-4 0v5.604h-3v-11h3v1.765c1.396-2.586 7-2.777 7 2.476v6.759z"/></svg>
                                        LinkedIn
                                    </button>
                                    <button 
                                        @click="copyUrl"
                                        class="w-full text-left px-4 py-2 text-sm text-slate-700 dark:text-slate-200 hover:bg-slate-50 dark:hover:bg-slate-700 flex items-center gap-2"
                                    >
                                        <span class="material-symbols-outlined text-[16px]">content_copy</span>
                                        {{ copyStatus || 'Copy Link' }}
                                    </button>
                                </div>
                            </div>
                        </div>
                </div>
            </main>
        </div>
    </div>
</template>
