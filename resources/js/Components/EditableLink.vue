<script setup>
import { ref, computed } from 'vue';
import { usePage } from '@inertiajs/vue3';
import { router } from '@inertiajs/vue3';

const props = defineProps({
    settingKeyText: { type: String, required: true },
    settingKeyUrl: { type: String, required: true },
    initialText: { type: String, default: 'Link' },
    initialUrl: { type: String, default: '#' },
    newTab: { type: Boolean, default: false }
});

const page = usePage();
const user = computed(() => page.props.auth?.user);

const text = ref(props.initialText);
const url = ref(props.initialUrl);
const isEditing = ref(false);
const status = ref('idle'); // idle, saving, saved, error

const openEdit = (e) => {
    if (!user.value) return; 
    e.preventDefault();
    isEditing.value = true;
};

const save = async () => {
    status.value = 'saving';
    try {
        await window.axios.post('/settings/update', {
            settings: {
                [props.settingKeyText]: text.value,
                [props.settingKeyUrl]: url.value
            }
        });
        status.value = 'saved';
        setTimeout(() => {
            if (status.value === 'saved') status.value = 'idle';
            isEditing.value = false;
        }, 500);
    } catch (e) {
        status.value = 'error';
        console.error(e);
    }
};

const cancel = () => {
    isEditing.value = false;
    status.value = 'idle';
};
</script>

<template>
    <a 
        :href="url" 
        :target="newTab ? '_blank' : '_self'" 
        class="relative inline-block group/editable" 
        @click="openEdit"
    >
        <span>{{ text }}</span>

        <!-- Edit Indicator (Pencil) -->
        <span 
            v-if="user" 
            class="absolute -top-3 -right-3 text-[10px] bg-white dark:bg-slate-800 text-slate-400 p-0.5 rounded-full shadow border border-slate-200 dark:border-slate-700 opacity-0 group-hover/editable:opacity-100 transition-opacity"
        >
             <span class="material-symbols-outlined text-[12px] block">edit</span>
        </span>

        <!-- Modal/Popover -->
        <Teleport to="body">
            <div v-if="isEditing" class="fixed inset-0 z-[999] flex items-center justify-center bg-black/50" @click.self="cancel">
                <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow-2xl w-[90%] max-w-sm transition-all animate-fadeIn" @click.stop>
                    <div class="flex justify-between items-center mb-4">
                        <h3 class="text-sm font-bold uppercase text-slate-400">Edit Link</h3>
                        <!-- Status in Modal -->
                        <span v-if="status === 'saving'" class="text-blue-500 animate-spin material-symbols-outlined text-sm">sync</span>
                        <span v-if="status === 'saved'" class="text-green-500 material-symbols-outlined text-sm">check_circle</span>
                        <span v-if="status === 'error'" class="text-red-500 material-symbols-outlined text-sm">error</span>
                    </div>
                    
                    <div class="space-y-4">
                        <div>
                            <label class="block text-xs font-bold mb-1 text-slate-600 dark:text-slate-300">Text</label>
                            <input v-model="text" type="text" class="w-full text-sm rounded-lg border-slate-300 dark:bg-slate-700 dark:border-slate-600 dark:text-white" />
                        </div>
                        <div>
                            <label class="block text-xs font-bold mb-1 text-slate-600 dark:text-slate-300">URL</label>
                            <input v-model="url" type="text" class="w-full text-sm rounded-lg border-slate-300 dark:bg-slate-700 dark:border-slate-600 dark:text-white" />
                        </div>
                        <div class="flex justify-end gap-2 pt-2">
                            <button @click="cancel" class="px-3 py-1.5 text-xs font-bold text-slate-500 hover:bg-slate-100 rounded">Cancel</button>
                            <button @click="save" :disabled="status === 'saving'" class="px-3 py-1.5 text-xs font-bold bg-blue-600 text-white hover:bg-blue-700 rounded disabled:opacity-50">
                                {{ status === 'saving' ? 'Saving...' : 'Save' }}
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </Teleport>
    </a>
</template>
