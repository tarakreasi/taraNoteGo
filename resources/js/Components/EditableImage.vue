<script setup>
import { ref, computed } from 'vue';
import { usePage } from '@inertiajs/vue3';

const props = defineProps({
    settingKey: {
        type: String,
        required: true
    },
    initialSrc: {
        type: String,
        default: null
    },
    alt: {
        type: String,
        default: 'Image'
    },
    imgClass: {
        type: String,
        default: ''
    }
});

const page = usePage();
const user = computed(() => page.props.auth?.user);
const currentSrc = ref(props.initialSrc);
const isUploading = ref(false);
const fileInput = ref(null);

const triggerUpload = () => {
    if (user.value) {
        fileInput.value.click();
    }
};

const handleFileChange = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    isUploading.value = true;
    const formData = new FormData();
    formData.append('image', file);

    try {
        // Upload 
        const res = await window.axios.post('/settings/upload', formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });
        
        const newUrl = res.data.url;
        currentSrc.value = newUrl;

        // Save Setting Key
        await window.axios.post('/settings/update', {
            settings: {
                [props.settingKey]: newUrl
            }
        });

    } catch (error) {
        console.error("Upload failed", error);
        alert("Failed to upload image.");
    } finally {
        isUploading.value = false;
    }
};
</script>

<template>
    <div class="relative group/edit leading-none" :class="{ 'cursor-pointer': !!user }" @click="triggerUpload">
        
        <!-- Hidden Input -->
        <input 
            v-if="user"
            ref="fileInput" 
            type="file" 
            accept="image/*" 
            class="hidden" 
            @change="handleFileChange"
        />

        <!-- Content -->
        <!-- If we have an image URL, show IMG -->
        <img 
            v-if="currentSrc" 
            :src="currentSrc" 
            :alt="alt" 
            :class="imgClass"
            class="w-full h-full object-cover"
        />
        
        <!-- Else (if fallback visual slot exists), show slot -->
        <slot v-else />

        <!-- Admin Overlay -->
        <div 
            v-if="user" 
            class="absolute inset-0 bg-blue-500/20 border-2 border-dashed border-blue-500 opacity-0 group-hover/edit:opacity-100 transition-opacity flex items-center justify-center rounded z-20 pointer-events-none"
        >
            <span class="bg-white text-blue-600 px-2 py-1 text-xs font-bold rounded shadow-sm">
                {{ isUploading ? 'Uploading...' : 'Click to Change Image' }}
            </span>
        </div>
    </div>
</template>
