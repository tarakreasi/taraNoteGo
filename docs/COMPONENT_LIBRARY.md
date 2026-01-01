# Component Library - TaraNote V6

**Design System:** V6 "Electric Earth"  
**Version:** Beta 0.9.80  
**Last Updated:** December 27, 2025



## Purpose

This document provides specifications for all reusable UI components in TaraNote. Each component includes visual examples, props documentation, accessibility notes, and usage guidelines.



## Table of Contents

1. [Buttons](#buttons)
2. [Form Elements](#form-elements)
3. [Cards](#cards)
4. [Navigation](#navigation)
5. [Modals & Toasts](#modals--toasts)
6. [Loading States](#loading-states)
7. [Empty States](#empty-states)



## Buttons

### Primary Button

**Usage:** Main call-to-action (submit forms, publish, create)

**Specs:**
- Background: `bg-primary`
- Text: `text-white`
- Hover: `hover:bg-primary-dark`
- Border radius: `rounded-lg` (8px)
- Padding: `px-6 py-3` (24px/12px)
- Font: `font-medium text-base`
- Min height: 44px (touch target)

**Example:**
```vue
<button class="bg-primary text-white px-6 py-3 rounded-lg font-medium hover:bg-primary-dark transition-colors">
  <i class="ri-check-line mr-2"></i> Publish Article
</button>
```

**Accessibility:**
- Focusable via keyboard
- Clear focus indicator
- Sufficient contrast (16.3:1)
- Descriptive button text



### Secondary Button

**Usage:** Alternative actions (cancel, back)

**Specs:**
- Background: `bg-white dark:bg-gray-800`
- Text: `text-gray-700 dark:text-gray-200`
- Border: `border border-gray-300 dark:border-gray-600`
- Hover: `hover:bg-gray-50 dark:hover:bg-gray-700`

**Example:**
```vue
<button class="bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-200 px-6 py-3 rounded-lg font-medium border border-gray-300 dark:border-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
  Cancel
</button>
```



### Ghost Button

**Usage:** Tertiary actions, icon buttons

**Specs:**
- Background: Transparent
- Text: `text-gray-600 dark:text-gray-400`
- Hover: `hover:bg-black/5 dark:hover:bg-white/5`
- Border: None
- Padding: `px-4 py-2`

**Example:**
```vue
<button class="text-gray-600 dark:text-gray-400 px-4 py-2 rounded-lg hover:bg-black/5 dark:hover:bg-white/5 transition-colors">
  <i class="ri-settings-3-line"></i>
  Settings
</button>
```



### Danger Button

**Usage:** Destructive actions (delete, remove)

**Specs:**
- Background: `bg-danger`
- Text: `text-white`
- Hover: `hover:bg-red-600`
- Padding: Same as primary

**Example:**
```vue
<button class="bg-danger text-white px-6 py-3 rounded-lg font-medium hover:bg-red-600 transition-colors">
  Delete Note
</button>
```

**Best Practice:** Always confirm destructive actions with a modal.



## Form Elements

### Text Input

**Specs:**
- Height: `h-12` (48px) - comfortable touch target
- Padding: `px-4 py-3`
- Border: `border border-gray-200 dark:border-white/10`
- Border radius: `rounded-lg`
- Focus: `focus:ring-2 focus:ring-primary/20 focus:border-primary`
- Background: `bg-white/50 dark:bg-white/5`

**Example:**
```vue
<input 
  type="text"
  placeholder="Search articles..."
  class="w-full h-12 bg-white/50 dark:bg-white/5 border border-gray-200 dark:border-white/10 rounded-lg px-4 py-3 text-sm focus:ring-2 focus:ring-primary/20 focus:border-primary transition-all placeholder:text-gray-400"
>
```

**Accessibility:**
- Associated label (`<label for="id">`)
- Placeholder is NOT the label
- Error state announced to screen readers
- Minimum 44px touch target



### Textarea

**Specs:**
- Min height: `h-32` (128px)
- Max height: `max-h-96` (384px)
- Resize: `resize-y` (vertical only)
- Rest: Same as text input

**Example:**
```vue
<textarea 
  class="w-full min-h-32 max-h-96 resize-y bg-white/50 dark:bg-white/5 border border-gray-200 dark:border-white/10 rounded-lg px-4 py-3 focus:ring-2 focus:ring-primary/20 focus:border-primary"
  placeholder="Write your note..."
></textarea>
```



### Select Dropdown

**Specs:**
- Height: 48px (same as input)
- Icon: Remixicon `ri-arrow-down-s-line`
- Background: Same as input
- Options: Light background, hover state

**Example:**
```vue
<select class="w-full h-12 bg-white/50 dark:bg-white/5 border border-gray-200 dark:border-white/10 rounded-lg px-4 py-3 focus:ring-2 focus:ring-primary/20 focus:border-primary">
  <option>All Notes</option>
  <option>Published</option>
  <option>Drafts</option>
</select>
```



### Checkbox

**Specs:**
- Size: `w-5 h-5` (20px)
- Border: `border-2 border-gray-300`
- Checked: `bg-primary border-primary`
- Rounded: `rounded` (4px)
- Focus: Ring effect

**Example:**
```vue
<input 
  type="checkbox" 
  class="w-5 h-5 text-primary border-2 border-gray-300 rounded focus:ring-2 focus:ring-primary/20"
>
<label class="ml-2 text-sm font-medium">Featured Article</label>
```



## Cards

### Glass Card

**Usage:** Content containers, panels

**Specs:**
- Background: `bg-white/60 dark:bg-white/10` (glassmorphism)
- Backdrop blur: `backdrop-blur-md` (12px)
- Border: `border border-white/20 dark:border-white/5`
- Border radius: `rounded-xl` (12px)
- Padding: `p-6`
- Shadow: `shadow-lg`

**Example:**
```vue
<div class="bg-white/60 dark:bg-white/10 backdrop-blur-md border border-white/20 dark:border-white/5 rounded-xl p-6 shadow-lg">
  <h3 class="font-display font-bold text-xl mb-2">Card Title</h3>
  <p class="text-gray-600 dark:text-gray-300">Card content goes here...</p>
</div>
```



### Article Card

**Usage:** Article list items

**Specs:**
- Layout: Vertical stack
- Hover: `hover:bg-white/80 dark:hover:bg-white/10`
- Cursor: `cursor-pointer`
- Transition: `transition-all`
- Padding: `p-4`

**Example:**
```vue
<div class="group p-4 rounded-lg cursor-pointer transition-all hover:bg-white/60 dark:hover:bg-white/10 hover:shadow-sm">
  <h3 class="font-semibold text-sm text-gray-900 dark:text-white line-clamp-2 mb-1">
    Article Title
  </h3>
  <div class="flex items-center gap-2 mt-3 text-[10px] text-gray-400 font-medium">
    <i class="ri-time-line text-[12px]"></i> 
    2h ago
  </div>
</div>
```



## Navigation

### Sidebar Item

**Specs:**
- Layout: Flex (icon + text + badge)
- Padding: `px-3 py-2`
- Border radius: `rounded-lg`
- Active: `bg-white/60 dark:bg-white/10 text-primary`
- Inactive: `text-gray-600 dark:text-gray-400`
- Hover: `hover:bg-black/5 dark:hover:bg-white/5`

**Example:**
```vue
<button class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-colors text-sm group text-left"
        :class="isActive ? 'bg-white/60 dark:bg-white/10 text-primary font-medium shadow-sm' : 'text-gray-600 dark:text-gray-400 hover:bg-black/5 dark:hover:bg-white/5'">
  <i class="ri-dashboard-line text-[20px]"></i>
  <span class="flex-1">All Notes</span>
  <span class="text-xs bg-gray-100 dark:bg-white/10 px-1.5 rounded text-gray-500">42</span>
</button>
```



### Breadcrumbs

**Specs:**
- Separator: Material `chevron_right`
- Text: `text-sm`
- Current: `text-primary font-medium`
- Links: `hover:text-primary

`

**Example:**
```vue
<nav class="flex flex-wrap gap-2 items-center text-sm text-gray-500 dark:text-gray-400">
  <Link href="/" class="hover:text-primary transition-colors">Home</Link>
  <span class="material-symbols-outlined text-[16px]">chevron_right</span>
  <Link href="/articles" class="hover:text-primary transition-colors">Articles</Link>
  <span class="material-symbols-outlined text-[16px]">chevron_right</span>
  <span class="text-primary font-medium">Current Page</span>
</nav>
```



## Modals & Toasts

### Modal Overlay

**Specs:**
- Background: `bg-black/50` (50% opacity)
- Backdrop blur: `backdrop-blur-sm`
- Z-index: `z-50`
- Position: Fixed, full viewport

**Modal Content:**
- Background: White/Dark
- Border radius: `rounded-2xl`
- Max width: `max-w-md`
- Shadow: `shadow-2xl`
- Padding: `p-6`

**Example:**
```vue
<!-- Overlay -->
<div class="fixed inset-0 bg-black/50 backdrop-blur-sm z-50 flex items-center justify-center">
  <!-- Modal -->
  <div class="bg-white dark:bg-gray-900 rounded-2xl shadow-2xl max-w-md w-full p-6">
    <h2 class="font-display font-bold text-2xl mb-4">Confirm Action</h2>
    <p class="text-gray-600 dark:text-gray-300 mb-6">Are you sure you want to delete this?</p>
    <div class="flex gap-3">
      <button class="flex-1 bg-gray-100 text-gray-700 px-4 py-2 rounded-lg">Cancel</button>
      <button class="flex-1 bg-danger text-white px-4 py-2 rounded-lg">Delete</button>
    </div>
  </div>
</div>
```

**Accessibility:**
- Focus trap (Tab cycles within modal)
- ESC key closes modal
- `role="dialog"` and `aria-modal="true"`
- Focus returned to trigger on close



## Loading States

### Spinner

**Specs:**
- Icon: Material `progress_activity`
- Animation: `animate-spin`
- Size: `text-4xl` (36px)
- Color: `text-primary`

**Example:**
```vue
<div class="flex justify-center items-center h-full">
  <span class="material-symbols-outlined animate-spin text-4xl text-primary opacity-50">
    progress_activity
  </span>
</div>
```



### Loading Text

**Example:**
```vue
<div class="text-xs text-gray-400">Loading notes...</div>
```

**Accessibility:**
- `aria-live="polite"` for screen readers
- Clear loading message



## Empty States

### No Results

**Specs:**
- Icon: Large, centered
- Text: Clear, helpful message
- CTA: Optional action button

**Example:**
```vue
<div class="flex flex-col items-center justify-center h-full text-center px-6 opacity-60">
  <span class="material-symbols-outlined text-5xl mb-2 text-gray-400">article</span>
  <p class="text-sm text-gray-500 dark:text-gray-400">No articles found</p>
  <button class="mt-4 text-primary text-sm font-medium">Create your first article</button>
</div>
```



## Best Practices

### Do's
- Use semantic HTML (`<button>`, `<nav>`, `<main>`)
- Provide clear focus indicators
- Maintain minimum 44px touch targets
- Use descriptive labels and placeholders
- Test keyboard navigation
- Ensure sufficient color contrast

### Don'ts
- Don't use `<div>` for buttons
- Don't rely on color alone for meaning
- Don't remove focus outlines
- Don't use placeholder as label
- Don't trap users in modals without escape
- Don't use arbitrary CSS values (use tokens)



**Version:** Beta 0.9.60 (UI Design V6)  
**Maintained by:** TaraKreasi (Tri Wantoro)  
**Last Updated:** December 21, 2025

> **For Developers:** Always use design tokens from `DESIGN_TOKENS.md` instead of hard-coded values. Every component should be keyboard-accessible and screen-reader friendly.
