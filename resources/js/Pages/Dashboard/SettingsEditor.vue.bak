<script setup>
import { ref, watch, onMounted } from 'vue';

const props = defineProps({
    initialSettings: {
        type: Object,
        default: () => ({})
    }
});

const emit = defineEmits(['close', 'update']);

const settings = ref({ ...props.initialSettings });
const isSaving = ref(false);
const selectedPage = ref('home'); // 'home', 'about', 'portfolio', 'footer'

// Image Upload State
const fileInput = ref(null);
const uploadField = ref(null);
const isUploading = ref(false);

const saveSettings = async () => {
    isSaving.value = true;
    try {
        await window.axios.post('/api/v1/admin/settings', settings.value);
        emit('update', settings.value);
    } catch (error) {
        console.error("Failed to save settings", error);
    } finally {
        setTimeout(() => {
            isSaving.value = false;
        }, 500);
    }
};

// Inline Editing Logic
const editingState = ref({
    field: null, // e.g., 'home_hero_title'
    value: '',
    originalValue: ''
});

const startEditing = (field) => {
    if (editingState.value.field === field) return;
    
    editingState.value = {
        field,
        value: settings.value[field] || '',
        originalValue: settings.value[field] || ''
    };

    // Auto-focus logic
    setTimeout(() => {
        const input = document.getElementById(`edit-${field}`);
        if (input) input.focus();
    }, 50);
};

const saveEdit = () => {
    if (editingState.value.field) {
        settings.value[editingState.value.field] = editingState.value.value;
        saveSettings(); // Auto-save on blur/enter
    }
    cancelEditing();
};

const cancelEditing = () => {
    editingState.value = { field: null, value: '', originalValue: '' };
};

// Image Upload Logic
const triggerUpload = (field) => {
    uploadField.value = field;
    if (fileInput.value) {
        fileInput.value.click();
    }
};

const handleUpload = async (event) => {
    const file = event.target.files[0];
    if (!file || !uploadField.value) return;

    const formData = new FormData();
    formData.append('image', file);
    
    isUploading.value = true;
    
    try {
        const response = await window.axios.post('/api/v1/admin/settings/upload', formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });
        
        // Update the setting with the new image URL
        settings.value[uploadField.value] = response.data.url;
        await saveSettings();
        
    } catch (error) {
        console.error("Upload failed", error);
        alert("Failed to upload image");
    } finally {
        isUploading.value = false;
        // Reset input
        event.target.value = '';
        uploadField.value = null;
    }
};

onMounted(() => {
    if (Object.keys(settings.value).length === 0) {
        // fetchSettings();
    }
});
</script>

<template>
    <div class="flex h-full flex-col bg-[#1e293b] text-slate-300 font-mono">
        <!-- Hidden File Input for Uploads -->
        <input 
            type="file" 
            ref="fileInput" 
            class="hidden" 
            accept="image/*"
            @change="handleUpload"
        >

        <!-- Header -->
        <header class="h-14 flex items-center justify-between px-6 border-b border-white/10 shrink-0 bg-[#0f172a]">
            <div class="flex items-center gap-4">
                <button @click="$emit('close')" class="hover:text-white transition-colors">
                    <span class="material-symbols-outlined">arrow_back</span>
                </button>
                <div class="flex items-center gap-2">
                    <span class="material-symbols-outlined text-cyan-400">architecture</span>
                    <h1 class="font-bold text-white tracking-widest uppercase text-sm">Wireframe Editor</h1>
                </div>
            </div>
            <div class="flex items-center gap-4">
                 <span v-if="isSaving" class="text-xs text-yellow-500 flex items-center gap-1">
                    <span class="material-symbols-outlined text-[14px] animate-spin">sync</span>
                    Saving...
                 </span>
                 <span v-else class="text-xs text-green-500 flex items-center gap-1">
                    <span class="material-symbols-outlined text-[14px]">check_circle</span>
                    Saved
                 </span>
            </div>
        </header>

        <div class="flex-1 flex overflow-hidden">
            <!-- Sidebar / Toolbox -->
            <aside class="w-64 border-r border-white/10 flex flex-col bg-[#1e293b]">
                <div class="p-4">
                    <h3 class="text-[10px] items-center text-slate-500 font-bold uppercase mb-4 tracking-wider flex gap-2">
                        <span class="material-symbols-outlined text-[14px]">layers</span> Pages
                    </h3>
                    <nav class="space-y-1">
                        <button 
                            v-for="page in ['home', 'about', 'portfolio', 'footer']" 
                            :key="page"
                            @click="selectedPage = page"
                            class="w-full text-left px-3 py-2 rounded text-xs font-bold uppercase tracking-wider transition-all border border-transparent"
                            :class="selectedPage === page ? 'bg-cyan-500/10 text-cyan-400 border-cyan-500/50 shadow-[0_0_15px_rgba(6,182,212,0.1)]' : 'hover:bg-white/5 text-slate-400'"
                        >
                            {{ page }}
                        </button>
                    </nav>
                </div>
            </aside>

            <!-- Main Canvas -->
            <main class="flex-1 overflow-y-auto p-8 relative bg-[#0f172a] bg-[radial-gradient(#1e293b_1px,transparent_1px)] [background-size:20px_20px]">
                
                <div v-if="!selectedPage" class="h-full flex flex-col items-center justify-center opacity-30 select-none pointer-events-none">
                    <span class="material-symbols-outlined text-9xl text-slate-600 mb-4">grid_4x4</span>
                    <p class="text-xl font-thin tracking-[0.5em] text-cyan-200">SELECT A BLUEPRINT</p>
                </div>

                <div v-else class="max-w-4xl mx-auto min-h-[800px] border border-white/20 bg-[#1e293b]/50 p-8 shadow-2xl relative backdrop-blur-sm">
                    <!-- Ruler Decorations -->
                    <div class="absolute top-0 left-0 w-full h-[1px] bg-cyan-500/30"></div>
                    <div class="absolute top-0 left-0 h-full w-[1px] bg-cyan-500/30"></div>
                    <div class="absolute top-[-20px] left-0 text-[10px] text-cyan-500/50 font-mono">0px</div>

                    <!-- --- PAGE: HOME --- -->
                    <template v-if="selectedPage === 'home'">
                        <!-- Hero Section Wireframe -->
                        <div class="border-2 border-dashed border-slate-600 p-8 mb-8 relative group hover:border-cyan-500/50 transition-colors">
                            <span class="absolute top-[-10px] left-4 bg-[#1e293b] px-2 text-[10px] text-cyan-500 uppercase font-bold tracking-wider">Hero Section</span>
                            
                            <div class="text-center space-y-6">
                                <!-- Title -->
                                <div class="relative">
                                    <input 
                                        v-if="editingState.field === 'home_hero_title'"
                                        :id="'edit-home_hero_title'"
                                        v-model="editingState.value"
                                        @blur="saveEdit"
                                        @keydown.enter="saveEdit"
                                        @keydown.esc="cancelEdit"
                                        class="w-full bg-[#0f172a] border border-cyan-500 text-white text-3xl font-bold text-center p-2 outline-none"
                                    >
                                    <h1 
                                        v-else
                                        @click="startEditing('home_hero_title')"
                                        class="text-3xl font-bold text-white border border-transparent hover:border-cyan-500/50 hover:bg-white/5 cursor-pointer p-2 rounded transition-all"
                                    >
                                        {{ settings.home_hero_title || '[Hero Title]' }}
                                    </h1>
                                </div>

                                <!-- Subtitle -->
                                <div class="relative">
                                    <textarea 
                                        v-if="editingState.field === 'home_hero_subtitle'"
                                        :id="'edit-home_hero_subtitle'"
                                        v-model="editingState.value"
                                        @blur="saveEdit"
                                        @keydown.esc="cancelEdit"
                                        rows="3"
                                        class="w-full bg-[#0f172a] border border-cyan-500 text-slate-300 text-lg text-center p-2 outline-none"
                                    ></textarea>
                                    <p 
                                        v-else
                                        @click="startEditing('home_hero_subtitle')"
                                        class="text-lg text-slate-400 border border-transparent hover:border-cyan-500/50 hover:bg-white/5 cursor-pointer p-2 rounded transition-all whitespace-pre-wrap"
                                    >
                                        {{ settings.home_hero_subtitle || '[Hero Subtitle]' }}
                                    </p>
                                </div>

                                <!-- Search Mockup -->
                                <div class="max-w-md mx-auto h-12 border border-slate-600 rounded-full flex items-center px-4 text-slate-500 select-none">
                                    <span class="material-symbols-outlined mr-2">search</span>
                                    <span class="italic">{{ settings.home_search_placeholder || 'Search placeholder...' }}</span>
                                    <button 
                                        @click="startEditing('home_search_placeholder')"
                                        class="ml-auto text-[10px] text-cyan-500 hover:underline uppercase"
                                    >Edit Placeholder</button>
                                     <input 
                                        v-if="editingState.field === 'home_search_placeholder'"
                                        :id="'edit-home_search_placeholder'"
                                        v-model="editingState.value"
                                        @blur="saveEdit"
                                        @keydown.enter="saveEdit"
                                        class="absolute top-0 left-0 w-full h-full opacity-0 cursor-text"
                                    >
                                </div>
                            </div>
                        </div>

                        <!-- Newsletter Wireframe -->
                        <div class="border-2 border-dashed border-slate-600 p-8 relative group hover:border-cyan-500/50 transition-colors">
                            <span class="absolute top-[-10px] left-4 bg-[#1e293b] px-2 text-[10px] text-purple-500 uppercase font-bold tracking-wider">Newsletter Section</span>
                             <div class="flex flex-col items-center gap-4">
                                <h2 @click="startEditing('newsletter_title')" class="text-xl font-bold hover:text-cyan-400 cursor-pointer">{{ settings.newsletter_title || '[Newsletter Title]' }}</h2>
                                 <input v-if="editingState.field === 'newsletter_title'" :id="'edit-newsletter_title'" v-model="editingState.value" @blur="saveEdit" @keydown.enter="saveEdit" class="w-full bg-[#0f172a] border border-cyan-500 text-white font-bold p-1 outline-none">
                                <p @click="startEditing('newsletter_subtitle')" class="text-sm text-slate-400 cursor-pointer hover:text-cyan-400">{{ settings.newsletter_subtitle || '[Newsletter Subtitle]' }}</p>
                                <button class="px-6 py-2 bg-slate-700 text-slate-300 rounded hover:bg-slate-600">{{ settings.newsletter_button_text || 'Subscribe' }} (Edit Button Text)</button>
                             </div>
                        </div>
                    </template>

                    <!-- --- PAGE: ABOUT --- -->
                    <template v-if="selectedPage === 'about'">
                         <!-- Greeting -->
                         <div class="border-2 border-dashed border-slate-600 p-8 mb-6 relative">
                            <span class="absolute top-[-10px] left-4 bg-[#1e293b] px-2 text-[10px] text-orange-500 uppercase font-bold tracking-wider">Intro</span>
                             <h1 @click="startEditing('about_greeting')" class="text-4xl font-bold text-white mb-2 cursor-pointer hover:text-cyan-400">{{ settings.about_greeting || '[Greeting]' }}</h1>
                             <input v-if="editingState.field === 'about_greeting'" :id="'edit-about_greeting'" v-model="editingState.value" @blur="saveEdit" @keydown.enter="saveEdit" class="w-full bg-[#0f172a] border border-cyan-500 text-white text-4xl font-bold p-1 outline-none mb-2">
                             
                             <h2 @click="startEditing('about_tagline')" class="text-xl text-cyan-400 mb-4 cursor-pointer hover:text-white">{{ settings.about_tagline || '[Tagline]' }}</h2>
                             <input v-if="editingState.field === 'about_tagline'" :id="'edit-about_tagline'" v-model="editingState.value" @blur="saveEdit" @keydown.enter="saveEdit" class="w-full bg-[#0f172a] border border-cyan-500 text-cyan-400 text-xl p-1 outline-none mb-4">
                             
                             <p @click="startEditing('about_description')" class="text-slate-400 leading-relaxed cursor-pointer hover:text-white whitespace-pre-wrap">{{ settings.about_description || '[Description Paragraph]' }}</p>
                             <textarea v-if="editingState.field === 'about_description'" :id="'edit-about_description'" v-model="editingState.value" @blur="saveEdit" class="w-full bg-[#0f172a] border border-cyan-500 text-slate-300 p-2 outline-none h-32"></textarea>
                         </div>

                         <!-- Principles Grid -->
                         <div class="grid grid-cols-3 gap-4">
                            <div v-for="i in [1,2,3]" :key="i" class="border border-slate-700 p-4">
                                <span class="text-[10px] text-slate-500 uppercase block mb-2">Principle {{ i }}</span>
                                <h3 @click="startEditing(`about_principle${i}_title`)" class="font-bold text-white mb-1 cursor-pointer hover:text-cyan-400">{{ settings[`about_principle${i}_title`] || '[Title]' }}</h3>
                                <input v-if="editingState.field === `about_principle${i}_title`" :id="`edit-about_principle${i}_title`" v-model="editingState.value" @blur="saveEdit" @keydown.enter="saveEdit" class="w-full bg-[#0f172a] border border-cyan-500 text-white font-bold text-sm p-1 outline-none mb-1">
                                
                                <p @click="startEditing(`about_principle${i}_desc`)" class="text-xs text-slate-400 cursor-pointer hover:text-white">{{ settings[`about_principle${i}_desc`] || '[Description]' }}</p>
                                <textarea v-if="editingState.field === `about_principle${i}_desc`" :id="`edit-about_principle${i}_desc`" v-model="editingState.value" @blur="saveEdit" class="w-full bg-[#0f172a] border border-cyan-500 text-slate-400 text-xs p-1 outline-none"></textarea>
                            </div>
                         </div>
                    </template>

                    <!-- --- PAGE: PORTFOLIO --- -->
                    <template v-if="selectedPage === 'portfolio'">
                         <!-- Hero Image Area -->
                         <div 
                            class="border-2 border-dashed border-slate-600 h-64 relative group hover:border-cyan-500 transition-colors cursor-pointer overflow-hidden mb-8"
                            @click="triggerUpload('portfolio_hero_image')"
                        >
                             <span class="absolute top-[-10px] left-4 bg-[#1e293b] px-2 text-[10px] text-cyan-500 uppercase font-bold tracking-wider z-10">Hero Image</span>
                             
                             <img v-if="settings.portfolio_hero_image" :src="settings.portfolio_hero_image" class="w-full h-full object-cover opacity-50 grayscale group-hover:grayscale-0 transition-all">
                             <div class="absolute inset-0 flex items-center justify-center flex-col" :class="settings.portfolio_hero_image ? 'opacity-0 group-hover:opacity-100 bg-black/60' : ''">
                                <span class="material-symbols-outlined text-4xl text-cyan-500 mb-2">image</span>
                                <span class="text-xs uppercase tracking-widest text-cyan-400">Click to Upload Image</span>
                             </div>
                        </div>

                        <!-- Tagline -->
                         <div class="border border-slate-600 p-6 text-center mb-8">
                             <span class="block text-[10px] text-slate-500 uppercase mb-2">Tagline</span>
                             <div class="relative">
                                <input 
                                    v-if="editingState.field === 'portfolio_hero_tagline'"
                                    :id="'edit-portfolio_hero_tagline'"
                                    v-model="editingState.value"
                                    @blur="saveEdit"
                                    @keydown.enter="saveEdit"
                                    class="w-full bg-[#0f172a] border border-cyan-500 text-white text-xl font-bold text-center p-2 outline-none"
                                >
                                <h2 
                                    v-else
                                    @click="startEditing('portfolio_hero_tagline')"
                                    class="text-xl font-bold text-white hover:text-cyan-400 cursor-pointer"
                                >
                                    {{ settings.portfolio_hero_tagline || '[Portfolio Tagline]' }}
                                </h2>
                             </div>
                         </div>
                         
                         <!-- Title Suffix -->
                          <div class="border border-slate-600 p-4 flex justify-between items-center">
                             <span class="text-[10px] text-slate-500 uppercase">Title Suffix (e.g. "John's Digital Garden")</span>
                             <span @click="startEditing('portfolio_title_suffix')" class="text-sm font-mono text-cyan-400 cursor-pointer hover:underline">{{ settings.portfolio_title_suffix || '[Suffix]' }}</span>
                              <input v-if="editingState.field === 'portfolio_title_suffix'" :id="'edit-portfolio_title_suffix'" v-model="editingState.value" @blur="saveEdit" @keydown.enter="saveEdit" class="bg-[#0f172a] border border-cyan-500 text-cyan-400 font-mono text-sm p-1 outline-none">
                          </div>

                    </template>
                     
                     <!-- --- PAGE: FOOTER --- -->
                     <template v-if="selectedPage === 'footer'">
                        <div class="mt-20 border-t border-slate-600 pt-8 text-center">
                            <span class="block text-[10px] text-slate-500 uppercase mb-4">Footer Wireframe</span>
                            <div class="flex justify-between items-center px-8">
                                <span @click="startEditing('footer_brand_name')" class="font-bold text-white cursor-pointer hover:text-cyan-400">{{ settings.footer_brand_name || 'Brand' }}</span>
                                <input v-if="editingState.field === 'footer_brand_name'" :id="'edit-footer_brand_name'" v-model="editingState.value" @blur="saveEdit" @keydown.enter="saveEdit" class="bg-[#0f172a] border border-cyan-500 text-white font-bold p-1 outline-none w-32">

                                <span @click="startEditing('footer_copyright_text')" class="text-xs text-slate-500 cursor-pointer hover:text-white">{{ settings.footer_copyright_text || 'Copyright' }}</span>
                                <input v-if="editingState.field === 'footer_copyright_text'" :id="'edit-footer_copyright_text'" v-model="editingState.value" @blur="saveEdit" @keydown.enter="saveEdit" class="bg-[#0f172a] border border-cyan-500 text-slate-500 text-xs p-1 outline-none w-48 text-right">
                            </div>
                        </div>
                     </template>

                </div>
            </main>
        </div>
    </div>
</template>

<style scoped>
/* Scrollbar for wireframe editor */
::-webkit-scrollbar {
    width: 8px;
    height: 8px;
}
::-webkit-scrollbar-track {
    background: #0f172a;
}
::-webkit-scrollbar-thumb {
    background: #334155;
    border-radius: 4px;
}
::-webkit-scrollbar-thumb:hover {
    background: #475569;
}
</style>
