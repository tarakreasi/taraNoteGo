# TaraNote - Development Roadmap (Retrospective)

> **NOTE: Portfolio Demonstration**: This roadmap is a **simulation/demonstration of project management skills**. It represents how a solo developer (Tri Wantoro) organizes work using agile-inspired methodology. **This is NOT actual professional sprint planning with a team** - it's a retrospective documentation created to showcase PM knowledge for portfolio purposes.

**Project Duration:** 3 Months (12 Weeks)  
**Developer:** Solo Fullstack Developer (Tri Wantoro / TaraKreasi)  
**Start Date:** September 23, 2025  
**End Date:** December 27, 2025 (UI Design V6.5 Release)  
**Status:** COMPLETED



## Execution Timeline

| Month | Phase | Focus Areas | Status |
| :--- | :--- | :--- | :--- |
| **Month 1 (Oct)** | **Core Architecture** | Auth, Database, Basic CRUD, Editor | Done |
| **Month 2 (Nov)** | **Frontend & Experience** | UI/UX Design, Public Views, Digital Garden | Done |
| **Month 3 (Dec)** | **System & Polish** | Admin Settings, Profiles, Testing, V6 Release | Done |



## Sprint Execution Log (2-Week Iterations)

### Month 1: Foundation (The Backend Sprint)

#### **Sprint 1: Inception & Auth (Sep 23 - Oct 6)**
**Goal:** Establish project structure and secure authentication.
- [x] **System Setup:** Initialize Laravel 12 + Vue 3 + Inertia stack.
- [x] **Database Schema:** Design `users`, `notes`, `notebooks` tables.
- [x] **Authentication:** Implement Login, Register, and Password Reset flows.
- [x] **Models:** Create Eloquent relationships (User has many Notes).

#### **Sprint 2: Core Engine (Oct 7 - Oct 20)**
**Goal:** Enable basic note-taking capabilities.
- [x] **Note CRUD:** Create, Read, Update, Delete functionality for notes.
- [x] **Rich Text Editor:** integrate Quill.js wrapper for Vue.
- [x] **Notebooks:** Implement organization logic (folders/categories).
- [x] **Auto-Save:** Build debounced auto-save mechanism (1.5s latency).



### Month 2: The Visual Layer (The Design Sprint)

#### **Sprint 3: Digital Garden (Oct 21 - Nov 3)**
**Goal:** Bridge the gap between private notes and public articles.
- [x] **Dual-View Architecture:** Separate `TaraNote.vue` (Private) vs `Home.vue` (Public).
- [x] **Public Routes:** Implement `/articles` and `/portfolio/{slug}`.
- [x] **Slug System:** SEO-friendly URLs for articles and profiles.
- [x] **Status Workflow:** Draft vs Published state management.

#### **Sprint 4: Design System V6 (Nov 4 - Nov 17)**
**Goal:** Establish the "Electric Earth" visual identity.
- [x] **UI Framework:** Implement TailwindCSS with custom config.
- [x] **Dark Mode:** Full dark/light theme support with system preference detection.
- [x] **Visuals:** Add Noise Textures, Glassmorphism panels, and hero gradients.
- [x] **Responsive:** Mobile-first refactoring for all public pages.



### Month 3: Polish & Release (The Launch Sprint)

#### **Sprint 5: Personalization (Nov 18 - Dec 1)**
**Goal:** Allow users to customize their corner of the internet.
- [x] **Settings Panel:** Build centralized dashboard for site configuration.
- [x] **Profile Management:** Update Name, Bio, Job Title, and Links.
- [x] **Social Integration:** Add Twitter/GitHub links to footer settings.
- [x] **Content Filters:** Implement "Evergreen", "Budding", "Seedling" taxonomy.

#### **Sprint 6: Final Polish & V6 Launch (Dec 2 - Dec 21)**
**Goal:** Production readiness and final refinements.
- [x] **Avatar Upload:** Implement file upload for user avatars.
- [x] **Landing Page:** Refine `Welcome.vue` with dynamic settings.
- [x] **Testing:** Write and pass 73 Feature/Unit tests (Coverage > 90%).
- [x] **Documentation:** Complete PRD, User Personas, and API Docs.
- [x] **UI Design V6 Release:** Final code freeze and deployment preparation.
- [x] **Editor Migration:** Replaced Quill.js with Tiptap for better stability (V6.5).



## Deliverables Checklist

- [x] **Source Code:** Laravel/Vue Monorepo
- [x] **Documentation:** `/docs` (9 Files)
- [x] **Test Suite:** PHPUnit (66 Tests)
- [x] **Design Assets:** UI Prototypes (HTML)

**Conclusion:** This roadmap demonstrates the ability to plan, execute, and retrospectively document a full-stack web application project using agile-inspired methodology. The structured approach showcases project management skills applicable to professional software development.
