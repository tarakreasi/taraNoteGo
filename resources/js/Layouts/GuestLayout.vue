<script setup>
import { Link } from '@inertiajs/vue3';
</script>

<template>
    <div class="min-h-screen w-full flex items-center justify-center bg-canvas-light dark:bg-canvas-dark text-slate-800 dark:text-slate-100 transition-colors duration-500 relative overflow-hidden">
        
        <!-- Noise Overlay -->
        <div class="fixed inset-0 z-[9999] opacity-[0.04] pointer-events-none mix-blend-overlay"
             style="background-image: url(&quot;data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noiseFilter'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.65' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noiseFilter)'/%3E%3C/svg%3E&quot;);">
        </div>

        <!-- Living Background: Animated Blobs -->
        <div class="fixed inset-0 overflow-hidden pointer-events-none">
             <!-- Blue Blob -->
            <div class="absolute top-0 left-1/4 w-[500px] h-[500px] bg-primary/20 dark:bg-primary/10 rounded-full mix-blend-multiply dark:mix-blend-screen filter blur-[80px] opacity-70 animate-[breathe_10s_infinite]"></div>
            
            <!-- Purple Blob -->
            <div class="absolute top-0 right-1/4 w-[500px] h-[500px] bg-purple-500/20 dark:bg-purple-500/10 rounded-full mix-blend-multiply dark:mix-blend-screen filter blur-[80px] opacity-70 animate-[breathe_10s_infinite] delay-1000"></div>
            
            <!-- Orange Blob (Subtle warmth) -->
            <div class="absolute -bottom-32 left-1/3 w-[600px] h-[600px] bg-accent-orange/10 dark:bg-accent-orange/5 rounded-full mix-blend-multiply dark:mix-blend-screen filter blur-[100px] opacity-50 animate-[float_12s_ease-in-out_infinite] delay-2000"></div>
        </div>
        
        <!-- Theme Toggle (Optional, usually global but good for Login) -->
        <!-- Logic is usually in a component, simplified here or removed if handled by user OS pref -->

        <!-- Slot Content (The Card) -->
        <div class="w-full max-w-md p-6 relative z-10 animate-[reveal-up_0.8s_ease-out]">
            <slot />
        </div>
    </div>
</template>
