# TaraNote - Portfolio Project Showcase
**Beta 0.9.80 | UI Design V6.5 "Zen Mode"**



## Overview

# TaraNote
## Personal Digital Garden & Note-Taking Platform

**What it is**: A portfolio project showcasing full-stack development skills  
**Dual Purpose**: 
- **Portfolio Showcase**: Demonstrates technical capabilities for potential employers/clients
- **Daily Productivity Tool**: Actually used every day for real note-taking and blogging

**Status**: Beta 0.9.80 (UI: V6.5 "Zen Mode")

> **Portfolio Project**: This is a personal portfolio project demonstrating full-stack development skills. Built by a solo fullstack developer (Tri Wantoro under TaraKreasi brand), not a commercial product with team support.

**Built by**: TaraKreasi (Solo Fullstack Developer: Tri Wantoro)  
**Contact**: ajarsinau@gmail.com  
**Website**: tarakreasi.com



## Project Purpose

### Why I Built This

**Primary Goal**: Demonstrate full-stack development capabilities
- Build a portfolio-quality web application
- Showcase modern tech stack (Laravel + Vue + Inertia)
- Practice UX/UI design (Design System V6)

**Secondary Goal**: Solve my own problem
- I needed a note-taking tool I could customize
- Wanted to publish notes as a blog without copy-pasting
- Learn by building something I'd actually use daily

**Result**: A fully functional tool that I (and others) can use for real work



## Technical Stack

### Modern Full-Stack Architecture

**Backend**:
- Laravel 12 (PHP 8.3+)
- MySQL Database
- RESTful API
- Laravel Breeze (Authentication)
- Laravel Policies (Authorization)

**Frontend**:
- Vue 3 (Composition API)
- Inertia.js (SPA without API)
- Vite (Build tool)
- Tailwind CSS + Custom Design System

**DevOps**:
- Git version control
- Composer & NPM
- PHPUnit testing
- Self-hostable (VPS deployment)



## Key Features

### 1. Dual-View Content Engine

**Private Workspace** (`/taranote`)
- OneNote-style WYSIWYG editor
- Organize notes into Notebooks
- Draft, edit, publish workflow

**Public Digital Garden** (`/articles`, `/portfolio`)
- Automatically generated blog
- Beautiful UI with V6 Design System
- SEO-optimized article pages

**The Magic**: One-click publish from draft → public article



### 2. Design System V6 "Electric Earth"

**Visual Features**:
- Glassmorphism (frosted glass panels)
- Noise texture overlay
- Ambient blob backgrounds
- Smooth scroll animations
- Dark mode support

**Result**: A modern, premium-feeling UI that showcases design skills



### 3. User Customization

**Settings Panel**:
- Upload avatar and hero images
- Customize portfolio page
- Set custom username (`/portfolio/yourname`)
- Configure public page titles and taglines

**Newsletter Subscription**:
- Email collection system
- Validation and duplicate handling
- Database storage (ready for email integration)



## Technical Highlights

### What This Project Demonstrates

**Full-Stack Development**
- Backend API & business logic
- Frontend SPA with reactive state
- Database design & migrations
- Authentication & authorization

**Software Architecture**
- MVC pattern (Laravel)
- Component-based UI (Vue)
- RESTful design
- Policy-based permissions

**Database Design**
- Relational schema (users, notes, notebooks, settings, subscribers)
- Foreign keys & constraints
- Migrations & seeders

**Security Best Practices**
- Password hashing (Bcrypt)
- CSRF protection
- Content sanitization (DOMPurify)
- XSS prevention
- SQL injection prevention (Eloquent ORM)

**Testing**
- PHPUnit feature tests
- Test coverage for critical flows
- Test-driven development practices

**UI/UX Design**
- Custom design system
- Responsive layouts
- Accessibility considerations
- Consistent branding

**Documentation**
- Comprehensive README
- API reference
- Security policy
- Privacy policy
- User guide
- Developer documentation



## Code Quality

### Development Practices

**Version Control**:
- Git with meaningful commits
- Feature branches
- Clear commit messages

**Code Standards**:
- Laravel Pint (PSR-12)
- ESLint for Vue
- Consistent code style

**Testing**:
- Unit tests for models
- Feature tests for controllers
- Test database with RefreshDatabase

**Documentation**:
- Inline code comments
- PHPDoc blocks
- Architecture Decision Records (ADR)
- Comprehensive user-facing docs



## Real-World Usage

### It's Not Just a Demo

**Daily Use**:
- I use TaraNote for my own note-taking
- Published articles serve as my blog
- Portfolio page showcases my writing

**Portfolio-Quality**:
- SSL/HTTPS support
- Database backups
- Error logging
- Performance optimization

**Self-Hostable**:
- Deploy to any VPS
- Full deployment documentation
- No vendor lock-in



## Project Stats

**Development Timeline**: ~3 months (part-time)  
**Lines of Code**: ~15,000+ (Backend + Frontend)  
**Test Coverage**: 73 passing tests (Feature + Unit)
**Documentation**: 3,000+ words across 20+ files  

**Tech Stack Depth**:
- Backend: Controllers, Models, Policies, Middleware, FormRequests
- Frontend: 15+ Vue components, Composables, Layouts
- Database: 6 tables, relationships, migrations
- Design: Custom V6 design system, 50+ Tailwind utilities



## Lessons Learned

### Technical Growth

**What I Learned**:
1. **Inertia.js**: Building SPAs without traditional REST APIs
2. **Vue 3 Composition API**: Reactive state management
3. **Laravel Policies**: Fine-grained authorization
4. **Design Systems**: Building consistent UI patterns
5. **Deployment**: VPS setup, SSL, domain configuration

**Challenges Overcome**:
- Balancing private/public views in same app
- Dynamic settings system (key-value store)
- Image upload & storage management
- Responsive design across devices
- Performance optimization (lazy loading, code splitting)



## Evolution Timeline

### V1-V3: Exploration (2025-09 to 2025-10)
- **V1 (Bootstrap)**: Ugly but functional MVP.
- **V2 (Bento Grid)**: Scrapped due to clutter.
- **V3 (OneNote Clone)**: Abandoned due to legal risk and lack of soul "imitation is not design".

### V4: TaraKreasi Branding (2025-11)
- Blue/Orange color scheme
- Material Icons
- Clean, minimal aesthetic

### V5: Technical Foundation (2025-11)
- Migrated to Vue 3 Composition API
- Implemented Inertia.js (SPA)
- Performance optimizations

### V6: Electric Earth (2025-12) ⭐ **Current**
- **Peak Design**: Glassmorphism, Noise Textures, Ambient Blobs
- Digital Garden concept fully realized
- Dark mode support



## Future Enhancements

### Roadmap (v1.0 and Beyond)

**Short-Term (Future Enhancements)**:
- Email automation for newsletters
- Server-side search (Laravel Scout)
- Comments system
- Mobile responsiveness improvements

**Medium-Term**:
- Markdown editor mode
- Export/import (Notion, Obsidian)
- Analytics dashboard
- Mobile apps (React Native)

**Long-Term**:
- Collaborative editing
- AI writing assistance
- Theme marketplace
- Plugin system

**Status**: Continuous improvement as I learn new techniques



## Use Cases

### Who Can Use TaraNote

**1. Personal Use** (Me)
- Daily note-taking
- Blog publishing
- Portfolio website

**2. Developers**
- Self-host for privacy
- Customize/extend codebase
- Learn from source code

**3. Writers/Creators**
- Private drafting workspace
- Public blog platform
- Newsletter management

**4. Students**
- Study notes organization
- Share learning publicly ("Digital Garden")
- Build knowledge base



## Why This Shows My Skills

### What Makes This Portfolio-Worthy

**1. Real-World Application**
- Not a todo list or weather app
- Solves an actual problem
- Used in production daily

**2. Full Feature Set**
- Authentication, authorization
- CRUD operations
- File uploads
- Email handling
- Custom settings

**3. Modern Stack**
- Latest Laravel & Vue versions
- Industry-standard tools
- Best practices

**4. Professional Polish**
- Beautiful UI (Design V6)
- Comprehensive docs
- Security considerations
- Testing coverage

**5. Open Source**
- Code available for review
- Clear README
- Contribution guidelines



## Technical Deep Dive

### Architecture Highlights

**Dual-View Pattern**:
- Same database, different UX
- Private: CRUD interface
- Public: Read-only, optimized for visitors
- Shared components for efficiency

**Settings Engine**:
- Key-value store for dynamic content
- No code changes needed for text updates
- Centralized configuration

**Content Publishing**:
- Draft → Published workflow
- Status management (Draft, Published, Archived)
- Slug generation for SEO

**User Profiles**:
- Custom usernames
- Avatar uploads
- Portfolio customization
- Public/private data separation



## Deployment

### Production Setup

**Current Hosting**:
- VPS (DigitalOcean/Hetzner)
- Ubuntu 22.04 LTS
- Nginx + PHP-FPM
- MySQL 8.0
- SSL via Let's Encrypt

**Infrastructure as Code**:
- Deployment scripts
- Database migrations
- Environment configuration
- Backup automation



## Open Source

### Code Availability

**Repository**: Available on request (private for now)  
**License**: MIT (open-source friendly)  
**Documentation**: Comprehensive setup guides  

**What You Can Do**:
- View source code
- Self-host your own instance
- Contribute improvements
- Learn from implementation



## Contact & Demo

### See It In Action

**Live Demo**: tarakreasi.com  
**Source Code**: Available upon request  
**Portfolio**: tarakreasi.com/portfolio  
**Contact**: ajarsinau@gmail.com  

**Links**:
- Email: ajarsinau@gmail.com
- Website: tarakreasi.com
- Contact Form: tarakreasi.com/contact



## Summary

# TaraNote: A Portfolio Project with Real Utility

### What It Demonstrates

**Full-Stack Expertise**: Laravel + Vue + Inertia  
**Modern Practices**: Testing, documentation, security  
**Design Skills**: Custom UI system, UX thinking  
**Problem-Solving**: Dual-view architecture, dynamic settings  
**Portfolio-Quality**: Deployed, maintained, actively used for personal productivity  

### Why It Matters

This isn't a toy project-it's a **fully functional web application** that:
- I use daily for real work
- Others could self-host and use
- Demonstrates industry-standard development practices
- Shows ability to ship complete products

**Bottom Line**: TaraNote showcases my ability to design, build, test, document, and deploy a modern web application from scratch.



*Last Updated: December 27, 2025 | Beta 0.9.80 (UI Design V6.5)*
