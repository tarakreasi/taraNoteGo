# Design Tokens - TaraNote V6

**Design System:** V6 "Electric Earth"  
**Version:** Beta 0.9.80  
**Last Updated:** December 27, 2025



## Purpose

This document defines the **design tokens** - the centralized design decisions that ensure visual consistency across TaraNote. All values are defined in `tailwind.config.js` and cascaded throughout the application.

**What are Design Tokens?**  
Named values representing design decisions (colors, spacing, typography, etc.). They ensure consistency and make system-wide changes easy.



## Color Tokens

### Primary Colors

| Token Name | Light Mode | Dark Mode | Usage |
|:-----------|:-----------|:----------|:------|
| `primary` | `#135bec` | `#3b82f6` | Primary actions, links, brand |
| `primary-dark` | `#0e4bb8` | `#2563eb` | Hover states, emphasis |
| `accent-orange` | `#FF5F1F` | `#fb923c` | Highlights, CTAs, warmth |
| `success` | `#10b981` | `#34d399` | Success states, positive actions |
| `warning` | `#f59e0b` | `#fbbf24` | Warning states, caution |
| `danger` | `#ef4444` | `#f87171` | Errors, destructive actions |

### Neutral Colors (Slate Scale)

| Token | Light | Dark | Usage |
|:------|:------|:-----|:------|
| `gray-50` | `#f8fafc` | N/A | Light mode background |
| `gray-100` | `#f1f5f9` | N/A | Subtle backgrounds |
| `gray-200` | `#e2e8f0` | N/A | Borders, dividers |
| `gray-300` | `#cbd5e1` | N/A | Disabled states |
| `gray-400` | `#94a3b8` | `#94a3b8` | Placeholder text |
| `gray-500` | `#64748b` | `#64748b` | Secondary text |
| `gray-600` | `#475569` | `#cbd5e1` | Body text (light mode) |
| `gray-700` | `#334155` | `#e2e8f0` | Headings (light mode) |
| `gray-800` | `#1e293b` | `#f1f5f9` | Primary text (light mode) |
| `gray-900` | `#0f172a` | N/A | Dark mode background |

### Glass Surface Tokens

| Token | Light Mode | Dark Mode |
|:------|:-----------|:----------|
| `glass-bg` | `rgba(255,255,255,0.6)` | `rgba(15,23,42,0.8)` |
| `glass-border` | `rgba(0,0,0,0.05)` | `rgba(255,255,255,0.1)` |
| `glass-blur` | `blur(20px)` | `blur(40px)` |



## Typography Tokens

### Font Families

```javascript
{
  display: ['Outfit', 'sans-serif'],  // Headings, brand
  body: ['Inter', 'sans-serif'],      // UI text, content
  mono: ['JetBrains Mono', 'monospace'] // Code
}
```

### Font Sizes

| Token | Size | Line Height | Usage |
|:------|:-----|:------------|:------|
| `text-xs` | 0.75rem (12px) | 1rem | Labels, captions |
| `text-sm` | 0.875rem (14px) | 1.25rem | Small UI text |
| `text-base` | 1rem (16px) | 1.5rem | Body text |
| `text-lg` | 1.125rem (18px) | 1.75rem | Large body |
| `text-xl` | 1.25rem (20px) | 1.75rem | Small headings |
| `text-2xl` | 1.5rem (24px) | 2rem | H3 |
| `text-3xl` | 1.875rem (30px) | 2.25rem | H2 |
| `text-4xl` | 2.25rem (36px) | 2.5rem | H1 (mobile) |
| `text-5xl` | 3rem (48px) | 1 | H1 (desktop) |

### Font Weights

| Token | Value | Usage |
|:------|:------|:------|
| `font-normal` | 400 | Body text |
| `font-medium` | 500 | UI elements, emphasis |
| `font-semibold` | 600 | Subheadings |
| `font-bold` | 700 | Headings |
| `font-black` | 900 | Hero titles, display |



## Spacing Tokens

### Spacing Scale (4px base unit)

| Token | Size | Usage |
|:------|:-----|:------|
| `0` | 0px | Reset |
| `1` | 0.25rem (4px) | Micro spacing |
| `2` | 0.5rem (8px) | Tiny gaps |
| `3` | 0.75rem (12px) | Small gaps |
| `4` | 1rem (16px) | Default gap |
| `5` | 1.25rem (20px) | Medium gap |
| `6` | 1.5rem (24px) | Large gap |
| `8` | 2rem (32px) | Section spacing |
| `10` | 2.5rem (40px) | Large sections |
| `12` | 3rem (48px) | Page sections |
| `16` | 4rem (64px) | Hero spacing |
| `20` | 5rem (80px) | Extra large |



## Border Radius Tokens

| Token | Size | Usage |
|:------|:-----|:------|
| `rounded-sm` | 0.125rem (2px) | Subtle rounding |
| `rounded` | 0.25rem (4px) | Default buttons |
| `rounded-md` | 0.375rem (6px) | Cards |
| `rounded-lg` | 0.5rem (8px) | Panels, modals |
| `rounded-xl` | 0.75rem (12px) | Large cards |
| `rounded-2xl` | 1rem (16px) | Hero elements |
| `rounded-3xl` | 1.5rem (24px) | Soft cards |
| `rounded-full` | 9999px | Circles, pills |



## Shadow Tokens

| Token | Value | Usage |
|:------|:------|:------|
| `shadow-sm` | `0 1px 2px 0 rgb(0 0 0 / 0.05)` | Subtle elevation |
| `shadow` | `0 1px 3px 0 rgb(0 0 0 / 0.1)` | Cards |
| `shadow-md` | `0 4px 6px -1px rgb(0 0 0 / 0.1)` | Dropdowns |
| `shadow-lg` | `0 10px 15px -3px rgb(0 0 0 / 0.1)` | Modals |
| `shadow-xl` | `0 20px 25px -5px rgb(0 0 0 / 0.1)` | Popovers |
| `shadow-2xl` | `0 25px 50px -12px rgb(0 0 0 / 0.25)` | Hero elements |

**Glow Shadows** (for glassmorphism):
- `shadow-primary/20`: Blue glow for primary elements
- `shadow-accent-orange/20`: Orange glow for CTAs



## Z-Index Scale

```javascript
{
  dropdown: 1000,
  sticky: 1020,
  fixed: 1030,
  modalBackdrop: 1040,
  modal: 1050,
  popover: 1060,
  tooltip: 1070,
  noiseOverlay: 9999  // Always on top
}
```



## Breakpoint Tokens

| Token | Min Width | Device |
|:------|:----------|:-------|
| `sm` | 640px | Mobile landscape |
| `md` | 768px | Tablet portrait |
| `lg` | 1024px | Tablet landscape / Laptop |
| `xl` | 1280px | Desktop |
| `2xl` | 1536px | Large desktop |

**Usage:**
```vue
<div class="text-base md:text-lg lg:text-xl">
  Responsive typography
</div>
```



## Animation Tokens

### Duration

| Token | Value | Usage |
|:------|:-----|:------|
| `duration-75` | 75ms | Instant feedback |
| `duration-100` | 100ms | Hover states |
| `duration-150` | 150ms | Quick transitions |
| `duration-200` | 200ms | Default transitions |
| `duration-300` | 300ms | Smooth transitions |
| `duration-500` | 500ms | Revealing content |
| `duration-700` | 700ms | Complex animations |
| `duration-1000` | 1000ms | Dramatic effects |

### Easing

| Token | Curve | Usage |
|:------|:------|:------|
| `ease-linear` | Linear | Continuous motion |
| `ease-in` | Cubic-bezier(0.4, 0, 1, 1) | Accelerating |
| `ease-out` | Cubic-bezier(0, 0, 0.2, 1) | Decelerating (DEFAULT) |
| `ease-in-out` | Cubic-bezier(0.4, 0, 0.2, 1) | Smooth start/end |



## Glassmorphism Tokens

### Glass Variants

**Deep Glass** (dark emphasis):
```css
background: rgba(255, 255, 255, 0.05);
backdrop-filter: blur(40px);
border: 1px solid rgba(255, 255, 255, 0.1);
```

**Milky Glass** (light subtle):
```css
background: rgba(255, 255, 255, 0.6);
backdrop-filter: blur(20px);
border: 1px solid rgba(0, 0, 0, 0.05);
```



## Accessibility Tokens

### Minimum Touch Targets

```javascript
{
  minTouchTarget: '44px',  // iOS HIG, WCAG requirement
  comfortableTap: '48px'   // Recommended
}
```

### Focus Ring

```css
outline: 2px solid currentColor;
outline-offset: 2px;
```



## Usage in Code

### Tailwind CSS Classes
```vue
<!-- Color -->
<button class="bg-primary text-white hover:bg-primary-dark">

<!-- Typography -->
<h1 class="font-display font-bold text-4xl md:text-5xl">

<!-- Spacing -->
<div class="p-6 gap-4 mb-8">

<!-- Glassmorphism -->
<div class="bg-white/60 dark:bg-white/5 backdrop-blur-xl border border-white/20">
```

### CSS Custom Properties
```css
:root {
  --color-primary: #135bec;
  --font-display: 'Outfit', sans-serif;
  --spacing-unit: 4px;
}
```

### Vue Composables
```javascript
import { useTheme } from '@/composables/useTheme';
const { isDark } = useTheme();
```



## Design Token Benefits

- **Consistency** - One source of truth for all design values  
- **Maintainability** - Change once, update everywhere  
- **Scalability** - Easy to add new tokens as system grows  
- **Accessibility** - Built-in compliance with standards  
- **Documentation** - Clear naming makes code self-documenting  



**Version:** Beta 0.9.60 (UI Design V6)  
**Maintained by:** TaraKreasi (Tri Wantoro)  
**Last Updated:** December 21, 2025

> **For Developers:** All tokens are defined in `tailwind.config.js`. Use Tailwind utilities instead of arbitrary values for consistency.
