# Product Requirements Document - V3.0 (OneNote Redesign)

**Date:** October 21, 2025  
**Version:** 3.0  
**Status:** Deprecated (See V6)

## 1. Overview
The V3 update aims to pivot TaraNote from a generic CRUD app to a desktop-class productivity tool. We will emulate the Microsoft OneNote interface to provide a familiar experience for power users.

## 2. Design Specifications
-   **Theme:** "Microsoft Purple" (#7719AA)
-   **Iconography:** Microsoft Office style icons.
-   **Typography:** Segoe UI.

## 3. New Features
-   **Three-Column Layout:** Notebooks -> Sections -> Page Content.
-   **Horizontal Tabs:** Sections displayed as tabs across the top.
-   **Bento Grid Dashboard:** A new home screen with widget-style entry points.

## 4. Risks & Constraints
-   **Licensing:** Need to ensure we don't infringe on MS trademarks with the UI.
-   **Complexity:** The horizontal tab logic requires complex state management in Vue.

## 5. Success Metrics
-   User feels like they are using a native desktop app in the browser.
