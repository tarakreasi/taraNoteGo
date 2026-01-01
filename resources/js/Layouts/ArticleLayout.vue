<script setup>
import { Link, router } from '@inertiajs/vue3';
import { ref } from 'vue';
import FloatingDock from '@/Components/FloatingDock.vue';
import AppFooter from '@/Components/AppFooter.vue';

// Simple debounce utility since lodash might not be present
const debounce = (fn, delay) => {
    let timeoutId;
    return (...args) => {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => fn(...args), delay);
    };
};

const search = ref('');

const handleSearch = debounce(() => {
    router.get(route('articles.index'), { 
        search: search.value 
    }, {
        preserveState: true,
        replace: true,
        preserveScroll: true
    });
}, 500);

</script>

<template>
    <div class="min-h-screen bg-white dark:bg-[#0F172A] text-[#333] dark:text-[#eee] font-sans flex flex-col">
        
        <!-- HEADER (Persistent) -->
        <header class="sticky top-0 z-40 bg-white/95 dark:bg-[#0F172A]/95 backdrop-blur-sm border-b border-slate-100 dark:border-slate-800 h-16 flex items-center px-4 md:px-6">
            <div class="max-w-[1200px] w-full mx-auto flex items-center justify-between">
                <!-- Logo -->
                <Link href="/" class="flex items-center gap-1 group">
                    <span class="font-extrabold text-2xl tracking-tight text-[#1a1a1a] dark:text-white group-hover:opacity-80 transition-opacity">
                        taraNote<span class="text-emerald-500">.</span>
                    </span>
                </Link>

                <!-- Search (Centered / Right) -->
                <div class="hidden md:flex flex-1 max-w-md mx-auto">
                    <div class="relative w-full group">
                        <span class="absolute left-3 top-1/2 -translate-y-1/2 material-symbols-outlined text-slate-400 text-[20px] cursor-pointer hover:text-emerald-500 group-focus-within:text-slate-600 transition-colors" @click="handleSearch">search</span>
                        <input 
                            v-model="search"
                            @input="handleSearch"
                            type="text" 
                            placeholder="Search articles..." 
                            class="w-full h-10 pl-10 pr-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-full focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 outline-none transition-all placeholder:text-slate-400 text-sm"
                        >
                    </div>
                </div>
            </div>
        </header>

        <!-- PAGE CONTENT CONTAINER -->
        <slot />

        <AppFooter class="mt-auto border-t border-slate-100 dark:border-slate-800" />
        
        <!-- FLOATING DOCK (Persistent) -->
        <FloatingDock />
        
    </div>
</template>
