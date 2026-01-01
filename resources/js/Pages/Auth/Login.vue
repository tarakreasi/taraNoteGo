<script setup>
import { Head, Link, useForm } from '@inertiajs/vue3';
import GuestLayout from '@/Layouts/GuestLayout.vue';

defineProps({
    canResetPassword: {
        type: Boolean,
    },
    status: {
        type: String,
    },
});

const form = useForm({
    email: '',
    password: '',
    remember: false,
});

const submit = () => {
    form.post(route('login'), {
        onFinish: () => form.reset('password'),
    });
};
</script>

<template>
    <GuestLayout>
        <Head title="Admin Login" />

        <div class="glass-panel rounded-3xl p-8 md:p-10 w-full relative overflow-hidden group">
            
            <!-- Top Shine Effect -->
            <div class="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-white/50 to-transparent opacity-50"></div>

            <!-- Content -->
            <div class="flex flex-col gap-8 relative z-10">
                
                <div class="text-center">
                    <div class="inline-flex items-center justify-center px-6 py-4 rounded-2xl bg-white/10 backdrop-blur-sm shadow-lg mb-8 animate-[float_6s_ease-in-out_infinite]">
                        <img src="/images/logo_tarakreasi.png" alt="TaraKreasi Logo" class="h-16 w-auto object-contain drop-shadow-md">
                    </div>
                    <h1 class="font-display text-3xl font-bold text-slate-900 dark:text-white mb-2 tracking-tight">
                        Welcome Back
                    </h1>
                    <p class="text-slate-500 dark:text-slate-400">Enter your coordinates to access the portal.</p>
                </div>

                <!-- Status Message -->
                <div v-if="status" class="mb-4 font-medium text-sm text-success text-center">
                    {{ status }}
                </div>

                <!-- Login Form -->
                <form @submit.prevent="submit" class="flex flex-col gap-5">

                    <!-- Email Input Group -->
                    <div class="space-y-1.5">
                        <label for="email" class="text-xs font-semibold uppercase tracking-wider text-slate-500 dark:text-slate-400 ml-1">Email Address</label>
                        <div class="relative group/input">
                            <div class="absolute inset-0 bg-gradient-to-r from-primary to-purple-500 rounded-xl opacity-0 group-focus-within/input:opacity-100 blur transition-opacity duration-500"></div>
                            <div class="relative flex items-center bg-white/50 dark:bg-[#0F172A]/50 border border-slate-200 dark:border-white/10 rounded-xl px-4 h-12 transition-all focus-within:bg-white dark:focus-within:bg-[#0F172A]">
                                <span class="material-symbols-outlined text-slate-400 mr-3">mail</span>
                                <input 
                                    id="email" 
                                    name="email"
                                    type="email" 
                                    v-model="form.email"
                                    class="w-full bg-transparent border-none focus:ring-0 text-slate-900 dark:text-white placeholder-slate-400 text-sm font-medium"
                                    placeholder="you@tarakreasi.com"
                                    required 
                                    autofocus 
                                    autocomplete="username"
                                >
                            </div>
                        </div>
                        <div v-if="form.errors.email" class="text-xs text-danger ml-1 mt-1 font-medium">
                            {{ form.errors.email }}
                        </div>
                    </div>

                    <!-- Password Input Group -->
                    <div class="space-y-1.5">
                        <div class="flex justify-between items-center ml-1">
                            <label for="password" class="text-xs font-semibold uppercase tracking-wider text-slate-500 dark:text-slate-400">Password</label>
                            <Link v-if="canResetPassword" :href="route('password.request')" class="text-xs font-medium text-primary hover:text-primary-hover transition-colors">
                                Forgot?
                            </Link>
                        </div>
                        <div class="relative group/input">
                            <div class="absolute inset-0 bg-gradient-to-r from-primary to-purple-500 rounded-xl opacity-0 group-focus-within/input:opacity-100 blur transition-opacity duration-500"></div>
                            <div class="relative flex items-center bg-white/50 dark:bg-[#0F172A]/50 border border-slate-200 dark:border-white/10 rounded-xl px-4 h-12 transition-all focus-within:bg-white dark:focus-within:bg-[#0F172A]">
                                <span class="material-symbols-outlined text-slate-400 mr-3">lock</span>
                                <input 
                                    id="password" 
                                    name="password"
                                    type="password" 
                                    v-model="form.password"
                                    class="w-full bg-transparent border-none focus:ring-0 text-slate-900 dark:text-white placeholder-slate-400 text-sm font-medium"
                                    placeholder="••••••••"
                                    required 
                                    autocomplete="current-password"
                                >
                            </div>
                        </div>
                        <div v-if="form.errors.password" class="text-xs text-danger ml-1 mt-1 font-medium">
                            {{ form.errors.password }}
                        </div>
                    </div>

                    <!-- Remember Me -->
                    <label class="flex items-center ml-1 cursor-pointer">
                        <input type="checkbox" v-model="form.remember" class="rounded border-slate-300 text-primary shadow-sm focus:ring-primary">
                        <span class="ms-2 text-xs text-slate-600 dark:text-slate-400">Remember me</span>
                    </label>

                    <!-- Submit Button -->
                    <button 
                        class="mt-2 w-full h-12 bg-primary hover:bg-primary-hover text-white font-display font-bold rounded-xl shadow-[0_0_20px_rgba(68,133,238,0.3)] hover:shadow-[0_0_30px_rgba(68,133,238,0.5)] active:scale-[0.98] transition-all duration-300 flex items-center justify-center gap-2 group/btn disabled:opacity-50 disabled:cursor-not-allowed"
                        :disabled="form.processing"
                    >
                        <span v-if="!form.processing">Login</span>
                        <span v-else>Processing...</span>
                        <span class="material-symbols-outlined text-[18px] group-hover/btn:translate-x-1 transition-transform">arrow_forward</span>
                    </button>

                </form>

                <!-- Footer / Divider
                <div class="relative py-2">
                    <div class="absolute inset-0 flex items-center">
                        <div class="w-full border-t border-slate-200 dark:border-white/10"></div>
                    </div>
                    <div class="relative flex justify-center text-sm">
                        <span class="px-4 bg-transparent backdrop-blur-xl text-slate-400">Or continue with</span>
                    </div>
                </div>
                 -->
                <!-- Link Register
                 <div class="text-center">
                    <Link :href="route('register')" class="text-sm text-slate-500 hover:text-primary transition-colors font-medium">
                        Don't have an account? <span class="font-bold underline text-slate-900 dark:text-slate-100">Register</span>
                    </Link>
                </div>
                -->
            </div>
        </div>

        <!-- Copyright -->
        <p class="text-center text-xs text-slate-400/60 mt-8">
            &copy; 2026 taraKreasi.com. the begining dream come true
        </p>

    </GuestLayout>
</template>