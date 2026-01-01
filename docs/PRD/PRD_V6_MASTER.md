# Product Requirements Document - V6.0 (Master)

**Application Version:** Beta 0.9.80 | **UI Design:** V6.5 "Zen Mode" | **Type:** Knowledge Management System + Portfolio | **Stack:** Laravel 12 + Vue 3 + MySQL

> **Project Type**: Portfolio project demonstrating full-stack skills. Also used daily as a personal productivity tool.

## 1. Project Overview
**TaraNote** is a "Living Digital Garden" that serves two distinct purposes:
1.  **Private Workspace (TaraNote):** A OneNote-style interface for organizing raw notes, ideas, and drafts into Notebooks.
2.  **Public Portfolio (Digital Garden):** A polished blog engine that showcases selected "Published" articles to the world.

## 2. User Persona
*   **Primary User (The Architect):** A developer or technical writer who needs a frictionless way to document their learning journey and share it as a professional portfolio.
*   **Secondary User (The Visitor):** Recruiters, peers, or learners visiting the site to gauge the author's expertise.

## 3. Core Features (UI Design V6 Scope)

### A. The Dual-View Content Engine
1.  **TaraNote View (Private):**
    *   **Structure:** Sidebar with Notebooks -> Sections -> Pages.
    *   **Function:** Rapid note-taking, drafting, organization.
    *   **Status:** Drafts are visible only here.
2.  **Digital Garden View (Public):**
    *   **Structure:** Hero Landing Page -> Article Grid -> Article Detail.
    *   **Function:** SEO-optimized portfolio display.
    *   **Status:** Only `PUBLISHED` notes appear here.

### B. Admin & Customization
1.  **Settings Panel:**
    *   **Profile:** Edit Name, Title, Bio, and **Avatar** (Upload/URL).
    *   **Home/Welcome:** Customize Hero text, Feature Stats, and Availability status.
    *   **Articles:** Configure search placeholders and "Digital Notes" section text.
    *   **Footer:** Manage social links and copyright text.
2.  **Notebook Management:**
    *   Create, Rename, Delete Notebooks.
    *   Assign "slugs" for URL routing (e.g., `/articles`, `/portfolio`).

### C. Content Management
1.  **Rich Text Editor:**
    -   **Tiptap** integration for headings, lists, code blocks, and links.
    -   **Auto-Save:** Saves changes automatically after 1.5s of inactivity.
2.  **Status Workflow:**
    -   Toggle between `DRAFT` and `PUBLISHED`.
    -   Set "Is Featured" to pin articles to the top.
3.  **Media:**
    -   Upload Cover Images for articles.
    -   Upload Avatars for user profiles.

### D. Design System V6.5 ("Zen Glass")
1.  **Visual Language:**
    -   **Minimalism:** Removed noise and blobs for a distraction-free environment.
    -   **Structure:** Strict 8px grid and no hard borders (background separation).
    -   **Typography:** Inter + JetBrains Mono for technical precision.
    -   **Colors:** Off-white (`#FAFAFA`) and Navy (`#0F172A`).
    -   **Icons:** Remixicon integration.

## 4. Non-Functional Requirements
*   **Performance:**
    -   Public pages load in < 1s (Target).
    -   Auto-save latency < 2s.
*   **Security:**
    *   **Authentication:** Laravel Breeze (Session-based).
    *   **Authorization:** Policy-based access control (Users can only edit their own notes).
    *   **Sanitization:** All WYSIWYG input is sanitized to prevent XSS.
*   **SEO:**
    *   Dynamic Meta Tags (Title, Description).
    *   Clean URLs (`/articles/refactoring-ui-logic`).

## 5. Exclusions (Post-V6 Roadmap)
*   Multi-user collaboration (Team implementations).
*   Real-time co-editing (Google Docs style).
*   Native Mobile App (PWA only).