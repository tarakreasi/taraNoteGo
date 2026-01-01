**Current Version**: Beta 0.9.60 (UI Design V6)

**TaraNote - Your Digital Garden**  
**Current Version: Beta 0.9.60** | **UI Design System: V6 "Electric Earth"**



## Beta 0.9.60 (December 21, 2025)

> **Beta Notice**: This is a pre-release version. Some features are still in development and may contain bugs. Production release (v1.0) is planned for Q2 2026.

### New Features
### New Features

#### Newsletter Subscription
You can now subscribe to receive updates directly from the Digital Garden! Look for the subscription form on:
- The **Landing Page** (bottom of the page)
- **Individual Article Pages** (after reading)

Simply enter your email and click Subscribe. You'll receive a confirmation message instantly!

#### Floating Navigation Bar
The navigation bar now "floats" above the content, making it easier to navigate between pages without scrolling back to the top. It stays visible as you scroll!

#### Green Portfolio Theme
The Portfolio page now features a lush **Emerald & Lime** color scheme that reinforces the "Digital Garden" metaphor. It feels more organic and alive!

### Improvements

- **Better Component Names**: Internal code structure improved for clarity (Welcome -> Home, Home -> Articles)
- **Smoother Animations**: Navigation transitions are now butter-smooth
- **Accessibility**: Better contrast and focus states for keyboard navigation

### Bug Fixes

- Fixed navbar positioning on mobile devices
- Improved form validation feedback
- Fixed dark mode color consistency on Portfolio page



## UI Design System V6 "Electric Earth" (December 2025)

### Complete UI Redesign

The entire visual language of TaraNote has been reimagined with the **V6 "Electric Earth" Design System**. This is a major aesthetic upgrade focused on creating a "Living Digital Garden" feel.

### What's New in V6 Design

#### Dual-View Engine
TaraNote now offers **two distinct experiences**:

1.  **Private Workspace** (`/taranote`)
    - OneNote-style note-taking interface
    - Organize notes into Notebooks
    - Draft, publish, or archive articles
    - WYSIWYG editor with live preview

2.  **Public Digital Garden** (`/articles`)
    - Beautiful, reader-friendly article display
    - Search and filter by topics
    - Trending articles sidebar
    - Topics cloud visualization

#### Electric Earth Design System
A completely redesigned UI featuring:
- **Glassmorphism**: Frosted glass panels with blur effects
- **Noise Texture Overlay**: Subtle film grain for a tactile feel
- **Ambient Blobs**: Floating color gradients that create depth
- **Smooth Animations**: Scroll-reveal effects and micro-interactions

#### Customizable Profiles
- Upload custom avatars
- Set a hero tagline and cover image
- Choose your custom username URL (`/portfolio/yourname`)
- Personalized "About Me" section

#### Settings Panel
Fine-tune your Digital Garden:
- **Welcome Page**: Customize hero text, stats, and features
- **Articles Page**: Set search placeholder and hero section
- **Portfolio**: Upload cover image and set tagline
- **Footer**: Social links, copyright text, and status message

#### Smart Article Management
- **Maturity Levels**: Mark notes as Seedling, Budding, or Evergreen
- **Cover Images**: Add visual appeal to your articles
- **Excerpts**: Automatic or custom summaries
- **Slug Generation**: SEO-friendly URLs

#### Search & Discovery
- Full-text search across all articles
- Filter by notebook/category
- Topics cloud for quick navigation
- Trending articles sidebar

### Performance
- **40% faster** page loads with optimized assets
- **Lazy loading** for images and components
- **Code splitting** for minimal initial bundle size

### Security
- Laravel Breeze authentication
- Policy-based authorization
- Content sanitization (DOMPurify)
- CSRF protection on all forms



## Upgrading to Beta 0.9.60

If you're upgrading from a previous alpha version:

1. **Database**: Run `php artisan migrate` to add new tables (`settings`, `subscribers`)
2. **Settings**: Populate default settings via the Settings Panel
3. **Avatars**: Re-upload user avatars (they're now stored in `storage/avatars/`)
4. **Notebooks**: Ensure notebooks have slugs (`articles` and `portfolio` are recommended)

See `docs/DEPLOYMENT.md` for full upgrade instructions.



## What's Next?

We're working on:
- **Comments System**: Enable discussions on articles
- **Analytics Dashboard**: Track views, popular articles, and engagement
- **Email Automation**: Send newsletters to subscribers
- **Mobile App**: Native iOS/Android apps for on-the-go note-taking
- **API v2**: GraphQL support for advanced integrations

See `docs/ROADMAP.md` for the full feature roadmap.



## Known Issues (Beta)

### Beta 0.9.60
- **Beta Status**: Not recommended for production use yet
- Newsletter emails are stored but not yet sent automatically (manual export required)
- Table of Contents on Article pages is static (not parsed from content)
- Search is client-side only (server-side search coming in v0.10)
- Some UI animations may be choppy on older devices

### Planned for v1.0 (Production)
- Email automation for newsletters
- Server-side search with Laravel Scout
- Performance optimizations
- Comprehensive error handling

### Workarounds
- To export newsletter emails: `php artisan tinker` -> `Subscriber::all()->pluck('email')`
- For server-side search: Use Laravel Scout (see `docs/CONTRIBUTING.md`)



## Thank You!

TaraNote wouldn't be possible without:
- The Laravel and Vue.js communities
- Our early adopters and testers
- The open-source ecosystem

**Feedback?** We'd love to hear from you! Email us at [ajarsinau@gmail.com](mailto:ajarsinau@gmail.com) or open an issue on GitHub.



**Documentation**: [tarakreasi.com/docs](https://tarakreasi.com/docs)  
**Support**: [ajarsinau@gmail.com](mailto:ajarsinau@gmail.com)  
**GitHub**: [github.com/tarakreasi/taranote](#)



*Last Updated: December 27, 2025 - Beta 0.9.80 (UI Design V6.5)*
