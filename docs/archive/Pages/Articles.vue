<script setup>
import { Head, Link } from '@inertiajs/vue3';
import { ref } from 'vue';
import { useScrollReveal } from '@/composables/useScrollReveal';
import ArticleLayout from '@/Layouts/ArticleLayout.vue';

defineOptions({ layout: ArticleLayout });

const props = defineProps({
    notes: Object,
    topics: Array,
    currentTopic: String,
    trendingNotes: Array,
    settings: Object,
    search: String
});

// Format Date
const formatDate = (dateString, type = 'short') => {
    if (!dateString) return '';
    const date = new Date(dateString);
    if (type === 'relative') {
        const now = new Date();
        const diff = now - date;
        const days = Math.floor(diff / (1000 * 60 * 60 * 24));
        if (days === 0) return 'Today';
        if (days === 1) return 'Yesterday';
        if (days < 7) return `${days} days ago`;
    }
    return date.toLocaleDateString('ja-JP', { // Using ja-JP style or similar short format for the look
        year: 'numeric',
        month: 'numeric',
        day: 'numeric'
    }).replace(/\//g, '/');
};

const formatTimeAgo = (dateString) => {
    if (!dateString) return '';
    const date = new Date(dateString);
    const now = new Date();
    const diffInSeconds = Math.floor((now - date) / 1000);
    
    if (diffInSeconds < 60) return 'Just now';
    if (diffInSeconds < 3600) return `${Math.floor(diffInSeconds / 60)} mins ago`;
    if (diffInSeconds < 86400) return `${Math.floor(diffInSeconds / 3600)} hours ago`;
    return `${Math.floor(diffInSeconds / 86400)} days ago`;
};


// Scroll reveal animations
// Scroll reveal animations
useScrollReveal({ threshold: 0.1, once: true });

// Share Logic
const shareToLinkedin = (note) => {
    const url = route('articles.show', note.slug);
    window.open(`https://www.linkedin.com/sharing/share-offsite/?url=${encodeURIComponent(url)}`, '_blank');
};

const copyUrl = async (note) => {
    try {
        const url = route('articles.show', note.slug);
        await navigator.clipboard.writeText(url);
        alert('Link copied to clipboard!'); 
    } catch (err) {
        console.error('Failed to copy', err);
    }
};
</script>

<template>
    <Head title="Articles" />

    <div class="flex-1 max-w-[1200px] w-full mx-auto px-4 sm:px-6 py-8">
        <div class="flex flex-col lg:flex-row gap-12">
            
            <!-- LEFT SIDEBAR: Categories -->
            <aside class="w-full lg:w-[260px] shrink-0 hidden lg:block">
                <div class="sticky top-24">
                    
                    <!-- Spaces Group -->
                    <nav class="flex flex-col gap-1">
                        <div class="flex items-center justify-between px-3 mb-2 cursor-pointer group">
                            <h3 class="font-bold text-lg text-slate-900 dark:text-white">Spaces</h3>
                            <span class="material-symbols-outlined text-slate-400 text-sm group-hover:text-slate-900 dark:group-hover:text-white transition-colors">expand_more</span>
                        </div>
                        
                        <div class="pl-0">
                            <Link 
                                v-for="topic in topics" 
                                :key="topic.id" 
                                :href="route('articles.index', { topic: topic.slug })"
                                class="flex items-center justify-between px-3 py-2 rounded-lg transition-colors text-sm font-medium"
                                :class="currentTopic === topic.slug 
                                    ? 'bg-slate-900 dark:bg-slate-800 text-white shadow-sm' 
                                    : 'text-slate-500 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800/50 hover:text-slate-900 dark:hover:text-white'"
                            >
                                <span>{{ topic.name }}</span>
                            </Link>
                        </div>
                    </nav>

                </div>
            </aside>

            <!-- MAIN AREA -->
            <main class="flex-1 min-w-0">
                
                <!-- MOBILE SPACES NAV (Visible on small screens) -->
                <div class="lg:hidden mb-8 overflow-x-auto pb-2 -mx-4 px-4 scrollbar-hide">
                    <div class="flex items-center gap-2">
                        <Link 
                            :href="route('articles.index')"
                            class="whitespace-nowrap px-4 py-2 rounded-full text-sm font-medium border transition-colors"
                            :class="!currentTopic 
                                ? 'bg-slate-900 dark:bg-white text-white dark:text-slate-900 border-slate-900 dark:border-white' 
                                : 'bg-white dark:bg-slate-800 text-slate-600 dark:text-slate-300 border-slate-200 dark:border-slate-700'"
                        >
                            All Spaces
                        </Link>
                        <Link 
                            v-for="topic in topics" 
                            :key="topic.id" 
                            :href="route('articles.index', { topic: topic.slug })"
                            class="whitespace-nowrap px-4 py-2 rounded-full text-sm font-medium border transition-colors"
                            :class="currentTopic === topic.slug 
                                ? 'bg-slate-900 dark:bg-white text-white dark:text-slate-900 border-slate-900 dark:border-white' 
                                : 'bg-white dark:bg-slate-800 text-slate-600 dark:text-slate-300 border-slate-200 dark:border-slate-700'"
                        >
                            {{ topic.name }}
                        </Link>
                    </div>
                </div>

                <!-- Section Header -->
                <div class="flex items-center gap-2 mb-6">
                    <h2 class="text-xl md:text-2xl font-bold text-slate-900 dark:text-white">
                        {{ currentTopic ? topics.find(t => t.slug === currentTopic)?.name : 'Recommended' }}
                    </h2>
                    <span class="material-symbols-outlined text-slate-400 text-xl">chevron_right</span>
                </div>

                <!-- Articles Grid -->
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                    <Link 
                        v-for="note in notes.data" 
                        :key="note.id"
                        :href="route('articles.show', note.slug)"
                        class="group flex flex-col gap-3 cursor-pointer"
                    >
                        <!-- Cover Image -->
                        <div class="aspect-[1.6/1] rounded-xl overflow-hidden bg-slate-100 dark:bg-slate-800 relative border border-slate-100 dark:border-slate-800">
                            <img 
                                v-if="note.cover_image" 
                                :src="note.cover_image" 
                                class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
                                alt="Article cover"
                            >
                            <!-- Note Icon Fallback -->
                            <div v-else class="w-full h-full flex items-center justify-center text-slate-300 dark:text-slate-600">
                                <span class="material-symbols-outlined text-[48px]">article</span>
                            </div>
                        </div>

                        <!-- Card Content -->
                        <div class="flex flex-col gap-2">
                            <!-- Title -->
                            <h3 class="font-bold text-lg leading-snug text-slate-900 dark:text-white group-hover:text-slate-600 dark:group-hover:text-slate-300 transition-colors line-clamp-2">
                                {{ note.title }}
                            </h3>

                            <!-- Author & Meta -->
                            <div class="flex items-center gap-2 mt-1">
                                <div class="size-5 rounded-full bg-slate-200 dark:bg-slate-700 overflow-hidden shrink-0">
                                    <img 
                                        v-if="note.user?.avatar_url" 
                                        :src="note.user.avatar_url" 
                                        class="w-full h-full object-cover"
                                    >
                                    <div v-else class="w-full h-full flex items-center justify-center text-[10px] font-bold text-slate-500">
                                        {{ note.user?.name?.charAt(0) || 'U' }}
                                    </div>
                                </div>
                                <span class="text-xs font-medium text-slate-900 dark:text-slate-300 truncate max-w-[100px]">
                                    {{ note.user?.name }}
                                </span>
                                <span class="text-xs text-slate-400">•</span>
                                <span class="text-xs text-slate-400">{{ formatTimeAgo(note.published_at) }}</span>
                            </div>

                            <!-- Footer Actions (Like, Bookmark) -->
                            <!-- Footer Actions -->
                            <div class="mt-auto pt-4 flex items-center justify-end relative">
                                <div class="relative group/share">
                                    <button @click.prevent class="flex items-center gap-1 text-slate-400 hover:text-sky-500 transition-colors">
                                        <span class="material-symbols-outlined text-[18px]">ios_share</span>
                                        <span class="text-xs font-bold">Share</span>
                                    </button>

                                    <!-- Dropdown -->
                                    <div class="absolute right-0 bottom-full mb-2 w-40 bg-white dark:bg-slate-800 rounded-lg shadow-xl border border-slate-100 dark:border-slate-700 py-1 invisible opacity-0 group-hover/share:visible group-hover/share:opacity-100 transition-all z-20">
                                        <button 
                                            @click.prevent="shareToLinkedin(note)"
                                            class="w-full text-left px-4 py-2 text-xs font-bold text-slate-700 dark:text-slate-200 hover:bg-slate-50 dark:hover:bg-slate-700 flex items-center gap-2"
                                        >
                                            Linked In
                                        </button>
                                        <button 
                                            @click.prevent="copyUrl(note)"
                                            class="w-full text-left px-4 py-2 text-xs font-bold text-slate-700 dark:text-slate-200 hover:bg-slate-50 dark:hover:bg-slate-700 flex items-center gap-2"
                                        >
                                            Copy Link
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </Link>
                </div>

                <!-- Empty State -->
                <div v-if="notes.data.length === 0" class="py-20 text-center">
                    <div class="inline-flex size-16 items-center justify-center rounded-full bg-slate-100 dark:bg-slate-800 mb-4">
                        <span class="material-symbols-outlined text-3xl text-slate-400">search_off</span>
                    </div>
                    <h3 class="text-lg font-bold text-slate-900 dark:text-white">No articles found</h3>
                    <p class="text-slate-500 dark:text-slate-400 max-w-md mx-auto mt-2">
                        Try adjusting your search or filter to find what you're looking for.
                    </p>
                </div>

                <!-- Pagination -->
                <div v-if="notes.data.length > 0" class="mt-12 flex justify-center gap-2">
                    <Link
                        v-for="(link, i) in notes.links"
                        :key="i"
                        :href="link.url || '#'"
                        v-html="link.label"
                        class="px-4 py-2 text-sm font-medium rounded-lg transition-colors"
                        :class="[
                            !link.url ? 'text-slate-400 cursor-not-allowed' : 
                            link.active ? 'bg-slate-900 dark:bg-white text-white dark:text-slate-900' : 
                            'text-slate-600 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-800'
                        ]"
                    />
                </div>

            </main>
        </div>

    </div>
</template>
