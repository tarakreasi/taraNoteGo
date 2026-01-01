<script setup>
import { Head, Link } from '@inertiajs/vue3';
import { ref, computed, onMounted } from 'vue';
import { useTheme } from '@/composables/useTheme';
import FloatingDock from '@/Components/FloatingDock.vue';

const props = defineProps({
    currentPath: String,
    displayName: String,
    content: String,
});

// Available docs list (hardcoded for now)
const allDocs = ref([
    { name: 'INDEX', title: 'Documentation Index', category: 'Getting Started' },
    { name: 'README', title: 'Project Overview', category: 'Getting Started' },
    { name: 'QUICKSTART', title: 'Quick Start Guide', category: 'Getting Started' },
    { name: 'FAQ', title: 'Frequently Asked Questions', category: 'Help' },
    { name: 'USER_GUIDE', title: 'User Guide', category: 'Help' },
    { name: 'TROUBLESHOOTING', title: 'Troubleshooting', category: 'Help' },
    { name: 'API_REFERENCE', title: 'API Reference', category: 'Development' },
    { name: 'ENGINEERING_STANDARDS', title: 'Engineering Standards', category: 'Development' },
    { name: 'FEATURE_FLAGS', title: 'Feature Flags', category: 'Development' },
    { name: 'TESTING', title: 'Testing Guide', category: 'Development' },
    { name: 'CONTRIBUTING', title: 'Contributing Guide', category: 'Development' },
    { name: 'UI/README', title: 'UI Gallery / Prototypes', category: 'Design' },
    { name: 'UI/story_v6.5', title: 'The Journey to Zen (Story)', category: 'Design' },
    { name: 'UI_UX_OVERVIEW', title: 'UI/UX Design Overview', category: 'Design' },
    { name: 'DESIGN_TOKENS', title: 'Design Tokens', category: 'Design' },
    { name: 'COMPONENT_LIBRARY', title: 'Component Library', category: 'Design' },
    { name: 'ACCESSIBILITY', title: 'Accessibility Guidelines', category: 'Design' },
    { name: 'PRD/README', title: 'PRD Overview & Archive', category: 'Product' },
    { name: 'PRD/PRD_V6_MASTER', title: 'PRD V6.5 Master', category: 'Product' },
    { name: 'PRD/USER_PERSONA', title: 'User Personas', category: 'Product' },
    { name: 'ROADMAP', title: 'Development Roadmap', category: 'Project Info' },
    { name: 'CHANGELOG', title: 'Changelog', category: 'Project Info' },
    { name: 'SECURITY', title: 'Security Policy', category: 'Legal' },
    { name: 'PRIVACY_POLICY', title: 'Privacy Policy', category: 'Legal' },
    { name: 'ATTRIBUTIONS', title: 'Asset Attributions', category: 'Legal' },
    { name: 'DEPLOYMENT', title: 'Deployment Guide', category: 'Operations' },
    { name: 'NGINX_CONFIG', title: 'Nginx Config Reference', category: 'Operations' },
    { name: 'BACKUP_RESTORE', title: 'Backup & Restore Strategy', category: 'Operations' },
]);

// State
const searchQuery = ref('');
const selectedCategory = ref(null);
const renderedHtml = computed(() => markdownToHtml(props.content));

// Theme
const { isDark, toggleTheme } = useTheme();

// Computed
const categories = computed(() => {
    return [...new Set(allDocs.value.map(doc => doc.category))];
});

const filteredDocs = computed(() => {
    let filtered = allDocs.value;
    
    // Filter by category
    if (selectedCategory.value) {
        filtered = filtered.filter(doc => doc.category === selectedCategory.value);
    }
    
    // Filter by search query
    if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase();
        filtered = filtered.filter(doc => 
            doc.title.toLowerCase().includes(query) ||
            doc.name.toLowerCase().includes(query)
        );
    }
    
    return filtered;
});

const currentCategoryName = computed(() => {
    return selectedCategory.value || 'All Documentation';
});

// Methods
const selectCategory = (category) => {
    selectedCategory.value = category;
    
    // Update URL without page reload
    const url = new URL(window.location);
    if (category) {
        url.searchParams.set('category', category);
    } else {
        url.searchParams.delete('category');
    }
    window.history.pushState({}, '', url);
};

// Simple markdown to HTML converter
const markdownToHtml = (markdown) => {
    if (!markdown) return '';
    
    let html = markdown;
    
    // Headers
    html = html.replace(/^### (.*$)/gim, '<h3>$1</h3>');
    html = html.replace(/^## (.*$)/gim, '<h2>$1</h2>');
    html = html.replace(/^# (.*$)/gim, '<h1>$1</h1>');
    
    // Bold
    html = html.replace(/\*\*(.*?)\*\*/gim, '<strong>$1</strong>');
    
    // Italic
    html = html.replace(/\*(.*?)\*/gim, '<em>$1</em>');
    
    // Links - smart handling for internal vs external
    html = html.replace(/\[([^\]]+)\]\(([^)]+)\)/gim, (match, text, href) => {
        const isExternal = href.startsWith('http') || href.startsWith('//');
        const target = isExternal ? 'target="_blank"' : '';
        const classes = isExternal ? 'text-primary hover:underline external-link' : 'text-primary hover:underline internal-link';
        return `<a href="${href}" ${target} class="${classes}">${text}</a>`;
    });
    
    // Code blocks - properly escape HTML
    html = html.replace(/```([a-z]*)\n([\s\S]*?)```/gim, (match, lang, code) => {
        const escapedCode = code
            .replace(/&/g, '&amp;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;')
            .replace(/"/g, '&quot;')
            .replace(/'/g, '&#039;');
        return `<pre><code class="language-${lang}">${escapedCode}</code></pre>`;
    });
    
    // Inline code - properly escape HTML
    html = html.replace(/`([^`]+)`/gim, (match, code) => {
        const escapedCode = code
            .replace(/&/g, '&amp;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;')
            .replace(/"/g, '&quot;')
            .replace(/'/g, '&#039;');
        return `<code class="bg-gray-100 dark:bg-gray-800 px-1.5 py-0.5 rounded text-sm">${escapedCode}</code>`;
    });
    
    // Lists
    html = html.replace(/^\* (.*$)/gim, '<li>$1</li>');
    html = html.replace(/(<li>.*<\/li>)/s, '<ul class="list-disc ml-6">$1</ul>');
    
    // Blockquotes
    html = html.replace(/^> (.*$)/gim, '<blockquote class="border-l-4 border-primary pl-4 italic my-4">$1</blockquote>');
    
    // Line breaks
    html = html.replace(/\n\n/g, '</p><p>');
    html = '<p>' + html + '</p>';
    
    // Horizontal rules
    html = html.replace(/^---$/gim, '<hr class="my-8 border-gray-200 dark:border-gray-700">');
    
    return html;
};

// Handle internal link clicks
const handleContentClick = (event) => {
    const link = event.target.closest('a.internal-link');
    if (link) {
        event.preventDefault();
        const href = link.getAttribute('href');
        
        // Resolve relative paths if needed, but for now simple checks
        // We rely on the browser to resolve the href relative to current URL, 
        // but since we preventDefault, we need to construct the absolute URL manually if it's relative
        // taking into account the current Inertia page URL.
        
        // Actually, simplest way for Inertia is:
        const url = new URL(href, window.location.href);
        
        // Only intercept if it's same origin
        if (url.origin === window.location.origin) {
            // Check if it's a direct file link (HTML, images, etc.)
            const staticExtensions = ['.html', '.png', '.jpg', '.jpeg', '.gif', '.svg', '.css', '.js'];
            if (staticExtensions.some(ext => href.toLowerCase().endsWith(ext))) {
                // Force open in new tab manually to be 100% reliable
                window.open(url.href, '_blank');
                return;
            }

            import('@inertiajs/vue3').then(({ router }) => {
                router.visit(url.href);
            });
        } else {
             window.open(url.href, '_blank');
        }
    }
};

onMounted(() => {
    // Check URL for initial category
    const params = new URLSearchParams(window.location.search);
    const category = params.get('category');
    if (category && categories.value.includes(category)) {
        selectedCategory.value = category;
    }
});
// Noise overlay style
const noiseOverlayStyle = "background-image: url('data:image/svg+xml,%3Csvg viewBox=\\'0 0 200 200\\' xmlns=\\'http://www.w3.org/2000/svg\\'%3E%3Cfilter id=\\'noiseFilter\\'%3E%3CfeTurbulence type=\\'fractalNoise\\' baseFrequency=\\'0.65\\' numOctaves=\\'3\\' stitchTiles=\\'stitch\\'/%3E%3C/filter%3E%3Crect width=\\'100%25\\' height=\\'100%25\\' filter=\\'url(%23noiseFilter)\\'/%3E%3C/svg%3E');";

</script>

<template>
    <div>
        <Head :title="`${displayName} - TaraNote Documentation`" />
        <FloatingDock />

        <div class="h-screen flex text-slate-800 dark:text-white bg-slate-50 dark:bg-[#0F172A] transition-colors duration-200 font-sans overflow-hidden relative">
            
            <!-- Noise Overlay (Subtle) -->
            <div class="fixed inset-0 z-0 opacity-[0.03] dark:opacity-[0.02] pointer-events-none"
                 :style="noiseOverlayStyle">
            </div>

            <!-- Ambient Background (Subtle) -->
            <div class="fixed inset-0 overflow-hidden pointer-events-none z-0"
                 style="background-image: radial-gradient(circle at 50% 0%, rgba(99, 102, 241, 0.03) 0%, transparent 50%);">
                 <!-- Primary Blob (Top Left) -->
                 <div class="absolute top-[-10%] left-[-10%] w-[40%] h-[40%] bg-indigo-500/5 blur-[120px] rounded-full animate-pulse"></div>
                 <!-- Accent Blob (Bottom Right) -->
                 <div class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-rose-500/5 blur-[120px] rounded-full animate-pulse" style="animation-delay: 2s"></div>
            </div>

            <!-- App Layout -->
            <div class="relative z-10 flex w-full h-full max-w-[1920px] mx-auto overflow-hidden">

                <!-- 1. LEFT SIDEBAR: Categories Navigation -->
                <aside class="w-[260px] flex flex-col bg-slate-50/50 dark:bg-[#0b1121] transition-all duration-300">
                    <!-- Brand -->
                    <div class="h-16 flex items-center px-4 gap-3 shrink-0">
                        <div class="size-8 rounded-lg bg-indigo-500 flex items-center justify-center text-white shadow-sm">
                            <span class="material-symbols-outlined text-[20px]">description</span>
                        </div>
                        <span class="font-sans font-bold text-base tracking-tight text-slate-800 dark:text-white">Documentation</span>
                    </div>

                    <!-- Navigation -->
                    <nav class="flex-1 px-3 space-y-0.5 overflow-y-auto custom-scrollbar">
                        <!-- All Docs -->
                        <button 
                             @click="selectCategory(null)"
                             class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all text-sm group text-left"
                             :class="selectedCategory === null 
                                ? 'bg-white dark:bg-white/5 text-slate-900 dark:text-white font-medium shadow-sm ring-1 ring-black/5 dark:ring-white/5' 
                                : 'text-slate-500 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-white/5 hover:text-slate-900 dark:hover:text-slate-200'"
                        >
                            <span class="material-symbols-outlined text-[20px]">library_books</span>
                            <span class="flex-1">All Documentation</span>
                            <span class="font-mono text-[10px] text-slate-400 opacity-80">{{ allDocs.length }}</span>
                        </button>

                        <!-- Spaces Header -->
                        <div class="px-3 mt-8 mb-2">
                            <span class="text-[11px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-wider font-sans">Categories</span>
                        </div>

                        <!-- Categories List -->
                        <button 
                            v-for="category in categories" 
                            :key="category"
                            @click="selectCategory(category)"
                            class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all text-sm group text-left relative"
                            :class="selectedCategory === category 
                                 ? 'bg-white dark:bg-white/10 ring-1 ring-indigo-500 text-slate-900 dark:text-white font-medium shadow-sm' 
                                : 'text-slate-500 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-white/5 hover:text-slate-700 dark:hover:text-slate-300'"
                        >
                            <span class="material-symbols-outlined text-[18px] shrink-0" :class="selectedCategory === category ? 'text-indigo-500 dark:text-indigo-400' : 'text-slate-400 opacity-80'">folder</span>
                            <span class="flex-1 truncate">{{ category }}</span>
                        </button>
                    </nav>
                </aside>

                <!-- 2. MIDDLE COLUMN: Docs List -->
                <div class="w-[320px] flex flex-col bg-white/50 dark:bg-[#0b1121]/80 transition-all duration-300">
                    
                    <!-- Header -->
                    <div class="h-14 flex items-center justify-between px-5 shrink-0">
                        <h2 class="font-sans font-bold text-[13px] uppercase tracking-wide opacity-80 text-slate-800 dark:text-slate-200 select-none truncate flex-1" :title="currentCategoryName">{{ currentCategoryName }}</h2>
                    </div>

                    <!-- Search -->
                    <div class="px-5 py-2 shrink-0">
                        <div class="relative group">
                            <span class="absolute left-3 top-1/2 -translate-y-1/2 material-symbols-outlined text-slate-400 text-[18px] group-focus-within:text-indigo-500 transition-colors">search</span>
                            <input 
                                type="text" 
                                name="search_docs"
                                id="search_docs"
                                aria-label="Search documentation"
                                v-model="searchQuery" 
                                placeholder="Search docs..."
                                class="w-full h-9 bg-slate-100 dark:bg-white/5 border-none rounded-lg pl-9 pr-3 text-sm focus:ring-1 focus:ring-indigo-500 transition-all placeholder:text-slate-400 text-slate-700 dark:text-slate-200"
                            >
                        </div>
                    </div>

                    <!-- Docs List -->
                    <div class="flex-1 overflow-y-auto custom-scrollbar px-3 font-sans">
                        <div v-if="filteredDocs.length === 0" class="flex flex-col items-center justify-center h-full text-center px-6 opacity-60">
                            <span class="material-symbols-outlined text-4xl mb-2 text-slate-300">description</span>
                            <p class="text-xs text-slate-500 dark:text-slate-400">No documentation found</p>
                        </div>

                        <div v-else class="space-y-1 py-2">
                            <Link 
                                v-for="doc in filteredDocs" 
                                :key="doc.name"
                                :href="selectedCategory ? route('docs.show', { path: doc.name, category: selectedCategory }) : route('docs.show', { path: doc.name })"
                                class="block group p-3 rounded-lg cursor-pointer transition-all duration-200"
                                :class="currentPath === doc.name 
                                    ? 'bg-white dark:bg-white/10 shadow-md ring-1 ring-black/5 dark:ring-white/5 z-10' 
                                    : 'hover:bg-white/60 dark:hover:bg-white/5 hover:shadow-sm'"
                            >
                                <h3 class="font-medium text-[13px] tracking-tight text-slate-800 dark:text-white line-clamp-2 mb-1.5 leading-snug"
                                    :class="{'text-indigo-600 dark:text-indigo-400': currentPath === doc.name}"
                                >
                                    {{ doc.title }}
                                </h3>
                                
                                <div class="flex items-center gap-2 mt-2 text-[11px] text-slate-400 font-medium">
                                    <span class="material-symbols-outlined text-[14px] opacity-70">folder_open</span> 
                                    {{ doc.category }}
                                </div>
                            </Link>
                        </div>
                    </div>
                </div>

                <!-- 3. RIGHT PANEL: Doc Reader -->
                <main class="flex-1 flex flex-col bg-slate-50 dark:bg-[#0F172A] relative transition-all duration-300 overflow-hidden border-l border-slate-100 dark:border-white/5">
                    
                    <!-- Doc Reader -->
                    <div class="flex-1 overflow-y-auto custom-scrollbar">
                        <article class="w-full max-w-7xl px-8 py-10 md:px-12 pb-40 xl:pr-[280px]">
                            <!-- Breadcrumbs -->
                            <nav class="flex flex-wrap gap-2 items-center text-xs mb-8 text-slate-400 font-sans tracking-wide">
                                <Link class="hover:text-slate-600 dark:hover:text-slate-300 transition-colors" :href="route('home')">Home</Link>
                                <span class="text-slate-300">/</span>
                                <Link class="hover:text-slate-600 dark:hover:text-slate-300 transition-colors" :href="route('docs.index')">Docs</Link>
                                <span class="text-slate-300">/</span>
                                <span class="text-indigo-500 font-medium">{{ displayName }}</span>
                            </nav>

                            <!-- Doc Header -->
                            <div class="mb-10 border-b border-slate-100 dark:border-white/5 pb-8">
                                <div class="flex items-center gap-2 mb-4">
                                    <span class="inline-flex items-center px-2 py-0.5 rounded text-[10px] font-bold uppercase tracking-wider bg-indigo-50 text-indigo-600 dark:bg-indigo-500/10 dark:text-indigo-400">
                                        Documentation
                                    </span>
                                </div>
                                <h1 class="font-sans text-3xl sm:text-4xl font-bold tracking-tight text-slate-900 dark:text-white leading-tight mb-2">
                                    {{ displayName }}
                                </h1>
                            </div>

                            <!-- Doc Content -->
                            <div class="prose prose-slate dark:prose-invert max-w-none prose-headings:font-sans prose-headings:font-bold prose-p:text-slate-600 dark:prose-p:text-slate-300 prose-a:text-indigo-500 hover:prose-a:text-indigo-600" v-html="renderedHtml" @click="handleContentClick"></div>
                        </article>
                    </div>
                </main>
            </div>
        </div>
        <!-- Watermark -->
        <!-- Watermark -->
        <img src="/images/logo_tarakreasi.png" class="fixed bottom-6 right-6 h-16 w-auto z-[100] pointer-events-none transition-all duration-300 dark:bg-white dark:rounded-xl dark:p-2 dark:shadow-lg shadow-sm opacity-90 hover:opacity-100">
    </div>
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
    background: rgba(148, 163, 184, 0.3); /* slate-400/30 */
    border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(148, 163, 184, 0.5); /* slate-400/50 */
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.1);
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.2);
}

/* Prose styling for markdown content */
.prose {
    line-height: 1.7;
}

.prose h1 {
    font-size: 2.25em;
    margin-top: 0;
    margin-bottom: 0.8888889em;
    line-height: 1.1111111;
    font-weight: 800;
}

.prose h2 {
    font-size: 1.5em;
    margin-top: 2em;
    margin-bottom: 1em;
    line-height: 1.3333333;
    font-weight: 700;
}

.prose h3 {
    font-size: 1.25em;
    margin-top: 1.6em;
    margin-bottom: 0.6em;
    line-height: 1.6;
    font-weight: 600;
}

.prose p {
    margin-top: 1.25em;
    margin-bottom: 1.25em;
}

.prose code {
    color: #4f46e5; /* indigo-600 */
    font-weight: 600;
    font-size: 0.875em;
}

.dark .prose code {
    color: #818cf8; /* indigo-400 */
}

.prose pre {
    background: #f1f5f9; /* slate-100 */
    padding: 1.25em;
    border-radius: 0.5rem;
    overflow-x: auto;
    margin: 1.7142857em 0;
}

.dark .prose pre {
    background: rgba(255, 255, 255, 0.05);
}

.prose pre code {
    color: inherit;
    font-weight: 400;
}
</style>
