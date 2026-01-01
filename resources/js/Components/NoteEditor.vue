<script setup>
import { useEditor, EditorContent } from '@tiptap/vue-3';
import StarterKit from '@tiptap/starter-kit';
import Placeholder from '@tiptap/extension-placeholder';
import Image from '@tiptap/extension-image';
import TaskList from '@tiptap/extension-task-list';
import TaskItem from '@tiptap/extension-task-item';
import { watch, onBeforeUnmount, ref, onMounted } from 'vue';

const props = defineProps({
    modelValue: {
        type: String,
        default: '',
    },
    placeholder: {
        type: String,
        default: 'Start writing...',
    },
    lastSavedAt: {
        type: [String, Object], // Accept String or Date object
        default: null
    },
    saveStatus: {
        type: String,
        default: ''
    }
});

const emit = defineEmits(['update:modelValue', 'save', 'request-create-space']);

const editor = useEditor({
    content: props.modelValue,
    extensions: [
        StarterKit,
        Placeholder.configure({
            placeholder: props.placeholder,
        }),
        Image,
        TaskList,
        TaskItem.configure({
            nested: true,
        }),
    ],
    editorProps: {
        attributes: {
            class: 'prose prose-lg dark:prose-invert max-w-none focus:outline-none min-h-[500px] outline-none px-8 py-6 mb-20',
        },
    },
    onUpdate: ({ editor }) => {
        emit('update:modelValue', editor.getHTML());
        emit('save'); // Notify parent to autosave
    },
});

// Watch for external content changes (e.g. switching notes)
watch(() => props.modelValue, (newValue) => {
    if (editor.value && newValue !== editor.value.getHTML()) {
        editor.value.commands.setContent(newValue, false);
    }
});

onBeforeUnmount(() => {
    if (editor.value) {
        editor.value.destroy();
    }
});

// --- Relative Time Logic ---
const timeAgo = ref('');
let timeInterval = null;

const updateTimeAgo = () => {
    if (!props.lastSavedAt) {
        timeAgo.value = '';
        return;
    }
    const savedDate = new Date(props.lastSavedAt);
    const now = new Date();
    const diffMs = now - savedDate;
    const diffSec = Math.floor(diffMs / 1000);
    const diffMin = Math.floor(diffSec / 60);

    if (diffSec < 10) {
        timeAgo.value = 'Just now';
    } else if (diffSec < 60) {
        timeAgo.value = 'Few seconds ago';
    } else if (diffMin < 60) {
        timeAgo.value = `${diffMin} min ago`;
    } else {
        timeAgo.value = savedDate.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
    }
};

onMounted(() => {
    updateTimeAgo();
    timeInterval = setInterval(updateTimeAgo, 30000); // Update every 30s
});

watch(() => props.lastSavedAt, () => {
    updateTimeAgo();
});

onBeforeUnmount(() => {
    if (timeInterval) clearInterval(timeInterval);
});

// --- Markdown Hints Popover ---
const showMarkdownHints = ref(false);

const markdownShortcuts = [
    { label: 'Heading 1', syntax: '# Space' },
    { label: 'Heading 2', syntax: '## Space' },
    { label: 'Bold', syntax: '**text**' },
    { label: 'Italic', syntax: '*text*' },
    { label: 'Bullet List', syntax: '- Space' },
    { label: 'Task List', syntax: '- [ ] Space' },
    { label: 'Code Block', syntax: '```Enter' },
];

</script>

<template>
    <div class="relative w-full flex flex-col h-full font-sans" v-if="editor">
        <!-- FIXED MINIMALIST TOOLBAR -->
        <div class="sticky top-0 z-10 w-full bg-white/90 dark:bg-[#0F172A]/90 backdrop-blur-md border-b border-slate-200 dark:border-white/5 px-6 py-2 flex items-center gap-1 shrink-0 overflow-x-auto no-scrollbar">
            
            <!-- History -->
            <div class="flex items-center gap-0.5 mr-3 border-r border-slate-200 dark:border-white/10 pr-3">
                <button @click="editor.chain().focus().undo().run()" :disabled="!editor.can().undo()" class="editor-btn" title="Undo">
                    <span class="material-symbols-outlined text-[18px]">undo</span>
                </button>
                <button @click="editor.chain().focus().redo().run()" :disabled="!editor.can().redo()" class="editor-btn" title="Redo">
                    <span class="material-symbols-outlined text-[18px]">redo</span>
                </button>
            </div>

            <!-- Text Formatting -->
            <div class="flex items-center gap-0.5 mr-3 border-r border-slate-200 dark:border-white/10 pr-3">
                <button @click="editor.chain().focus().toggleBold().run()" :class="{ 'is-active': editor.isActive('bold') }" class="editor-btn" title="Bold">
                    <span class="material-symbols-outlined text-[18px]">format_bold</span>
                </button>
                <button @click="editor.chain().focus().toggleItalic().run()" :class="{ 'is-active': editor.isActive('italic') }" class="editor-btn" title="Italic">
                    <span class="material-symbols-outlined text-[18px]">format_italic</span>
                </button>
                <button @click="editor.chain().focus().toggleStrike().run()" :class="{ 'is-active': editor.isActive('strike') }" class="editor-btn" title="Strikethrough">
                    <span class="material-symbols-outlined text-[18px]">strikethrough_s</span>
                </button>
                 <button @click="editor.chain().focus().toggleCode().run()" :class="{ 'is-active': editor.isActive('code') }" class="editor-btn" title="Code">
                    <span class="material-symbols-outlined text-[18px]">code</span>
                </button>
            </div>

            <!-- Headings -->
            <div class="flex items-center gap-0.5 mr-3 border-r border-slate-200 dark:border-white/10 pr-3">
                <button @click="editor.chain().focus().toggleHeading({ level: 1 }).run()" :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }" class="editor-btn" title="Heading 1">
                    <span class="material-symbols-outlined text-[18px]">title</span>
                </button>
                <button @click="editor.chain().focus().toggleHeading({ level: 2 }).run()" :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }" class="editor-btn" title="Heading 2">
                    <span class="font-bold text-[13px]">H2</span>
                </button>
                 <button @click="editor.chain().focus().toggleHeading({ level: 3 }).run()" :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }" class="editor-btn" title="Heading 3">
                    <span class="font-bold text-[11px]">H3</span>
                </button>
            </div>

            <!-- Lists -->
            <div class="flex items-center gap-0.5">
                <button @click="editor.chain().focus().toggleBulletList().run()" :class="{ 'is-active': editor.isActive('bulletList') }" class="editor-btn" title="Bullet List">
                    <span class="material-symbols-outlined text-[18px]">format_list_bulleted</span>
                </button>
                <button @click="editor.chain().focus().toggleOrderedList().run()" :class="{ 'is-active': editor.isActive('orderedList') }" class="editor-btn" title="Ordered List">
                    <span class="material-symbols-outlined text-[18px]">format_list_numbered</span>
                </button>
                <button @click="editor.chain().focus().toggleTaskList().run()" :class="{ 'is-active': editor.isActive('taskList') }" class="editor-btn" title="Task List">
                    <span class="material-symbols-outlined text-[18px]">check_box</span>
                </button>
                <button @click="editor.chain().focus().toggleCodeBlock().run()" :class="{ 'is-active': editor.isActive('codeBlock') }" class="editor-btn" title="Code Block">
                    <span class="material-symbols-outlined text-[18px]">data_object</span>
                </button>
            </div>

             <div class="ml-auto flex items-center gap-2 text-[10px] text-slate-400 uppercase tracking-widest font-medium">
                <span v-if="editor.storage.characterCount">{{ editor.storage.characterCount.characters() }} CHARS</span>
             </div>
        </div>

        <editor-content :editor="editor" class="flex-1 overflow-y-auto custom-scrollbar relative bg-slate-50 dark:bg-[#0F172A]" />

        <!-- NEW REFINED FOOTER -->
        <div class="absolute bottom-0 left-0 right-0 bg-white/95 dark:bg-[#0F172A]/95 backdrop-blur border-t border-slate-200 dark:border-white/5 py-1 px-4 flex items-center justify-between text-[11px] text-slate-500 select-none z-20">
            <!-- Left: Status -->
            <div class="w-1/3 text-left flex items-center gap-2">
                <div v-if="saveStatus === 'Saving...'" class="flex items-center text-blue-500">
                    <span class="material-symbols-outlined text-[12px] animate-spin mr-1">sync</span>
                    Saving
                </div>
                <!-- Use timeAgo ref -->
                <div v-else-if="timeAgo" class="flex items-center">
                    <span class="material-symbols-outlined text-[14px] mr-1 text-slate-400">cloud_done</span>
                    <span class="text-slate-400">Saved {{ timeAgo }}</span>
                </div>
            </div>

            <!-- Center: Shortcuts -->
            <div class="flex-1 text-center font-mono text-slate-400 hidden md:block opacity-70">
                <span class="mx-2">Ctrl+S Save</span>
                <span class="mx-2 text-slate-300">|</span>
                <span class="mx-2">Ctrl+Enter Publish</span>
            </div>

            <!-- Right: Hints & New Space -->
            <div class="w-1/3 flex items-center justify-end gap-3">
                <div class="relative">
                    <button 
                        @click="showMarkdownHints = !showMarkdownHints"
                        class="flex items-center gap-1 hover:text-slate-800 dark:hover:text-white transition-colors focus:outline-none"
                    >
                        <span class="material-symbols-outlined text-[14px]">keyboard_command_key</span>
                        <span>Shortcuts</span>
                    </button>

                    <!-- Markdown Hints Popover -->
                    <div v-if="showMarkdownHints" class="absolute bottom-full right-0 mb-2 w-56 bg-white dark:bg-slate-900 rounded-lg shadow-xl shadow-slate-200/50 dark:shadow-none border border-slate-200 dark:border-slate-800 overflow-hidden transform origin-bottom-right transition-all p-2 z-50">
                        <div class="text-[9px] uppercase font-bold text-slate-400 mb-2 px-2 pt-1 tracking-wider">Markdown Shortcuts</div>
                        <div class="flex flex-col gap-0.5">
                            <div v-for="hint in markdownShortcuts" :key="hint.label" class="flex items-center justify-between px-2 py-1.5 rounded hover:bg-slate-50 dark:hover:bg-white/5">
                                <span class="text-slate-600 dark:text-slate-300 text-[11px]">{{ hint.label }}</span>
                                <code class="bg-slate-100 dark:bg-white/10 px-1 py-0.5 rounded text-[10px] font-mono text-slate-500 dark:text-slate-400">{{ hint.syntax }}</code>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="w-[1px] h-3 bg-slate-200 dark:bg-white/10"></div>

                <button 
                    @click="$emit('request-create-space')"
                    class="flex items-center gap-1 hover:text-slate-800 dark:hover:text-white transition-colors font-medium"
                >
                    <span class="material-symbols-outlined text-[14px]">add</span>
                    Space
                </button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.editor-btn {
    @apply p-1.5 rounded-lg text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-white/5 transition-all flex items-center justify-center min-w-[32px] min-h-[32px];
}
.editor-btn.is-active {
    @apply bg-primary/10 text-primary dark:text-primary;
}
.editor-btn:disabled {
    @apply opacity-30 cursor-not-allowed;
}

/* Scoped Overrides for Tiptap */
:deep(.ProseMirror) {
    outline: none;
    min-height: 100%;
}
:deep(.ProseMirror p.is-editor-empty:first-child::before) {
  color: #94a3b8;
  content: attr(data-placeholder);
  float: left;
  height: 0;
  pointer-events: none;
}
:deep(.dark .ProseMirror p.is-editor-empty:first-child::before) {
    color: #64748b;
}
:deep(ul[data-type="taskList"]) {
  list-style: none;
  padding: 0;
}
:deep(ul[data-type="taskList"] li) {
  display: flex;
  align-items: flex-start;
  margin-bottom: 0.5rem;
}
:deep(ul[data-type="taskList"] li > label) {
  margin-right: 0.75rem;
  margin-top: 0.2rem;
  user-select: none;
  flex-shrink: 0;
}
:deep(ul[data-type="taskList"] li > div) {
  flex: 1;
}
</style>
