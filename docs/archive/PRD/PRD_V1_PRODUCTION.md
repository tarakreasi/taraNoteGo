# PRD: TaraNote V1.0

**Current**: Beta 0.9.60 (UI V6)  
**Target**: V1.0 Production  
**Timeline**: ~2 months (solo)  
**Last Updated**: Dec 21, 2025



## What We Have (Beta 0.9.80)

The app works. It's tested (64 tests passing), documented, and I use it daily. The UI is solid-V6 "Electric Earth" looks professional. Code is clean (no technical debt markers in app code).

**The problem**: It's still just me using it. V1.0 means making it ready for others.



## What's Missing

### Must-Fix (Can't launch without these)
1. **Analytics** - I have no idea if anyone's actually using the public site. Need Plausible or similar.
2. **Email** - Newsletter signups go to database and... nothing. Need Mailgun/SendGrid integration.
3. **Performance** - Page loads feel slow (2-3s). Should be under 1.5s. Need caching.
4. **Search** - Can't find my own articles. Embarrassing. Laravel Scout + Meilisearch.
5. **Mobile** - Works on desktop, breaks on phones. Not acceptable.

### Would Be Nice
6. **Markdown** - Some people hate WYSIWYG. Fair enough.
7. **Export** - Let people take their data out. Reduces "vendor lock-in" feeling even for self-hosted.
8. **Comments** - Articles feel dead without them. Commento is privacy-friendly.
9. **Analytics Dashboard** - I want to see growth charts in admin panel, not just raw Plausible.



## The Plan

### Month 1: Infrastructure
**Goal**: Make the app fast and measurable.

**Analytics (3 days)**
- Add Plausible script to `app.blade.php`
- Track: `note_created`, `note_published`, `article_viewed`
- Show basic metrics in admin dashboard (MAU, top articles)

**Email System (5 days)**
- Mailgun API integration
- Weekly digest email (latest articles)
- Unsubscribe link (GDPR requirement)
- Admin panel to compose/send newsletters

**Performance (4 days)**
- Laravel cache for article lists (5min timeout)
- Image lazy loading
- Code splitting in Vite
- Database indexes on `notes.status` and `published_at`
- Target: Lighthouse score > 90

### Month 2: Features People Want

**Search (3 days)**
- Laravel Scout + Meilisearch
- Search box on articles page
- Highlights matching text
- Filter by notebook

**Markdown Editor (5 days)**
- Toggle in settings: WYSIWYG vs Markdown
- Split-pane preview
- CodeMirror for syntax highlighting
- Store `content_type` in DB ('html' or 'markdown')

**Export (2 days)**
- Export single note as Markdown or PDF
- Export all notes as ZIP
- Compatible with Obsidian/Notion import

### Month 3: Polish

**Comments (2 days)**
- Integrate Commento (self-hosted, no tracking)
- Show below published articles
- Email me when someone comments

**Mobile Fix (4 days)**
- Test on iPhone SE, Pixel 5, iPad Mini
- Fix sidebar overlap on small screens
- Make settings panel scrollable
- Resize portfolio images

**Analytics Dashboard (3 days)**
- New page: `/dashboard/analytics`
- Charts: articles over time, top posts, subscriber count
- Use ApexCharts (matches V6 design)



## Timeline

| Phase | Days | What's Done |
|-------|------|-------------|
| Month 1 | 12 | Analytics, Email, Performance |
| Month 2 | 10 | Search, Markdown, Export |
| Month 3 | 9 | Comments, Mobile, Dashboard |
| Testing | 5 | Docs, final QA |
| **Total** | **36** | Launch |

Realistically: 6-8 weeks if I work nights/weekends.



## Success Looks Like

**Launch Day**:
- App is live at taranote.com (or subdomain)
- Documentation says "V1.0"
- Tests still passing (target: 85+ tests)

**3 Months Later**:
- 10+ people self-hosting (tracked via GitHub issues)
- 100+ GitHub stars (means developers care)
- 50+ newsletter subscribers (organic)
- No critical bugs reported



## Risks

**"What if nobody uses it?"**  
Then it's still a good portfolio piece. The code quality and documentation speak for themselves.

**"What if it costs too much to host?"**  
Start with free tier (Railway, Fly.io). Only scale if needed.

**"What if I don't have time?"**  
Cut Phase 3. Launch with just Phase 1+2.



## Notes to Self

- Don't over-engineer. V1.0 doesn't need to be perfect.
- Focus on making it work for 10 users, not 10,000.
- Keep the "Digital Garden" concept-it's what makes this different from Notion clones.
- Remember: this is a portfolio project first. It needs to show skills, not make money.



**Next**: Start with Analytics integration. 3 days max.
