# Research & Design Findings (UI 2026)

This document outlines the research and specifications that drove the evolution of the TaraNote UI towards the "Liquid Glass" and "Minimalism 3.0" aesthetic, specifically for the v6 and v5.2 iterations.



## 1. Aesthetic Direction: "Liquid Glass"

The core finding was to move away from flat design towards a material that feels like **frosted glass** suspended in a 3D space.

### Key Principles
- **Materiality**: Use `backdrop-filter: blur(20px)` to create depth.
- **Boundaries**: Replace heavy drop shadows with subtle **Ambient Glows** and thin semi-transparent white borders (`1px solid rgba(255,255,255,0.4)`).
- **Whitespace**: "Active Whitespace" strategy-padding should be 2x standard to create breathing room.
- **Shape**: "Squiress" (Super-ellipse) corners with large radii (e.g., `24px`).



## 2. Color System "Electric Earth"

Research indicated a need to balance professional clarity with organic warmth.

| Token | Color | Hex | Role |
| :--- | :--- | :--- | :--- |
| **Canvas** | Cloud Dancer | `#F0F2F5` | Reduces eye strain compared to pure white. |
| **Ink** | Deep Charcoal | `#2B2B2B` | Softer than pure black for text. |
| **Primary** | Electric Orange | `#FF5F1F` | High-energy accent for CTAs and interactions. |
| **Secondary** | Neo-Mint | `#A8E6CF` | Subtle accent for health/status indicators. |

**Usage Rule:**
> "Use Tarakreasi Orange (#FF5F1F) for all main actions (CTA), hover states, and 'pulse' indicators. Backgrounds should remain clean and neutral."



## 3. Layout Strategy: The Bento Grid

Moving away from the traditional list-view blog, the research favored a modular, grid-based approach.

### The 4x3 Home Hero Grid
A viewport-filling grid (90vh) constructed as follows:
1.  **2x2 (Left)**: Value Proposition / Large Typography ("We Design Future").
2.  **1x1 (Top Right)**: Visual Showreel / Loop.
3.  **1x1 (Mid Right)**: Live Statistics ("Living Numbers").
4.  **2x1 (Bottom)**: Giant CTA Pill Button (#FF7A01).

### The "Living Portfolio" (Digital Garden)
Instead of chronological lists, content is organized by **Growth Status**:
- **Seedlings**: Experimental ideas.
- **Budding**: Case studies in progress.
- **Evergreen**: Finished, polished major projects.



## 4. Navigation: The Floating Dock

**Decision:** Remove the top static navbar.
**Solution:** Implement a **Floating Dock** at the bottom of the screen (macOS style).
- **Benefit:** Makes the upper canvas improved for content and feels more "app-like".
- **Glass Effect:** The dock itself uses the `glass-panel` token.



## 5. Interaction Design "Generative Feel"

The UI should feel anticipatory and alive.
- **Hover**: Elements shouldn't just change color; they should "float" (transform Z-axis) and change glass opacity.
- **Idle State**: After 3 seconds of inactivity, the UI can gently pulse or offer a suggestion ("Time for a break?").



*Verified Spec based on 2026 Trend Analysis.*
