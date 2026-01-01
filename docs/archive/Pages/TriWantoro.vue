<script setup>
import { Head, usePage } from '@inertiajs/vue3';
import { ref, onMounted, computed } from "vue";
import { useTheme } from "@/composables/useTheme";
import EditableText from "@/Components/EditableText.vue";
import EditableImage from "@/Components/EditableImage.vue";
// import EditableLink from "@/Components/EditableLink.vue";
import FloatingDock from "@/Components/FloatingDock.vue";

const page = usePage();
const user = computed(() => page.props.auth?.user);

// Hero Image Toggle Logic
const heroImageIndex = ref(0);
const activeHeroKey = computed(() => heroImageIndex.value === 0 ? 'hero_image' : 'hero_image_2');
const toggleHeroImage = () => {
    heroImageIndex.value = (heroImageIndex.value + 1) % 2;
};

const props = defineProps({
    settings: {
        type: Object,
        default: () => ({})
    }
});

// Theme Logic
const { isDark, toggleTheme } = useTheme();
const theme = computed(() => isDark.value ? 'dark' : 'light');

// Helper to safely get setting or fallback
const getSetting = (key, fallback = '') => props.settings[key] || fallback;

// Inject JSON-LD (Updated to use settings if updated, but for SEO static load is usually fine. 
// We can make it dynamic if we really want, but let's stick to the script for now.)
onMounted(() => {
    const script = document.createElement('script');
    script.type = 'application/ld+json';
    script.text = JSON.stringify({
        "@context": "https://schema.org",
        "@type": "Person",
        "name": "Tri Wantoro",
        "jobTitle": getSetting('hero_role', 'Full Stack Engineer'),
        "url": "https://tarakreasi.com",
        "email": getSetting('contact_email', 'ajarsinau@gmail.com'),
        "image": getSetting('hero_image') ? (window.location.origin + getSetting('hero_image')) : (window.location.origin + '/favicon.ico'),
        "sameAs": [
            getSetting('contact_linkedin', 'https://linkedin.com/in/twantoro'),
            getSetting('github_card_url', 'https://github.com/tarakreasi')
        ]
    });
    document.head.appendChild(script);
});

// Data Mappers for Lists (to simplify template)
const timelineItems = [1, 2, 3, 4].map(i => ({
    id: i,
    yearKey: `timeline_${i}_year`,
    titleKey: `timeline_${i}_title`,
    descKey: `timeline_${i}_desc`
}));

const techStackItems = [1, 2, 3, 4, 5].map(i => ({
    id: i,
    nameKey: `tech_${i}_name`,
    iconKey: `tech_${i}_icon` // Icon URL is stored in setting
}));

</script>

<template>
    <Head>
        <title>Tri Wantoro — {{ getSetting('hero_role') }}</title>
        <!-- Meta tags preserved but could be dynamic too -->
        <meta name="description" content="Tri Wantoro - Full Stack Engineer. Bridging Physical Ops with Digital Solutions." />
    </Head>

    <div :class="theme" class="min-h-screen transition-colors duration-500 ease-in-out bg-[#F5F7FA] dark:bg-[#0F172A] text-slate-800 dark:text-slate-100 font-['Noto_Sans_JP'] selection:bg-blue-100 selection:text-blue-900 scroll-smooth">
        
        <!-- Vertical Dock (Overrides global layout/nav if needed, or we just render it here) -->
        <!-- Note: If Layout is used, we might need to hide default dock. Assuming TriWantoro is standalone page layout. -->
        <FloatingDock :vertical="true" />

        <!-- HEADER / LOGO (Moved to Hero) -->

        <!-- HERO -->
        <section class="max-w-6xl mx-auto px-6 pt-24 pb-16 md:pt-28 md:pb-20 grid md:grid-cols-12 gap-10 items-center pl-24 transition-all"> 
            <!-- Added pl-24 to account for vertical dock on left -->
            
            <div class="md:col-span-7 animate-fadeIn space-y-6">
                <!-- Name -->
                <h1 class="text-3xl font-bold tracking-tighter text-slate-900 dark:text-white">
                    Tri Wantoro<span class="text-blue-600 dark:text-blue-400">.</span>
                </h1>

                <!-- Role Badge -->
                <div class="inline-flex items-center px-3 py-1 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-700 dark:text-blue-300 text-xs font-bold uppercase tracking-wider mb-2 border border-blue-200 dark:border-blue-800">
                    <EditableText setting-key="hero_role" :initial-value="getSetting('hero_role')" />
                </div>
                
                <!-- Main Title -->
                <h2 class="text-4xl md:text-5xl lg:text-6xl font-extrabold leading-tight text-slate-900 dark:text-white tracking-tight">
                    <EditableText setting-key="hero_title_1" :initial-value="getSetting('hero_title_1')" /> 
                    <EditableText setting-key="hero_title_2" :initial-value="getSetting('hero_title_2')" class="text-slate-500 dark:text-slate-400" /><br />
                    with <EditableText setting-key="hero_title_3" :initial-value="getSetting('hero_title_3')" class="bg-clip-text text-transparent bg-gradient-to-r from-blue-600 to-indigo-600" />.
                </h2>

                <!-- Description -->
                <div class="space-y-3 max-w-lg">
                    <p class="text-lg md:text-xl font-medium text-slate-700 dark:text-slate-300 leading-relaxed">
                        <EditableText setting-key="hero_desc_1" :initial-value="getSetting('hero_desc_1')" tag="p" />
                    </p>
                    <p class="text-slate-500 dark:text-slate-400 leading-relaxed text-base">
                        <EditableText setting-key="hero_desc_2" :initial-value="getSetting('hero_desc_2')" tag="p" />
                    </p>
                </div>
                
                <!-- Links -->
                <div class="pt-6 flex flex-wrap gap-4">
                    <a 
                        :href="getSetting('hero_btn_work_url', '#work')"
                        class="inline-block px-12 py-4 text-lg font-bold rounded-xl hover:brightness-110 shadow-lg hover:shadow-xl transform hover:-translate-y-1 transition-all duration-200 bg-blue-600 text-white"
                        style="background-color: #2563eb; color: white;" 
                    >
                        {{ getSetting('hero_btn_work_text', 'View Work') }}
                    </a>
                    
                    <a 
                        :href="getSetting('hero_btn_thoughts_url', '/articles')"
                        class="inline-block px-12 py-4 text-lg font-bold rounded-xl hover:brightness-110 shadow-lg hover:shadow-xl transform hover:-translate-y-1 transition-all duration-200 bg-indigo-500 text-white"
                        style="background-color: #6366f1; color: white;"
                    >
                        {{ getSetting('hero_btn_thoughts_text', 'Read Notes') }}
                    </a>
                </div>
            </div>

            <!-- Hero Image -->
            <div class="md:col-span-5 flex flex-col items-center md:items-end">
                <div class="relative group w-full max-w-[360px]">
                    <div class="absolute -inset-4 bg-gradient-to-br from-blue-100 to-slate-100 dark:from-slate-800 dark:to-slate-900 rounded-[2rem] -z-10 transform rotate-3 group-hover:rotate-0 transition-transform duration-500"></div>
                    
                    <!-- ADMIN VIEW: Split Dual Layout -->
                    <div v-if="user" class="grid grid-cols-2 gap-4">
                        <!-- Image 1 -->
                        <div class="relative rounded-2xl overflow-hidden shadow-xl bg-slate-200 dark:bg-slate-800 aspect-[3/4] group/img1">
                            <EditableImage 
                                setting-key="hero_image" 
                                :initial-src="getSetting('hero_image')"
                                alt="Main Photo"
                                img-class="w-full h-full object-cover grayscale group-hover/img1:grayscale-0 transition duration-500"
                            />
                            <div class="absolute bottom-0 left-0 right-0 bg-black/50 text-white text-[10px] uppercase font-bold tracking-wider py-1.5 text-center backdrop-blur-md">
                                Main Photo
                            </div>
                        </div>
                        <!-- Image 2 -->
                        <div class="relative rounded-2xl overflow-hidden shadow-xl bg-slate-200 dark:bg-slate-800 aspect-[3/4] group/img2">
                            <EditableImage 
                                setting-key="hero_image_2" 
                                :initial-src="getSetting('hero_image_2') || getSetting('hero_image')"
                                alt="Alt Photo"
                                img-class="w-full h-full object-cover grayscale group-hover/img2:grayscale-0 transition duration-500"
                            />
                            <div class="absolute bottom-0 left-0 right-0 bg-black/50 text-white text-[10px] uppercase font-bold tracking-wider py-1.5 text-center backdrop-blur-md">
                                Alt Photo
                            </div>
                        </div>
                    </div>

                    <!-- GUEST VIEW: Single Toggleable Layout -->
                    <div 
                        v-else
                        class="relative rounded-2xl overflow-hidden shadow-2xl bg-slate-200 dark:bg-slate-800 aspect-[3/4] cursor-pointer"
                        @click="toggleHeroImage"
                    >
                        <img 
                            :src="getSetting(activeHeroKey) || getSetting('hero_image')" 
                            alt="Hero Image"
                            class="w-full h-full object-cover grayscale hover:grayscale-0 transition duration-700 ease-out transform hover:scale-105"
                        />
                        <div class="absolute inset-0 bg-blue-900/10 dark:bg-blue-900/20 hover:opacity-0 transition-opacity duration-500 pointer-events-none"></div>
                    </div>
                </div>
                
                <!-- Download CV -->
                <div class="mt-8 flex justify-center w-full max-w-[360px]">
                    <a 
                        :href="getSetting('hero_cv_url', '#')"
                        target="_blank"
                        class="group relative inline-flex items-center gap-3 px-8 py-3.5 rounded-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 font-bold text-sm tracking-widest uppercase hover:border-blue-500 dark:hover:border-blue-400 hover:text-blue-600 dark:hover:text-blue-400 transition-all duration-300 shadow-lg hover:shadow-xl hover:-translate-y-0.5"
                    >
                        <span class="material-symbols-outlined text-xl group-hover:translate-y-0.5 transition-transform">download</span>
                        <span>{{ getSetting('hero_cv_text', 'Download CV') }}</span>
                    </a>
                </div>
            </div>
        </section>

        <!-- TIMELINE -->
        <section id="timeline" class="py-16 bg-white dark:bg-slate-900 border-y border-slate-100 dark:border-slate-800 pl-24">
            <div class="max-w-4xl mx-auto px-6">
                <div class="mb-12 text-center md:text-left">
                    <h3 class="text-3xl font-bold text-slate-900 dark:text-white">
                        <EditableText setting-key="timeline_heading" :initial-value="getSetting('timeline_heading')" />
                    </h3>
                    <p class="text-slate-500 dark:text-slate-400 mt-2">
                        <EditableText setting-key="timeline_subheading" :initial-value="getSetting('timeline_subheading')" />
                    </p>
                </div>

                <div class="relative border-l-2 border-slate-100 dark:border-slate-800 pl-16 space-y-12">
                    <div v-for="(item, index) in timelineItems" :key="item.id" class="animate-slideUp relative group" :style="{ animationDelay: (index * 100) + 'ms' }">
                        <div class="absolute -left-[67px] top-1.5 h-7 w-7 rounded-full border-[6px] border-white dark:border-slate-900 bg-blue-600 dark:bg-blue-500 z-10 shadow-sm transition-transform group-hover:scale-110"></div>
                        
                        <div>
                             <div class="flex items-center gap-3 mb-3">
                                <span class="font-mono text-sm font-bold text-blue-600 dark:text-blue-400 bg-blue-50 dark:bg-blue-900/30 px-2 py-0.5 rounded">
                                    <EditableText :setting-key="item.yearKey" :initial-value="getSetting(item.yearKey)" />
                                </span>
                                <h4 class="text-xl font-bold text-slate-900 dark:text-white">
                                    <EditableText :setting-key="item.titleKey" :initial-value="getSetting(item.titleKey)" />
                                </h4>
                             </div>
                            <p class="text-slate-600 dark:text-slate-300 text-base leading-relaxed max-w-2xl">
                                <EditableText :setting-key="item.descKey" :initial-value="getSetting(item.descKey)" tag="p" />
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- PORTFOLIO -->
        <section id="work" class="max-w-6xl mx-auto px-6 py-16 pl-24">
            <div class="mb-12">
                <h3 class="text-3xl font-bold text-slate-900 dark:text-white">
                    <EditableText setting-key="work_heading" :initial-value="getSetting('work_heading')" />
                </h3>
                <p class="text-slate-500 dark:text-slate-400 mt-2">
                    <EditableText setting-key="work_subheading" :initial-value="getSetting('work_subheading')" />
                </p>
            </div>
            
            <div class="space-y-6">
                <!-- Project 1 (TaraNote) -->
                <div class="p-12 rounded-[2rem] bg-gradient-to-br from-blue-50/80 to-indigo-50/80 dark:from-blue-900/30 dark:to-indigo-900/30 backdrop-blur-sm border-2 border-blue-200/60 dark:border-blue-700/40 hover:border-blue-300/80 dark:hover:border-blue-600/60 hover:shadow-2xl transition-all duration-300 group">
                    <div class="grid md:grid-cols-2 gap-12 items-center">
                        <div class="space-y-6 pr-4">
                            <div class="flex items-center gap-4 mb-2">
                                <span class="material-symbols-outlined text-4xl text-blue-600 dark:text-blue-400">edit_note</span>
                                <h4 class="text-3xl font-bold text-slate-900 dark:text-white">
                                    <EditableText setting-key="project_1_name" :initial-value="getSetting('project_1_name')" />
                                </h4>
                            </div>
                            <p class="text-base text-slate-600 dark:text-slate-300 leading-relaxed">
                                <EditableText setting-key="project_1_desc" :initial-value="getSetting('project_1_desc')" />
                            </p>
                            <div class="flex flex-wrap gap-2 pt-2 text-xs font-bold text-blue-600 dark:text-blue-400">
                                <EditableText setting-key="project_1_tags" :initial-value="getSetting('project_1_tags')" />
                            </div>
                            <div class="mt-6">
                                <a 
                                    :href="getSetting('project_1_url', '/taraNote/taraNote')"
                                    class="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-blue-100 dark:bg-blue-900/40 text-blue-700 dark:text-blue-300 font-bold text-sm tracking-wide hover:bg-blue-200 dark:hover:bg-blue-900/60 transition-all border border-blue-200 dark:border-blue-700/50"
                                >
                                    {{ getSetting('project_1_cta', 'View Case Study') }}
                                </a>
                            </div>
                        </div>
                        
                        <!-- Visual -->
                        <div class="relative h-64 bg-white dark:bg-slate-800 rounded-2xl overflow-hidden border border-slate-200 dark:border-slate-700 shadow-lg transform group-hover:scale-105 transition-transform duration-500">
                            <EditableImage setting-key="project_1_image" :initial-src="getSetting('project_1_image')">
                                <!-- Fallback Visual -->
                                <div class="absolute inset-0 p-6 flex gap-3 opacity-60">
                                    <div class="w-1/4 bg-slate-200 dark:bg-slate-600 rounded-lg"></div>
                                    <div class="w-3/4 space-y-3">
                                        <div class="h-6 bg-slate-200 dark:bg-slate-600 rounded w-full"></div>
                                        <div class="h-32 bg-slate-200 dark:bg-slate-600 rounded w-full"></div>
                                        <div class="grid grid-cols-2 gap-3">
                                            <div class="h-24 bg-slate-200 dark:bg-slate-600 rounded"></div>
                                            <div class="h-24 bg-slate-200 dark:bg-slate-600 rounded"></div>
                                        </div>
                                    </div>
                                </div>
                            </EditableImage>
                        </div>
                    </div>
                </div>

                <!-- Projects 2 & 3 -->
                <div class="grid md:grid-cols-2 gap-6">
                    <!-- Project 2 (TaraTask) -->
                    <div class="p-8 rounded-[2rem] bg-gradient-to-br from-orange-50/80 to-pink-50/80 dark:from-orange-900/30 dark:to-pink-900/30 backdrop-blur-sm border-2 border-orange-200/60 dark:border-orange-700/40 hover:border-orange-300/80 dark:hover:border-orange-600/60 hover:shadow-xl transition-all duration-300 group flex flex-col h-[420px]">
                        <div class="relative z-10 mb-auto">
                            <div class="flex items-center gap-3 mb-4">
                                <span class="material-symbols-outlined text-3xl text-orange-600 dark:text-orange-400">dashboard</span>
                                <h4 class="text-2xl font-bold text-slate-900 dark:text-white">
                                    <EditableText setting-key="project_2_name" :initial-value="getSetting('project_2_name')" />
                                </h4>
                            </div>
                            <p class="text-slate-600 dark:text-slate-300 leading-relaxed mb-4">
                                <EditableText setting-key="project_2_desc" :initial-value="getSetting('project_2_desc')" />
                            </p>
                            <div class="text-xs font-bold text-orange-800 dark:text-orange-200 bg-orange-100 dark:bg-orange-900/50 px-3 py-1.5 rounded-full inline-block">
                                <EditableText setting-key="project_2_tags" :initial-value="getSetting('project_2_tags')" />
                            </div>
                            <div class="mt-4">
                                <a 
                                    :href="getSetting('project_2_url', '/taraTask')"
                                    class="inline-flex items-center gap-1 text-sm font-bold text-orange-600 dark:text-orange-400 hover:text-orange-700 dark:hover:text-orange-300 transition-colors"
                                >
                                    {{ getSetting('project_2_cta', 'View Project') }}
                                </a>
                            </div>
                        </div>
                        
                        <div class="relative h-36 w-full mt-4 bg-white dark:bg-slate-800 rounded-xl overflow-hidden border border-orange-200 dark:border-orange-700 shadow-sm transform group-hover:scale-105 transition-transform duration-500">
                             <EditableImage setting-key="project_2_image" :initial-src="getSetting('project_2_image')">
                                <div class="absolute inset-0 p-4 flex gap-2 opacity-60">
                                    <div class="w-1/3 bg-orange-200 dark:bg-orange-800/50 rounded p-2 space-y-1">
                                        <div class="h-3 bg-orange-300 dark:bg-orange-700 rounded"></div>
                                        <div class="h-8 bg-white dark:bg-slate-700 rounded"></div>
                                        <div class="h-8 bg-white dark:bg-slate-700 rounded"></div>
                                    </div>
                                    <div class="w-1/3 bg-pink-200 dark:bg-pink-800/50 rounded p-2 space-y-1">
                                        <div class="h-3 bg-pink-300 dark:bg-pink-700 rounded"></div>
                                        <div class="h-8 bg-white dark:bg-slate-700 rounded"></div>
                                    </div>
                                    <div class="w-1/3 bg-purple-200 dark:bg-purple-800/50 rounded p-2 space-y-1">
                                        <div class="h-3 bg-purple-300 dark:bg-purple-700 rounded"></div>
                                        <div class="h-8 bg-white dark:bg-slate-700 rounded"></div>
                                    </div>
                                </div>
                             </EditableImage>
                        </div>
                    </div>

                    <!-- Project 3 (TaraQueue) -->
                    <div class="p-8 rounded-[2rem] bg-gradient-to-br from-green-50/80 to-teal-50/80 dark:from-green-900/30 dark:to-teal-900/30 backdrop-blur-sm border-2 border-green-200/60 dark:border-green-700/40 hover:border-green-300/80 dark:hover:border-green-600/60 hover:shadow-xl transition-all duration-300 group flex flex-col h-[420px]">
                        <div class="relative z-10 mb-auto">
                            <div class="flex items-center gap-3 mb-4">
                                <span class="material-symbols-outlined text-3xl text-green-600 dark:text-green-400">queue_music</span>
                                <h4 class="text-2xl font-bold text-slate-900 dark:text-white">
                                    <EditableText setting-key="project_3_name" :initial-value="getSetting('project_3_name')" />
                                </h4>
                            </div>
                            <p class="text-slate-600 dark:text-slate-300 leading-relaxed mb-4">
                                <EditableText setting-key="project_3_desc" :initial-value="getSetting('project_3_desc')" />
                            </p>
                            <div class="text-xs font-bold text-green-800 dark:text-green-200 bg-green-100 dark:bg-green-900/50 px-3 py-1.5 rounded-full inline-block">
                                <EditableText setting-key="project_3_tags" :initial-value="getSetting('project_3_tags')" />
                            </div>
                            <div class="mt-4">
                                <a 
                                    :href="getSetting('project_3_url', '/taraQueue')"
                                    class="inline-flex items-center gap-1 text-sm font-bold text-green-600 dark:text-green-400 hover:text-green-700 dark:hover:text-green-300 transition-colors"
                                >
                                    {{ getSetting('project_3_cta', 'View Demo') }}
                                </a>
                            </div>
                        </div>

                        <div class="relative h-36 w-full mt-4 bg-white dark:bg-slate-800 rounded-xl overflow-hidden border border-green-200 dark:border-green-700 shadow-sm transform group-hover:scale-105 transition-transform duration-500">
                             <EditableImage setting-key="project_3_image" :initial-src="getSetting('project_3_image')">
                                <div class="absolute inset-0 p-6 flex flex-col justify-center opacity-60">
                                    <div class="space-y-2">
                                        <div class="flex items-center gap-2 bg-green-200 dark:bg-green-800/50 rounded p-2">
                                            <div class="w-8 h-8 bg-green-400 dark:bg-green-600 rounded-full"></div>
                                            <div class="flex-1 h-3 bg-green-300 dark:bg-green-700 rounded"></div>
                                        </div>
                                        <div class="flex items-center gap-2 bg-teal-200 dark:bg-teal-800/50 rounded p-2">
                                            <div class="w-8 h-8 bg-teal-400 dark:bg-teal-600 rounded-full"></div>
                                            <div class="flex-1 h-3 bg-teal-300 dark:bg-teal-700 rounded"></div>
                                        </div>
                                    </div>
                                </div>
                             </EditableImage>
                        </div>
                    </div>
                </div>

                <!-- GitHub & Tech -->
                <div class="grid md:grid-cols-2 gap-6">
                    <!-- GitHub Card (Refactored for Image Support) -->
                    <div class="relative p-8 rounded-[2rem] bg-slate-900 dark:bg-slate-900 text-white flex flex-col items-center justify-center gap-6 hover:scale-[1.02] hover:shadow-2xl transition duration-300 h-[300px] md:h-[380px] group overflow-hidden border border-slate-800 dark:border-slate-700">
                        
                        <!-- Background Image (Editable) -->
                        <div class="absolute inset-0 z-0">
                            <EditableImage 
                                setting-key="github_card_image" 
                                :initial-src="getSetting('github_card_image')"
                                img-class="w-full h-full object-cover opacity-50 group-hover:opacity-100 transition-opacity duration-500"
                            >
                                <!-- Default Gradient if no image -->
                                <div class="w-full h-full bg-gradient-to-br from-slate-800 to-slate-950 dark:from-slate-900 dark:to-black"></div>
                            </EditableImage>
                        </div>

                        <!-- Dark Overlay for readability -->
                        <div class="absolute inset-0 bg-black/30 dark:bg-white/10 z-0 pointer-events-none"></div>

                         <!-- Content Overlay -->
                         <div class="relative z-10 text-center flex flex-col items-center">
                            <!-- Title Link -->
                            <a 
                                :href="getSetting('github_card_url', '#')"
                                target="_blank"
                                class="text-6xl md:text-7xl font-bold block mb-4 hover:text-blue-400 dark:hover:text-blue-600 transition-colors drop-shadow-lg"
                            >
                                {{ getSetting('github_card_title', 'GitHub') }}
                            </a>
                            
                            <!-- CTA Link -->
                             <a
                                :href="getSetting('github_card_url', '#')"
                                target="_blank"
                                class="inline-flex items-center gap-2 text-sm font-bold uppercase tracking-widest hover:text-blue-400 dark:hover:text-blue-600 transition-colors bg-black/20 dark:bg-white/20 px-4 py-2 rounded-full backdrop-blur-md"
                             >
                                <span>{{ getSetting('github_card_cta', 'Follow') }}</span>
                                <span class="material-symbols-outlined text-lg">arrow_outward</span>
                             </a>
                        </div>
                    </div>
                    
                    <!-- RE-IMPLEMENTING GITHUB CARD TO BE CLEANER -->
                    <!-- I will make the whole card a normal link to the URL setting, but add small edit controls if admin? 
                         Or just use EditableLink for the whole block? 
                         EditableLink replaces the Anchor. So the content inside is the "Text".
                         That might destroy the layout.
                         
                         BETTER: Make the card a <div>. Inside, put an Anchor for the whole area?
                         Or just EditableText for Title.
                         And the URL? Maybe an "Edit URL" button in corner for admin.
                         
                         Let's stick to simple:
                         Card is <a>. Href computed from setting.
                         Title is EditableText (click stopPropagation).
                         CTA is EditableText.
                         
                         But how to edit the href? 
                         Check the "Edit Link" component I made. It edits both Text and Url.
                         I'll use it for the CTA button essentially?
                         
                         Let's just make the CTA button an EditableLink.
                    -->
                    
                    <!-- Tech Stack -->
                    <div class="p-8 rounded-[2rem] bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 h-[300px] md:h-[380px] flex flex-col justify-center items-center relative overflow-hidden">
                        <div class="absolute top-0 w-full h-1 bg-gradient-to-r from-blue-500 via-indigo-500 to-purple-500"></div>
                        <h4 class="text-sm font-bold text-slate-400 mb-10 uppercase tracking-[0.3em]">
                            <EditableText setting-key="tech_stack_heading" :initial-value="getSetting('tech_stack_heading')" />
                        </h4>
                        <div class="flex items-center justify-center gap-8 md:gap-10 flex-wrap relative z-10">
                            <div v-for="tech in techStackItems" :key="tech.id" class="flex flex-col items-center">
                                <EditableImage :setting-key="tech.iconKey" :initial-src="getSetting(tech.iconKey)" class="w-14 h-14 md:w-16 md:h-16 grayscale hover:grayscale-0 transition-all duration-300 hover:scale-110 drop-shadow-sm mb-2" />
                                <EditableText :setting-key="tech.nameKey" :initial-value="getSetting(tech.nameKey)" class="text-xs font-bold text-slate-500" />
                            </div>
                        </div>
                    </div>
                </div>
                <!-- Terminal / CLI Philosophy -->
                <div class="mt-6 p-8 md:p-12 rounded-[2rem] bg-slate-900 text-white border border-slate-800 relative overflow-hidden group shadow-2xl">
                    <div class="absolute top-0 right-0 w-64 h-64 bg-green-500/10 rounded-full blur-[100px] pointer-events-none"></div>
                    <div class="absolute bottom-0 left-0 w-64 h-64 bg-blue-500/10 rounded-full blur-[100px] pointer-events-none"></div>
                    
                    <div class="grid md:grid-cols-2 gap-12 items-stretch relative z-10">
                        <!-- Left: Terminal Visual -->
                        <div class="bg-black/80 backdrop-blur-md rounded-xl border border-slate-700/50 shadow-inner overflow-hidden font-mono text-sm md:text-base h-full flex flex-col">
                            <!-- Window Controls -->
                            <div class="bg-slate-800/50 px-4 py-2 flex gap-2 border-b border-slate-700/50 shrink-0">
                                <div class="w-3 h-3 rounded-full bg-red-500/80"></div>
                                <div class="w-3 h-3 rounded-full bg-yellow-500/80"></div>
                                <div class="w-3 h-3 rounded-full bg-green-500/80"></div>
                            </div>
                            <!-- Interaction -->
                            <div class="p-6 space-y-4 text-slate-300 flex-1 flex flex-col justify-center">
                                <div>
                                    <div class="flex gap-2">
                                        <span class="text-green-400 font-bold">tri@tarakreasi:~$</span>
                                        <span class="text-white">
                                            <EditableText setting-key="term_cmd_1" :initial-value="getSetting('term_cmd_1', 'whoami')" />
                                        </span>
                                    </div>
                                    <div class="pl-4 text-blue-300 mt-1">
                                        > <EditableText setting-key="term_out_1" :initial-value="getSetting('term_out_1', 'Hybrid Engineer who loves the CLI.')" />
                                    </div>
                                </div>
                                <div>
                                    <div class="flex gap-2">
                                        <span class="text-green-400 font-bold">tri@tarakreasi:~$</span>
                                        <span class="text-white">
                                            <EditableText setting-key="term_cmd_2" :initial-value="getSetting('term_cmd_2', 'status')" />
                                        </span>
                                    </div>
                                    <div class="pl-4 text-green-300 mt-1 flex items-center gap-2">
                                        > <EditableText setting-key="term_out_2" :initial-value="getSetting('term_out_2', 'Ready to deploy.')" />
                                        <span class="cursor-blink w-2.5 h-5 bg-slate-400 inline-block"></span>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <!-- Right: Quotes -->
                        <div class="space-y-8 h-full flex flex-col justify-center pb-16">
                            <h4 class="text-2xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-green-400 to-blue-500 mb-2">
                                <EditableText setting-key="term_heading" :initial-value="getSetting('term_heading', 'Terminal Native')" />
                            </h4>
                            <ul class="space-y-6">
                                <li class="flex items-center gap-4 group/item">
                                    <span class="flex-shrink-0 w-8 h-8 rounded-lg bg-slate-800 flex items-center justify-center text-green-400 group-hover/item:text-white group-hover/item:bg-green-500 transition-all duration-300">
                                        <span class="material-symbols-outlined text-lg">terminal</span>
                                    </span>
                                    <p class="text-lg md:text-xl font-medium text-slate-300 group-hover/item:text-white transition-colors">
                                        <EditableText setting-key="term_quote_1" :initial-value="getSetting('term_quote_1', 'Feeling at home in ~/$')" />
                                    </p>
                                </li>
                                <li class="flex items-center gap-4 group/item">
                                    <span class="flex-shrink-0 w-8 h-8 rounded-lg bg-slate-800 flex items-center justify-center text-blue-400 group-hover/item:text-white group-hover/item:bg-blue-500 transition-all duration-300">
                                        <span class="material-symbols-outlined text-lg">admin_panel_settings</span>
                                    </span>
                                    <p class="text-lg md:text-xl font-medium text-slate-300 group-hover/item:text-white transition-colors">
                                        <EditableText setting-key="term_quote_2" :initial-value="getSetting('term_quote_2', 'Sudo access is my love language.')" />
                                    </p>
                                </li>
                                <li class="flex items-center gap-4 group/item">
                                    <span class="flex-shrink-0 w-8 h-8 rounded-lg bg-slate-800 flex items-center justify-center text-purple-400 group-hover/item:text-white group-hover/item:bg-purple-500 transition-all duration-300">
                                        <span class="material-symbols-outlined text-lg">dns</span>
                                    </span>
                                    <p class="text-lg md:text-xl font-medium text-slate-300 group-hover/item:text-white transition-colors">
                                        <EditableText setting-key="term_quote_3" :initial-value="getSetting('term_quote_3', 'SSH is my daily commute.')" />
                                    </p>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- PHILOSOPHY -->
        <section id="philosophy" class="bg-[#EFF6FF] dark:bg-[#1E293B] text-center py-20 px-6 transition-colors duration-500 border-t border-blue-100 dark:border-slate-800 pl-24">
            <div class="max-w-4xl mx-auto space-y-10">
                <span class="text-blue-600 dark:text-blue-400 font-bold uppercase tracking-[0.2em] text-xs bg-white dark:bg-slate-800 px-4 py-2 rounded-full shadow-sm">
                    <EditableText setting-key="philosophy_badge" :initial-value="getSetting('philosophy_badge')" />
                </span>
                
                <p class="text-2xl md:text-4xl font-medium leading-relaxed text-slate-800 dark:text-slate-100 font-display">
                    "<EditableText setting-key="philosophy_quote_1" :initial-value="getSetting('philosophy_quote_1')" /><br class="hidden md:block" />
                    <EditableText setting-key="philosophy_quote_2" :initial-value="getSetting('philosophy_quote_2')" /><br class="hidden md:block" />
                    <span class="text-blue-600 dark:text-blue-400">
                        <EditableText setting-key="philosophy_quote_highlight" :initial-value="getSetting('philosophy_quote_highlight')" />
                    </span>"
                </p>
                
                <div class="w-16 h-1.5 bg-blue-500/20 mx-auto rounded-full"></div>
                
                <div class="space-y-2">
                     <p class="text-xl md:text-2xl text-slate-500 dark:text-slate-400 font-serif italic">
                        <EditableText setting-key="philosophy_jp" :initial-value="getSetting('philosophy_jp')" />
                    </p>
                    <p class="text-xs md:text-sm text-slate-400 uppercase tracking-widest">
                        <EditableText setting-key="philosophy_en" :initial-value="getSetting('philosophy_en')" />
                    </p>
                </div>
            </div>
        </section>

        <!-- CTA -->
        <section id="contact" class="bg-white dark:bg-slate-900 py-20 text-center transition-colors duration-500 border-t border-slate-100 dark:border-slate-800 pl-24">
            <div class="max-w-4xl mx-auto px-6">
                <h3 class="text-3xl md:text-4xl font-bold text-slate-900 dark:text-slate-100 mb-8 tracking-tight">
                    <EditableText setting-key="contact_heading" :initial-value="getSetting('contact_heading')" />
                </h3>
                <div class="flex flex-col md:flex-row justify-center gap-4">
                    <a 
                        :href="'mailto:' + getSetting('contact_email')"
                        class="inline-block px-8 py-3.5 rounded-xl font-bold text-base shadow-xl hover:brightness-110 transition-all transform hover:-translate-y-1 bg-blue-600 text-white"
                        style="background-color: #2563eb; color: white;"
                    >
                        Email Me
                    </a>
                     <!-- Note: Helper logic for mailto? Just let user edit the full URL 'mailto:...' -->

                    <a
                        :href="getSetting('contact_linkedin')"
                        target="_blank"
                        class="inline-block px-8 py-3.5 rounded-xl font-bold text-base shadow-lg hover:brightness-110 transition-all transform hover:-translate-y-1 bg-[#0077b5] text-white"
                        style="background-color: #0077b5; color: white;"
                    >
                        LinkedIn
                    </a>
                </div>
                <div class="mt-16 pt-6 border-t border-slate-100 dark:border-slate-800 text-sm text-slate-400 dark:text-slate-500 flex flex-col md:flex-row justify-between items-center gap-4">
                    <EditableText setting-key="footer_copyright" :initial-value="getSetting('footer_copyright')" />
                    <EditableText setting-key="footer_tech" :initial-value="getSetting('footer_tech')" />
                </div>
            </div>
        </section>
    </div>
</template>

<style scoped>
@keyframes fadeIn {
  0% { opacity: 0; transform: translateY(20px); }
  100% { opacity: 1; transform: translateY(0); }
}
.animate-fadeIn { animation: fadeIn 0.8s ease-out forwards; }

@keyframes slideUp {
  0% { opacity: 0; transform: translateY(30px); }
  100% { opacity: 1; transform: translateY(0); }
}
.animate-slideUp { animation: slideUp 0.6s cubic-bezier(0.16, 1, 0.3, 1) both; }

.cursor-blink {
    animation: blink 1s step-end infinite;
}
@keyframes blink {
    0%, 100% { opacity: 1; }
    50% { opacity: 0; }
}
</style>
