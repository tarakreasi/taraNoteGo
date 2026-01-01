<script setup>
import { ref, watch, onBeforeUnmount, onMounted } from 'vue';
import NoteEditor from '@/Components/NoteEditor.vue';

const props = defineProps({
    note: { type: Object, default: null },
    saveStatus: { type: String, default: '' }
});

const emit = defineEmits(['save', 'publish', 'delete']);

// Local State
const title = ref('');
const content = ref('');
const status = ref('DRAFT');
const isFeatured = ref(false);
const notebookId = ref(null);
const fileInput = ref(null);
const coverImageFile = ref(null); // The actual file object
const coverImagePreview = ref(null); // The URL for display
const lastSavedAt = ref(null);

// Debounce Timer
let autoSaveTimeout = null;

// Sync Valid Note Data
watch(() => props.note, (newNote, oldNote) => {
    if (newNote) {
        // Only update content/title if switching to a DIFFERENT note
        // or if it's the first load (oldNote is undefined)
        if (!oldNote || newNote.id !== oldNote.id) {
            title.value = newNote.title || '';
            content.value = newNote.content || '';
            status.value = newNote.status || 'DRAFT';
            isFeatured.value = !!newNote.is_featured;
            notebookId.value = newNote.notebook_id;
            coverImagePreview.value = newNote.cover_image || null;
            coverImageFile.value = null; // Reset file input
            lastSavedAt.value = newNote.updated_at;
        } else {
            // Same note updated (e.g. after save) - only sync metadata that shouldn't conflict with typing
            // We TRUST local state for content/title while editing
            lastSavedAt.value = newNote.updated_at;
            // Optionally sync status/featured if needed, but local acts as source of truth usually
            if (newNote.cover_image !== undefined) coverImagePreview.value = newNote.cover_image;
        }
    }
}, { immediate: true });

// Helper to gather data
const getFormData = () => {
    return {
        id: props.note?.id,
        title: title.value,
        content: content.value,
        status: status.value,
        is_featured: isFeatured.value,
        notebook_id: notebookId.value,
        cover_image: coverImageFile.value // File or null
    };
};

// Auto-Save Trigger
const triggerAutoSave = () => {
    if (!props.note) return;
    
    // We emit an event to notify parent something changed (for UI feedback like "Unsaved")
    // But technically we want to debounce the actual save call.
    // The parent controls the API call.
    // Let's handle debounce here to avoid spamming parent.
    
    clearTimeout(autoSaveTimeout);
    autoSaveTimeout = setTimeout(() => {
        emit('save', { data: getFormData(), isAutoSave: true });
    }, 1500);
};

// Immediate Save (e.g. Publish)
const triggerImmediateSave = (overrideStatus = null) => {
    if (overrideStatus) status.value = overrideStatus;
    clearTimeout(autoSaveTimeout);
    emit('save', { data: getFormData(), isAutoSave: false });
};

// Watchers for Auto-Save
watch([title, isFeatured], () => {
    triggerAutoSave();
});

// Shortcuts
const handleKeydown = (e) => {
    // Ctrl + S
    if ((e.ctrlKey || e.metaKey) && e.key === 's') {
        e.preventDefault();
        triggerImmediateSave();
    }
    // Ctrl + Enter
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        e.preventDefault();
        triggerImmediateSave('PUBLISHED');
    }
};

onMounted(() => {
    window.addEventListener('keydown', handleKeydown);
});

onBeforeUnmount(() => {
    window.removeEventListener('keydown', handleKeydown);
    clearTimeout(autoSaveTimeout);
});

// Content Save (from Tiptap)
const onEditorSave = (newContent) => {
    // NoteEditor v-model updates 'content' automatically? 
    // If v-model is used: <NoteEditor v-model="content" />
    // But we also need to trigger save.
    // NoteEditor @save event is fired usually on user pause/typing.
    
    triggerAutoSave();
};

// Image Upload
const triggerImageUpload = () => {
    fileInput.value.click();
};

const handleImageUpload = (event) => {
    const file = event.target.files[0];
    if (file) {
        coverImageFile.value = file;
        
        // Create preview
        const reader = new FileReader();
        reader.onload = (e) => {
            coverImagePreview.value = e.target.result;
        };
        reader.readAsDataURL(file);

        // Auto save immediately (or quickly)
        setTimeout(() => triggerImmediateSave(), 1000);
    }
};

const deleteNote = () => {
    if (props.note) emit('delete', props.note.id);
};
</script>

<template>
    <main class="flex-1 flex flex-col bg-slate-50 dark:bg-[#0F172A] relative transition-all duration-300 border-l border-slate-100 dark:border-white/5">
        
        <!-- Refined Toolbar -->
        <header v-if="note" class="h-14 flex items-center justify-between px-8 shrink-0 bg-transparent transition-all">
            <!-- Breadcrumb -->
            <div class="flex items-center gap-2 text-xs text-slate-400 font-sans tracking-wide">
                <span class="hover:text-slate-600 dark:hover:text-slate-300 cursor-pointer transition-colors">Home</span>
                <span class="text-slate-200 dark:text-slate-700">/</span>
                <span v-if="note.notebook" class="hover:text-slate-600 dark:hover:text-slate-300 cursor-pointer transition-colors">{{ note.notebook.name }}</span>
                <span v-else class="text-slate-300">Uncategorized</span>
                <span class="text-slate-200 dark:text-slate-700">/</span>
                <span class="text-slate-600 dark:text-slate-300 font-medium max-w-[200px] truncate select-all">{{ title || 'Untitled' }}</span>
                
                <!-- Status Badge (Minimalist) -->
                <span 
                    class="ml-3 text-[10px] font-mono uppercase tracking-wider opacity-70"
                    :class="{
                        'text-orange-500': status === 'DRAFT',
                        'text-indigo-500': status === 'PUBLISHED'
                    }"
                >
                    {{ status }}
                </span>
            </div>

            <div class="flex items-center gap-3">
                 <span v-if="saveStatus" class="mr-2 text-[10px] font-mono text-slate-400 tracking-wide uppercase" :class="{'text-blue-500': saveStatus.includes('Saving'), 'text-green-500': saveStatus.includes('Saved')}">
                     <span v-if="saveStatus.includes('Saving')" class="animate-pulse mr-1">●</span>
                     {{ saveStatus === '✓ Saved' ? 'Saved' : saveStatus }}
                 </span>

                <button 
                    @click="isFeatured = !isFeatured"
                    class="size-8 rounded-full transition-colors flex items-center justify-center hover:bg-slate-100 dark:hover:bg-white/5"
                    :class="isFeatured 
                        ? 'text-yellow-500 dark:text-yellow-400' 
                        : 'text-slate-300 hover:text-yellow-500'"
                    title="Toggle Featured"
                >
                    <span class="material-symbols-outlined text-[18px]">{{ isFeatured ? 'star' : 'star_border' }}</span>
                </button>

                <!-- Preview Link (Commented - route archived) -->
                <!--
                <a 
                    v-if="note?.slug"
                    :href="route('articles.show', note.slug)" 
                    target="_blank"
                    class="size-8 rounded-full text-slate-300 hover:text-indigo-500 dark:hover:text-indigo-400 hover:bg-slate-100 dark:hover:bg-white/5 transition-colors flex items-center justify-center"
                    title="Preview Article"
                >
                    <span class="material-symbols-outlined text-[18px]">visibility</span>
                </a>
                -->

                <button 
                    @click="triggerImmediateSave('PUBLISHED')"
                    class="h-7 px-3 rounded text-[11px] font-bold uppercase tracking-wider transition-all flex items-center gap-2"
                    :class="status === 'PUBLISHED' 
                        ? 'bg-slate-100 dark:bg-white/10 text-slate-500 hover:text-slate-800 dark:hover:text-white' 
                        : 'bg-indigo-500 text-white hover:bg-indigo-600 shadow-sm'"
                >
                    {{ status === 'PUBLISHED' ? 'Update' : 'Publish' }}
                </button>
                

            </div>
        </header>

        <!-- Empty State -->
        <div v-if="!note" class="flex-1 flex flex-col items-center justify-center text-gray-400 opacity-40">
            <div class="size-16 rounded-2xl bg-slate-100 dark:bg-white/5 flex items-center justify-center mb-4">
                <span class="material-symbols-outlined text-3xl opacity-50">edit_note</span>
            </div>
            <h2 class="text-sm font-medium text-gray-900 dark:text-white mb-1">Select a note</h2>
            <p class="text-xs">Select from list or create new</p>
        </div>

        <!-- Editor Surface -->
        <div v-else class="flex-1 overflow-hidden flex flex-col">
            <div class="flex-1 overflow-y-auto custom-scrollbar px-12 py-8">
                <div class="w-full max-w-7xl xl:pr-[280px]">
                    
                    <!-- Title Area: Increase Negative Space -->
                    <input 
                        type="text" 
                        v-model="title"
                        placeholder="Untitled Page"
                        class="w-full bg-transparent text-4xl font-sans font-bold text-slate-800 dark:text-slate-100 border-none focus:ring-0 placeholder-gray-300 dark:placeholder-gray-700 mb-6 p-0 leading-tight tracking-tight"
                    >

                    <!-- Cover Image UI (Refined) -->
                    <div class="mb-10 group relative">
                        <input 
                            type="file" 
                            ref="fileInput" 
                            class="hidden" 
                            accept="image/*"
                            @change="handleImageUpload"
                        >
                        
                        <!-- Preview Area -->
                        <div v-if="coverImagePreview" class="relative w-full h-56 md:h-72 rounded-sm overflow-hidden group/image cursor-pointer shadow-sm hover:shadow-md transition-shadow" @click="triggerImageUpload">
                            <img :src="coverImagePreview" class="w-full h-full object-cover opacity-90 group-hover/image:opacity-100 transition-opacity duration-500" alt="Cover">
                            <div class="absolute inset-0 bg-black/10 opacity-0 group-hover/image:opacity-100 transition-opacity flex items-center justify-center">
                                <span class="flex items-center gap-2 px-3 py-1.5 bg-white/90 dark:bg-black/80 backdrop-blur-sm rounded text-slate-800 dark:text-white text-xs font-bold uppercase tracking-wide">
                                    Change Cover
                                </span>
                            </div>
                        </div>

                        <!-- Add Button (if no image) -->
                        <button 
                            v-else 
                            @click="triggerImageUpload"
                            class="flex items-center gap-2 text-xs font-bold uppercase tracking-wider text-slate-300 hover:text-indigo-500 transition-colors px-1 py-1"
                        >
                            <span class="material-symbols-outlined text-[18px]">add_photo_alternate</span>
                            Add Cover
                        </button>
                    </div>

                    <!-- Tiptap Container -->
                    <div class="h-[calc(100%-80px)]">
                         <NoteEditor 
                            v-model="content" 
                            :save-status="saveStatus"
                            :last-saved-at="lastSavedAt"
                            @save="onEditorSave"
                            @request-create-space="$emit('request-create-space')"
                         />
                    </div>
                </div>
            </div>
        </div>
    </main>
</template>
