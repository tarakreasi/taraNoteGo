# TaraNote UI v3 - Electric Orange Edition

**Design Philosophy:** Minimalism 3.0 + Liquid Glass with vibrant Electric Orange accents



## Color Palette

### Primary Colors
```css
--bg-color: #F8FAFC;           /* Slate 50 - Clean background */
--text-primary: #1E293B;        /* Slate 800 - Main text */
--accent-orange: #FF5F1F;       /* Electric Orange - Primary CTA */
--accent-blue: #4485EE;         /* Tara Blue - Secondary accent */
```

### Glass Effects
```css
--glass-bg: rgba(255, 255, 255, 0.75);
--glass-border: rgba(68, 133, 238, 0.25);  /* Blue-tinted */
--shadow-ambient: 0 12px 40px -10px rgba(68, 133, 238, 0.15);
```



## Typography

**Font Family:** Inter (sans-serif)
**Base Size:** 18px (comfortable reading)

### Scale
- Hero Title: `text-6xl` (60px) -> `text-8xl` (96px) on desktop
- Subtitle: `text-2xl` (24px) -> `text-3xl` (30px) on desktop
- Body: `text-lg` (18px)
- Small: `text-base` (16px)

### Special Effects
- **Orange Gradient Text:** `linear-gradient(135deg, #FF5F1F, #FF8C00)`
- **Letter Spacing:** `-0.03em` for titles (tight, modern)
- **Line Height:** `1.1` for headings



## Components

### 1. Glass Panel
```css
.glass-panel {
    background: rgba(255, 255, 255, 0.75);
    backdrop-filter: blur(24px);
    border: 1px solid rgba(68, 133, 238, 0.25);
    border-radius: 32px;
    box-shadow: 0 12px 40px -10px rgba(68, 133, 238, 0.15);
}

.glass-panel:hover {
    transform: translateY(-6px) scale(1.01);
    background: rgba(255, 255, 255, 0.85);
    border-color: rgba(68, 133, 238, 0.4);
}
```

**Usage:** Navigation bar, article cards, footer



### 2. Liquid Button (Orange)
```css
.liquid-button {
    background: linear-gradient(to-r, #fb923c, #f97316);
    color: white;
    padding: 1rem 2rem;
    border-radius: 20px;
    box-shadow: 0 4px 15px rgba(68, 133, 238, 0.3);
}
```

**Applied to:** Subscribe button (with orange gradient override)



### 3. Badge
```css
.badge-blue {
    background: #E0F2FE;
    color: #0284C7;
    padding: 0.5rem 1rem;
    border-radius: 99px;
    font-size: 0.85rem;
    font-weight: 700;
    text-transform: uppercase;
}
```

**Usage:** Topic tags, category labels



### 4. Orange Gradient Text
```css
.text-gradient-orange {
    background: linear-gradient(135deg, #FF5F1F, #FF8C00);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
```

**Applied to:** Hero title "code & servers"



## Layout Structure

### Container Widths
- **Maximum:** `1400px` (generous, comfortable)
- **Padding:** `px-6 lg:px-20` (responsive)
- **Vertical Spacing:** `py-20 lg:py-32` (active whitespace)

### Grid System
- **Articles:** 3 columns on desktop (`lg:grid-cols-3`)
- **Gap:** `gap-10 lg:gap-16` (spacious)



## Page Sections

### 1. Navigation Bar
- Sticky glass panel
- Logo (left), Dashboard/Login + Subscribe button (right)
- Orange gradient subscribe button
- Height: `py-6` (24px padding)

### 2. Hero Section
- Centered layout
- Orange/blue gradient badge with glowing dot
- Massive title (96px on desktop) with orange gradient
- Large subtitle (30px)
- Topic filter buttons (pill-shaped)

### 3. Article Grid
- Glass panel cards
- Blue-tinted borders
- Hover: lift + scale effect
- Orange "Read Story" link on hover
- Badge for category

### 4. Footer
- Glass background
- Multi-column layout
- Navigation links
- Brand information



## Interactions & Animations

### Hover States
```css
/* Cards */
transform: translateY(-6px) scale(1.01);
border-color: rgba(68, 133, 238, 0.4);

/* Buttons */
transform: translateY(-2px);
box-shadow: 0 8px 25px rgba(68, 133, 238, 0.4);

/* Links */
color: #FF5F1F; /* Orange on hover */
```

### Transitions
- **Duration:** `0.3s - 0.4s`
- **Easing:** `cubic-bezier(0.2, 0.8, 0.2, 1)` (smooth, natural)



## Special Features

### 1. Engineering Blog Badge
- Orange to blue gradient background
- Glowing orange dot with shadow
- Gradient text (orange -> blue)

### 2. Subscribe Button
- Full orange gradient (`from-orange-400 to-orange-500`)
- Hover: darker gradient (`from-orange-500 to-orange-600`)
- Pill-shaped (fully rounded)

### 3. Background Gradient
- Radial gradient from top center
- Blue glow: `rgba(68, 133, 238, 0.1)`
- Subtle, doesn't overpower content



## Implementation Files

### CSS
- **File:** `resources/css/app.css`
- **Variables:** `:root` section
- **Utilities:** `.glass-panel`, `.liquid-button`, `.badge-blue`, `.text-gradient-orange`

### Vue Component
- **File:** `resources/js/Pages/ArticleList.vue`
- **Sections:** Nav, Hero, Articles Grid, Footer



## Design Principles

1. **Active Whitespace** - Generous spacing (2x normal)
2. **Liquid Glass** - Glassmorphism with blur effects
3. **Electric Orange** - Vibrant, energetic primary color
4. **Comfortable Sizing** - 18px base, large touch targets
5. **Smooth Interactions** - Subtle, natural animations
6. **Brand Consistency** - TaraKreasi orange + blue identity



## Browser Support

- **Backdrop Filter:** Modern browsers (Chrome 76+, Safari 9+, Firefox 103+)
- **Gradient Text:** All modern browsers with `-webkit-` prefix
- **Flexbox/Grid:** Universal support



## Future Enhancements (Optional)

- [ ] Add ambient orange glow orbs in background
- [ ] Implement card hover with orange border
- [ ] Add micro-animations on scroll
- [ ] Dark mode variant
- [ ] Reduce motion preference support



**Version:** 3.0  
**Last Updated:** December 2025  
**Status:** Implemented
