<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';

const props = defineProps({
    notes: { type: Array, default: () => [] },
    selectedNoteId: { type: [Number, String], default: null },
    searchQuery: { type: String, default: '' },
    filterStatus: { type: String, default: 'ALL' },
    currentNotebookName: { type: String, default: 'All Notes' },
    isLoading: { type: Boolean, default: false }
});

const emit = defineEmits([
    'update:searchQuery',
    'update:filterStatus',
    'select-note',
    'create-page',
    'rename-note', // { id, type: 'TITLE'|'SLUG', value }
    'request-create-space',
    'delete-note'
]);

// --- LOCAL STATE ---
const searchInput = ref(null);

// Inline Editing for Notes
const editingState = ref({
    id: null,
    type: null, // 'NOTE_TITLE', 'NOTE_SLUG'
    value: '',
    originalValue: ''
});

// Inline Deleting State
const deletingNoteId = ref(null);

const startDeleting = (id) => {
    deletingNoteId.value = id;
};

const cancelDeleting = () => {
    deletingNoteId.value = null;
};

const confirmDeleting = (id) => {
    emit('delete-note', id);
    deletingNoteId.value = null;
};

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

    emit('rename-note', { id, type, value });
    cancelEditing();
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



// Keyboard Shortcuts
const handleKeydown = (e) => {
    // Ctrl + K or Cmd + K
    if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
        e.preventDefault();
        if (searchInput.value) {
            searchInput.value.focus();
        }
    }
};

onMounted(() => {
    window.addEventListener('keydown', handleKeydown);
});

onBeforeUnmount(() => {
    window.removeEventListener('keydown', handleKeydown);
});

const vFocus = {
    mounted: (el) => el.focus()
};
</script>

<template>
    <div class="w-[320px] flex flex-col bg-white/50 dark:bg-[#0b1121]/80 transition-all duration-300">
        <!-- Header -->
        <div class="h-16 flex items-center justify-between px-5 shrink-0">
            <h2 class="font-sans font-bold text-[13px] text-slate-800 dark:text-slate-100 select-none truncate pr-2 uppercase tracking-wide opacity-80" :title="currentNotebookName">{{ currentNotebookName }}</h2>
            <button 
                @click="$emit('create-page')"
                class="size-8 rounded-full bg-slate-900 dark:bg-white text-white dark:text-slate-900 flex items-center justify-center hover:bg-slate-700 dark:hover:bg-slate-200 transition-colors shadow-sm shrink-0"
                title="Add Note"
            >
                <span class="material-symbols-outlined text-[20px]">add</span>
            </button>
        </div>

        <!-- Search -->
        <div class="px-5 py-2 shrink-0">
            <div class="relative group">
                <span class="absolute left-3 top-1/2 -translate-y-1/2 material-symbols-outlined text-slate-400 text-[18px] group-focus-within:text-indigo-500 transition-colors">search</span>
                <input 
                    ref="searchInput"
                    type="text" 
                    name="search_notes"
                    id="search_notes"
                    :value="searchQuery"
                    @input="$emit('update:searchQuery', $event.target.value)"
                    placeholder="Search..."
                    class="w-full h-10 bg-slate-100 dark:bg-white/5 border-none rounded-lg pl-9 pr-14 text-sm focus:ring-1 focus:ring-indigo-500 transition-all placeholder:text-slate-400 text-slate-700 dark:text-slate-200"
                >
                <!-- Ctrl+K Badge -->
                <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none flex items-center opacity-60">
                    <kbd class="hidden sm:inline-block border border-slate-200 dark:border-white/10 rounded px-1.5 py-0.5 text-[10px] font-bold text-slate-400 font-mono bg-white dark:bg-white/5">⌘K</kbd>
                </div>
            </div>
        </div>

        <!-- Filters -->
        <div class="px-5 pb-3 flex gap-2 shrink-0 overflow-x-auto no-scrollbar py-2">
            <button 
                @click="$emit('update:filterStatus', 'ALL')" 
                class="px-3 py-1 rounded-full text-[11px] font-semibold tracking-wide whitespace-nowrap transition-all border border-transparent"
                :class="filterStatus === 'ALL' ? 'bg-slate-200 dark:bg-white/10 text-slate-900 dark:text-white' : 'text-slate-500 hover:bg-slate-100 dark:hover:bg-white/5'"
            >All</button>
            <button 
                @click="$emit('update:filterStatus', 'DRAFT')" 
                class="px-3 py-1 rounded-full text-[11px] font-semibold tracking-wide whitespace-nowrap transition-all border border-transparent"
                :class="filterStatus === 'DRAFT' ? 'bg-orange-50 dark:bg-orange-500/10 text-orange-600 dark:text-orange-400' : 'text-slate-500 hover:bg-slate-100 dark:hover:bg-white/5'"
            >Draft</button>
            <button 
                @click="$emit('update:filterStatus', 'PUBLISHED')" 
                class="px-3 py-1 rounded-full text-[11px] font-semibold tracking-wide whitespace-nowrap transition-all border border-transparent"
                :class="filterStatus === 'PUBLISHED' ? 'bg-indigo-50 dark:bg-indigo-500/10 text-indigo-600 dark:text-indigo-400' : 'text-slate-500 hover:bg-slate-100 dark:hover:bg-white/5'"
            >Published</button>
        </div>

        <!-- Notes List -->
        <div class="flex-1 overflow-y-auto custom-scrollbar px-3 pb-3">
            <div v-if="isLoading" class="flex justify-center items-center h-full">
                <span class="material-symbols-outlined animate-spin text-2xl text-slate-300">progress_activity</span>
            </div>
            
            <div v-else-if="notes.length === 0" class="flex flex-col items-center justify-center h-full text-center px-4 opacity-60">
                <span class="material-symbols-outlined text-4xl mb-3 text-slate-300">edit_note</span>
                <p class="text-xs font-medium text-slate-500">No notes found</p>
                <div class="mt-2 text-[11px] text-slate-400 flex flex-col gap-1">
                    <button @click="$emit('create-page')" class="text-slate-600 hover:underline font-medium">Create new</button>
                </div>
            </div>

            <div v-else class="space-y-1 py-1">
                <div 
                    v-for="note in notes" 
                    :key="note.id"
                    @click="$emit('select-note', note)"
                    class="group p-3 rounded-xl cursor-pointer transition-all duration-200 border border-transparent relative"
                    :class="selectedNoteId === note.id 
                        ? 'bg-white dark:bg-white/10 shadow-sm ring-1 ring-black/5 dark:ring-white/5 z-10' 
                        : 'hover:bg-white/60 dark:hover:bg-white/5'"
                >
                    <!-- Inline Edit Title -->
                    <div v-if="editingState.id === note.id && editingState.type === 'NOTE_TITLE'" class="mb-1">
                         <input 
                            v-model="editingState.value"
                            @blur="saveEditing"
                            @keydown.enter="saveEditing"
                            @keydown.esc="cancelEditing"
                            @click.stop
                            v-focus
                            type="text"
                            class="w-full bg-white dark:bg-black/20 border-slate-300 text-sm px-2 py-1 rounded focus:ring-1 focus:ring-indigo-500 font-bold"
                        >
                    </div>
                    <h3 
                        v-else
                        @dblclick.stop="startEditing(note.id, 'NOTE_TITLE', note.title)"
                        class="font-medium text-[13px] text-slate-800 dark:text-slate-100 line-clamp-2 mb-1 leading-snug tracking-tight"
                        :class="selectedNoteId === note.id ? 'text-slate-900' : 'text-slate-700'"
                        title="Double-click to rename"
                    >
                        {{ note.title || 'Untitled' }}
                    </h3>
                    
                    <div class="flex items-center justify-between text-[11px] text-slate-400 font-medium h-5">
                        <div class="flex items-center gap-2">
                            <span class="font-mono text-[10px] opacity-70">{{ formatRelativeTime(note.updated_at) }}</span>
                            <!-- Status Indicators (Dot only for minimalist) -->
                            <span 
                                v-if="note.status === 'DRAFT'"
                                class="size-1.5 rounded-full bg-orange-400"
                                title="Draft"
                            ></span>
                            <span 
                                v-else
                                class="size-1.5 rounded-full bg-indigo-400"
                                title="Published"
                            ></span>
                        </div>
                        
                        <div class="flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                             <button @click.stop="startEditing(note.id, 'NOTE_TITLE', note.title)" class="p-1 hover:bg-slate-200 dark:hover:bg-white/10 rounded text-slate-400 hover:text-indigo-500 transition-colors" title="Rename">
                                <span class="material-symbols-outlined text-[14px]">edit</span>
                            </button>
                            <button @click.stop="startEditing(note.id, 'NOTE_SLUG', note.slug)" class="p-1 hover:bg-slate-200 dark:hover:bg-white/10 rounded text-slate-400 hover:text-indigo-500 transition-colors" title="Link slug">
                                <span class="material-symbols-outlined text-[14px]">link</span>
                            </button>
                            
                            <!-- Inline Delete -->
                            <div v-if="deletingNoteId === note.id" class="flex items-center bg-white dark:bg-slate-800 shadow-sm rounded border border-slate-200 dark:border-white/10 overflow-hidden transform scale-110 origin-right transition-all" @click.stop>
                                <button @click.stop="confirmDeleting(note.id)" class="px-1.5 py-0.5 text-[10px] font-bold text-white bg-red-500 hover:bg-red-600 transition-colors">DEL?</button>
                                <button @click.stop="cancelDeleting" class="px-1 py-0.5 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 hover:bg-slate-100 dark:hover:bg-white/5 transition-colors">
                                    <span class="material-symbols-outlined text-[14px] leading-none block">close</span>
                                </button>
                            </div>
                            <button v-else @click.stop="startDeleting(note.id)" class="p-1 hover:bg-red-50 dark:hover:bg-red-500/10 rounded text-slate-400 hover:text-red-500 transition-colors" title="Delete">
                                 <span class="material-symbols-outlined text-[14px]">delete</span>
                            </button>
                        </div>
                    </div>
                    
                    <!-- Inline Edit: Slug Popover -->
                    <div v-if="editingState.id === note.id && editingState.type === 'NOTE_SLUG'" class="absolute right-2 top-full mt-1 w-56 bg-white dark:bg-slate-800 rounded-lg shadow-xl border border-slate-200 dark:border-white/10 p-3 z-50">
                            <label class="block text-[10px] font-bold text-slate-400 uppercase mb-1">URL Slug</label>
                            <input 
                                v-model="editingState.value" 
                                @keydown.enter="saveEditing"
                                @keydown.esc="cancelEditing"
                                v-focus
                                @click.stop
                                type="text" 
                                class="w-full text-xs bg-slate-50 dark:bg-black/20 border border-slate-200 dark:border-white/10 rounded px-2 py-1.5 focus:ring-1 focus:ring-indigo-500"
                            >
                    </div>

                </div>
            </div>
        </div>
    </div>
</template>
