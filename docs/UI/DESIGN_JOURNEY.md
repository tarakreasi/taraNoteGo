# The Design Journey: From Concept to "Electric Earth"

**A Chronological History of TaraNote's Interface Evolution**

> "Design is not just what it looks like and feels like. Design is how it works." - Steve Jobs

This document chronicles the complete design history of TaraNote, from its humble beginnings as a Bootstrap template to its current state as a bespoke "Living Digital Garden".



## Era 1: The Foundation (V1 & V2)
*September - October 2025*

In the beginning, the goal was simple functionality. Aesthetics were secondary to getting the code to work.

### V1: The "Bootstrap" Era
**Codename:** *Init*  
**Philosophy:** Pure Functionality.  
**Tech:** Standard Bootstrap CSS.

V1 was the raw MVP. It looked like a generic admin panel because it *was*. The focus was entirely on database schema and basic CRUD operations.
- **[View Admin Dashboard](./v1/admin_dashboard.html)** (Raw HTML)
- **[View Blog Home](./v1/blog_home.html)** (Raw HTML)

### V2: The "Bento Grid" Experiment
**Codename:** *Gridlock*  
**Philosophy:** Modern Layouts.  
**Tech:** CSS Grid.

V2 tried to be trendy. We implemented a "Bento Grid" layout for the dashboard. While visually interesting, it proved confusing for deep content hierarchies. It was a failure in UX, but a success in learning CSS Grid.
- **[View Admin Dashboard](./v2/admin_dashboard.html)** (Raw HTML)



## Era 2: The Imitation (V3)
*October 2025*

### V3: The OneNote Clone
**Codename:** *Copycat*  
**Philosophy:** "Good artists copy."  
**Tech:** Complex Flexbox layouts.

We abandoned creativity to solve the UX problem. "If Microsoft OneNote works, let's just build that." We successfully replicated the OneNote interface (Sections on top, Pages on left).
- **Outcome:** It worked perfectly but felt "soulless". It had no identity. It was just a worse version of a Microsoft product.
- *Note: No functional HTML prototypes remain effectively from this era; they were refactored into V4.*



## Era 3: The Awakening (V4)
*November 2025*

### V4: The Pivot (TaraKreasi)
**Codename:** *Identity*  
**Philosophy:** "Be Yourself."  
**Tech:** Vue.js + Tailwind CSS.

This was the turning point. We stripped away the purple OneNote branding and established our own: **Electric Blue (#4485EE)** and **Sunset Orange (#FF5F1F)**. We moved navigation to a sidebar and simplified the layout.
- **[Notebook Interface Study](./v4/notebook_legacy.html)** (Interactive Vue Prototype)
- **[Dashboard Layout Study](./v4/dashboard_legacy.html)** (Static HTML)



## Era 4: The Refinement (V5)
*Late November 2025*

### V5: The Vue Transition
**Codename:** *Structure*  
**Philosophy:** Component-Driven Design.

V5 was about porting V4's design into real Vue components. The "Prototypes" became the actual application code.

### V5.2: The Static Export
A snapshot of the codebase as it stood before the V6 redesign.
- **[Dashboard Snapshot](./v5.2/dashboard.html)**
- **[Article List](./v5.2/article_list.html)**



## Era 5: The Masterpiece (V6)
*December 2025 (Current)*

### V6: Electric Earth
**Codename:** *Living Garden*  
**Philosophy:** Glassmorphism, Ambient Intelligence, Organic Feel.  
**Tech:** Tailwind CSS + Backdrop Blur + CSS Animations.

V6 is the current, active design system. It is no longer just a "tool"-it is a place. Usage of "Glassmorphism" creates depth, while "Ambient Blobs" make the interface feel alive. It differentiates the **Private Workspace** (Structured) from the **Public Garden** (Organic).

**Key Prototypes (High Fidelity):**
- **[The Dashboard](./v6/dashboard.html)**: The command center.
- **[The Login Portal](./v6/login.html)**: First impressions matter.
- **[The Landing Page](./v6/home.html)**: The public face.
- **[The Article View](./v6/article.html)**: The reading experience.



## Conclusion
TaraNote's design journey reflects the developer's growth:
1.  **V1/V2**: Can I build it? (Engineering)
2.  **V3**: Can I make it usable? (UX)
3.  **V4**: Can I give it soul? (Branding)
4.  **V6**: Can I make it art? (Aesthetics)
