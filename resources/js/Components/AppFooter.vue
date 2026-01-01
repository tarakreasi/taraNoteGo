<script setup>
import { Link } from '@inertiajs/vue3';

const props = defineProps({
    settings: {
        type: Object,
        default: () => ({})
    }
});
</script>

<template>
    <footer class="mt-20 border-t border-slate-200 dark:border-white/5 bg-white/50 dark:bg-black/20 backdrop-blur-lg">
        <div class="max-w-[1200px] mx-auto px-4 sm:px-6 lg:px-8 py-8">
            <div class="flex flex-col md:flex-row justify-between items-center gap-4 text-xs text-slate-500">
                <p>{{ settings?.footer_copyright_text || '© 2026 TaraKreasi Notes. Built with ❤️ and ☕.' }}</p>
                <p class="flex items-center gap-2">
                    <span class="size-2 rounded-full bg-green-500 animate-pulse"></span>
                    {{ settings?.footer_status_text || 'All systems operational' }}
                </p>
            </div>
        </div>
    </footer>
</template>
