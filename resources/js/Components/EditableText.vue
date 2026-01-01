<script setup>
import { ref, computed } from 'vue';
import { usePage } from '@inertiajs/vue3';

const props = defineProps({
    settingKey: {
        type: String,
        required: true
    },
    initialValue: {
        type: String,
        default: ''
    },
    tag: {
        type: String,
        default: 'span'
    }
});

const page = usePage();
const user = computed(() => page.props.auth?.user);
const content = ref(props.initialValue);
const status = ref('idle'); // idle, saving, saved, error

const onBlur = async (e) => {
    if (!user.value) return;
    
    // Sanitize input: simply taking innerText here.
    // Ideally we might want to trim or handle newlines.
    const newValue = e.target.innerText;
    
    if (newValue === content.value) return; // No change

    // Optimistic UI update
    const oldValue = content.value;
    content.value = newValue; 
    status.value = 'saving';

    try {
        await window.axios.post('/settings/update', {
            settings: {
                [props.settingKey]: newValue
            }
        });
        status.value = 'saved';
        setTimeout(() => {
            if (status.value === 'saved') status.value = 'idle';
        }, 2000);
    } catch (error) {
        console.error("Failed to save setting:", error);
        status.value = 'error';
        // Optional: revert? For now, keep user edit so they don't lose work, but show error.
    }
};
</script>

<template>
    <span class="relative group/editable" :class="$attrs.class">
        <component
            :is="tag"
            :contenteditable="!!user"
            @blur="onBlur"
            class="outline-none transition-all duration-200 decoration-blue-300 underline-offset-4 caret-blue-500 dark:caret-white"
            :class="{
                'hover:bg-blue-50/50 dark:hover:bg-blue-900/20 cursor-text rounded px-1 -mx-1 border border-transparent hover:border-dashed hover:border-blue-300 dark:hover:border-blue-600': !!user,
                'opacity-50': status === 'saving',
                'bg-red-50 dark:bg-red-900/20': status === 'error'
            }"
            v-text="content"
        />

        <!-- Hover Indicator (Pencil) -->
        <span 
            v-if="user && status === 'idle'" 
            class="absolute -top-3 -right-3 text-[10px] bg-white dark:bg-slate-800 text-slate-400 p-0.5 rounded-full shadow border border-slate-200 dark:border-slate-700 opacity-0 group-hover/editable:opacity-100 transition-opacity pointer-events-none z-50 flex items-center justify-center w-5 h-5 text-center leading-none"
        >
            <span class="material-symbols-outlined text-[12px]">edit</span>
        </span>

        <!-- Status Indicators (Inline) -->
        <span v-if="status === 'saving'" class="absolute -top-3 -right-3 text-blue-500 animate-spin z-50 bg-white dark:bg-slate-800 rounded-full p-0.5 shadow flex items-center justify-center w-5 h-5">
            <span class="material-symbols-outlined text-[14px]">sync</span>
        </span>
        <span v-if="status === 'saved'" class="absolute -top-3 -right-3 text-green-500 z-50 bg-white dark:bg-slate-800 rounded-full p-0.5 shadow flex items-center justify-center w-5 h-5">
             <span class="material-symbols-outlined text-[14px]">check</span>
        </span>
        <span v-if="status === 'error'" class="absolute -top-3 -right-3 text-red-500 z-50 bg-white dark:bg-slate-800 rounded-full p-0.5 shadow flex items-center justify-center w-5 h-5" title="Failed to save">
             <span class="material-symbols-outlined text-[14px]">error</span>
        </span>
    </span>
</template>
