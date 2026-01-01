<script setup>
import { Head, usePage } from '@inertiajs/vue3';
import { ref, computed } from "vue";

const props = defineProps({
    settings: {
        type: Object,
        default: () => ({})
    }
});

const getSetting = (key, fallback = '') => props.settings[key] || fallback;
</script>

<template>
    <Head>
        <title>Test - Tri Wantoro</title>
    </Head>

    <div class="min-h-screen bg-white dark:bg-slate-900 flex items-center justify-center">
        <div class="text-center">
            <h1 class="text-4xl font-bold text-slate-900 dark:text-white mb-4">
                Emergency Test Page
            </h1>
            <p class="text-lg text-slate-600 dark:text-slate-300">
                If you can see this, the page is rendering correctly.
            </p>
            <div class="mt-8 p-4 bg-blue-50 dark:bg-blue-900/20 rounded-lg">
                <p class="text-sm font-mono">
                    Settings loaded: {{ Object.keys(props.settings).length }} keys
                </p>
            </div>
        </div>
    </div>
</template>
