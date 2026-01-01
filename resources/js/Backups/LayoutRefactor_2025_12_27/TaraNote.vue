<script setup>
import { Head, Link } from '@inertiajs/vue3';
import { ref, computed } from 'vue';
import DOMPurify from 'isomorphic-dompurify';
import { useTheme } from '@/composables/useTheme';
import FloatingDock from '@/Components/FloatingDock.vue';

const props = defineProps({
    notes: Array,
    notebooks: Array,
});

// State
const searchQuery = ref('');
const selectedNotebookId = ref(null);
const selectedArticle = ref(null);
const showMobileSidebar = ref(false); // Mobile state for notebook drawer

// Theme
const { isDark, toggleTheme } = useTheme();

// Computed
const filteredNotes = computed(() => {
    let filtered = props.notes || [];
    
    // Filter by notebook
    if (selectedNotebookId.value) {
        filtered = filtered.filter(note => note.notebook_id === selectedNotebookId.value);
    }
    
    // Filter by search query
    if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase();
        filtered = filtered.filter(note => 
            note.title?.toLowerCase().includes(query) ||
            note.content?.toLowerCase().includes(query)
        );
    }
    
    return filtered;
});

const currentNotebookName = computed(() => {
    if (!selectedNotebookId.value) return 'All Notes';
    return props.notebooks.find(n => n.id === selectedNotebookId.value)?.name || 'All Notes';
});

// Methods
const selectNotebook = (id) => {
    selectedNotebookId.value = id;
    selectedArticle.value = null; // Clear selection when changing notebook
    showMobileSidebar.value = false; // Close mobile drawer on selection
};

const openArticle = (article) => {
    selectedArticle.value = article;
};

const formatRelativeTime = (dateString) => {
    if (!dateString) return '-';
    const date = new Date(dateString);
    const now = new Date();
    const diffHours = Math.floor((now - date) / (1000 * 60 * 60));
    const diffDays = Math.floor(diffHours / 24);
    
    if (diffHours < 1) return 'Just now';
    if (diffHours < 24) return `${diffHours}h ago`;
    if (diffDays === 1) return 'Yesterday';
    if (diffDays < 7) return `${diffDays}d ago`;
    return new Date(dateString).toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
};

const formatDate = (dateString) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
    });
};

const calculateReadTime = (content) => {
    if (!content) return '1 min read';
    const wordsPerMinute = 200;
    const textContent = content.replace(/<[^>]*>?/gm, '');
    const wordCount = textContent.split(/\s+/).length;
    const minutes = Math.ceil(wordCount / wordsPerMinute);
    return `${minutes} min read`;
};

// Noise overlay style
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


</script>

<template>
    <Head title="TaraNote" />
    <FloatingDock />

    <div class="h-screen md:h-screen sm:h-[100dvh] flex text-slate-800 dark:text-white bg-slate-50 dark:bg-[#0F172A] transition-colors duration-200 font-sans overflow-hidden relative">
        
        <!-- Noise Overlay (Subtle) -->
        <div class="fixed inset-0 z-0 opacity-[0.03] dark:opacity-[0.02] pointer-events-none"
             :style="noiseOverlayStyle">
        </div>

        <!-- Ambient Background (Subtle) -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none z-0"
             style="background-image: radial-gradient(circle at 50% 0%, rgba(99, 102, 241, 0.03) 0%, transparent 50%);">
             <div class="absolute top-[-10%] left-[-10%] w-[40%] h-[40%] bg-indigo-500/5 blur-[120px] rounded-full animate-pulse"></div>
             <div class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-rose-500/5 blur-[120px] rounded-full animate-pulse" style="animation-delay: 2s"></div>
        </div>
        
        <!-- App Layout -->
        <div class="relative z-10 flex w-full h-full max-w-[1920px] mx-auto overflow-hidden">
            <!-- 1. LEFT SIDEBAR: Notebooks Navigation -->
            <!-- Desktop: Always visible. Mobile: Drawer/Overlaid -->
            <aside 
                class="fixed inset-y-0 left-0 z-30 w-[260px] flex flex-col bg-slate-50/40 dark:bg-[#0b1121]/60 backdrop-blur-xl border-r border-white/20 dark:border-white/5 transition-transform duration-300 md:relative md:translate-x-0"
                :class="showMobileSidebar ? 'translate-x-0 shadow-2xl' : '-translate-x-full'"
            >
                <!-- Brand -->
                <div class="h-16 flex items-center px-4 gap-3 shrink-0">
                    <div class="size-8 rounded-lg bg-indigo-500 flex items-center justify-center text-white shadow-sm shadow-indigo-500/30">
                        <span class="material-symbols-outlined text-[20px]">edit_note</span>
                    </div>
                    <span class="font-sans font-bold text-base tracking-tight text-slate-800 dark:text-white">TaraNote</span>
                    
                    <!-- Mobile Close -->
                    <button @click="showMobileSidebar = false" class="ml-auto md:hidden p-1 rounded-md hover:bg-black/5 dark:hover:bg-white/10">
                        <span class="material-symbols-outlined text-[20px]">close</span>
                    </button>
                </div>
                
                <!-- Navigation -->
                <nav class="flex-1 px-3 space-y-0.5 overflow-y-auto custom-scrollbar pb-20 md:pb-0">
                     <!-- All Articles -->
                    <button 
                         @click="selectNotebook(null)"
                         class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all text-sm group text-left"
                         :class="selectedNotebookId === null 
                            ? 'bg-white/80 dark:bg-white/10 text-slate-900 dark:text-white font-medium shadow-sm ring-1 ring-black/5 dark:ring-white/5' 
                            : 'text-slate-500 dark:text-slate-400 hover:bg-white/40 dark:hover:bg-white/5 hover:text-slate-900 dark:hover:text-slate-200'"
                    >
                        <span class="material-symbols-outlined text-[20px]">description</span>
                        <span class="flex-1">All Notes</span>
                        <span class="font-mono text-[10px] text-slate-400 opacity-80">{{ notes.length }}</span>
                    </button>


                    <!-- Spaces Header -->
                    <div class="px-3 mt-8 mb-2">
                        <span class="text-[11px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-wider font-sans">Spaces</span>
                    </div>

                    <!-- Notebooks List -->
                    <button 
                        v-for="notebook in notebooks" 
                        :key="notebook.id"
                        @click="selectNotebook(notebook.id)"
                        class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all text-sm group text-left relative"
                        :class="selectedNotebookId === notebook.id 
                            ? 'bg-white/80 dark:bg-white/10 ring-1 ring-indigo-500 text-slate-900 dark:text-white font-medium shadow-sm' 
                            : 'text-slate-500 dark:text-slate-400 hover:bg-white/40 dark:hover:bg-white/5 hover:text-slate-700 dark:hover:text-slate-300'"
                    >
                        <span class="material-symbols-outlined text-[18px] shrink-0" :class="selectedNotebookId === notebook.id ? 'text-indigo-500 dark:text-indigo-400' : 'text-slate-400 opacity-80'">folder</span>
                        <span class="flex-1 truncate">{{ notebook.name }}</span>
                        <span class="font-mono text-[10px] text-slate-400 opacity-60 ml-2">{{ notebook.notes_count || 0 }}</span>
                    </button>
                </nav>
            </aside>

            <!-- 2. MIDDLE COLUMN: Articles List -->
            <div 
                class="flex flex-col bg-white/30 dark:bg-[#0b1121]/40 backdrop-blur-md transition-all duration-300 w-full md:w-[320px] pt-16 md:pt-0 border-r border-white/20 dark:border-white/5"
                :class="{'hidden md:flex': selectedArticle}"
            >
                
                <!-- Header -->
                <div class="h-14 flex items-center justify-between px-5 shrink-0">
                     <!-- Mobile Toggle Sidebar -->
                    <button @click="showMobileSidebar = true" class="md:hidden mr-3 p-1 -ml-2 rounded-md hover:bg-black/5 dark:hover:bg-white/10 text-slate-500">
                        <span class="material-symbols-outlined">menu</span>
                    </button>

                    <h2 class="font-sans font-bold text-[13px] uppercase tracking-wide opacity-80 text-slate-800 dark:text-slate-200 select-none truncate flex-1" :title="currentNotebookName">{{ currentNotebookName }}</h2>
                </div>

                <!-- Search -->
                <div class="px-5 py-2 shrink-0">
                    <div class="relative group">
                        <span class="absolute left-3 top-1/2 -translate-y-1/2 material-symbols-outlined text-slate-400 text-[18px] group-focus-within:text-indigo-500 transition-colors">search</span>
                        <input 
                            type="text" 
                            v-model="searchQuery" 
                            placeholder="Search articles..."
                            class="w-full h-9 bg-white/50 dark:bg-white/5 border-none rounded-lg pl-9 pr-3 text-sm focus:ring-1 focus:ring-indigo-500 transition-all placeholder:text-slate-400 text-slate-700 dark:text-slate-200 shadow-sm"
                        >
                    </div>
                </div>

                <!-- Articles List -->
                <div class="flex-1 overflow-y-auto custom-scrollbar px-3 pb-20 md:pb-0 font-sans">
                    <div v-if="filteredNotes.length === 0" class="flex flex-col items-center justify-center h-full text-center px-6 opacity-60">
                        <span class="material-symbols-outlined text-4xl mb-2 text-slate-300">article</span>
                        <p class="text-xs text-slate-500 dark:text-slate-400">No articles found</p>
                    </div>

                    <div v-else class="space-y-1 py-2">
                        <div 
                            v-for="article in filteredNotes" 
                            :key="article.id"
                            @click="openArticle(article)"
                            class="group p-3 rounded-lg cursor-pointer transition-all duration-200 relative"
                            :class="selectedArticle?.id === article.id 
                                ? 'bg-white/90 dark:bg-white/15 shadow-md ring-1 ring-black/5 dark:ring-white/10 z-10 translate-x-1' 
                                : 'hover:bg-white/60 dark:hover:bg-white/5 hover:shadow-sm'"
                        >
                            <h3 class="font-medium text-[13px] tracking-tight text-slate-800 dark:text-white line-clamp-2 mb-1.5 leading-snug"
                                :class="{'text-indigo-600 dark:text-indigo-400': selectedArticle?.id === article.id}"
                            >
                                {{ article.title || 'Untitled Article' }}
                            </h3>
                            
                            <div class="flex items-center gap-2 mt-2 text-[11px] text-slate-400 font-medium">
                                <span class="font-mono">{{ formatRelativeTime(article.published_at || article.updated_at) }}</span>
                                
                                <span class="ml-auto size-1.5 rounded-full bg-green-400 shadow-sm shadow-green-400/50"></span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- 3. RIGHT PANEL: Article Preview/Reader -->
            <main 
                class="flex-1 flex flex-col bg-slate-50 dark:bg-[#0F172A] relative transition-all duration-300 overflow-hidden w-full md:w-auto"
                :class="selectedArticle ? 'fixed inset-0 z-50 md:static md:z-auto' : 'hidden md:flex'"
            >
                
                <!-- No Article Selected -->
                <div v-if="!selectedArticle" class="flex-1 hidden md:flex items-center justify-center">
                    <div class="text-center py-20 opacity-40">
                        <div class="size-16 rounded-2xl bg-white/50 dark:bg-white/5 flex items-center justify-center mb-4 mx-auto backdrop-blur-sm">
                            <span class="material-symbols-outlined text-3xl opacity-50">article</span>
                        </div>
                        <h2 class="text-sm font-medium text-gray-900 dark:text-white mb-1">Select an article</h2>
                        <p class="text-xs text-gray-500 dark:text-gray-400">Choose from the list to read</p>
                    </div>
                </div>

                <!-- Article Reader -->
                <div v-else class="flex-1 overflow-y-auto custom-scrollbar">
                    
                    <!-- Mobile Back Button -->
                    <div class="sticky top-0 left-0 right-0 z-10 flex items-center p-4 bg-white/80 dark:bg-[#0F172A]/80 backdrop-blur-md border-b border-slate-100 dark:border-white/5 md:hidden">
                        <button @click="selectedArticle = null" class="flex items-center gap-2 text-slate-600 dark:text-slate-300">
                            <span class="material-symbols-outlined">arrow_back</span>
                            <span class="font-bold text-sm">Back to List</span>
                        </button>
                    </div>

                    <article class="w-full max-w-4xl mx-auto px-6 py-8 md:px-12 md:py-10">
                        <!-- Cover Image -->
                        <div v-if="selectedArticle.cover_image" class="w-full h-56 md:h-80 mb-8 rounded-2xl overflow-hidden shadow-lg shadow-black/5 relative">
                            <img :src="selectedArticle.cover_image" :alt="selectedArticle.title" class="w-full h-full object-cover">
                            <div class="absolute inset-0 bg-gradient-to-t from-black/10 to-transparent"></div>
                        </div>

                        <!-- Breadcrumbs -->
                        <nav class="flex flex-wrap gap-2 items-center text-xs mb-8 text-slate-400 font-sans tracking-wide">
                            <Link class="hover:text-slate-600 dark:hover:text-slate-300 transition-colors" href="/">Home</Link>
                            <span class="text-slate-300">/</span>
                            <Link class="hover:text-slate-600 dark:hover:text-slate-300 transition-colors" :href="route('taranote')">TaraNote</Link>
                            <span class="text-slate-300">/</span>
                            <span class="text-indigo-500 font-medium">{{ selectedArticle.notebook?.name || 'General' }}</span>
                        </nav>

                        <!-- Article Header -->
                        <header class="mb-10">
                            <h1 class="font-sans text-3xl sm:text-4xl font-bold tracking-tight text-slate-900 dark:text-white leading-tight mb-6">
                                {{ selectedArticle.title }}
                            </h1>
                            <!-- Author Meta -->
                            <div class="flex items-center justify-between border-b border-slate-200/60 dark:border-white/10 pb-6">
                                <div class="flex items-center gap-3">
                                    <div class="relative">
                                        <div class="size-10 rounded-full bg-cover bg-center ring-2 ring-white dark:ring-white/10 shadow-sm"
                                            :style="{ backgroundImage: `url(${selectedArticle.user?.avatar_url || 'https://via.placeholder.com/150'})` }">
                                        </div>
                                    </div>
                                    <div class="flex flex-col">
                                        <span class="text-sm font-bold text-slate-800 dark:text-slate-200">
                                            {{ selectedArticle.user?.name }}
                                        </span>
                                        <div class="flex items-center gap-2 text-[11px] text-slate-500 font-mono">
                                            <span>{{ formatDate(selectedArticle.published_at || selectedArticle.created_at) }}</span>
                                            <span class="text-slate-300">•</span>
                                            <span>{{ calculateReadTime(selectedArticle.content) }}</span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </header>

                        <!-- Article Body -->
                        <div class="prose prose-slate dark:prose-invert max-w-none prose-headings:font-sans prose-headings:font-bold prose-p:text-slate-600 dark:prose-p:text-slate-300 prose-a:text-indigo-500 hover:prose-a:text-indigo-600 prose-img:rounded-2xl prose-img:shadow-md" v-html="DOMPurify.sanitize(selectedArticle.content || '')"></div>
                    </article>
                </div>
            </main>
        </div>
    </div>
    <img src="/images/logo_tarakreasi.png" class="fixed bottom-6 right-6 h-16 w-auto z-[100] pointer-events-none transition-all duration-300 dark:bg-white dark:rounded-xl dark:p-2 dark:shadow-lg shadow-sm opacity-90 hover:opacity-100">
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
    height: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(156, 163, 175, 0.3);
    border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(156, 163, 175, 0.5);
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.1);
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.2);
}
</style>
