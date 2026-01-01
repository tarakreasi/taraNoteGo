<script setup>
import { ref, computed } from 'vue';
import axios from 'axios';

const props = defineProps({
    settings: {
        type: Object,
        default: () => ({})
    },
    user: {
        type: Object,
        default: () => ({})
    },
    isSaving: {
        type: Boolean,
        default: false
    },
    saveStatus: {
        type: String,
        default: ''
    }
});

const emit = defineEmits(['update:settings', 'update:user', 'save', 'save-profile']);

const localSettings = computed({
    get: () => props.settings,
    set: (val) => emit('update:settings', val)
});

const localUser = computed({
    get: () => props.user,
    set: (val) => emit('update:user', val)
});

const selectedPage = ref('home'); // 'home', 'about', 'portfolio', 'footer', 'profile'

const tabs = [
    { id: 'profile', label: 'My Profile', icon: 'account_circle' },
    { id: 'home', label: 'Home', icon: 'home' }, 
    { id: 'articles', label: 'Articles', icon: 'library_books' },
    { id: 'about', label: 'About', icon: 'person' },
    { id: 'portfolio', label: 'Portfolio', icon: 'folder_open' },
    { id: 'contact', label: 'Contact', icon: 'mail' },
    { id: 'footer', label: 'Footer', icon: 'web_stories' },
];

const save = () => {
    if (selectedPage.value === 'profile') {
        emit('save-profile', localUser.value);
    } else {
        emit('save');
    }
};

const fileInputAvatar = ref(null);
const isUploading = ref(false);

const triggerAvatarUpload = () => {
    fileInputAvatar.value.click();
};

const handleAvatarUpload = async (event) => {
    const file = event.target.files[0];
    if (!file) return;

    isUploading.value = true;
    const formData = new FormData();
    formData.append('image', file);

    try {
        // Get CSRF token from meta tag
        const token = document.querySelector('meta[name="csrf-token"]')?.getAttribute('content');
        
        // Use web route instead of API route (web session auth vs API token auth)
        const response = await axios.post('/settings/upload', formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
                'X-CSRF-TOKEN': token,
                'Accept': 'application/json',
            }
        });
        localUser.value.avatar = response.data.url;
    } catch (error) {
        console.error('Upload failed:', error);
        const errorMsg = error.response?.data?.message || error.message || 'Failed to upload image';
        alert(`Upload failed: ${errorMsg}`);
    } finally {
        isUploading.value = false;
        // Reset input so same file can be selected again if needed
        event.target.value = '';
    }
};

</script>

<template>
    <div class="flex h-full flex-col bg-slate-50/50 dark:bg-[#0F172A]/50 backdrop-blur-xl">
        <!-- Header -->
        <header class="h-16 flex items-center justify-between px-6 border-b border-gray-200 dark:border-white/5 shrink-0">
            <div class="flex items-center gap-3">
                <div class="p-2 rounded-lg bg-primary/10 text-primary dark:text-blue-400">
                    <span class="material-symbols-outlined text-[20px]">tune</span>
                </div>
                <h2 class="font-display font-bold text-lg text-gray-900 dark:text-white">
                    {{ selectedPage === 'profile' ? 'Edit Profile' : 'Site Configuration' }}
                </h2>
            </div>
            
            <div class="flex items-center gap-4">
                <span v-if="saveStatus" class="text-xs font-mono transition-colors" :class="saveStatus.includes('Saved') ? 'text-green-500' : 'text-gray-400'">
                    {{ saveStatus }}
                </span>

                <button 
                    @click="save"
                    class="flex items-center gap-2 px-4 py-2 bg-primary hover:bg-primary-dark text-white rounded-xl shadow-lg shadow-primary/20 transition-all font-medium text-sm group"
                    :disabled="isSaving"
                >
                    <span v-if="isSaving" class="material-symbols-outlined text-[18px] animate-spin">sync</span>
                    <span v-else class="material-symbols-outlined text-[18px] group-hover:scale-110 transition-transform">save</span>
                    {{ isSaving ? 'Saving...' : 'Save Changes' }}
                </button>
            </div>
        </header>

        <div class="flex-1 flex overflow-hidden">
            <!-- Sidebar Navigation -->
            <aside class="w-64 border-r border-gray-200 dark:border-white/5 flex flex-col bg-white/40 dark:bg-[#0F172A]/20 backdrop-blur-md">
                
                <!-- User Avatar (Iconic Login) -->
                <div class="p-6 flex flex-col items-center border-b border-gray-200 dark:border-white/5 mb-2">
                    <div class="relative group cursor-pointer" @click="selectedPage = 'profile'">
                        <div class="size-20 rounded-full bg-cover bg-center ring-4 ring-white dark:ring-white/10 shadow-xl transition-transform group-hover:scale-105"
                             :style="{ backgroundImage: `url(${localUser.avatar || 'https://placehold.co/150'})` }">
                        </div>
                        <div class="absolute bottom-0 right-0 p-1.5 bg-primary rounded-full text-white shadow-lg border-2 border-white dark:border-[#0F172A]">
                            <span class="material-symbols-outlined text-[14px] font-bold block">edit</span>
                        </div>
                    </div>
                    <h3 class="mt-3 font-bold text-gray-900 dark:text-white">{{ localUser.name }}</h3>
                    <p class="text-xs text-gray-500 dark:text-gray-400 truncate max-w-[180px]">{{ localUser.email }}</p>
                </div>

                <nav class="p-4 space-y-1">
                    <button 
                        v-for="tab in tabs" 
                        :key="tab.id"
                        @click="selectedPage = tab.id"
                        class="w-full flex items-center gap-3 px-4 py-3 rounded-xl text-sm font-medium transition-all duration-200 border border-transparent"
                        :class="selectedPage === tab.id 
                            ? 'bg-white shadow-sm dark:bg-white/10 text-primary dark:text-white border-gray-100 dark:border-white/5' 
                            : 'text-gray-600 dark:text-gray-400 hover:bg-white/40 dark:hover:bg-white/5 hover:text-gray-900 dark:hover:text-gray-200'"
                    >
                        <span class="material-symbols-outlined text-[20px]" :class="selectedPage === tab.id ? 'text-primary dark:text-white' : 'text-gray-400'">{{ tab.icon }}</span>
                        {{ tab.label }}
                        <span v-if="selectedPage === tab.id" class="ml-auto w-1.5 h-1.5 rounded-full bg-accent-orange"></span>
                    </button>
                </nav>
            </aside>

            <!-- Main Content Area -->
            <main class="flex-1 overflow-y-auto custom-scrollbar p-8 bg-gray-50/30 dark:bg-[#0F172A]/30">
                <div class="max-w-3xl mx-auto pb-20">
                    
                    <!-- PROFILE SETTINGS -->
                    <div v-if="selectedPage === 'profile'" class="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">account_box</span> Basic Info
                            </h3>
                            <div class="grid grid-cols-2 gap-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Display Name</label>
                                    <input v-model="localUser.name" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Portfolio URL Slug</label>
                                    <div class="relative">
                                        <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 text-xs">/portfolio/</span>
                                        <input v-model="localUser.username" placeholder="username" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl pl-20 pr-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all font-mono">
                                    </div>
                                    <p class="text-[10px] text-gray-400">Leave blank to use name. Only letters, numbers, and dashes.</p>
                                </div>
                                <div class="space-y-2 col-span-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Title / Job</label>
                                    <input v-model="localUser.hero_tagline" placeholder="e.g. Writer & Gardener" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                </div>
                                <div class="space-y-2 col-span-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Short Bio</label>
                                    <textarea v-model="localUser.bio" rows="4" placeholder="Tell us about yourself..." class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all resize-none"></textarea>
                                </div>
                            </div>
                        </section>

                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">image</span> Portfolio Images
                            </h3>
                            <div class="space-y-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Avatar URL</label>
                                    <div class="flex gap-4 items-center">
                                         <input v-model="localUser.avatar" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all" placeholder="https://...">
                                         <div class="size-10 rounded-full bg-cover bg-center shrink-0 border border-gray-200 dark:border-white/10" :style="{ backgroundImage: `url(${localUser.avatar || 'https://placehold.co/150'})` }"></div>
                                         
                                         <!-- Upload Button -->
                                         <div class="relative">
                                            <button 
                                                @click="triggerAvatarUpload" 
                                                class="flex items-center justify-center size-10 rounded-xl bg-gray-100 dark:bg-white/5 hover:bg-primary/10 hover:text-primary transition-colors border border-gray-200 dark:border-white/10 text-gray-500"
                                                title="Upload Image"
                                                :disabled="isUploading"
                                            >
                                                <span v-if="isUploading" class="material-symbols-outlined text-[20px] animate-spin">sync</span>
                                                <span v-else class="material-symbols-outlined text-[20px]">upload</span>
                                            </button>
                                            <input 
                                                ref="fileInputAvatar" 
                                                type="file" 
                                                class="hidden" 
                                                accept="image/*"
                                                @change="handleAvatarUpload"
                                            >
                                         </div>
                                    </div>
                                    <p class="text-xs text-gray-400">Currently only supporting external URLs.</p>
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Hero Cover Image</label>
                                    <input v-model="localUser.hero_image" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all" placeholder="https://...">
                                </div>
                                <div v-if="localUser.hero_image" class="w-full h-40 rounded-xl bg-cover bg-center border border-gray-200 dark:border-white/10" :style="{ backgroundImage: `url(${localUser.hero_image})` }"></div>
                            </div>
                        </section>
                    </div>

                    <!-- WELCOME (LANDING) SETTINGS -> RENAMED TO HOME -->
                    <div v-if="selectedPage === 'home'" class="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <div class="absolute top-0 left-0 w-1 h-full bg-gradient-to-b from-primary to-transparent opacity-0 group-hover:opacity-100 transition-opacity rounded-l-2xl"></div>
                            
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">home</span> Home (Landing)
                            </h3>
                            
                            <div class="space-y-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Hero Title</label>
                                    <input 
                                        v-model="localSettings.welcome_hero_title" 
                                        type="text" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all placeholder:text-gray-400"
                                        placeholder="e.g. Deep dives into Code & Creation"
                                    >
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Hero Subtitle</label>
                                    <textarea 
                                        v-model="localSettings.welcome_hero_subtitle" 
                                        rows="3" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all resize-none placeholder:text-gray-400"
                                        placeholder="Brief description..."
                                    ></textarea>
                                </div>
                            </div>
                        </section>

                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">grid_view</span> Digital Notes Section
                            </h3>
                            <div class="space-y-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Section Title</label>
                                    <input v-model="localSettings.welcome_notes_title" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Section Subtitle</label>
                                    <textarea v-model="localSettings.welcome_notes_subtitle" rows="2" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all resize-none"></textarea>
                                </div>
                            </div>
                        </section>

                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">feature_search</span> Stats & Features
                            </h3>
                            <div class="grid grid-cols-2 gap-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Feature 1 Title (Latest)</label>
                                    <input v-model="localSettings.welcome_feature_1_title" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Feature 2 Title (Garden)</label>
                                    <input v-model="localSettings.welcome_feature_2_title" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                </div>
                                <div class="space-y-2 col-span-2">
                                     <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Feature 3 Title (Availability)</label>
                                    <input v-model="localSettings.welcome_feature_3_title" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                </div>
                                <div class="space-y-2 col-span-2">
                                     <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Availability Status Text</label>
                                     <textarea v-model="localSettings.welcome_availability_text" rows="2" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all resize-none"></textarea>
                                </div>
                            </div>
                        </section>

                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">mail</span> Newsletter
                            </h3>
                            <div class="space-y-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Section Title</label>
                                    <input v-model="localSettings.newsletter_title" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Description</label>
                                    <textarea v-model="localSettings.newsletter_subtitle" rows="2" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all resize-none"></textarea>
                                </div>
                            </div>
                        </section>
                    </div>

                    <!-- ABOUT SETTINGS -->
                    <div v-if="selectedPage === 'about'" class="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">person</span> Introduction
                            </h3>
                            <div class="space-y-6">
                                <div class="grid grid-cols-2 gap-6">
                                    <div class="space-y-2">
                                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Greeting</label>
                                        <input v-model="localSettings.about_greeting" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                    </div>
                                    <div class="space-y-2">
                                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Role / Tagline</label>
                                        <input v-model="localSettings.about_tagline" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                    </div>
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Bio Description</label>
                                    <textarea v-model="localSettings.about_description" rows="6" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all resize-none"></textarea>
                                </div>
                            </div>
                        </section>

                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">diamond</span> Core Principles
                            </h3>
                            <div class="grid gap-4">
                                <div v-for="i in 3" :key="i" class="p-4 rounded-xl bg-gray-50/50 dark:bg-white/5 border border-gray-200 dark:border-white/5 hover:bg-white dark:hover:bg-white/10 transition-colors">
                                    <span class="text-[10px] font-bold uppercase text-primary mb-3 block tracking-wider">Principle {{ i }}</span>
                                    <div class="space-y-3">
                                        <input v-model="localSettings[`about_principle${i}_title`]" :placeholder="`Principle ${i} Title`" type="text" class="w-full bg-transparent border-b border-gray-200 dark:border-white/10 px-0 py-2 text-sm font-bold focus:ring-0 focus:border-primary transition-all placeholder:font-normal">
                                        <textarea v-model="localSettings[`about_principle${i}_desc`]" :placeholder="`Description for principle ${i}`" rows="2" class="w-full bg-transparent border-none px-0 py-0 text-sm text-gray-500 focus:ring-0 resize-none"></textarea>
                                    </div>
                                </div>
                            </div>
                        </section>
                    </div>

                    <!-- ARTICLES (DIGITAL GARDEN) SETTINGS -->
                    <!-- ARTICLES (DIGITAL GARDEN) SETTINGS -->
                    <div v-if="selectedPage === 'articles'" class="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">library_books</span> Articles Page
                            </h3>
                            <div class="space-y-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Page Title</label>
                                    <input 
                                        v-model="localSettings.articles_hero_title" 
                                        type="text" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all placeholder:text-gray-400"
                                        placeholder="e.g. Explore Knowledge"
                                    >
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Page Subtitle</label>
                                    <textarea 
                                        v-model="localSettings.articles_hero_subtitle" 
                                        rows="3" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all resize-none placeholder:text-gray-400"
                                        placeholder="Brief description..."
                                    ></textarea>
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Search Placeholder</label>
                                    <div class="relative">
                                        <span class="absolute left-4 top-1/2 -translate-y-1/2 material-symbols-outlined text-gray-400 text-[18px]">search</span>
                                        <input 
                                            v-model="localSettings.articles_search_placeholder" 
                                            type="text" 
                                            class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl pl-11 pr-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all placeholder:text-gray-400"
                                            placeholder="Search text..."
                                        >
                                    </div>
                                </div>
                            </div>
                        </section>
                    </div>

                    <!-- PORTFOLIO SETTINGS -->
                    <div v-if="selectedPage === 'portfolio'" class="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">folder_open</span> Portfolio Header
                            </h3>
                            <div class="space-y-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Title Suffix</label>
                                    <input 
                                        v-model="localSettings.portfolio_title_suffix" 
                                        type="text" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                                        placeholder="e.g. John's Digital Garden"
                                    >
                                    <p class="text-xs text-gray-400">Appears in browser tab title.</p>
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Hero Tagline</label>
                                    <input 
                                        v-model="localSettings.portfolio_hero_tagline" 
                                        type="text" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                                        placeholder="e.g. Selected Works"
                                    >
                                </div>
                            </div>
                        </section>
                    </div>

                    <!-- CONTACT SETTINGS -->
                    <div v-if="selectedPage === 'contact'" class="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">mail</span> Contact Information
                            </h3>
                            <div class="space-y-6">
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Contact Email</label>
                                    <input 
                                        v-model="localSettings.contact_email" 
                                        type="email" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                                        placeholder="hello@example.com"
                                    >
                                </div>
                                <div class="space-y-2">
                                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Office Location</label>
                                    <input 
                                        v-model="localSettings.contact_location" 
                                        type="text" 
                                        class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                                        placeholder="e.g. Jakarta, Indonesia"
                                    >
                                </div>
                            </div>
                        </section>
                    </div>

                    <!-- FOOTER SETTINGS -->
                    <div v-if="selectedPage === 'footer'" class="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <section class="group relative bg-white dark:bg-[#1E293B]/50 border border-gray-200 dark:border-white/5 rounded-2xl p-8 shadow-sm hover:shadow-md hover:border-primary/20 transition-all duration-300 backdrop-blur-sm">
                            <h3 class="text-xs font-bold uppercase text-gray-400 tracking-widest mb-6 flex items-center gap-2">
                                <span class="material-symbols-outlined text-[16px]">web_stories</span> Global Footer
                            </h3>
                            <div class="space-y-6">
                                <div class="grid grid-cols-2 gap-6">
                                    <div class="space-y-2">
                                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Brand Name</label>
                                        <input v-model="localSettings.footer_brand_name" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                    </div>
                                    <div class="space-y-2">
                                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Copyright Text</label>
                                        <input v-model="localSettings.footer_copyright_text" type="text" class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all">
                                    </div>
                                    <div class="space-y-2 col-span-2 md:col-span-1">
                                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Status Text</label>
                                        <input 
                                            v-model="localSettings.footer_status_text" 
                                            type="text" 
                                            class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                                            placeholder="e.g. All systems operational"
                                        >
                                    </div>
                                    <div class="space-y-2 md:col-span-1">
                                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Twitter / X URL</label>
                                        <input 
                                            v-model="localSettings.footer_social_twitter" 
                                            type="text" 
                                            class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                                            placeholder="https://x.com/username"
                                        >
                                    </div>
                                    <div class="space-y-2 md:col-span-1">
                                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">GitHub URL</label>
                                        <input 
                                            v-model="localSettings.footer_social_github" 
                                            type="text" 
                                            class="w-full bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl px-4 py-3 text-sm focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                                            placeholder="https://github.com/username"
                                        >
                                    </div>
                                </div>
                            </div>
                        </section>
                    </div>

                </div>
            </main>
        </div>
    </div>
</template>

<style scoped>
/* Custom Scrollbar */
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
</style>
