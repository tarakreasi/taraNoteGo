<script setup>
import { Link, usePage } from '@inertiajs/vue3';
import { useTheme } from '@/composables/useTheme';
import { computed } from 'vue';

const { isDark, toggleTheme } = useTheme();
const page = usePage();

const user = computed(() => page.props.auth?.user);

const props = defineProps({
    vertical: {
        type: Boolean,
        default: false
    }
});

const navItems = computed(() => [
    {
        name: 'Home',
        icon: 'home',
        route: '/',
        active: page.url === '/' || (page.url.startsWith('/taraNote') && !page.url.includes('/docs')),
    },
    {
        name: 'Dashboard',
        icon: 'dashboard',
        route: '/dashboard',
        active: page.url.startsWith('/dashboard'),
        show: !!user.value,
    },
    {
        name: 'Login',
        icon: 'login',
        route: '/login',
        active: page.url.startsWith('/login'),
        show: !user.value,
    },
    {
        name: 'Docs',
        icon: 'menu_book',
        route: '/docs',
        active: page.url.includes('/docs'),
    }
]);
</script>

<template>
    <div 
        class="fixed z-[100]"
        :class="vertical 
            ? 'left-6 top-1/2 -translate-y-1/2' 
            : 'bottom-6 left-1/2 -translate-x-1/2'"
    >
        <div 
            class="flex items-center p-2 bg-slate-200/50 dark:bg-white/10 backdrop-blur-xl border border-white/20 dark:border-white/10 rounded-2xl shadow-xl shadow-black/10 transition-all hover:bg-white/50 dark:hover:bg-white/20"
             :class="vertical 
                ? 'flex-col gap-3 min-w-[64px]' 
                : 'gap-2 hover:scale-[1.02]'"
        >
            
            <Link 
                v-for="item in navItems.filter(i => i.show !== false)" 
                :key="item.name"
                :href="item.route"
                class="group relative flex items-center justify-center size-12 rounded-xl transition-all duration-300"
                :class="[
                    item.active ? 'bg-primary text-white shadow-lg shadow-primary/30' : 'text-slate-600 dark:text-slate-300 hover:bg-white/50 dark:hover:bg-white/10',
                    vertical ? 'hover:translate-x-1' : 'hover:-translate-y-1'
                ]"
            >
                <span class="material-symbols-outlined text-[24px]">{{ item.icon }}</span>
                
                <!-- Tooltip -->
                <div 
                    class="absolute px-2 py-1 bg-slate-900 dark:bg-white text-white dark:text-slate-900 text-[10px] font-bold rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap pointer-events-none shadow-lg z-50"
                    :class="vertical
                        ? 'left-full ml-3 top-1/2 -translate-y-1/2'
                        : '-top-12 left-1/2 -translate-x-1/2'"
                >
                    {{ item.name }}
                    
                    <!-- Arrow -->
                    <div 
                        class="absolute border-4 border-transparent"
                        :class="vertical
                            ? 'top-1/2 -translate-y-1/2 -left-2 border-r-slate-900 dark:border-r-white'
                            : '-bottom-1 left-1/2 -translate-x-1/2 border-t-slate-900 dark:border-t-white'"
                    ></div>
                </div>
            </Link>

            <div 
                class="bg-slate-200 dark:bg-white/10"
                :class="vertical ? 'w-8 h-[1px] my-1' : 'w-[1px] h-8 mx-1'"
            ></div>

            <!-- Theme Toggle -->
            <button
                @click="toggleTheme"
                class="group relative flex items-center justify-center size-12 rounded-xl transition-all duration-300 text-slate-600 dark:text-slate-300 hover:bg-white/50 dark:hover:bg-white/10"
                :class="vertical ? 'hover:translate-x-1' : 'hover:-translate-y-1'"
            >
                <span v-if="!isDark" class="material-symbols-outlined text-[24px]">light_mode</span>
                <span v-else class="material-symbols-outlined text-[24px]">dark_mode</span>

                <!-- Tooltip -->
                 <div 
                    class="absolute px-2 py-1 bg-slate-900 dark:bg-white text-white dark:text-slate-900 text-[10px] font-bold rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap pointer-events-none shadow-lg z-50"
                    :class="vertical
                        ? 'left-full ml-3 top-1/2 -translate-y-1/2'
                        : '-top-12 left-1/2 -translate-x-1/2'"
                >
                    {{ isDark ? 'Light Mode' : 'Dark Mode' }}
                    
                     <!-- Arrow -->
                    <div 
                        class="absolute border-4 border-transparent"
                        :class="vertical
                            ? 'top-1/2 -translate-y-1/2 -left-2 border-r-slate-900 dark:border-r-white'
                            : '-bottom-1 left-1/2 -translate-x-1/2 border-t-slate-900 dark:border-t-white'"
                    ></div>
                </div>
            </button>

        </div>
    </div>
</template>
