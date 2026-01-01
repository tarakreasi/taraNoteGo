# TaraNote UI/UX Design Overview

**Design System:** V6 "Electric Earth"  
**Version:** Beta 0.9.80  
**Designer:** TaraKreasi (Tri Wantoro)  
**Last Updated:** December 27, 2025



## Executive Summary

TaraNote's UI/UX design demonstrates a **modern, accessible, and user-centered approach** to building a dual-purpose note-taking and blogging platform. The V6 "Electric Earth" design system combines cutting-edge visual aesthetics (glassmorphism, ambient animations) with solid UX principles (clarity, consistency, accessibility).

**Key Achievement:** Successfully balancing **two distinct user experiences** within one application:
- **Private Workspace** - Productivity-focused, organized, functional
- **Public Garden** - Beautiful, engaging, reader-friendly



## Design Philosophy

### Core Principles

**1. Clarity Over Cleverness**  
Every design decision prioritizes user understanding over visual flair. Complex features are progressively disclosed.

**2. Living & Breathing Interface**  
Moving beyond static pages to create an "alive" experience through subtle animations, glassmorphism, and ambient elements.

**3. Accessibility First**  
WCAG 2.1 AA compliance from the start, not an afterthought. Every component is keyboard-navigable and screen-reader friendly.

**4. Consistent Metaphors**  
"Digital Garden" concept permeates the entire design - Spaces (notebooks), Cards (notes), Streams (lists).

**5. Performance Matters**  
Beautiful doesn't mean slow. Optimized animations, lazy loading, and efficient rendering.



## Design Vision: "Electric Earth"

### Concept

The V6 design system merges two seemingly opposing aesthetics:

**Electric** (Technology):
- Glassmorphism and blur effects
- Clean typography (Outfit + Inter)
- Precise geometric layouts
- Smooth, calculated animations

**Earth** (Organic):
- Ambient blob backgrounds
- Noise texture overlays
- Natural color palette (emerald accents)
- Breathing, floating movements

**Result:** An interface that feels **modern yet warm**, **digital yet tangible**, **structured yet alive**.



## User Experience Strategy

### Target Audiences

**Primary: Solo Creators & Developers**
- Need both private workspace and public portfolio
- Value customization and control
- Appreciate clean, modern aesthetics

**Secondary: Students & Learners**
- Organize study notes privately
- Share knowledge publicly (digital garden)
- Need simple, intuitive interface

**Tertiary: Writers & Bloggers**
- Draft content in private editor
- Publish as beautiful blog posts
- Minimal friction from draft -> publish

### Key User Flows

**1. Note Creation -> Publishing** (Most Critical)
```
Dashboard -> Create Note -> Write (Tiptap Editor) -> Publish -> Public Article
```
**Pain Points Solved:**
- One-click publish (no copy-pasting)
- Auto-generated slug for SEO
- Cover image handling
- Status management (Draft/Published)

**2. Portfolio Customization**
```
Settings -> Upload Avatar/Hero -> Set Username -> Customize Text -> View Portfolio
```
**Pain Points Solved:**
- No code knowledge required
- Live preview of changes
- File upload with validation
- Persistent settings

**3. Reader Discovery**
```
Landing Page -> Articles -> Search/Filter -> Read Article -> Author Portfolio
```
**Pain Points Solved:**
- Clear navigation
- Fast search (client-side)
- Beautiful reading experience
- Author discovery



## Visual Design System

### Color Strategy

**Semantic Color Naming:**
```
primary:     #135bec (Blue)  - Actions, links, focus states
accent:      #FF5F1F (Orange) - Highlights, CTAs, notifications
success:     #10b981 (Emerald)- Success states, positive actions
danger:      #ef4444 (Red)    - Destructive actions, errors
neutral:     Gray scale       - Text hierarchy, backgrounds
```

**Dark Mode Support:**
All colors have light/dark variants with **minimum 4.5:1 contrast ratio** (WCAG AA).

### Typography System

**Display (Headings):** Outfit
- Geometric, friendly, modern
- Used for: H1-H3, brand name, CTAs
- Weights: 700 (Bold), 900 (Black)

**Body (Content):** Inter
- Highly legible, neutral
- Used for: Paragraphs, UI text, forms
- Weights: 400 (Regular), 500 (Medium), 600 (Semibold)

**Scale:** 1.25 (major third)
- Base: 16px
- h1: 3rem (48px)
- h2: 2.25rem (36px)
- h3: 1.5rem (24px)
- body: 1rem (16px)

### Glassmorphism Implementation

**Two Glass Types:**

**Deep Glass** (Dark mode/emphasis):
- Blur: 40px
- Background: `rgba(255,255,255,0.05)`
- Border: `rgba(255,255,255,0.1)`
- Shadow: Large, soft

**Milky Glass** (Light mode/subtle):
- Blur: 20px  
- Background: `rgba(255,255,255,0.6)`
- Border: `rgba(0,0,0,0.05)`
- Shadow: Medium, crisp



## Animation & Interaction

### Motion Principles

**Purposeful Animation:**
Every animation serves a purpose:
- Feedback (button clicks, form submission)
- Attention (new notifications, errors)
- Continuity (page transitions, reveals)

**Performance Budget:**
- Maximum 60fps on all interactions
- Animations under 500ms for micro-interactions
- Lazy-load heavy effects (blobs, noise)

### Key Animations

**1. Scroll Reveal** (`IntersectionObserver`)
- Fade + slide up
- Threshold: 0.1
- Duration: 0.8s ease-out

**2. Ambient Blobs**
- Gentle float animation
- 6s duration, infinite loop
- Transform-only (GPU accelerated)

**3. Glassmorphism**
- Backdrop-filter: blur() (hardware accelerated)
- CSS-only, no JavaScript



## Responsive Design

### Breakpoints

```
sm:  640px  - Mobile landscape
md:  768px  - Tablet portrait
lg:  1024px - Tablet landscape / small laptop
xl:  1280px - Desktop
2xl: 1536px - Large desktop
```

### Mobile-First Approach

**Base Styles:** Mobile (320px+)
- Single column layouts
- Touch-friendly (44px+ targets)
- Simplified navigation (hamburger future)

**Progressive Enhancement:**
- md: Two-column layouts appear
- lg: Three-column dashboard
- xl: Wider content areas, more spacing

### Touch Targets

**Minimum:** 44x44px (iOS HIG, WCAG)
**Implemented:**
- Buttons: 44px height minimum
- Form inputs: 48px height
- Navigation links: 44px tap area



## Accessibility Compliance

### WCAG 2.1 AA Standards

**Color Contrast**
- Text: Minimum 4.5:1 ratio
- Large text: Minimum 3:1 ratio
- Tested with WebAIM contrast checker

**Keyboard Navigation**
- All interactive elements focusable
- Visual focus indicators (2px outline)
- Logical tab order maintained

**Screen Readers**
- Semantic HTML (`<nav>`, `<main>`, `<article>`)
- ARIA labels where needed
- Alt text for all images

**Form Accessibility**
- Labels associated with inputs
- Error messages announced
- Validation feedback visible

**See:** [ACCESSIBILITY.md](./ACCESSIBILITY.md) for detailed compliance documentation.



## Design System Tools

### Figma Component Library
*Status:* Prototypes created as HTML/CSS (no Figma yet)  
**Future:** Export to Figma for team collaboration

### Design Tokens
All values centralized in:
- `tailwind.config.js` (source of truth)
- CSS custom properties for theming
- Vue composables for dynamic values

**See:** [DESIGN_TOKENS.md](./DESIGN_TOKENS.md)

### Component Library
Reusable Vue components with:
- Props documentation
- Accessibility notes
- Usage examples
- Do's and Don'ts

**See:** [COMPONENT_LIBRARY.md](./COMPONENT_LIBRARY.md)



## Design Metrics & Success

### Qualitative Goals

- **Visually Impressive** - Wow factor on first impression  
- **Easy to Navigate** - Intuitive information architecture  
- **Accessible** - Usable by everyone  
- **Performant** - Fast load times, smooth animations  
- **Maintainable** - Clear design system, reusable components

### Future Metrics (Post-Launch)

- Time to first publish (Dashboard -> Published Article)
- Reader engagement (time on article page)
- Portfolio customization rate (% users who customize)
- Mobile vs desktop usage split
- Accessibility tool usage (keyboard navigation %)



## Evolution Timeline

### V1: The Bootstrap MVP (2025-09)
- Basic CRUD functionality
- Standard Bootstrap 5 styling
- "Ugly but functional"

### V2: The Bento Grid (2025-10)
- Explored dashboard-heavy layouts
- High information density
- **Outcome:** Scrapped (too cluttered)

### V3: The OneNote Clone (2025-10)
- Attempted to mimic Microsoft OneNote pixel-for-pixel
- **Outcome:** Abandoned (legally risky, lacked soul)

### V4: TaraKreasi Branding (2025-11)
- Established the Blue/Orange identity
- Switched to Remixicon
- Clean, corporate aesthetic

### V5: Technical Foundation (2025-11)
- Migrated to Vue 3 Composition API
- Implemented Inertia.js
- Performance focus

### V6: Electric Earth (2025-12) [Current]
- **Peak Design**: Glassmorphism, Noise Textures, Ambient Blobs
- Digital Garden concept fully realized
- Dark mode support



## Next Steps & Roadmap

### Short-Term Improvements
- Add keyboard shortcuts documentation
- Create onboarding flow for new users
- Improve mobile navigation (hamburger menu)
- Add accessibility testing results

### Medium-Term
- Markdown editor mode (alternative to WYSIWYG)
- Theme customization (user chooses accent colors)
- Component playground (for developers)
- Figma component library

### Long-Term
- Design system API documentation
- Contribution guidelines for designers
- A/B testing framework
- Advanced animation library



## Resources

**Documentation:**
- [Design Tokens](./DESIGN_TOKENS.md) - Centralized design variables
- [Component Library](./COMPONENT_LIBRARY.md) - Vue component specs
- [Accessibility](./ACCESSIBILITY.md) - WCAG compliance details
- [UI v6 README](./UI/v6/README.md) - Technical implementation

**External References:**
- Material Design Guidelines
- Apple Human Interface Guidelines
- WCAG 2.1 Specification
- Inclusive Components by Heydon Pickering



**Version:** Beta 0.9.80 (UI Design V6.5)  
**Designer:** TaraKreasi (Solo Fullstack Designer/Developer: Tri Wantoro)  
**Last Updated:** December 27, 2025

> **Portfolio Note:** This UI/UX documentation demonstrates professional design thinking, accessibility awareness, and systematic design approach suitable for senior-level designer roles.
