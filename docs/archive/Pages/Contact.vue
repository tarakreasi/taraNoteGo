<script setup>
import { Head } from '@inertiajs/vue3';
import GuestLayout from '@/Layouts/GuestLayout.vue';
import AppNavbar from '@/Components/AppNavbar.vue';
import AppFooter from '@/Components/AppFooter.vue';

const props = defineProps({
    settings: {
        type: Object,
        default: () => ({})
    }
});

// Noise overlay style
const noiseOverlayStyle = "background-image: url('data:image/svg+xml,%3Csvg viewBox=\\'0 0 200 200\\' xmlns=\\'http://www.w3.org/2000/svg\\'%3E%3Cfilter id=\\'noiseFilter\\'%3E%3CfeTurbulence type=\\'fractalNoise\\' baseFrequency=\\'0.65\\' numOctaves=\\'3\\' stitchTiles=\\'stitch\\'/%3E%3C/filter%3E%3Crect width=\\'100%25\\' height=\\'100%25\\' filter=\\'url(%23noiseFilter)\\'/%3E%3C/svg%3E');";
</script>

<template>
    <div class="min-h-screen bg-canvas-light dark:bg-canvas-dark text-slate-800 dark:text-white transition-colors duration-200 font-body relative overflow-x-hidden">
        <Head title="Contact" />
        
        <!-- Noise Overlay -->
        <div class="fixed inset-0 z-[9999] opacity-[0.15] dark:opacity-[0.07] pointer-events-none" :style="noiseOverlayStyle"></div>

        <div class="relative flex h-auto min-h-screen w-full flex-col overflow-x-hidden">
            <AppNavbar />

            <main class="layout-container flex h-full grow flex-col pt-24 z-10">
                <div class="px-4 md:px-10 flex flex-1 justify-center py-10 relative">
                    <!-- Ambient Contact Aura -->
                    <div class="absolute top-0 left-0 w-[500px] h-[500px] bg-primary/10 rounded-full blur-[120px] opacity-40 pointer-events-none -z-10 -translate-x-1/2 -translate-y-1/2"></div>
                    <!-- Ambient Contact Aura (Updated Blob) -->
                    <div class="absolute top-[-10%] left-[50%] -translate-x-1/2 w-[60%] h-[60%] bg-primary/10 blur-[120px] rounded-full pointer-events-none z-0"></div>

                    <div class="layout-content-container flex flex-col max-w-[1200px] flex-1 relative z-10">
                        <!-- Page Heading -->
                        <div class="flex flex-wrap justify-between gap-3 p-4 mb-8">
                            <div class="flex min-w-72 flex-col gap-3">
                                <p class="font-display text-gray-900 dark:text-white text-4xl md:text-5xl lg:text-6xl font-black leading-tight tracking-tight">
                                    Get in Touch
                                </p>
                                <p class="text-[#887763] dark:text-[#b0a69a] text-lg font-normal leading-normal max-w-2xl">
                                    Have a project in mind or just want to say hi? We'd love to hear from you. Fill out the
                                    form below or reach out directly via email.
                                </p>
                            </div>
                        </div>

                        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 p-4">
                            <!-- Left Column: Contact Info -->
                            <div class="lg:col-span-5 flex flex-col gap-8">
                                <!-- Contact Cards Section -->
                                <div class="flex flex-col gap-4">
                                    <h2 class="text-primary dark:text-white tracking-light text-[24px] font-bold leading-tight pb-2">
                                        Direct Contact
                                    </h2>
                                    <!-- Email Item -->
                                    <div class="glass-panel flex items-center gap-4 px-5 py-4 min-h-20 justify-between rounded-xl bg-white/50 dark:bg-black/20 backdrop-blur-md border border-white/20 dark:border-white/5 shadow-sm">
                                        <div class="flex items-center gap-4 overflow-hidden">
                                            <div class="text-primary dark:text-primary flex items-center justify-center rounded-lg bg-[#f0f4f8] dark:bg-[#3e3427]/50 shrink-0 size-12">
                                                <span class="material-symbols-outlined text-[24px]">mail</span>
                                            </div>
                                            <div class="flex flex-col overflow-hidden">
                                                <span class="text-xs text-[#887763] dark:text-[#b0a69a] font-medium uppercase tracking-wider">Email</span>
                                                <p class="text-primary dark:text-white text-base font-semibold leading-normal truncate">
                                                    {{ settings.contact_email || 'hello@tarakreasi.com' }}
                                                </p>
                                            </div>
                                        </div>
                                    </div>
                                    <!-- Location Item -->
                                    <div class="glass-panel flex items-center gap-4 px-5 py-4 min-h-20 justify-between rounded-xl bg-white/50 dark:bg-black/20 backdrop-blur-md border border-white/20 dark:border-white/5 shadow-sm">
                                        <div class="flex items-center gap-4 overflow-hidden">
                                            <div class="text-primary dark:text-primary flex items-center justify-center rounded-lg bg-[#f0f4f8] dark:bg-[#3e3427]/50 shrink-0 size-12">
                                                <span class="material-symbols-outlined text-[24px]">location_on</span>
                                            </div>
                                            <div class="flex flex-col overflow-hidden">
                                                <span class="text-xs text-[#887763] dark:text-[#b0a69a] font-medium uppercase tracking-wider">Office</span>
                                                <p class="text-primary dark:text-white text-base font-semibold leading-normal truncate">
                                                    {{ settings.contact_location || 'Jakarta, Indonesia (WIB)' }}
                                                </p>
                                            </div>
                                        </div>
                                        <div class="shrink-0">
                                            <div class="flex size-7 items-center justify-center" title="Online now">
                                                <div class="size-3 rounded-full bg-[#078810] animate-pulse"></div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <!-- Social Links -->
                                <div class="flex flex-col gap-4 mt-4">
                                    <h2 class="text-primary dark:text-white tracking-light text-[24px] font-bold leading-tight pb-2">
                                        Follow Us
                                    </h2>
                                    <div class="flex gap-4">
                                        <a v-if="settings.footer_social_twitter" :href="settings.footer_social_twitter" target="_blank" class="glass-panel flex items-center justify-center size-12 rounded-xl text-primary dark:text-white hover:text-[#181511] hover:border-primary transition-all duration-300 bg-white/50 dark:bg-black/20 backdrop-blur-md border border-white/20 dark:border-white/5">
                                            <svg aria-hidden="true" class="size-6" fill="currentColor" viewbox="0 0 24 24">
                                                <path d="M8.29 20.251c7.547 0 11.675-6.253 11.675-11.675 0-.178 0-.355-.012-.53A8.348 8.348 0 0022 5.92a8.19 8.19 0 01-2.357.646 4.118 4.118 0 001.804-2.27 8.224 8.224 0 01-2.605.996 4.107 4.107 0 00-6.993 3.743 11.65 11.65 0 01-8.457-4.287 4.106 4.106 0 001.27 5.477A4.072 4.072 0 012.8 9.713v.052a4.105 4.105 0 003.292 4.022 4.095 4.095 0 01-1.853.07 4.108 4.108 0 003.834 2.85A8.233 8.233 0 012 18.407a11.616 11.616 0 006.29 1.84">
                                                </path>
                                            </svg>
                                        </a>
                                        <a v-if="settings.footer_social_github" :href="settings.footer_social_github" target="_blank" class="glass-panel flex items-center justify-center size-12 rounded-xl text-primary dark:text-white hover:text-[#181511] hover:border-primary transition-all duration-300 bg-white/50 dark:bg-black/20 backdrop-blur-md border border-white/20 dark:border-white/5">
                                            <svg aria-hidden="true" class="size-6" fill="currentColor" viewbox="0 0 24 24">
                                                <path fill-rule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clip-rule="evenodd"></path>
                                            </svg>
                                        </a>
                                    </div>
                                </div>
                                <!-- Decorative Image -->
                                <div class="mt-4 rounded-xl overflow-hidden shadow-sm border border-[#f4f2f0] dark:border-[#3e3427] h-48 relative group">
                                    <img alt="Abstract minimalist architectural detail" class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDafY0pi9QCVFWlc1QsRDa3jG1rlI_nZoYTUu45dEbZNTcptyQvAjoV4yOHmg3ahnIDHMvDLyB65KX3kcokymhYyeRHtsmDmB7gjANl4WBXFacm9mHSoCwSGEtxMuQY6vo41db4n8Y-Wya6Mw1XRuoei0GtvoSOZbCvIu6IHCUCKIGNFWNJu4aVhf3TKoj88F25iEt7-dTNMfYcEhkkairZ-wy_vV41g7adu1lZQE3w92sOnTnrAATVaL7-S0pM-Bo358U2QTKX2jdZ" />
                                    <div class="absolute inset-0 bg-primary/10 dark:bg-black/20 mix-blend-overlay"></div>
                                </div>
                            </div>
                            <!-- Right Column: Contact Form -->
                            <div class="lg:col-span-7">
                                <div class="glass-panel rounded-2xl p-8 md:p-10 bg-white/60 dark:bg-black/40 backdrop-blur-xl border border-white/20 dark:border-white/5 shadow-2xl">
                                    <h2 class="text-primary dark:text-white text-2xl font-bold mb-6">Send a Message</h2>
                                    <form @submit.prevent class="flex flex-col gap-6">
                                        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                                            <div class="flex flex-col gap-2">
                                                <label class="text-sm font-medium text-primary dark:text-[#b0a69a]" for="name">Your Name</label>
                                                <input class="w-full h-12 rounded-lg border border-[#d6d0ca] dark:border-[#52483d] bg-[#fdfcfc] dark:bg-[#211a11]/50 px-4 text-[#181511] dark:text-white placeholder-[#887763] dark:placeholder-[#6d6459] focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all outline-none" id="name" placeholder="John Doe" type="text" />
                                            </div>
                                            <div class="flex flex-col gap-2">
                                                <label class="text-sm font-medium text-primary dark:text-[#b0a69a]" for="email">Email Address</label>
                                                <input class="w-full h-12 rounded-lg border border-[#d6d0ca] dark:border-[#52483d] bg-[#fdfcfc] dark:bg-[#211a11]/50 px-4 text-[#181511] dark:text-white placeholder-[#887763] dark:placeholder-[#6d6459] focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all outline-none" id="email" placeholder="john@example.com" type="email" />
                                            </div>
                                        </div>
                                        <div class="flex flex-col gap-2">
                                            <label class="text-sm font-medium text-primary dark:text-[#b0a69a]" for="subject">Subject</label>
                                            <div class="relative">
                                                <select class="w-full h-12 rounded-lg border border-[#d6d0ca] dark:border-[#52483d] bg-[#fdfcfc] dark:bg-[#211a11]/50 px-4 text-[#181511] dark:text-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all outline-none appearance-none cursor-pointer" id="subject">
                                                    <option disabled="" selected="" value="">Select a topic</option>
                                                    <option value="project">New Project Inquiry</option>
                                                    <option value="collab">Collaboration</option>
                                                    <option value="feedback">Feedback</option>
                                                    <option value="other">Other</option>
                                                </select>
                                                <div class="absolute right-4 top-1/2 -translate-y-1/2 pointer-events-none text-[#887763] dark:text-[#b0a69a]">
                                                    <span class="material-symbols-outlined">expand_more</span>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="flex flex-col gap-2">
                                            <label class="text-sm font-medium text-primary dark:text-[#b0a69a]" for="message">Message</label>
                                            <textarea class="w-full rounded-lg border border-[#d6d0ca] dark:border-[#52483d] bg-[#fdfcfc] dark:bg-[#211a11]/50 p-4 text-[#181511] dark:text-white placeholder-[#887763] dark:placeholder-[#6d6459] focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all outline-none resize-none" id="message" placeholder="Tell us about your project..." rows="5"></textarea>
                                            <div class="flex justify-end">
                                                <span class="text-xs text-[#887763] dark:text-[#6d6459]">0/500 characters</span>
                                            </div>
                                        </div>
                                        <div class="pt-2">
                                            <button class="w-full md:w-auto px-8 h-12 bg-primary hover:bg-[#d9964a] text-white text-base font-bold rounded-lg shadow-lg shadow-primary/20 hover:shadow-xl transition-all duration-200 flex items-center justify-center gap-2 group transform hover:-translate-y-0.5" type="submit">
                                                <span>Send Message</span>
                                                <span class="material-symbols-outlined transition-transform group-hover:translate-x-1 text-sm font-bold">send</span>
                                            </button>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </main>

            <AppFooter :settings="settings" />
        </div>
    </div>
</template>
