import { ref, computed } from 'vue';

export function useNotes(props) {
    const searchQuery = ref('');
    const selectedNotebookId = ref(null);
    const selectedArticle = ref(null);

    const filteredNotes = computed(() => {
        let filtered = props.notes || [];

        // Filter by notebook
        if (selectedNotebookId.value) {
            filtered = filtered.filter(note => note.notebook_id === selectedNotebookId.value);
        }

        // Filter by search query (Client-side refinement)
        if (searchQuery.value) {
            const query = searchQuery.value.toLowerCase();
            filtered = filtered.filter(note =>
                note.title?.toLowerCase().includes(query) ||
                note.content?.toLowerCase().includes(query)
            );
        }

        return filtered;
    });

    const currentNotebookName = computed(() => {
        if (!selectedNotebookId.value) return 'All Notes';
        return props.notebooks.find(n => n.id === selectedNotebookId.value)?.name || 'All Notes';
    });

    const selectNotebook = (id) => {
        selectedNotebookId.value = id;
        selectedArticle.value = null;
    };

    const openArticle = (article) => {
        selectedArticle.value = article;
    };

    return {
        searchQuery,
        selectedNotebookId,
        selectedArticle,
        filteredNotes,
        currentNotebookName,
        selectNotebook,
        openArticle
    };
}
