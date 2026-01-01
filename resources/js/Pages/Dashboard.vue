<script setup>
import { Head } from '@inertiajs/vue3';
import { ref, onMounted, watch, computed } from 'vue';
import { useTheme } from '@/composables/useTheme';
import SettingsPanel from '@/Pages/Dashboard/SettingsPanel.vue';
import FloatingDock from '@/Components/FloatingDock.vue';
import Sidebar from '@/Pages/Dashboard/Sidebar.vue';
import NoteList from '@/Pages/Dashboard/NoteList.vue';
import EditorSection from '@/Pages/Dashboard/EditorSection.vue';

// --- STATE ---
const notes = ref([]);
const notebooks = ref([]);
const notebooksLoading = ref(true);
const isLoading = ref(true); // for notes
const sidebarRef = ref(null);

// Filter & Search
const searchQuery = ref('');
const filterStatus = ref('ALL');
const selectedNotebookId = ref(null);

// View Mode
const viewMode = ref('EDITOR'); // 'EDITOR' or 'SETTINGS'

// Editor State (just the selected note reference and status)
const selectedNote = ref(null);
const saveStatus = ref('');

// --- DATA FETCHING ---
const fetchNotes = async () => {
    isLoading.value = true;
    try {
        const response = await window.axios.get('/api/v1/admin/notes', {
            params: {
                search: searchQuery.value,
                status: filterStatus.value !== 'ALL' ? filterStatus.value : null,
                notebook_id: selectedNotebookId.value
            }
        });
        notes.value = response.data.data; 
    } catch (error) {
        console.error("Error fetching notes:", error);
    } finally {
        isLoading.value = false;
    }
};

const fetchNotebooks = async () => {
    notebooksLoading.value = true;
    try {
        const response = await window.axios.get('/api/v1/admin/notebooks');
        notebooks.value = response.data.data.sort((a, b) => a.id - b.id);
    } catch (error) {
        console.error("Error fetching notebooks:", error);
        notebooksLoading.value = false;
    } finally {
        notebooksLoading.value = false;
    }
};

// --- SEARCH & FILTER WATCHERS ---
let searchTimeout = null;
watch(searchQuery, () => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => fetchNotes(), 500);
});

const onSetFilter = (status) => {
    filterStatus.value = status;
    fetchNotes();
};

const onSelectNotebook = (id) => {
    selectedNotebookId.value = id;
    fetchNotes();
};

// --- LIFECYCLE ---
onMounted(() => {
    fetchNotebooks();
    fetchNotes();

    // Check query param
    const params = new URLSearchParams(window.location.search);
    if (params.get('view') === 'settings') {
        viewMode.value = 'SETTINGS';
        selectedNote.value = null;
    }
});

// --- ACTIONS: SIDEBAR (Notebooks) ---
const handleCreateNotebook = async (name) => {
    try {
        await window.axios.post('/api/v1/admin/notebooks', { name });
        await fetchNotebooks();
    } catch (error) {
        console.error(error);
        alert("Failed to create space");
    }
};

const handleRenameNotebook = async ({ id, type, value }) => {
    try {
        if (type === 'NOTEBOOK_NAME') {
            await window.axios.put(`/api/v1/admin/notebooks/${id}`, { name: value });
        } else if (type === 'NOTEBOOK_SLUG') {
             await window.axios.put(`/api/v1/admin/notebooks/${id}`, { slug: value });
        }
        await fetchNotebooks();
    } catch (error) {
         console.error("Failed to rename notebook:", error);
         alert("Failed to save changes.");
    }
};

const handleDeleteNotebook = async (id) => {
    try {
        await window.axios.delete(`/api/v1/admin/notebooks/${id}`);
        await fetchNotebooks();
        if (selectedNotebookId.value === id) {
            onSelectNotebook(null);
        }
    } catch (error) {
        console.error(error);
        alert(error.response?.data?.message || "Failed to delete space.");
    }
};

const handleRequestCreateSpace = () => {
    if (sidebarRef.value) {
        sidebarRef.value.startCreatingSpace();
    }
};

// --- ACTIONS: NOTELIST (Notes) ---
const handleCreatePage = async () => {
    const targetNotebookId = selectedNotebookId.value || (notebooks.value.length > 0 ? notebooks.value[0].id : null);
    try {
        const response = await window.axios.post('/api/v1/admin/notes', {
            title: 'Untitled Page',
            notebook_id: targetNotebookId,
            content: '',
            status: 'DRAFT'
        });
        fetchNotes();
        // Select the new note
        // The list will reload, so we need to find it or set it.
        // The API returns the new note, so we can set it directly.
        // But we should also push it to notes list locally to avoid reload lag?
        // fetchNotes() handles reload.
        selectedNote.value = response.data.data;
    } catch (error) {
        console.error(error);
        alert("Failed to create page");
    }
};

const handleRenameNote = async ({ id, type, value }) => {
    try {
        const payload = type === 'NOTE_TITLE' ? { title: value } : { slug: value };
        const response = await window.axios.patch(`/api/v1/admin/notes/${id}`, payload);
        
        // Update local list
        const updated = response.data.data;
        const index = notes.value.findIndex(n => n.id === updated.id);
        if (index !== -1) notes.value[index] = updated;
        
        // Update selected note if matches
        if (selectedNote.value?.id === updated.id) {
            selectedNote.value = updated;
        }
    } catch (error) {
        console.error("Failed to rename note:", error);
        alert("Failed to save changes.");
    }
};

const handleSelectNote = (note) => {
    selectedNote.value = note;
    saveStatus.value = '';
};

// --- ACTIONS: EDITOR ---
const handleSaveNote = async ({ data, isAutoSave }) => {
    saveStatus.value = 'Saving...';
    try {
        const formData = new FormData();
        formData.append('title', data.title || 'Untitled Page');
        formData.append('content', data.content || '');
        if (data.notebook_id) formData.append('notebook_id', data.notebook_id);
        formData.append('status', data.status);
        formData.append('_method', 'PUT'); // Fake PUT
        formData.append('is_featured', data.is_featured ? '1' : '0');

        if (data.cover_image) {
            formData.append('cover_image', data.cover_image);
        }

        const response = await window.axios.post(`/api/v1/admin/notes/${data.id}`, formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });

        saveStatus.value = '✓ Saved';
        
        // Update local state
        const updatedNote = response.data.data;
        const index = notes.value.findIndex(n => n.id === updatedNote.id);
        if (index !== -1) notes.value[index] = updatedNote;
        
        // Keep selectedNote in sync
        if (selectedNote.value?.id === updatedNote.id) {
            selectedNote.value = updatedNote;
        }

        if (isAutoSave) {
            setTimeout(() => {
                if (saveStatus.value === '✓ Saved') saveStatus.value = '';
            }, 2000);
        }
    } catch (error) {
        console.error("Error saving note:", error);
        saveStatus.value = '✗ Error';
        if (!isAutoSave) alert('Failed to save note');
    }
};

const handleDeleteNote = async (id) => {
    // Confirmation handled in NoteList.vue inline

    try {
        await window.axios.delete(`/api/v1/admin/notes/${id}`);
        if (selectedNote.value?.id === id) {
            selectedNote.value = null;
        }
        fetchNotes();
        fetchNotebooks(); // Count might change
    } catch (error) {
        console.error("Delete failed:", error);
        alert('Failed to delete');
    }
};

// --- ACTIONS: SETTINGS (PROFILE) ---
const handleSaveProfile = async (userData) => {
    saveStatus.value = 'Saving Profile...';
    try {
        // Prepare payload (only send allowed fields)
        const payload = {
            name: userData.name,
            username: userData.username,
            email: userData.email,
            bio: userData.bio,
            hero_tagline: userData.hero_tagline,
            // avatar and hero_image are URLs already uploaded via SettingsPanel
            avatar: userData.avatar,
            hero_image: userData.hero_image,
        };

        await window.axios.patch(route('profile.update'), payload);
        saveStatus.value = '✓ Profile Saved';
        
        // Update local auth user state if Inertia doesn't do it automatically
        // Since we are not doing a full reload, we rely on the response or just assume success.
        // Inertia 'usePage' props might get stale until next visit, 
        // but 'handleSaveProfile' receives 'userData' which comes from 'localUser' in SettingsPanel
        // which was already mutated by v-model. 
        // So the UI is optimistic.
    } catch (error) {
        console.error("Profile save failed:", error);
        saveStatus.value = '✗ Error Saving Profile';
        alert("Failed to save profile.");
    } finally {
        setTimeout(() => {
            if (saveStatus.value === '✓ Profile Saved') saveStatus.value = '';
        }, 3000);
    }
};

// --- COMPUTED ---
const currentNotebookName = computed(() => {
    if (!selectedNotebookId.value) return 'All Notes';
    return notebooks.value.find(n => n.id === selectedNotebookId.value)?.name || 'All Notes';
});

// This filteredNotes logic is currently handled by Backend API (fetchNotes)
// but local search filtering exists in original code?
// Original used `fetchNotes` for search, but also had `filteredNotes` computed.
// Let's check original:
// "const response = await window.axios.get('/api/v1/admin/notes', ...)" -> `notes.value = response.data.data`
// `filteredNotes` computed: "let filtered = notes.value ... if filterStatus ... if searchQuery"
// Wait, if API does filtering, local filtering is redundant unless API returns ALL.
// Original `fetchNotes` params: `search: searchQuery.value`.
// So the API DOES key filtering.
// BUT `filteredNotes` computed ALSO filters locally?
// Line 470 `filteredNotes` computed.
// It seems hybrid/redundant or for instant feedback? 
// If `notes.value` already comes from API filtered, then local filter might be empty.
// Ah, `fetchNotes` used `searchQuery` and `filterStatus`.
// So `notes` IS ALREADY filtered.
// So `filteredNotes` computed might be double filtering or doing nothing roughly.
// EXCEPT: "if (searchQuery.value) ... filtered = filtered.filter(...)".
// If the API returns filtered results, local filter is redundant.
// I will just pass `notes` to NoteList. The API handles it.
// The original code passed `filteredNotes` to the template.
// I should rely on `notes.value` since `fetchNotes` is called on changes.



const { isDark } = useTheme();

// --- STYLE: NOISE OVERLAY ---
const noiseOverlayStyle = {
    backgroundImage: `url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noiseFilter'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.65' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noiseFilter)' opacity='0.5'/%3E%3C/svg%3E")`,
};
</script>

<template>
    <div>
        <Head title="Dashboard" />
        <FloatingDock />

        <div class="min-h-screen bg-slate-50 dark:bg-[#0B1121] transition-colors duration-300 relative overflow-hidden flex flex-col h-screen">
            
            <!-- Watermark -->
            <!-- Watermark -->
            <img src="/images/logo_tarakreasi.png" class="fixed bottom-6 right-6 h-16 w-auto z-[100] pointer-events-none transition-all duration-300 dark:bg-white dark:rounded-xl dark:p-2 dark:shadow-lg shadow-sm opacity-90 hover:opacity-100">
            
            <!-- Noise Overlay (Subtle) -->
            <div class="fixed inset-0 z-0 opacity-[0.03] dark:opacity-[0.02] pointer-events-none"
                 :style="noiseOverlayStyle">
            </div>

            <!-- Ambient Background (Subtle) -->
            <div class="fixed inset-0 overflow-hidden pointer-events-none z-0"
                 style="background-image: radial-gradient(circle at 50% 0%, rgba(99, 102, 241, 0.03) 0%, transparent 50%);">
                 <!-- Primary Blob (Top Left) -->
                 <div class="absolute top-[-10%] left-[-10%] w-[40%] h-[40%] bg-indigo-500/5 blur-[120px] rounded-full animate-pulse"></div>
                 <!-- Accent Blob (Bottom Right) -->
                 <div class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-rose-500/5 blur-[120px] rounded-full animate-pulse" style="animation-delay: 2s"></div>
            </div>

            <!-- App Layout -->
            <div class="relative z-10 flex w-full h-full max-w-[1920px] mx-auto overflow-hidden">

                <!-- 1. LEFT SIDEBAR -->
                <Sidebar 
                    ref="sidebarRef"
                    :notebooks="notebooks"
                    :selected-notebook-id="selectedNotebookId"
                    :view-mode="viewMode"
                    :notebooks-loading="notebooksLoading"
                    @update:view-mode="viewMode = $event"
                    @select-notebook="onSelectNotebook"
                    @create-notebook="handleCreateNotebook"
                    @rename-notebook="handleRenameNotebook"
                    @delete-notebook="handleDeleteNotebook"
                />

                <!-- 2. MIDDLE COLUMN -->
                <NoteList 
                    v-if="viewMode === 'EDITOR'"
                    :notes="notes"
                    :selected-note-id="selectedNote?.id"
                    :search-query="searchQuery"
                    :filter-status="filterStatus"
                    :current-notebook-name="currentNotebookName"
                    :is-loading="isLoading"
                    @update:search-query="searchQuery = $event"
                    @update:filter-status="onSetFilter"
                    @select-note="handleSelectNote"
                    @create-page="handleCreatePage"
                    @request-create-space="handleRequestCreateSpace"
                    @rename-note="handleRenameNote"
                    @delete-note="handleDeleteNote"
                />

                <!-- 3. MAIN AREA -->
                <main v-if="viewMode === 'SETTINGS'" class="flex-1 flex flex-col bg-white/50 dark:bg-[#0F172A]/50 relative transition-all duration-300 backdrop-blur-sm">
                    <SettingsPanel 
                        :user="$page.props.auth.user"
                        :save-status="saveStatus"
                        @close="viewMode = 'EDITOR'" 
                        @save-profile="handleSaveProfile"
                    />
                </main>

                <EditorSection 
                    v-else
                    :note="selectedNote"
                    :save-status="saveStatus"
                    @save="handleSaveNote"
                    @delete="handleDeleteNote"
                    @request-create-space="handleRequestCreateSpace"
                />

            </div>
        </div>
    </div>
</template>

<style>
/* Custom Scrollbar for dashboard containers */
.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(156, 163, 175, 0.2);
    border-radius: 3px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(156, 163, 175, 0.4);
}

/* Hide scrollbar for horiz filters */
.no-scrollbar::-webkit-scrollbar {
    display: none;
}
.no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>