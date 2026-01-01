<script setup>
import { Head, Link, usePage } from '@inertiajs/vue3';
import { ref } from 'vue';
import DeleteUserForm from './Partials/DeleteUserForm.vue';
import UpdatePasswordForm from './Partials/UpdatePasswordForm.vue';
import UpdateProfileInformationForm from './Partials/UpdateProfileInformationForm.vue';

const props = defineProps({
    mustVerifyEmail: {
        type: Boolean,
    },
    status: {
        type: String,
    },
});

const user = usePage().props.auth.user;
const activeTab = ref('profile'); // profile, password, delete

// Helper to format date if needed, or other utils
</script>

<template>
    <Head title="Profile" />
    
    <!-- Noise Overlay -->
    <div class="fixed inset-0 z-[9999] opacity-[0.15] dark:opacity-[0.07] pointer-events-none" style="background-image: url('data:image/svg+xml,%3Csvg viewBox=\'0 0 200 200\' xmlns=\'http://www.w3.org/2000/svg\'%3E%3Cfilter id=\'noiseFilter\'%3E%3CfeTurbulence type=\'fractalNoise\' baseFrequency=\'0.65\' numOctaves=\'3\' stitchTiles=\'stitch\'/%3E%3C/filter%3E%3Crect width=\'100%25\' height=\'100%25\' filter=\'url(%23noiseFilter)\'/%3E%3C/svg%3E');"></div>

    <!-- Ambient Gradient Background -->
    <div class="fixed inset-0 bg-canvas-light dark:bg-canvas-dark transition-colors duration-300 font-display overflow-hidden h-screen flex">
        <!-- Ambient Blobs -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none z-0">
             <div class="absolute top-[10%] left-[20%] w-[500px] h-[500px] bg-primary/5 dark:bg-primary/10 blur-[120px] rounded-full animate-[breathe_10s_infinite]"></div>
             <div class="absolute bottom-[20%] right-[10%] w-[600px] h-[600px] bg-accent-orange/5 dark:bg-accent-orange/10 blur-[100px] rounded-full animate-[float_15s_ease-in-out_infinite]"></div>
        </div>
        
        <!-- SIDEBAR (Copied from Dashboard.vue for consistency) -->
        <aside class="hidden md:flex flex-col w-72 bg-brand-blue h-full p-4 justify-between shadow-xl z-20">
            <div class="flex flex-col gap-8">
                <a href="/" class="flex flex-col items-center justify-center gap-2 px-2 py-4 group mb-4">
                    <div class="bg-white/10 backdrop-blur-md border border-white/20 px-6 py-3 rounded-2xl shadow-lg group-hover:bg-white/20 transition-all group-hover:scale-105 duration-300">
                        <img src="/images/logo_tarakreasi.png" alt="TaraKreasi Logo" class="h-10 object-contain opacity-90 group-hover:opacity-100 transition-opacity">
                    </div>
                </a>

                <!-- Navigation -->
                <nav class="flex flex-col gap-2">
                    <Link :href="route('dashboard')" class="hover:bg-white/10 text-white/90 hover:text-white flex items-center gap-3 px-4 py-3 rounded-xl transition-all group">
                        <span class="material-symbols-outlined font-bold">dashboard</span>
                        <span class="text-sm font-bold">Dashboard</span>
                    </Link>
                    
                    <div class="px-4 pt-2 pb-1 text-white/60 text-xs font-bold uppercase tracking-wider">Account</div>
                    <Link :href="route('profile.edit')" class="bg-white shadow-lg shadow-blue-900/10 text-brand-blue flex items-center gap-3 px-4 py-3 rounded-xl transition-all group">
                        <span class="material-symbols-outlined font-bold">person</span>
                        <span class="text-sm font-bold">Profile</span>
                    </Link>

                    <a href="/dashboard?view=settings" class="hover:bg-white/10 text-white/90 hover:text-white flex items-center gap-3 px-4 py-3 rounded-xl transition-all group">
                        <span class="material-symbols-outlined font-bold">settings</span>
                        <span class="text-sm font-bold">Site Settings</span>
                    </a>
                </nav>
            </div>

            <!-- User Profile & Actions -->
            <div class="flex flex-col gap-2 p-3 rounded-2xl bg-white/10 backdrop-blur-md border border-white/20 shadow-lg">
                <div class="flex items-center gap-3 px-1 mb-1">
                    <div class="h-10 w-10 shrink-0 rounded-full bg-white flex items-center justify-center text-brand-blue font-bold shadow-sm">
                        {{ user.name.substring(0,2).toUpperCase() }}
                    </div>
                    <div class="flex flex-col overflow-hidden">
                        <span class="text-sm font-bold text-white truncate">{{ user.name }}</span>
                        <span class="text-[10px] text-blue-100/80 truncate">{{ user.email }}</span>
                    </div>
                </div>
                 <div class="grid grid-cols-2 gap-2 mt-1">
                    <Link :href="route('dashboard')" class="flex items-center justify-center gap-1.5 py-2 px-2 rounded-xl bg-white/10 hover:bg-white/20 text-white text-[11px] font-bold transition-all hover:scale-105 border border-white/5">
                        <span class="material-symbols-outlined text-[16px]">arrow_back</span>
                        <span>Back</span>
                    </Link>
                    <Link :href="route('logout')" method="post" as="button" class="flex items-center justify-center gap-1.5 py-2 px-2 rounded-xl bg-white/10 hover:bg-rose-500/80 hover:border-rose-400/50 text-white text-[11px] font-bold transition-all hover:scale-105 border border-white/5">
                        <span class="material-symbols-outlined text-[16px]">logout</span>
                        <span>Logout</span>
                    </Link>
                </div>
            </div>
        </aside>

        <!-- MAIN CONTENT -->
        <main class="flex-1 flex flex-col h-full overflow-hidden relative bg-[#F4F7FE] dark:bg-slate-900">
             <div class="flex-1 overflow-y-auto p-4 sm:p-8">
                <div class="w-full max-w-[960px] mx-auto flex flex-col gap-6">
                    
                    <!-- Profile Header Section (From code.html) -->
                    <section class="flex flex-col items-center bg-white dark:bg-slate-800 rounded-3xl p-8 shadow-sm border border-slate-100 dark:border-slate-700">
                        <div class="flex flex-col items-center gap-4">
                            <!-- Avatar -->
                            <div class="relative">
                                <div class="bg-center bg-no-repeat bg-cover rounded-full h-32 w-32 border-4 border-slate-50 dark:border-slate-700 shadow-sm flex items-center justify-center bg-brand-blue text-white text-4xl font-bold" 
                                     :style="user.avatar ? `background-image: url('${user.avatar}')` : ''">
                                     <span v-if="!user.avatar">{{ user.name.substring(0,2).toUpperCase() }}</span>
                                </div>
                                <div class="absolute bottom-1 right-1 bg-emerald-500 border-2 border-white dark:border-slate-800 rounded-full w-6 h-6"></div>
                            </div>
                            <!-- Bio Text -->
                            <div class="flex flex-col items-center justify-center text-center max-w-lg">
                                <h1 class="text-2xl font-bold leading-tight tracking-[-0.015em] dark:text-white">{{ user.name }}</h1>
                                <p class="text-brand-blue font-medium text-base mt-1">{{ user.email }}</p>
                                <p class="text-slate-500 dark:text-slate-400 text-base mt-3 leading-relaxed">
                                    {{ user.bio || "No bio yet. Tell us about yourself!" }}
                                </p>
                            </div>
                        </div>
                    </section>

                    <!-- Navigation Tabs -->
                    <nav class="border-b border-slate-200 dark:border-slate-700">
                        <div class="flex gap-8 px-4">
                            <button @click="activeTab = 'profile'" 
                                :class="activeTab === 'profile' ? 'border-brand-blue text-slate-800 dark:text-white' : 'border-transparent text-slate-500 hover:text-slate-700 dark:text-slate-400'"
                                class="flex items-center justify-center border-b-[3px] pb-3 pt-2 transition-colors">
                                <p class="text-sm font-bold leading-normal tracking-[0.015em]">Profile Settings</p>
                            </button>
                            <button @click="activeTab = 'password'" 
                                :class="activeTab === 'password' ? 'border-brand-blue text-slate-800 dark:text-white' : 'border-transparent text-slate-500 hover:text-slate-700 dark:text-slate-400'"
                                class="flex items-center justify-center border-b-[3px] pb-3 pt-2 transition-colors">
                                <p class="text-sm font-bold leading-normal tracking-[0.015em]">Password</p>
                            </button>
                            <button @click="activeTab = 'delete'" 
                                :class="activeTab === 'delete' ? 'border-rose-500 text-rose-600' : 'border-transparent text-slate-500 hover:text-rose-600 dark:text-slate-400'"
                                class="flex items-center justify-center border-b-[3px] pb-3 pt-2 transition-colors">
                                <p class="text-sm font-bold leading-normal tracking-[0.015em]">Danger Zone</p>
                            </button>
                        </div>
                    </nav>

                    <!-- Tab Content -->
                    <section class="py-4">
                        
                        <!-- PROFILE TAB -->
                        <div v-if="activeTab === 'profile'" class="bg-white dark:bg-slate-800 rounded-3xl border border-slate-100 dark:border-slate-700 p-8 shadow-sm">
                            <UpdateProfileInformationForm
                                :must-verify-email="mustVerifyEmail"
                                :status="status"
                                class="max-w-xl"
                            />
                        </div>

                         <!-- PASSWORD TAB -->
                        <div v-if="activeTab === 'password'" class="bg-white dark:bg-slate-800 rounded-3xl border border-slate-100 dark:border-slate-700 p-8 shadow-sm">
                            <UpdatePasswordForm class="max-w-xl" />
                        </div>

                         <!-- DELETE TAB -->
                        <div v-if="activeTab === 'delete'" class="bg-white dark:bg-slate-800 rounded-3xl border border-slate-100 dark:border-slate-700 p-8 shadow-sm border-l-4 border-l-rose-500">
                            <DeleteUserForm class="max-w-xl" />
                        </div>

                    </section>

                </div>
             </div>
        </main>
    </div>
</template>
