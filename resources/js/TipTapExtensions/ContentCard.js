import { Node, mergeAttributes } from '@tiptap/core'

export default Node.create({
    name: 'contentCard',

    group: 'block',

    content: 'block+',

    draggable: true,

    addAttributes() {
        return {
            class: {
                default: 'bg-white p-6 rounded-2xl border border-slate-200 dark:border-slate-700 shadow-sm mb-4',
            },
            style: {
                default: '',
                parseHTML: element => element.getAttribute('style'),
            }
        }
    },

    parseHTML() {
        return [
            {
                tag: 'div[data-type="contentCard"]',
            },
        ]
    },

    renderHTML({ HTMLAttributes }) {
        return ['div', mergeAttributes(HTMLAttributes, { 'data-type': 'contentCard' }), 0]
    },
})
