<script setup>
import { ref, nextTick, computed } from 'vue';
import { usePage, router } from '@inertiajs/vue3';

// User Data
const page = usePage();
const user = computed(() => page.props.auth?.user);

const handleLogout = () => {
    router.post(route('logout'));
};

const props = defineProps({
    notebooks: {
        type: Array,
        default: () => []
    },
    selectedNotebookId: {
        type: [Number, String],
        default: null
    },
    viewMode: {
        type: String,
        default: 'EDITOR' // 'EDITOR' or 'SETTINGS'
    },
    notebooksLoading: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits([
    'update:viewMode',
    'select-notebook',
    'create-notebook',
    'rename-notebook', // { id, type: 'NAME'|'SLUG', value }
    'delete-notebook'  // id
]);

// --- LOCAL STATE ---

// Creating Space
const isCreatingSpace = ref(false);
const newSpaceName = ref('');
const newSpaceInput = ref(null);

const startCreatingSpace = async () => {
    isCreatingSpace.value = true;
    newSpaceName.value = '';
    await nextTick();
    if (newSpaceInput.value) newSpaceInput.value.focus();
};

const cancelCreatingSpace = () => {
    isCreatingSpace.value = false;
    newSpaceName.value = '';
};

const saveNewSpace = () => {
    if (!newSpaceName.value.trim()) {
        cancelCreatingSpace();
        return;
    }
    emit('create-notebook', newSpaceName.value);
    cancelCreatingSpace();
};

// Inline Editing for Notebooks
const editingState = ref({
    id: null,
    type: null, // 'NOTEBOOK_NAME', 'NOTEBOOK_SLUG'
    value: '',
    originalValue: ''
});

const startEditing = (id, type, currentValue) => {
    editingState.value = {
        id,
        type,
        value: currentValue,
        originalValue: currentValue
    };
};

const cancelEditing = () => {
    editingState.value = { id: null, type: null, value: '', originalValue: '' };
};

const saveEditing = () => {
    const { id, type, value, originalValue } = editingState.value;
    
    if (value === originalValue) {
        cancelEditing();
        return;
    }

    if (!value.trim()) {
        alert("Value cannot be empty");
        return;
    }

    emit('rename-notebook', { id, type, value });
    cancelEditing();
};

// Delete Confirmation
const deleteConfirmationId = ref(null);

const requestDelete = (id) => {
    deleteConfirmationId.value = id;
};

const confirmDelete = (id) => {
    emit('delete-notebook', id);
    deleteConfirmationId.value = null;
};

// Validations for displaying delete option
const canDelete = (notebook) => {
    // Only allow if empty? The original code had: (!notebook.notes_count || notebook.notes_count === 0)
    return (!notebook.notes_count || notebook.notes_count === 0);
};

// Directives
const vFocus = {
    mounted: (el) => el.focus()
};

// Navigation Helpers
const onSwitchToSettings = () => {
    emit('update:viewMode', 'SETTINGS');
    emit('select-notebook', null); // Optional: clear selection
};

const onSelectAllNotes = () => {
    emit('update:viewMode', 'EDITOR');
    emit('select-notebook', null);
};

const onSelectNotebook = (id) => {
    emit('update:viewMode', 'EDITOR');
    emit('select-notebook', id);
};

defineExpose({
    startCreatingSpace
});
</script>

<template>
    <aside class="w-[260px] flex flex-col bg-slate-50/50 dark:bg-[#0b1121]/50 transition-all duration-300">
        <!-- Brand -->
        <div class="h-16 flex items-center px-4 gap-3 shrink-0">
            <div class="size-8 rounded-lg bg-slate-900 dark:bg-white flex items-center justify-center text-white dark:text-slate-900 shadow-sm">
                <span class="material-symbols-outlined text-[20px]">dashboard</span>
            </div>
            <span class="font-sans font-medium text-sm tracking-tight flex-1 text-slate-700 dark:text-slate-200">TaraNote</span>
            
            <!-- Settings Trigger -->
            <button 
                @click="onSwitchToSettings"
                class="size-8 rounded-md hover:bg-slate-200 dark:hover:bg-white/10 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors flex items-center justify-center duration-200"
                title="Settings"
                :class="viewMode === 'SETTINGS' ? 'bg-slate-200 dark:bg-white/10 text-slate-900 dark:text-white' : ''"
            >
                <span class="material-symbols-outlined text-[18px]">settings</span>
            </button>
        </div>

        <!-- Menu -->
        <nav class="flex-1 px-3 py-2 space-y-1 overflow-y-auto custom-scrollbar">
            
            <!-- All Notes Item -->
            <button 
                 @click="onSelectAllNotes"
                 class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all duration-200 text-sm group text-left"
                 :class="(viewMode === 'EDITOR' && selectedNotebookId === null) 
                    ? 'bg-white dark:bg-white/10 text-slate-900 dark:text-white font-medium shadow-sm' 
                    : 'text-slate-500 dark:text-slate-400 hover:bg-white/60 dark:hover:bg-white/5 hover:text-slate-900 dark:hover:text-slate-200'"
            >
                <span class="material-symbols-outlined text-[20px]" :class="(viewMode === 'EDITOR' && selectedNotebookId === null) ? 'text-indigo-500 dark:text-indigo-400' : 'text-slate-400 group-hover:text-slate-500'">description</span>
                <span class="flex-1">All Notes</span>
                <slot name="count"></slot>
            </button>
            
            <!-- New Space Button (Refined) -->
            <button 
                @click="startCreatingSpace"
                class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all duration-200 text-sm group text-left mt-1 text-slate-500 dark:text-slate-400 hover:bg-white/60 dark:hover:bg-white/5 hover:text-slate-900 dark:hover:text-slate-200"
            >
                 <span class="material-symbols-outlined text-[20px] text-slate-400 group-hover:text-indigo-500 transition-colors">add_circle</span>
                <span class="flex-1">New Space</span>
            </button>
            
            <!-- Spaces Header -->
            <div class="px-3 mt-8 mb-2 flex items-center justify-between group">
                <span class="text-[11px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-wider font-sans">Spaces</span>
            </div>

            <!-- Inline Create Input -->
            <div v-if="isCreatingSpace" class="px-2 mb-1">
                <div class="flex items-center gap-2 px-3 py-2 rounded-lg bg-white dark:bg-white/5 ring-1 ring-indigo-500/30 shadow-sm">
                    <span class="material-symbols-outlined text-[18px] text-indigo-500">folder_open</span>
                    <input 
                        ref="newSpaceInput"
                        v-model="newSpaceName"
                        @keydown.enter="saveNewSpace"
                        @keydown.esc="cancelCreatingSpace"
                        @blur="saveNewSpace"
                        type="text" 
                        class="flex-1 bg-transparent border-none text-sm p-0 text-slate-900 dark:text-slate-200 placeholder-slate-400 focus:ring-0"
                        placeholder="Space name..."
                    >
                </div>
            </div>

            <!-- Spaces List -->
            <div v-if="notebooksLoading" class="px-4 py-2 text-xs text-slate-400">Loading...</div>
            <div 
                v-for="notebook in notebooks" 
                :key="notebook.id"
                class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all duration-200 text-sm group text-left relative"
                :class="[
                    (editingState.id === notebook.id) ? 'bg-white dark:bg-white/10 ring-1 ring-indigo-500' :
                    (viewMode === 'EDITOR' && selectedNotebookId === notebook.id) ? 'bg-white dark:bg-white/10 text-slate-900 dark:text-white font-medium shadow-sm' : 
                    'text-slate-500 dark:text-slate-400 hover:bg-white/60 dark:hover:bg-white/5 hover:text-slate-700 dark:hover:text-slate-300'
                ]"
                @click="onSelectNotebook(notebook.id)"
            >
                <span class="material-symbols-outlined text-[18px] shrink-0 transition-colors" :class="selectedNotebookId === notebook.id ? 'text-indigo-500 dark:text-indigo-400' : 'text-slate-400 opacity-80 group-hover:text-slate-500'">folder</span>
                
                <!-- Inline Edit: Space Name -->
                <div v-if="editingState.id === notebook.id && editingState.type === 'NOTEBOOK_NAME'" class="flex-1 min-w-0">
                     <input 
                        v-model="editingState.value"
                        @blur="saveEditing"
                        @keydown.enter="saveEditing"
                        @keydown.esc="cancelEditing"
                        @click.stop
                        v-focus
                        type="text"
                        class="w-full bg-transparent border-none text-sm px-0 py-0 focus:ring-0 p-0 text-slate-900 dark:text-white"
                    >
                </div>
                <span 
                    v-else 
                    @dblclick.stop="startEditing(notebook.id, 'NOTEBOOK_NAME', notebook.name)"
                    class="flex-1 truncate select-none cursor-pointer"
                    title="Double-click to rename"
                >{{ notebook.name }}</span>

                <!-- Inline Edit: Slug Trigger & Delete Option -->
                <div class="hidden group-hover:flex items-center gap-1">
                     <span 
                        @click.stop="startEditing(notebook.id, 'NOTEBOOK_SLUG', notebook.slug)" 
                        class="material-symbols-outlined text-[14px] text-slate-300 hover:text-indigo-500 cursor-pointer transition-colors" 
                        title="Edit Slug"
                    >link</span>
                    
                    <!-- Delete Option (Only if empty) -->
                    <span 
                        v-if="canDelete(notebook) && deleteConfirmationId !== notebook.id"
                        @click.stop="requestDelete(notebook.id)" 
                        class="material-symbols-outlined text-[14px] text-slate-300 hover:text-red-500 cursor-pointer transition-colors" 
                        title="Delete Space"
                    >delete</span>
                </div>

                <!-- Large Inline Confirmation Overlay -->
                <div 
                    v-if="deleteConfirmationId === notebook.id" 
                    class="absolute inset-0 z-20 flex items-center justify-between bg-slate-50 dark:bg-slate-800 rounded-lg px-2 animate-in fade-in duration-200"
                    @click.stop
                >
                    <span class="text-[11px] font-bold text-slate-500 uppercase tracking-wider">Delete?</span>
                    <div class="flex items-center gap-1.5">
                        <button 
                            @click.stop="confirmDelete(notebook.id)" 
                            class="px-2 py-0.5 bg-red-500 hover:bg-red-600 text-white text-[10px] font-bold rounded shadow-sm uppercase tracking-wide transition-colors"
                        >Yes</button>
                        <button 
                            @click.stop="deleteConfirmationId = null" 
                            class="px-2 py-0.5 bg-white dark:bg-slate-700 hover:bg-slate-200 dark:hover:bg-slate-600 text-slate-600 dark:text-slate-300 border border-slate-200 dark:border-slate-600 text-[10px] font-bold rounded shadow-sm uppercase tracking-wide transition-colors"
                        >Cancel</button>
                    </div>
                </div>

                <!-- Inline Edit: Slug Popover -->
                <div v-if="editingState.id === notebook.id && editingState.type === 'NOTEBOOK_SLUG'" class="absolute left-0 bottom-full mb-2 w-60 bg-white dark:bg-slate-800 rounded-lg shadow-xl border border-slate-200 dark:border-white/10 p-3 z-50">
                    <label class="block text-[10px] font-bold text-slate-400 uppercase mb-2">Space Slug</label>
                    <input 
                        v-model="editingState.value" 
                        @keydown.enter="saveEditing"
                        @keydown.esc="cancelEditing"
                        v-focus
                        type="text" 
                        class="w-full text-xs bg-slate-50 dark:bg-black/20 border border-slate-200 dark:border-white/10 rounded px-2 py-1.5 focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500"
                    >
                </div>

                <span v-if="notebook.notes_count > 0" class="text-[10px] text-slate-400 shrink-0 font-mono">{{ notebook.notes_count }}</span>
            </div>
            
        </nav>

        <!-- User Footer (Logout) -->
        <div class="p-3 border-t border-slate-200 dark:border-white/5 bg-slate-100/50 dark:bg-black/10">
            <div class="flex items-center gap-3 p-2 rounded-lg hover:bg-slate-200 dark:hover:bg-white/5 transition-colors group">
                <!-- Avatar -->
                <div class="size-8 rounded-full bg-slate-300 dark:bg-slate-700 overflow-hidden shrink-0 border border-slate-100 dark:border-white/10">
                     <img 
                        v-if="user?.avatar" 
                        :src="user.avatar" 
                        class="w-full h-full object-cover"
                        alt="User"
                    />
                    <div v-else class="w-full h-full flex items-center justify-center text-xs font-bold text-slate-500 dark:text-slate-400">
                        {{ user?.name?.charAt(0) || 'U' }}
                    </div>
                </div>
                
                <div class="flex-1 min-w-0">
                    <div class="text-sm font-medium text-slate-900 dark:text-slate-200 truncate">{{ user?.name }}</div>
                    <div class="text-[10px] text-slate-500 dark:text-slate-400 truncate">{{ user?.email }}</div>
                </div>

                <button 
                    @click="handleLogout" 
                    title="Log Out"
                    class="size-8 flex items-center justify-center rounded-md hover:bg-white dark:hover:bg-white/10 text-slate-400 hover:text-red-500 transition-all shadow-sm opacity-0 group-hover:opacity-100 transform translate-x-2 group-hover:translate-x-0"
                >
                    <span class="material-symbols-outlined text-[18px]">logout</span>
                </button>
            </div>
        </div>
    </aside>
</template>
