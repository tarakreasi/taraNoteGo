<script setup>
import { Head, Link } from '@inertiajs/vue3';
import { ref, onMounted } from 'vue';

const props = defineProps({
    note: Object, // Optional, for edit mode
});

// --- STATE ---
// Local ID for handling "Create then Update" flow
const noteId = ref(props.note?.id || null);
const title = ref(props.note?.title || '');
const content = ref(props.note?.content || '');
const selectedNotebookId = ref(props.note?.notebook_id || null);
const noteStatus = ref(props.note?.status || 'DRAFT'); // DRAFT, PUBLISHED
const noteSlug = ref(props.note?.slug || null); // Store Slug
const notebooks = ref([]);
const isSaving = ref(false);
const saveStatus = ref('Saved'); // 'Saved', 'Saving...', 'Unsaved Changes'

// --- API ---
const fetchNotebooks = async () => {
    try {
        const response = await window.axios.get('/api/v1/admin/notebooks');
        notebooks.value = response.data.data;
        // Default to first notebook if none selected and creating new
        if (!selectedNotebookId.value && notebooks.value.length > 0) {
            selectedNotebookId.value = notebooks.value[0].id;
        }
    } catch (error) {
        console.error("Error fetching notebooks:", error);
    }
};

// --- AUTO SAVE ---
let debounceTimer = null;

const debouncedAutoSave = () => {
    saveStatus.value = 'Saving...';
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => {
        saveNote('DRAFT', true); // Auto-save always as current status or DRAFT? Usually draft for auto-save.
    }, 2000); // 2 seconds delay
};

import { watch } from 'vue';
watch([title, content, selectedNotebookId], () => {
    saveStatus.value = 'Unsaved Changes';
    debouncedAutoSave();
});

const saveNote = async (status, isAutoSave = false) => {
    isSaving.value = true;
    if (isAutoSave) saveStatus.value = 'Saving...';

    try {
        const payload = {
            title: title.value,
            content: content.value,
            notebook_id: selectedNotebookId.value,
            status: status
        };

        let response;
        if (noteId.value) {
            // Update
            response = await window.axios.put(`/api/v1/admin/notes/${noteId.value}`, payload);
        } else {
            // Create
            response = await window.axios.post('/api/v1/admin/notes', payload);
            // Update local ID so next save is an update
            noteId.value = response.data.data.id;
            // Optionally update URL without reload (HTML5 History API)
            // window.history.replaceState({}, '', `/notes/${noteId.value}/edit`);
        }
        
        noteStatus.value = status; // Update local status
        noteSlug.value = response.data.data.slug; // Update slug from response
        saveStatus.value = 'Saved';
        
        if (!isAutoSave) {
             alert('Note published successfully!');
             // Redirect to dashboard logic could go here if requested
             // window.location.href = '/dashboard'; 
        }

    } catch (error) {
        console.error("Error saving note:", error);
        saveStatus.value = 'Error Saving';
        if (!isAutoSave) alert('Failed to save note.');
    } finally {
        isSaving.value = false;
    }
};

const previewNote = async () => {
    // Determine target status: keep current if known, else DRAFT
    const targetStatus = noteStatus.value || 'DRAFT';
    
    // Save first to ensure latest content is available
    await saveNote(targetStatus, true); // Treat as auto-save to avoid "Published" alert
    
    if (noteSlug.value) {
        // Open article in new tab
        // Use full URL or route helper if available in script. Assuming window.route or manual construct.
        // Since we are in script setup, we might need to manually construct if `route()` isn't globally available in JS context (usually it is via Ziggy).
        // Safest is to just use window.open
        window.open(`/articles/${noteSlug.value}`, '_blank');
    } else {
        alert("Could not generate preview link.");
    }
};

// --- LIFECYCLE ---
onMounted(() => {
    fetchNotebooks();
});

// Auto-resize textarea
const resizeTextarea = (event) => {
    event.target.style.height = 'auto';
    event.target.style.height = event.target.scrollHeight + 'px';
};
</script>

<template>
    <Head title="Editor" />

    <div class="flex h-screen w-full flex-col overflow-hidden bg-[#F8FAFC] dark:bg-[#0F172A] font-display text-text-main antialiased selection:bg-primary/20 selection:text-primary relative transition-colors duration-500">
        
        <!-- ZEN MODE BACKGROUND BLOBS -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none z-0">
             <!-- Blue Blob -->
            <div class="absolute top-0 left-1/4 w-[500px] h-[500px] bg-blue-500/10 dark:bg-blue-500/5 rounded-full mix-blend-multiply dark:mix-blend-screen filter blur-[80px] opacity-70 animate-[breathe_10s_infinite]"></div>
            <!-- Purple Blob -->
            <div class="absolute top-0 right-1/4 w-[500px] h-[500px] bg-purple-500/10 dark:bg-purple-500/5 rounded-full mix-blend-multiply dark:mix-blend-screen filter blur-[80px] opacity-70 animate-[breathe_10s_infinite] delay-1000"></div>
            <!-- Orange Blob -->
            <div class="absolute -bottom-32 left-1/3 w-[600px] h-[600px] bg-orange-500/10 dark:bg-orange-500/5 rounded-full mix-blend-multiply dark:mix-blend-screen filter blur-[100px] opacity-50 animate-[float_12s_ease-in-out_infinite] delay-2000"></div>
        </div>

        <!-- HEADER -->
        <header class="relative z-20 flex h-16 w-full shrink-0 items-center justify-between border-b border-white/20 bg-white/70 dark:bg-[#0F172A]/70 backdrop-blur-xl dark:border-white/10 px-6 transition-colors">
            <div class="flex items-center gap-4">
                <Link href="/taraNote" class="flex size-10 items-center justify-center rounded-full hover:bg-black/5 dark:hover:bg-white/10 transition-colors text-slate-700 dark:text-slate-200">
                    <span class="material-symbols-outlined">arrow_back</span>
                </Link>
                <div class="flex flex-col">
                    <div class="flex items-center gap-2">
                         <span class="font-extrabold text-lg tracking-tight text-[#1a1a1a] dark:text-white">
                            taraNote<span class="text-emerald-500">.</span> Editor
                        </span>
                    </div>
                    <div class="flex items-center gap-2">
                        <span class="text-[10px] text-slate-500 dark:text-slate-400">{{ title || 'Untitled Post' }}</span>
                        <span class="text-[9px] px-1.5 py-0.5 rounded-full bg-slate-100 dark:bg-white/10 text-slate-500 dark:text-slate-300 font-bold uppercase">{{ saveStatus }}</span>
                    </div>
                </div>
            </div>
            
            <div class="flex items-center gap-4">
                <button @click="previewNote" class="flex items-center gap-2 rounded-full px-4 py-2 text-sm font-medium text-slate-700 hover:bg-black/5 dark:text-slate-200 dark:hover:bg-white/10 transition-colors">
                    <span class="material-symbols-outlined text-[20px] text-indigo-500">visibility</span>
                    <span>Preview</span>
                </button>
            </div>
        </header>

        <div class="flex flex-1 overflow-hidden relative z-10">
            <!-- MAIN EDITOR AREA -->
            <main class="flex flex-1 flex-col overflow-y-auto bg-transparent">
                <div class="mx-auto w-full max-w-4xl px-8 py-10 md:px-12 pb-40">
                    <!-- Title Input -->
                    <div class="mb-10">
                        <textarea 
                            v-model="title"
                            class="w-full resize-none overflow-hidden bg-transparent p-0 text-5xl md:text-6xl font-bold leading-tight text-slate-900 dark:text-white placeholder:text-slate-300 dark:placeholder:text-slate-600 focus:outline-none focus:ring-0 transition-colors" 
                            @input="resizeTextarea" 
                            placeholder="Enter a captivating title..." 
                            rows="1"
                        ></textarea>
                    </div>

                    <!-- Toolbar (Floating Glass) -->
                    <div class="sticky top-0 z-50 mb-8 flex flex-wrap items-center gap-2 p-2 rounded-2xl bg-white/60 dark:bg-[#0F172A]/60 backdrop-blur-md border border-white/20 dark:border-white/10 shadow-sm w-fit mx-auto transition-all">
                        <button class="rounded-xl p-2 text-slate-600 dark:text-slate-300 hover:bg-indigo-50 dark:hover:bg-indigo-500/20 hover:text-indigo-600 dark:hover:text-indigo-400 transition-all" title="Bold"><span class="material-symbols-outlined text-[20px]">format_bold</span></button>
                        <button class="rounded-xl p-2 text-slate-600 dark:text-slate-300 hover:bg-indigo-50 dark:hover:bg-indigo-500/20 hover:text-indigo-600 dark:hover:text-indigo-400 transition-all" title="Italic"><span class="material-symbols-outlined text-[20px]">format_italic</span></button>
                        <div class="mx-2 h-4 w-px bg-slate-300 dark:bg-white/20"></div>
                        <button class="rounded-xl p-2 text-slate-600 dark:text-slate-300 hover:bg-indigo-50 dark:hover:bg-indigo-500/20 hover:text-indigo-600 dark:hover:text-indigo-400 transition-all" title="Heading 1"><span class="material-symbols-outlined text-[20px]">title</span></button>
                        <button class="rounded-xl p-2 text-slate-600 dark:text-slate-300 hover:bg-indigo-50 dark:hover:bg-indigo-500/20 hover:text-indigo-600 dark:hover:text-indigo-400 transition-all" title="List"><span class="material-symbols-outlined text-[20px]">format_list_bulleted</span></button>
                    </div>

                    <!-- Content Area -->
                    <div class="min-h-[500px] w-full text-lg leading-relaxed text-slate-700 dark:text-slate-300 font-body">
                        <textarea 
                            v-model="content"
                            class="w-full h-full min-h-[500px] bg-transparent border-none focus:ring-0 p-0 resize-none text-lg text-slate-700 dark:text-slate-300 placeholder:text-slate-400"
                            placeholder="Start writing your story here..."
                        ></textarea>
                    </div>
                </div>
            </main>

            <!-- RIGHT SETTINGS PANEL (Glass) -->
            <aside class="w-[320px] shrink-0 border-l border-white/20 bg-white/40 dark:bg-[#0F172A]/40 backdrop-blur-md dark:border-white/5 overflow-y-auto hidden lg:block transition-all">
                <div class="flex flex-col gap-8 p-6">
                    
                    <!-- Notebook Selection -->
                    <div class="flex flex-col gap-3">
                        <label class="text-xs font-bold uppercase tracking-wider text-slate-400 dark:text-slate-500">Notebook</label>
                        <div class="relative">
                            <select v-model="selectedNotebookId" class="w-full appearance-none rounded-xl border-none bg-white/50 dark:bg-black/20 p-3 pr-10 text-sm font-medium text-slate-700 dark:text-slate-200 shadow-sm ring-1 ring-black/5 dark:ring-white/10 focus:ring-2 focus:ring-indigo-500 outline-none transition-all">
                                <option :value="null" disabled>Select Notebook</option>
                                <option v-for="book in notebooks" :key="book.id" :value="book.id">{{ book.name }}</option>
                            </select>
                            <span class="pointer-events-none absolute right-3 top-1/2 -translate-y-1/2 text-slate-400">
                                <span class="material-symbols-outlined">expand_more</span>
                            </span>
                        </div>
                    </div>

                    <!-- Cover Image (Mockup) -->
                    <div class="flex flex-col gap-3">
                        <label class="text-xs font-bold uppercase tracking-wider text-slate-400 dark:text-slate-500">Cover Image</label>
                        <div class="group relative flex aspect-video w-full cursor-pointer flex-col items-center justify-center gap-2 rounded-2xl border-2 border-dashed border-slate-300 dark:border-slate-700 bg-white/30 dark:bg-white/5 transition-all hover:border-indigo-500 hover:bg-indigo-50/50 dark:hover:border-indigo-400 dark:hover:bg-indigo-500/10">
                            <div class="flex size-10 items-center justify-center rounded-full bg-slate-100 dark:bg-white/10 text-slate-400 group-hover:bg-indigo-500 group-hover:text-white transition-colors">
                                <span class="material-symbols-outlined text-[20px]">cloud_upload</span>
                            </div>
                            <p class="text-xs font-medium text-slate-500 dark:text-slate-400 group-hover:text-indigo-500 transition-colors">Click to upload</p>
                        </div>
                    </div>

                    <!-- Action Buttons -->
                    <div class="mt-auto pt-8">
                        <div class="flex flex-col gap-4">
                            <!-- Status Toggle -->
                            <div class="flex w-full overflow-hidden rounded-full bg-slate-200 dark:bg-black/20 p-1">
                                <button @click="saveNote('DRAFT')" 
                                    :class="noteStatus === 'DRAFT' ? 'bg-white dark:bg-slate-700 text-indigo-600 dark:text-white shadow-sm' : 'text-slate-500 hover:text-slate-700 dark:hover:text-slate-300'"
                                    class="flex-1 rounded-full py-2 text-xs font-bold transition-all">
                                    DRAFT
                                </button>
                                <button @click="saveNote('PUBLISHED')"
                                    :class="noteStatus === 'PUBLISHED' ? 'bg-white dark:bg-slate-700 text-indigo-600 dark:text-white shadow-sm' : 'text-slate-500 hover:text-slate-700 dark:hover:text-slate-300'"
                                    class="flex-1 rounded-full py-2 text-xs font-bold transition-all">
                                    PUBLISH
                                </button>
                            </div>
                            
                            <!-- Main Publish Button -->
                            <button @click="saveNote('PUBLISHED')" :disabled="isSaving" class="relative overflow-hidden rounded-full py-4 px-8 text-base font-bold text-white shadow-lg shadow-orange-500/30 bg-gradient-to-r from-orange-400 to-orange-600 hover:from-orange-500 hover:to-orange-700 active:scale-[0.98] transition-all flex items-center justify-center gap-2 w-full disabled:opacity-70 disabled:cursor-not-allowed">
                                <span v-if="isSaving && saveStatus !== 'Saving...'" class="material-symbols-outlined animate-spin text-[20px]">progress_activity</span>
                                <span v-else class="material-symbols-outlined">publish</span>
                                {{ isSaving && saveStatus !== 'Saving...' ? 'PUBLISHING...' : 'PUBLISH NOW' }}
                            </button>
                        </div>
                    </div>

                </div>
            </aside>
        </div>
    </div>
</template>

<style scoped>
/* Scoped styles if needed */
</style>
