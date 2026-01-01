# TaraNote Documentation

Welcome to TaraNote's official documentation!

> **Project Nature**: TaraNote is a **portfolio project** built to showcase full-stack development skills (Laravel 12, Vue 3, Inertia.js, UI Design V6). It's also my **daily productivity tool** for note-taking and blogging. While functional and production-ready, it's primarily a learning/demonstration project rather than a commercial product.

> **Open Source Disclaimer**: This project is provided "as-is" under the MIT License with **NO WARRANTIES OR GUARANTEES**. Maintenance is done on a **best-effort basis** as a personal hobby project. There is **no SLA, no guaranteed support**, and **no liability** for data loss, bugs, or security issues. **Use at your own risk.**



## Main Documentation

### API Documentation
- **[API_NOTES.md](./API_NOTES.md)** - Complete Notes API documentation (v4/v5 legacy, v6 pending)
  - All endpoints (GET, POST, PUT, DELETE)
  - Authentication & authorization
  - Data models & relationships
  - Error handling & best practices
  - Code examples & testing guide

## Core Concepts

TaraNote manages content as **Notes**, which can be presented in two distinct ways depending on the context:

### 1. TaraNote Workspace (`TaraNote.vue`)
- **Concept:** Structured, hierarchical organization (OneNote-style).
- **View:** `TaraNote.vue` - A split-pane interface with Notebooks on the left and content on the right.
- **Purpose:** Knowledge management, structured learning, and deep diving into topics.
- **Content:** Displayed as "Notes" in a notebook hierarchy.

### 2. Digital Garden / Public Portfolio (`Article.vue` / `Portfolio.vue`)
- **Concept:** A curated public showcase or "Digital Garden".
- **View:** `Article.vue` (Single Article), `Portfolio.vue` (User Garden), `Home.vue` (Garden Entrance).
- **Purpose:** Public showcasing, blogging, and linear reading of finished thoughts.
- **Content:** Displayed as "Articles".

*Note: Both views consume the same underlying `Note` data model from the database.*



### Custom Design System (v6)
- **[UI_UX_OVERVIEW.md](./UI_UX_OVERVIEW.md)** - Professional design approach & vision (NEW)
- **[DESIGN_TOKENS.md](./DESIGN_TOKENS.md)** - Centralized design variables (NEW)
- **[COMPONENT_LIBRARY.md](./COMPONENT_LIBRARY.md)** - V6 Component specifications (NEW)
- **[ACCESSIBILITY.md](./ACCESSIBILITY.md)** - WCAG 2.1 AA Compliance (NEW)
- **[Archive (Legacy Docs)](./archive/README.md)**: Old V2/V3 specs and scrapped ideas.
- **[DESIGN_SYSTEM.md](./DESIGN_SYSTEM.md)** - Legacy v4/v5 design system goals
  - Design philosophy & principles
  - Color palette & typography
  - Component specifications
  - Layout system
  - Material Icons usage
  - Implementation details

### Database
- **[DATABASE_SCHEMA.md](./DATABASE_SCHEMA.md)** - Database structure
  - Tables: users, notebooks, notes
  - Relationships & constraints
  - Business logic

### Deployment
- **[DEPLOYMENT.md](./DEPLOYMENT.md)** - Production deployment guide
  - Server requirements
  - Installation steps
  - Configuration
  - Web server setup

### Testing
- **[TESTING.md](./TESTING.md)** - Testing guide
  - Running tests
  - Test structure
  - Feature & unit tests

### Quick Start
- **[QUICKSTART.md](./QUICKSTART.md)** - Get started quickly
  - Installation
  - Development setup
  - Common tasks
  - Troubleshooting

### Version History
- **[CHANGELOG.md](./CHANGELOG.md)** - All version changes
  - v6.0: Living Digital Garden (Glassmorphism)
  - v5.2: Static Vue Export
  - v4.0: TaraKreasi redesign
  - v3.0: OneNote-style (deprecated)
  - v2.0: Bento Grid
  - Migration guides

### Documentation Index
- **[INDEX.md](./INDEX.md)** - Complete documentation index
  - Quick navigation
  - Links to all resources
  - Getting started guides

### Product Requirements (Historical)
- **[PRD/](./PRD/)** - Initial requirements & specifications
  - Product requirements document
  - Initial API specifications
  - Component specifications
  - *Note: For historical reference only*



## Directory Structure

```
docs/
 ├── README.md              # This file (The Hub)
 ├── API_NOTES.md           # Notes API documentation
 ├── DESIGN_SYSTEM.md       # v4 Design system
 ├── DATABASE_SCHEMA.md     # Database structure
 ├── DEPLOYMENT.md          # Deployment guide
 ├── TESTING.md             # Testing guide
 ├── QUICKSTART.md          # Quick start guide
 ├── CHANGELOG.md           # Version history
 ├── INDEX.md               # Documentation index
 ├── PRD/                   # Product Requirements (Initial specs)
 │   ├── README.md
 │   ├── PRD_V6_MASTER.md   # Current Master PRD
 │   ├── archive/           # History (V1, V3, V4)
 │   └── USER_PERSONA.md
 ├── UI/                    # UI Documentation & Prototypes
 │   ├── README.md          # UI Master Index
 │   ├── v6/                # CURRENT: Living Digital Garden
 │   ├── v5.2/              # Static Vue Export
 │   ├── v5/                # Refined UI
 │   ├── v4/                # Clean Blue/Orange (prev 'prototypes')
 │   ├── v3/                # OneNote-style (Deprecated)
 │   ├── v1/                # Legacy Designs
 │   ├── workshop/          # Component Demos & Workshop
 │   └── research/          # UX Research
 └── archive/               # Archived/unused files
     ├── colab/             # Colab Module (archived)
     ├── WYSIWYG/           # WYSIWYG docs (archived)
     ├── investor/          # Investor materials (archived)
     ├── SC/                # Screenshots
     └── *.md               # Old specs
 ```



## Quick Links

### For Developers
- [API Documentation](./API_NOTES.md) - Complete API reference
- [Database Schema](./DATABASE_SCHEMA.md) - Database structure
- [Testing Guide](./TESTING.md) - How to run tests
- [Quick Start](./QUICKSTART.md) - Get up and running
- [Design System](./DESIGN_SYSTEM.md) - UI components & styles

### For DevOps
- [Deployment Guide](./DEPLOYMENT.md) - Production deployment
- [Database Schema](./DATABASE_SCHEMA.md) - Database setup

### For Designers
- [Color Palette](./DESIGN_SYSTEM.md#color-palette)
- [Typography](./DESIGN_SYSTEM.md#typography)
- [Components](./DESIGN_SYSTEM.md#components)
- [Icons](./DESIGN_SYSTEM.md#icons)

### For Product Managers
- [Changelog](./CHANGELOG.md) - What's new
- [Portfolio Showcase](./PROJECT_SHOWCASE.md) - Skills demonstration
- [Design Philosophy](./UI_UX_OVERVIEW.md) - UX strategy
- [Documentation Index](./INDEX.md)



## Current Version

**TaraNote Beta 0.9.70** - "Zen Garden" (V6.5)
 
### Key Features (v6.5)
- **Japanese Minimalism** - Zen aesthetic
- **Strict 8px Grid** - Precision layout
- **Mono Typography** - Technical clarity
- **Adaptive Theming** - Light/Dark support

### Key Features (v4/v5)
- **TaraKreasi Branding** - Blue & Orange colors
- **Open Source Icons** - Material Icons
- **Complete API** - RESTful with auth
- **Full Tests** - 62 tests passing



## Contributing

When updating documentation:

1. **API Changes** -> Update [API_NOTES.md](./API_NOTES.md)
2. **UI Changes** -> Update [DESIGN_SYSTEM.md](./DESIGN_SYSTEM.md)
3. **New Version** -> Update [CHANGELOG.md](./CHANGELOG.md)
4. **Breaking Changes** -> Add migration guide



## External Resources

- **Material Icons:** https://fonts.google.com/icons
- **Inter Font:** https://rsms.me/inter/
- **Quill.js:** https://quilljs.com/
- **Laravel:** https://laravel.com/docs
- **Vue 3:** https://vuejs.org/
- **Inertia.js:** https://inertiajs.com/



## Archived Files

Old prototypes, demos, and deprecated documentation are in [archive/](./archive/)



**Version:** Beta 0.9.70 (UI Design V6.5)  
**Last Updated:** December 24, 2025  
**Maintained by:** TaraKreasi Team
