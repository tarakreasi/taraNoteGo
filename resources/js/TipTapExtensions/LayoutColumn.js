import { Node, mergeAttributes } from '@tiptap/core'

export default Node.create({
    name: 'layoutColumn',

    content: 'block+',

    isolating: true,

    parseHTML() {
        return [
            {
                tag: 'div[data-type="layoutColumn"]',
            },
        ]
    },

    renderHTML({ HTMLAttributes }) {
        return ['div', mergeAttributes(HTMLAttributes, {
            'data-type': 'layoutColumn',
            class: 'min-w-0 flex flex-col gap-2' // min-w-0 needed for grid child truncation
        }), 0]
    },
})
