# Changelog

A detailed history of TaraNote's evolution, tracking both technical milestones and the major UI Design System iterations.

## [v1.0.0] - Dec 27, 2025
**Production Release**

Final preparations for public GitHub release, focusing on documentation consistency, cleanup, and tone refinement.

### Final Polish (The "Zen" Update)
- **Layout Standardization**: Implemented a consistent left-aligned layout across Dashboard, Docs, and Home views, with a responsive right margin to balance the sidebar.
- **Visual Refinements**: Enhanced the Floating Dock with a "Thicker Glass" aesthetic (distinct styles for Light/Dark modes).
- **Content**: Seeded deep-dive articles on "Technical Architecture" (Laravel 12, Vue 3) and "Design Philosophy" (Jobs To Be Done).
- **Identity**: Updated administrative identity and resolved database seeding conflicts.

### Setup Script Overhaul (`run.sh`)
- **Enhanced Error Handling**: Complete rewrite with production-grade error detection and recovery.
- **Automatic NPM Conflict Resolution**: 
    - Script now automatically detects `ERESOLVE` peer dependency errors (Vite 7 vs @vitejs/plugin-vue compatibility).
    - Automatic retry with `--legacy-peer-deps` flag - zero manual intervention required.
- **Verbose Output**: 
    - Added colorful progress indicators.
    - Real-time color-coded feedback (Red=Error, Yellow=Warning, Green=Success, Cyan=Info).
    - Clear error boxes with actionable solutions.
- **Better Visibility**:
    - Displays detected versions (PHP, Composer, Node.js, NPM).
    - Shows database type being configured (MySQL/SQLite).
    - Professional success banner with credentials and server URLs.
- **Database Seeder Cleanup**: Removed all emoji/icon prefixes from notebook names for professional appearance.

### Documentation & Tone
- **Narrative Refinement**: Updated `README.md` and `PROJECT_SUMMARY.md` to adopt a "Humanist, Humble, Professional" tone.
- **Version Sync**: Synchronized all documentation to **Beta 0.9.80**.
- **Archives**: created `/docs/archive` to house legacy logic (`TriWantoro`, `Articles`, `Contact`) while keeping core architecture clean.

### Architecture & Backend (Refactor)
- **Controller Split**: Separated `AdminNoteController` and `AdminNotebookController` for better separation of concerns.
- **Type Safety**: Enforced `declare(strict_types=1)` and explicit return types (`JsonResponse`) across API controllers.
- **Validation**: Migrated inline validation to dedicated `FormRequest` classes (`StoreNoteRequest`, `UpdateNoteRequest`).
- **Route Cleanup**: Clarified public entry points. Root (`/`) now redirects cleanly to `/taraNote`.
- **Security Scan**: Verified codebase is free of hardcoded secrets before release.
- **Redirects**: Simplified post-login flow to land directly on the `taraNote` writer interface.

## [Beta 0.9.80] - Dec 24, 2025
**The "TriWantoro" Portfolio Refactor**

Focused on enhancing the personal branding page with dynamic editing and a refined layout.

### UI & Features
- **Vertical Navigation**: Implemented a vertical floating dock for better screen real estate.
- **Inline Editing**: 
    - Full WYSIWYG editing for all text, images, and links on `TriWantoro.vue`.
    - Seamless admin experience with inline status indicators (Sync/Check) instead of popups.
    - **Header**: Moved Branding to Hero section for better hierarchy.
- **GitHub Card**: Added background image support via `EditableImage`.

### Project Landing Pages
- **New Dedicated Pages**: Created `/taraNote/taraNote`, `/taraTask`, and `/taraQueue` landing pages.
- **Design Consistency**:
    - **Centered Layout**: Refactored to match Home Page alignment (`mx-auto` + `pl-24`).
    - **Dark Mode**: Full "Electric Earth" dark theme support.
    - **Floating Dock**: Consistent vertical navigation.
- **UX**: Added "Back to Portfolio" navigation and tech stack footers.

### Cleanup
- **Archiving**: Moved legacy pages (`Home`, `About`, `Portfolio`) to `docs/archive/legacy_code`.
- **Tests**: Removed obsolete `PortfolioTest` (replaced by single-page logic).
- **Docs**: Created `ARCHIVE_LOG.md`.

**The "Japanese Engineer" UI Refinement**

A major aesthetic overhaul focusing on minimalism, precision, and visual calm ("Zen Mode").

### UI Design System V6.5 (The Polish)
- **Concept**: A shift from "Electric Earth" (noise/blobs) to a simpler, distraction-free interface inspired by Japanese utility software.
- **Visuals**:
    - **Palette**: Strict Off-White (`#FAFAFA`) and Navy (`#0F172A`). Removed specific gradient blobs.
    - **Grid**: Enforced strict **8px grid** spacing.
    - **Reduction**: Removed noise overlays and glassmorphism borders in favor of clean negative space.
    - **Floating Dock Clearance**: Implemented a **"Deep Bottom Padding"** strategy (`pb-40` / 160px) on all main scroll containers. This ensures that the bottom-most content (e.g., the last paragraph of an article) can always be scrolled *above* the Floating Dock, preventing visual obstruction.
- **Affected Areas**:
    - **Dashboard**: Simplified Sidebar (no borders), Ghost Search inputs, Technical Typography (`font-mono` for metadata).
    - **TaraNote / Docs**: Aligned with the dashboard's new minimal look.

### Documentation
- **Updated**: All major documentation to reflect the V6.5 state.
- **Indexed**: Ensured `/docs` is perfectly synced with the file system.

## [Beta 0.9.60] - Dec 21, 2025
**The "Electric Earth" Polish & Legal Hardening**

This sprint focused on refining the V6 "Electric Earth" design system for mobile devices, enhancing the "Spaces" (Notebooks) experience, and solidifying the project's legal/security standing.

### UI Design System V6 (Refinements)
- **Mobile Responsiveness Overhaul**:
    - Complete mobile refactor for Public Views (`Home`, `About`, `Portfolio`, `Article`).
    - Implemented a new mobile-specific Navigation Menu with scroll-locking logic.
    - Optimized sidebar behavior for small screens (collapsible/overlay).
    - Adjusted padding and typography scaling for readability on phones.
- **Spaces (Notebooks) UI V6**:
    - **Neon Branding**: Active/Editing states now glow with a neon orange border (`border-orange-500`) and shadow.
    - **Glassmorphism Sidebar**: Unified glass effect across the sidebar and headers (`bg-white/5` + blur).
    - **Inline Management**: 
        - Replaced browser prompts with **Inline Input Fields** for creating/renaming spaces.
        - Added interactive stylistic feedback (neon borders) during editing.
    - **Unified Delete UX**: 
        - Replaced native `confirm()` with a custom **Inline Permission UI**.
        - Added safety checks: Only empty spaces (0 notes) can be deleted.
        - "DELETE" button acts as a confirmed destructive action.

### Legal & Security
- **Comprehensive Documentation Audit**:
    - Deep scan of all `/docs` for patent/license infringement risks (Clean).
    - Verified version consistency ("Beta 0.9.60") across 20+ files.
- **Security Policy**:
    - Established **Best Effort** response policy in `SECURITY.md`.
    - Clarified disclaimer: "Aspirational targets, not guaranteed SLAs".
- **Legal Compliance**:
    - Created `LEGAL_COMPLIANCE.md` declaring patent safety and license compatibility.
    - Added `ATTRIBUTIONS.md` for proper credit of open-source assets.

### Bug Fixes & Tech Debt
- **Tech Debt**: Removed unused legacy CSS from the V3 era.
- **Fix**: Resolved syntax errors in `TaraNote.vue` caused by refactoring.
- **Fix**: Corrected column spacing issues in `Docs.vue` grid layout.
- **Fix**: Standardized "Spaces" terminology across Dashboard/Docs/TaraNote views.

## [Beta 0.9.0] - Dec 10, 2025
**The "Digital Garden" Launch (UI V6 Core)**

The major leap to the "Electric Earth" design language.

### UI Design System V6 (Launch)
- **Concept**: A "Living Digital Garden" merging technology (Electric) with nature (Earth).
- **Core Visuals**:
    - **Glassmorphism**: Heavy use of `backdrop-blur-xl` and translucent layers.
    - **Noise Texture**: Custom SVG noise overlay for tactile depth.
    - **Ambient Blobs**: Animated CSS gradients that breathe in the background.
    - **Dark Mode**: First-class dark mode support with semantic color mapping.

### New Features
- **Public Portfolio**: Dedicated `/portfolio/{username}` pages for authors.
- **Dual-View Architecture**:
    - **Private**: OneNote-style productivity view.
    - **Public**: Medium-style reading view.
- **Settings Engine**: Dynamic key-value store for site customization (Title, Tagline, Avatar).

## [Beta 0.8.0] - Nov 25, 2025
**The "Technical Foundation" (UI V5)**

Transitioning from prototype to robust application structure.

### UI Design System V5 (Refinement)
- **Focus**: Performance and Structure.
- **Changes**:
    - Cleaned up the V4 "TaraKreasi" layout.
    - Migrated fully to **Vue 3 Composition API**.
    - Introduced **Inertia.js** to replace the chaos of V4's manual API calls.

### Technical Milestones
- **Rich Text Editor**: Integrated `Quill.js` with custom toolbar and sanitization.
- **Auto-Save**: Implemented `lodash.debounce` (1.5s) for worry-free writing.
- **Status Workflow**: Added `Draft` vs `Published` states to Notes.

## [Alpha 0.4.0] - Nov 10, 2025
**The "TaraKreasi" Rebrand (UI V4)**

Finding our own identity after the "Clone" phase.

### UI Design System V4 (TaraKreasi)
- **Concept**: "Clean & Corporate".
- **Colors**: Introduced the signature **Electric Blue** (#4485EE) and **Sunset Orange** (#FF5F1F).
- **Layout**: Moved from Top-Navbar to Sidebar Navigation.
- **Icons**: Switched to **Material Symbols** (Apache 2.0) to avoid copyright issues.

## [Alpha 0.3.0] - Oct 20, 2025
**The "OneNote" Experiment (UI V3)**

A lesson in what *not* to do.

### UI Design System V3 (The Clone)
- **Concept**: "Pixel-Perfect OneNote Clone".
- **Visuals**: Copied Microsoft's vivid purple (#7719AA) and tab layout.
- **Outcome**: **Abandoned**. 
    - *Reason*: Felt lifeless and legally risky. Lacked a unique soul.
    - *Lesson*: Inspiration is good, imitation is bad.

## [Alpha 0.2.0] - Oct 05, 2025
**The "Bento" Grid (UI V2)**

Exploring layout possibilities.

### UI Design System V2
- **Concept**: Dashboard-heavy "Bento Grid" layout.
- **Visuals**: Dense information density, widget-based UI.
- **Outcome**: **Scrapped**. Too cluttered for a writing app.

## [Alpha 0.1.0] - Sept 23, 2025
**The "Bootstrap" MVP (UI V1)**

Where it all started.

### UI Design System V1
- **Tech**: Laravel Blade + Bootstrap 5.
- **Features**: Basic CRUD (Create, Read, Update, Delete).
- **Visuals**: Default Bootstrap styling. Ugly but functional.
- **Goal**: Prove the backend architecture works.

## Versioning Strategy
- **Beta 0.9.x**: Feature-complete, polishing for V1.0.
- **UI Vx**: Major Design System iterations (V1=Bootstrap, V3=OneNote, V4=TaraKreasi, V6=Electric Earth).
