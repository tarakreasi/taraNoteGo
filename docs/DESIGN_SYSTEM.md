# Design System: Zen Glass (V6.5)

**Philosophy:** "Whisper, don't shout."
**Core Value:** Cognitive Load Reduction (Customer Service Mindset)

---

## 1. The "Why" (Empathy for the User)

As a former Customer Service representative, I learned that a "powerful" interface is useless if it exhausts the user. 

Early versions of TaraNote ("Electric Earth") were beautiful but loud-gradients screaming for attention. Users don't come to a note-taking app to admire the UI; they come to think.

**Zen Glass** is the engineering response to that problem:
*   **Goal:** Zero distraction.
*   **Method:** High-blur glassmorphism that hints at context without demanding focus.
*   **Result:** An interface that feels "breathable".

---

## 2. Core Tokens

### Color Palette (The Slate Scale)
We strictly avoid pure black (`#000000`) and pure white (`#FFFFFF`) to reduce eye strain during long writing sessions.

*   **Primary Action:** `#135bec` (Blue) - Used sparingly for "Save" or "Publish".
*   **Accent:** `#FF5F1F` (Orange) - Used ONLY for "Destructive/Attention" states.
*   **Background:** `slate-50` (Light) / `slate-900` (Dark) - Soft, expansive surfaces.

### Typography (Readability First)
*   **Headings:** `Outfit` - Geometric, modern, friendly.
*   **Body:** `Inter` - The gold standard for screen readability. Tall x-height reduces reading fatigue.
*   **Code:** `JetBrains Mono` - Ligatures allow for quicker logic parsing.

---

## 3. The "Zen Glass" Effect

Unlike standard flat design, Zen Glass uses layering to create depth without clutter.

### CSS Recipe for "Milky Glass"
This is the standard card background used throughout the app:

```css
.card-glass {
    /* 1. Low Opacity Background */
    background: rgba(255, 255, 255, 0.6);
    
    /* 2. Heavy Blur (The "Zen" factor) */
    backdrop-filter: blur(20px);
    
    /* 3. Subtle Border (The Definition) */
    border: 1px solid rgba(255, 255, 255, 0.1);
}
```

**Why this matters:**
Standard "Flat Design" feels cold. "Neuomorphism" feels heavy. **Zen Glass** feels organic-it changes slightly as you scroll over content, giving a subconscious sense of location without needing explicit grid lines.

---

## 4. Interaction Design (The "Hardware" Feel)

Inspired by physical buttons (tactile feedback), every interaction in TaraNote follows the **Input-Process-Output** model.

*   **Hover:** `duration-150` ease-out. It acknowledges presence immediately.
*   **Click:** `scale-95`. A physical "press" feeling.
*   **Loading:** Skeleton loaders, never spinners (perception of speed).

---

## 5. Accessibility (a11y)

**"Empathy is not a feature."**

*   **Contrast:** All text satisfies WCAG AA.
*   **Touch Targets:** Minimum 44px (iOS standard).
*   **Focus Rings:** High-visibility outline for keyboard navigation.

---

**References:**
*   [Design Tokens (JSON/Tailwind Config)](./DESIGN_TOKENS.md)
*   [The Story of Version 6.5](./UI/story_v6.5.md)
