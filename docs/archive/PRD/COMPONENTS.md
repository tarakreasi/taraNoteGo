# Component Library (UI Design V6)

This document provides a reference for the key Vue components in `resources/js/Components` and `Pages`.

## 1. Core Layouts

### AppLayout (Public)
**Path:** `resources/js/Layouts/AppLayout.vue`
- **Description**: The master layout for public-facing pages (Home, Articles, Article, Portfolio).
- **Features**: 
    - Includes `AppNavbar` and `AppFooter`.
    - Handles the "Noise Overlay" globally.
    - Manages dark mode toggling.

### AuthenticatedLayout (Admin)
**Path:** `resources/js/Layouts/AuthenticatedLayout.vue`
- **Description**: The layout for the Admin Dashboard and Private Workspace.
- **Features**: Authentication guard, Dashboard Sidebar.



## 2. Public Components

### AppNavbar
**Path:** `resources/js/Components/AppNavbar.vue`
- **Description**: Responsive navigation bar for public pages.
- **Props**: `user` (for auth state), `transparent` (for hero overlay).
- **Features**: Mobile menu drawer, User dropdown.

### AppFooter
**Path:** `resources/js/Components/AppFooter.vue`
- **Description**: Global footer with social links and copyright.
- **Props**: `settings` (Object containing `footer_*` keys).
- **Usage**: Automatically injected with settings from `HandleInertiaRequests`.

### ArticleCard
**Path:** `resources/js/Components/ArticleCard.vue` (Implied/Integrated in `Articles.vue`)
- **Description**: Card component for displaying article preview.
- **Features**: Cover image with hover zoom, tags pill, read time calculation.



## 3. Dashboard Components

### SettingsPanel
**Path:** `resources/js/Pages/Dashboard/SettingsPanel.vue`
- **Description**: The central configuration hub for the application.
- **Tabs**:
    - **Profile**: Avatar upload, Name, Bio.
    - **Home (Welcome)**: Landing page hero & sections.
    - **Articles**: Articles index settings.
    - **Footer**: Global footer settings.
- **Methods**: `handleAvatarUpload`, `save`, `saveProfile`.

### Editor
**Path:** `resources/js/Components/Editor.vue`
- **Description**: Wrapper around Quill.js for Rich Text Editing.
- **Props**: `modelValue` (String content).
- **Events**: `update:modelValue`, `save` (on Ctrl+S).



## 4. UI Primitives

### ApplicationLogo
**Path:** `resources/js/Components/ApplicationLogo.vue`
- **Description**: The TaraNote SVG brand mark.

### Modal
**Path:** `resources/js/Components/Modal.vue`
- **Description**: Reusable accessible modal dialog.
- **Props**: `show` (Boolean), `maxWidth` (String).

### PrimaryButton
**Path:** `resources/js/Components/PrimaryButton.vue`
- **Description**: Standard action button with primary branding colors.
