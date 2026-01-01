# Technical Workshop (Prototypes)

**Status:** Experimental | **Tech:** HTML + Tailwind CSS (No Vue)

## Overview
The `demo/` folder serves as the "technical workshop" for the project. Before implementing complex UI features in the Vue application, we first build them here as isolated, static HTML files. This allows us to iterate on layout and CSS without the overhead of the build process.

## Active Experiments

### 1. [Landing Page (index.html)](./index.html)
The prototype for the V6 Landing Page (`Home.vue`).
-   **Focus:** Sticky glass navbar, Hero typography, Newsletter component styling.

### 2. [Dashboard (dashboard.html)](./dashboard.html)
The prototype for the V6 Dashboard.
-   **Focus:** Three-column layout, sidebar states, and editor integration (Legacy Quill).

### 3. [Login Portal (login.html)](./login.html)
The prototype for the Auth screens.
-   **Focus:** Ambient background blobs and glass card centering.

> **Note:** These files are independent of the main app. They load dependencies via CDN for portability.
