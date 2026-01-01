# TaraNote: Architectural Overview

**Version**: Beta 0.9.80  
**UI**: V6.5 "Zen Mode"  
**Status**: Functional, tested, and in active use.

> A portfolio project that serves as a personal note-taking tool. Built to demonstrate full-stack engineering capabilities and solve a personal need for organization.



## What Is This?

A unified platform combining private note-taking with public publishing.

**Private Role**: Organize thoughts and drafts in notebooks (similar to OneNote).  
**Public Role**: Publish selected notes as blog posts (similar to Medium).

The workflow is simple: transition a note from private to public with a single status change.



## Project Goals

**Primary Goal**: To create a comprehensive portfolio piece that demonstrates the ability to architect and build production-ready software, moving beyond basic CRUD applications.

**Secondary Goal**: To solve a personal need for managing knowledge. I wanted a platform that allowed me to:
- Organize unstructured notes privately.
- Share polished insights publicly.
- Retain data ownership through self-hosting.

Building a tool for daily personal use has provided valuable feedback loops for improving the user experience and code quality.



## Technical Overview

### Statistics
- **73 automated tests** (Feature + Unit) ensuring stability.
- **6 models**, **19 controllers**, **25 Vue pages**.
- **15,000+ lines of code** (Full Stack).
- **Clean Architecture**: Focused on maintainability and separation of concerns.

### Key Features
- **Authentication**: Secure login, registration, and password recovery.
- **Rich Text Editor**: Quill.js integration with 1.5s auto-save.
- **Organization**: Notebook-based structure for content management.
- **Publishing Workflow**: Seamless transition from Draft to Published.
- **Public Portfolio**: dedicated author pages and article listings.
- **Customization**: User-defined avatars, slugs, and site meta-settings.
- **Newsletter**: Architecture ready for email service integration.
- **Theme Support**: First-class dark/light mode support.

### Technology Stack
**Backend**: Laravel 12, MySQL, RESTful API architecture.  
**Frontend**: Vue 3, Inertia.js, Tailwind CSS.  
**Build System**: Vite.  
**Deployment**: VPS-ready (Nginx + PHP-FPM).



## Documentation Structure

```
docs/
├── README.md              # Documentation hub
├── QUICKSTART.md          # Setup guide
├── USER_GUIDE.md          # User manual
├── DEV_HANDBOOK.md        # Developer setup & standards
├── API_REFERENCE.md       # API endpoints & behavior
├── DATABASE_SCHEMA.md     # Data structure
├── TESTING.md             # Testing guide
├── DEPLOYMENT.md          # Deployment procedures
├── CHANGELOG.md           # Version history
├── ROADMAP.md             # Development timeline
├── PROJECT_SHOWCASE.md    # Portfolio presentation
├── LEGAL_COMPLIANCE.md    # Licensing & compliance audit
├── SECURITY.md            # Security policy
├── PRIVACY_POLICY.md      # Data handling policy
├── CONTRIBUTING.md        # Contribution guidelines
├── PRD/                   # Product Requirements Documents
│   ├── PRD_V1_PRODUCTION.md  # Future roadmap
│   ├── PRD_V6_MASTER.md      # Current specifications
│   └── archive/              # Legacy requirements
└── UI/                    # Design documentation
    ├── v6/                # Current Design System
    ├── v4/, v3/, v1/      # Legacy iterations
    └── research/          # UX research notes
```



## Design Evolution

**V1** (Bootstrap): Functional proof-of-concept focused on backend logic.
**V2** (Bento Grid): Explored dashboard-dense layouts; simplified for better focus.
**V3** (OneNote-inspired): Experimented with familiar tab layouts; iterated for uniqueness.
**V4** (TaraKreasi): Established a distinct branding identity.
**V5** (Technical Foundation): Refactoring to modern Vue 3 Composition API & Inertia.js.
**V6** (Zen Mode): Current iteration. Focuses on clarity, negative space, and a refined "Zen" aesthetic using glassmorphism and subtle textures.

Each iteration has been a stepping stone in understanding product design and user experience.



## Development Journey

**Timeline**: Sept - Dec 2025  
**Team**: Solo Developer  
**Methodology**: Iterative development sprints.

### Phase 1: Backend Foundation
- Established Laravel architecture.
- Implemented Authentication (Breeze).
- Built core CRUD for Notes.
- Integrated Rich Text Editor.

### Phase 2: Frontend & Public Interface
- Developed public portfolio views.
- Built the Settings configuration engine.
- Implemented Newsletter logic.
- Refined UI based on V4 design system.

### Phase 3: Polish & Reliability
- Refined UI to V6.5 standards.
- Achieved 73 passing tests.
- Comprehensive documentation.
- Legal and security compliance checks.

See [ROADMAP.md](ROADMAP.md) for a detailed breakdown.



## Quality Assurance

**64 Tests Passing** (PHPUnit)

**Testing Scope**:
- **Authentication**: Login, Register, Password Reset flows.
- **Core Logic**: Note creation, updates, deletion, and publishing.
- **Authorization**: Ensuring users can only access their own data.
- **Public Access**: Verifying validity of public routes.
- **Admin Features**: Settings and Notebook management.
- **Integrity**: Checking documentation links and static routes.

Run tests via: `php artisan test`



## Security Measures

**Implemented Protections**:
- **Passwords**: Bcrypt hashing (work factor 12).
- **Sessions**: Secure, encrypted Laravel sessions.
- **Forms**: CSRF protection on all mutating requests.
- **Content**: DOMPurify sanitization for rich text.
- **Database**: Eloquent ORM protections against SQL injection.

**Note**: As a personal portfolio project, this has not undergone an external penetration test. Please refer to [SECURITY.md](SECURITY.md) for full details.



## Legal & Compliance

**License**: MIT License (Open Source).  
**Trademarks**: "TaraNote" and "TaraKreasi" are trademarks of the project owner.  
**Attribution**: Material Symbols (Apache 2.0), Google Fonts (OFL).

**Compliance**: Not a commercial entity. Use involves acceptance of terms in `LICENSE.md`.

All third-party dependencies are compliant with the project's licensing. See [LEGAL_COMPLIANCE.md](LEGAL_COMPLIANCE.md).



## Known Limitations

**Future Improvements (V1.0)**:
- **Analytics**: Integration of privacy-focused analytics.
- **Email Delivery**: Connection to transactional email services.
- **Search**: Enhanced search capabilities (e.g., Laravel Scout).
- **Markdown**: Support for native markdown editing.
- **Interaction**: Comments and social features.
- **Mobile**: Further optimizations for small screens.

These features are tracked in the [V1.0 Roadmap](PRD/PRD_V1_PRODUCTION.md).



## Reflections

**Technical Growth**:
- Deepened understanding of Inertia.js for modern monoliths.
- Mastered Vue 3 Composition API patterns.
- Utilized Laravel Policies for robust authorization.
- Implemented a consistent design system from scratch.

**Project Insights**:
- Comprehensive documentation pays dividends in maintenance.
- Automated testing provides the confidence to refactor aggressively.
- Building for personal use aligns incentives for quality.
- Disciplined scoping is essential for delivery.



## Next Steps

See [PRD_V1_PRODUCTION.md](PRD/PRD_V1_PRODUCTION.md) for the forward-looking plan.

**Focus Areas**:
1. Analytics (Usage insights)
2. Email Infrastructure (Communication)
3. Performance Tuning (Caching, Optimization)
4. Advanced Search



## Deployment Configuration

Currently self-hosted on:
- **OS**: Ubuntu 22.04 LTS
- **Server**: Nginx + PHP 8.3-FPM
- **Database**: MySQL 8.0
- **Security**: Let's Encrypt SSL

Deployment Guide: [DEPLOYMENT.md](DEPLOYMENT.md)



## Contact

**Tri Wantoro** - Full-Stack Developer

Email: ajarsinau@gmail.com  
Website: tarakreasi.com  
Location: Indonesia



## Acknowledgments

This project stands on the shoulders of the open-source community:
- **Laravel** team for an exceptional framework.
- **Vue.js** community for a reactive frontend experience.
- **Tailwind Labs** for the utility-first CSS methodology.
- **Inertia.js** for bridging the gap between backend and frontend.



**Completed**: December 27, 2025  
**Developer**: Tri Wantoro (TaraKreasi)  
**Status**: Beta 0.9.80 - Functional, Tested, Documented

