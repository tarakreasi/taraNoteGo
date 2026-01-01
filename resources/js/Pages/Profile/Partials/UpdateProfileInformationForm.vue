<script setup>
import InputError from '@/Components/InputError.vue';
import InputLabel from '@/Components/InputLabel.vue';
import PrimaryButton from '@/Components/PrimaryButton.vue';
import TextInput from '@/Components/TextInput.vue';
import { Link, useForm, usePage } from '@inertiajs/vue3';

defineProps({
    mustVerifyEmail: {
        type: Boolean,
    },
    status: {
        type: String,
    },
});

const user = usePage().props.auth.user;

const form = useForm({
    name: user.name,
    email: user.email,
    bio: user.bio,
    avatar: user.avatar,
    hero_tagline: user.hero_tagline,
    hero_image: user.hero_image,
});
</script>

<template>
    <section>
        <header class="mb-8">
            <h2 class="text-xl font-bold text-slate-800 dark:text-white">
                Profile Information
            </h2>

            <p class="mt-2 text-sm text-slate-500 dark:text-slate-400">
                Update your account's profile information and details.
            </p>
        </header>

        <form
            @submit.prevent="form.patch(route('profile.update'))"
            class="mt-6 space-y-8"
        >
             <!-- Portfolio Customization -->
             <div class="p-6 bg-brand-blue/5 dark:bg-brand-blue/10 rounded-2xl border border-brand-blue/10 dark:border-brand-blue/20">
                <h3 class="text-lg font-bold text-brand-blue mb-4 flex items-center gap-2">
                    <span class="material-symbols-outlined">palette</span>
                    Portfolio Customization
                </h3>
                
                <div class="space-y-6">
                    <!-- Hero Image -->
                    <div>
                        <InputLabel for="hero_image" value="Hero Background Image URL" />
                        <TextInput
                            id="hero_image"
                            type="text"
                            class="mt-1 block w-full"
                            v-model="form.hero_image"
                            placeholder="https://images.unsplash.com/photo-..."
                        />
                        <p class="mt-2 text-xs text-slate-500 dark:text-slate-400">
                             A large, high-quality image URL for your portfolio header background.
                        </p>
                        <InputError class="mt-2" :message="form.errors.hero_image" />
                    </div>

                    <!-- Tagline -->
                     <div>
                        <InputLabel for="hero_tagline" value="Hero Tagline" />
                        <TextInput
                            id="hero_tagline"
                            type="text"
                            class="mt-1 block w-full"
                            v-model="form.hero_tagline"
                            placeholder="Exploring the intersection of design and code."
                        />
                         <p class="mt-2 text-xs text-slate-500 dark:text-slate-400">
                            A short sentence describing your digital garden. Replaces the default text.
                        </p>
                        <InputError class="mt-2" :message="form.errors.hero_tagline" />
                    </div>
                </div>
             </div>
            <!-- Avatar Section -->
            <div class="flex flex-col sm:flex-row gap-6 items-start p-4 bg-slate-50 dark:bg-white/5 rounded-2xl border border-slate-100 dark:border-white/5">
                <div class="shrink-0 relative group">
                    <div class="h-20 w-20 rounded-full bg-white dark:bg-slate-700 flex items-center justify-center text-xl font-bold text-brand-blue shadow-md border-2 border-white dark:border-slate-600 overflow-hidden">
                        <img v-if="form.avatar" :src="form.avatar" alt="Avatar Preview" class="h-full w-full object-cover" @error="e => e.target.style.display='none'">
                        <span v-else>{{ form.name?.substring(0,2).toUpperCase() }}</span>
                    </div>
                </div>
                
                <div class="flex-1 w-full relative">
                    <InputLabel for="avatar" value="Avatar URL" class="mb-1.5" />
                    <TextInput
                        id="avatar"
                        type="text"
                        class="mt-1 block w-full bg-white dark:bg-slate-900"
                        v-model="form.avatar"
                        placeholder="https://example.com/my-photo.jpg"
                    />
                    <p class="mt-2 text-xs text-slate-500 dark:text-slate-400">
                        Paste a URL to an image to update your avatar.
                    </p>
                    <InputError class="mt-2" :message="form.errors.avatar" />
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- Name -->
                <div>
                    <InputLabel for="name" value="Display Name" />
                    <TextInput
                        id="name"
                        type="text"
                        class="mt-1 block w-full"
                        v-model="form.name"
                        required
                        autofocus
                        autocomplete="name"
                        placeholder="Your full name"
                    />
                    <InputError class="mt-2" :message="form.errors.name" />
                </div>

                <!-- Email -->
                <div>
                    <InputLabel for="email" value="Email Address" />
                    <TextInput
                        id="email"
                        type="email"
                        class="mt-1 block w-full"
                        v-model="form.email"
                        required
                        autocomplete="username"
                        placeholder="name@example.com"
                    />
                    <InputError class="mt-2" :message="form.errors.email" />
                </div>
            </div>

            <!-- Verification Warning -->
            <div v-if="mustVerifyEmail && user.email_verified_at === null" class="p-4 bg-amber-50 dark:bg-amber-900/20 rounded-xl border border-amber-100 dark:border-amber-700/30">
                <p class="text-sm text-amber-800 dark:text-amber-200">
                    Your email address is unverified.
                    <Link
                        :href="route('verification.send')"
                        method="post"
                        as="button"
                        class="underline font-bold hover:text-amber-900 dark:hover:text-white transition-colors"
                    >
                        Click here to re-send the verification email.
                    </Link>
                </p>

                <div
                    v-show="status === 'verification-link-sent'"
                    class="mt-2 text-sm font-medium text-green-600 dark:text-green-400"
                >
                    A new verification link has been sent to your email address.
                </div>
            </div>

            <!-- Bio -->
            <div>
                <InputLabel for="bio" value="Bio" />
                <textarea
                    id="bio"
                    class="mt-1 block w-full border-gray-300 dark:border-gray-700 dark:bg-slate-900 dark:text-gray-300 focus:border-brand-blue dark:focus:border-brand-blue focus:ring-brand-blue dark:focus:ring-brand-blue rounded-xl shadow-sm transition-all"
                    v-model="form.bio"
                    rows="4"
                    placeholder="Tell us a little about yourself... (Supports Markdown)"
                ></textarea>
                <p class="mt-2 text-xs text-slate-500 dark:text-slate-400 flex justify-between">
                    <span>Write a short bio to display on your public profile.</span>
                    <span :class="{'text-rose-500': (form.bio?.length || 0) > 255}">{{ form.bio?.length || 0 }} / 255</span>
                </p>
                <InputError class="mt-2" :message="form.errors.bio" />
            </div>

            <div class="flex items-center gap-4 pt-4 border-t border-slate-100 dark:border-white/5">
                <PrimaryButton :disabled="form.processing">
                    Save Changes
                </PrimaryButton>

                <Transition
                    enter-active-class="transition ease-in-out duration-300"
                    enter-from-class="opacity-0 translate-x-[-10px]"
                    enter-to-class="opacity-100 translate-x-0"
                    leave-active-class="transition ease-in-out duration-300"
                    leave-from-class="opacity-100 translate-x-0"
                    leave-to-class="opacity-0 translate-x-[-10px]"
                >
                    <div v-if="form.recentlySuccessful" class="flex items-center gap-2 text-sm text-emerald-600 dark:text-emerald-400 font-medium bg-emerald-50 dark:bg-emerald-500/10 px-3 py-1.5 rounded-full">
                        <span class="material-symbols-outlined text-[18px]">check_circle</span>
                        <span>Saved successfully.</span>
                    </div>
                </Transition>
            </div>
        </form>
    </section>
</template>
