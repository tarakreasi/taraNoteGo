# TaraNote Design System V6.5 "Zen Glass"

**UI Design Version:** V6.5 ("Zen Mode") | **Application Version:** Beta 0.9.80
**Codename:** "Zen Garden"
**Philosophy:** Precision / Calm / Functional / Tactile

---

## 1. Design Philosophy: "The Anchor & The Void"

The interface prioritizes **mental flow** and **clarity**. It draws inspiration from Japanese technical precision and "Zen" aesthetics - removing all unnecessary visual noise (borders, harsh shadows, gradients) to leave only the content and essential structure.

### Core Principles
1.  **Left-Aligned Stability**: Unlike centralized layouts that "float" in a void, V6.5 anchors content to the left (`pl-0` to `pl-80`). This creates a stable reading axis that feels grounded, reminiscent of technical manuals or IDEs.
2.  **The Void (Negative Space)**: We use massive white space (padding > 120px) to give content room to breathe. The "void" is not empty; it is a tool for focus.
3.  **Thicker Glass**: Interface elements (Floating Dock, Sidebars) are not just transparent; they simulated "thick", frosted glass. This implies weight and permanence, treating tools as physical objects sitting *on top* of the content layer.

---

## 2. Layout System: The Trinity

The application uses a strict three-column layout (The Trinity) for the dashboard and writing experience.

### Column 1: The "Ghost" Sidebar (260px)
*   **Role**: Navigation context (Spaces/Notebooks).
*   **Behavior**: Fixed position.
*   **Visuals**:
    *   **No Borders**: It floats next to the content without a hard `border-r`. Distinction is created via background nuance (`bg-slate-50/40` vs `bg-white`).
    *   **Navigation**: Items are simple text with subtle icons. Active state is indicated by a soft background (`bg-indigo-50` / `ring-1 ring-indigo-500/20`).

### Column 2: The Stream (320px)
*   **Role**: List of items (Notes/Articles) within the selected Space.
*   **Behavior**: Scrollable independent of the Sidebar.
*   **Visuals**: Slightly denser density. Card-based items with title, date, and status.

### Column 3: The Canvas (Fluid)
*   **Role**: The primary reading/writing area.
*   **Behavior**: Fluid width (`max-w-7xl`), left-aligned but balanced.
*   **The "Zen" Constraint**: Text columns never exceed `75ch` (characters) for optimal readability, even heavily wide screens.

---

## 3. The Floating Dock System

The core navigation mechanism for the application, designed to separate "Tools" from "Content".

### Visual Construction
*   **Location**: Fixed at `bottom-6`, horizontally centered relative to the viewport.
*   **Material**: "Thicker Glass" aesthetic.
    *   **Light Mode**: `bg-slate-200/50` + `backdrop-blur-xl` + `border-white/20`.
    *   **Dark Mode**: `bg-white/10` + `backdrop-blur-xl` + `border-white/10`.
    *   **Shadow**: Deep, soft shadow (`shadow-xl shadow-black/10`) to lift it off the canvas.

### Clearance Strategy (Critical)
*   **Problem**: A fixed bottom dock can obscure content at the end of a document.
*   **Solution**: "Deep Bottom Padding".
*   **Implementation**: All main scroll containers enforce a **`pb-40` (160px)** bottom padding.
*   **Result**: The user can scroll the document until the final element is completely *above* the dock, ensuring zero obstruction.

---

## 4. Typography & Color

### Color Palette: "Monochrome + Indigo"
A restrained palette that relies on variations of Slate for depth, using Indigo only for focused interaction.

| Token | Light Mode (Hex) | Dark Mode (Hex) | Usage |
| :--- | :--- | :--- | :--- |
| **Primary** | `#4F46E5` (Indigo 600) | `#6366F1` (Indigo 500) | Active states, primary links |
| **Background** | `#FAFAFA` (Slate 50) | `#0F172A` (Slate 900) | Main canvas (The Void) |
| **Surface** | `#FFFFFF` (White) | `#0B1121` (Deep Navy) | Panels, Cards |
| **Text Primary** | `#1E293B` | `#F8FAFC` | Headings |
| **Text Muted** | `#64748B` | `#94A3B8` | Meta data |

### Typography Stack
*   **Headings**: `Inter` (Sans-serif). Bold, tight tracking (`tracking-tight`).
*   **Body**: `Inter` (Sans-serif). High readability, `leading-relaxed` (1.7).
*   **Data/Meta**: `JetBrains Mono`. Used for dates, IDs, and versions to convey semantic precision.

---

## 5. Animation Dynamics

We reject "bouncy" or "flashy" animations in favor of "Subtle Shifts".

*   **Hover**: `transition-all duration-300`. Elements do not "pop"; they gently shift state.
*   **Page Load**: No aggressive skeleton screens. Content fades in smoothly.
*   **The "Breathe" Effect**: Background blobs (if present) act as ambient light, moving slowly (10s+ duration) to make the screen feel "alive" without distracting.

---

## 6. Component Specification

### The "Zen" Input
*   **State:** Default border is transparent (`border-none`).
*   **Focus:** Smooth transition to `ring-1 ring-indigo-500`.
*   **Metaphor:** "Writing on paper, not filling a form."

### Status Indicators
*   **Minimal Dots**: Instead of heavy badges, use small colored dots (`w-2 h-2 rounded-full`) to indicate status (Online, Draft, Published).
*   **Mono Labels**: Status text is small (`text-[10px]`) and monospace to feel like a technical spec.

---

## 7. App Interface Taxonomy (The Garden)

The dashboard moves away from traditional file-system naming to organic terms:

*   **Spaces** (previously Notebooks): The high-level environments or categories.
*   **Stream** (previously Notes List): The active flow of ideas and content.
*   **Card** (previously Note): The individual unit of knowledge.

*This taxonomy reinforces the "Digital Garden" philosophy where ideas grow in spaces and flow in streams.*

