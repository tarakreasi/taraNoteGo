# Frequently Asked Questions (FAQ)

**TaraNote User FAQ**  
**Version: Beta 0.9.80** | **UI: V6.5 "Zen Mode" Design System**

> **[BETA NOTICE]**: TaraNote is currently in beta testing. Some features may be incomplete or contain bugs. Use at your own risk for now.



## General Questions

### What is TaraNote?

TaraNote is a **Digital Garden** platform that combines private note-taking with public knowledge sharing. Think of it as a cross between Notion (for private notes) and Medium (for public articles).



### Is TaraNote free?

Yes! TaraNote Beta is **100% free and open-source**. You can:
- Self-host it on your own server
- Use it for personal or commercial projects

**Future**: Premium features and hosted plans may be introduced after v1.0 (production release).



### What's the difference between a "Note" and an "Article"?

-**Note**: A private draft in your workspace (`/taranote`)
- **Article**: A published note visible on your public Digital Garden (`/articles`)

When you click "Publish" on a note, it becomes an article.



### Can I use TaraNote offline?

Not yet. TaraNote is a web application that requires an internet connection. **Offline mode** is planned for v1.0 (production release).



## Account & Privacy

### How do I create an account?

1. Visit the Landing Page
2. Click **"Get Started"** or **"Register"**
3. Enter your name, email, and password
4. Check your email for verification (if enabled)



### Is my data private?

**Yes**. Your notes are private until you explicitly publish them. See our [Privacy Policy](PRIVACY_POLICY.md) for details.

We **never**:
- Sell your data
- Share your content without permission
- Train AI models on your notes



### Can I delete my account?

Yes. Go to **Settings -> Account -> Delete Account**.

> **Warning**: This action is **permanent** and cannot be undone. All your notes, articles, and data will be deleted.



### How do I export my data?

Go to **Settings -> Export Data** (coming in V6.2).

You'll receive a **JSON file** with all your notes, notebooks, and metadata.



## Writing & Publishing

### How do I create a note?

1. Log in to your account
2. Go to **TaraNote** (private workspace)
3. Click **"New Note"**
4. Write your content using the WYSIWYG editor
5. Click **"Save Draft"**



### How do I publish an article?

1. Open a draft note
2. Click **"Publish"**
3. The note becomes an article visible at `/articles/{slug}`

**Tip**: Set a cover image and excerpt before publishing for better presentation!



### Can I unpublish an article?

Yes. Click **"Unpublish"** in the editor to convert it back to a draft.



### What is the "Maturity" level?

Maturity indicates how polished your article is:

- [Seedling]: Early ideas, rough drafts
- [Budding]: Work in progress, partially complete
- [Evergreen]: Fully polished, comprehensive articles

This helps readers set expectations.



### How do I add images to my articles?

1. In the editor, click the **Image** button
2. Upload an image (max 5MB, `.jpg`, `.png`, `.gif`, `.webp`)
3. The image is inserted into your article

**Cover Images**: Set via **Settings -> Cover Image** in the note editor.



### Can I use Markdown?

Not yet. TaraNote uses a **WYSIWYG editor** (What You See Is What You Get).

**Markdown support** is planned for V6.2 (you'll be able to toggle between WYSIWYG and Markdown mode).



### How do I organize my notes?

Use **Notebooks** (categories):

1. Create a new notebook (e.g., "Tutorials", "Personal")
2. Assign notes to notebooks when creating/editing
3. Filter notes by notebook in the sidebar



## Public Profile & Portfolio

### What is my "Digital Garden"?

Your Digital Garden is the **public-facing** version of your notes. It includes:
- **Articles Page** (`/articles`): All your published articles
- **Portfolio Page** (`/portfolio/{username}`): Your author profile + articles



### How do I customize my portfolio?

1. Go to **Settings -> Portfolio**
2. Upload a **Hero Cover Image**
3. Set your **Hero Tagline** (e.g., "Full-Stack Developer")
4. Choose a custom **Username** (e.g., `/portfolio/johndoe`)



### Can I use a custom domain?

Not yet. Custom domains (e.g., `blog.yourname.com`) are planned for **Premium users** in V6.3.



## Newsletter & Subscriptions

### How do I subscribe to a newsletter?

1. Visit any **Landing Page** or **Article Page**
2. Scroll to the **Newsletter** section
3. Enter your email and click **"Subscribe"**

You'll receive a confirmation message instantly.



### How often will I receive emails?

**Maximum once per week** (when the author publishes new content).



### How do I unsubscribe?

Click the **"Unsubscribe"** link at the bottom of any newsletter email.

Or email us at ajarsinau@gmail.com with your email address.



## Technical & Troubleshooting

### What browsers are supported?

TaraNote works best on:
- **Chrome** / **Edge** (v90+)
- **Firefox** (v88+)
- **Safari** (v14+)

**Mobile**: iOS Safari, Chrome for Android.



### Why is the editor lagging?

If the WYSIWYG editor feels slow:
1. **Check your connection**: Slow internet can cause delays
2. **Reduce article length**: Very long articles (>10,000 words) may lag
3. **Clear browser cache**: `Ctrl+Shift+Delete` -> Clear Cache

Still slow? Report it to ajarsinau@gmail.com.



### Images aren't loading. What do I do?

1. **Check file size**: Max 5MB per image
2. **Check format**: Only `.jpg`, `.png`, `.gif`, `.webp` are allowed
3. **Try a different browser**: Some browsers block uploads
4. **Clear cache**: `Ctrl+F5` to hard refresh



### I forgot my password. How do I reset it?

1. Click **"Forgot Password?"** on the login page
2. Enter your email address
3. Check your inbox for a password reset link
4. Click the link and set a new password

**No email?** Check your spam folder or email ajarsinau@gmail.com.



### Can I collaborate with others?

Not yet. **Collaborative editing** (multiple authors per note) is planned for V7.0.



## Features & Roadmap

### Will there be a mobile app?

Yes! Native **iOS** and **Android** apps are planned for **v1.0** (production release).



### Can I import notes from other apps?

**Coming soon**. We're planning importers for:
- **Notion** (v0.10)
- **Obsidian** (v0.10)
- **Evernote** (v1.0)
- **Google Docs** (v1.0)



### Is there a dark mode?

Yes! Click the **Moon/Sun icon** in the top-right corner to toggle between light and dark mode.



### Can I add comments to articles?

Not yet. **Comments system** is planned for V6.1.



### Will TaraNote support AI writing assistance?

**Maybe**. We're exploring:
- AI-powered grammar checking
- Auto-generated summaries/excerpts
- Content suggestions

This would be **opt-in** and privacy-preserving (no data sent to third parties).



## Pricing & Premium

### Will TaraNote always be free?

The **core features** (note-taking, publishing, portfolio) will always be free and open-source.

**Premium features** (post v1.0) may include:
- Custom domains (`blog.you.com`)
- Advanced analytics dashboard
- Priority support
- White-label branding removal

**Price estimate**: TBD (not finalized, likely $5-15/month).

> **[NO COMMITMENT]:** Premium pricing and features are **speculative only** and for demonstration purposes. As a hobby/portfolio project, there is **no guarantee** these will be implemented. TaraNote may remain 100% free forever, or development may stop. **No promises, no obligations.**



## Community & Support

### How do I report a bug?

1. **GitHub Issues**: [github.com/tarakreasi/taranote/issues](#)
2. **Email**: ajarsinau@gmail.com
3. **Discord** (coming soon): Join our community

Please include:
- Browser and OS version
- Steps to reproduce
- Screenshots (if applicable)



### How can I contribute?

See [CONTRIBUTING.md](CONTRIBUTING.md) for:
- How to set up the dev environment
- Coding standards
- How to submit a pull request



### Where can I get help?

- **Documentation**: [tarakreasi.com/docs](https://tarakreasi.com/docs)
- **Email Support**: ajarsinau@gmail.com
- **Community Forum** (coming soon)



## Didn't find your answer?

**Contact us**: ajarsinau@gmail.com

We typically respond within 24-48 hours.



**Last Updated:** December 27, 2025 (Beta 0.9.80 - UI Design V6.5)
