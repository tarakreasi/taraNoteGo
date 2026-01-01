import { Node, mergeAttributes } from '@tiptap/core'

export default Node.create({
    name: 'layoutWrapper',

    group: 'block',

    content: 'layoutColumn+',

    isolating: true,

    addAttributes() {
        return {
            type: {
                default: 'two-col',
                parseHTML: element => element.getAttribute('data-layout-type'),
                renderHTML: attributes => ({ 'data-layout-type': attributes.type }),
            },
        }
    },

    parseHTML() {
        return [
            {
                tag: 'div[data-type="layoutWrapper"]',
            },
        ]
    },

    renderHTML({ HTMLAttributes }) {
        let classes = 'grid gap-4 mb-4 mt-4 w-full'

        // Check type to determine grid columns
        // We render attributes to specific class names for Tailwind to pick up
        if (HTMLAttributes['data-layout-type'] === 'three-col') {
            classes += ' grid-cols-1 md:grid-cols-3'
        } else {
            classes += ' grid-cols-1 md:grid-cols-2'
        }

        return ['div', mergeAttributes(HTMLAttributes, { 'data-type': 'layoutWrapper', class: classes }), 0]
    },
})
