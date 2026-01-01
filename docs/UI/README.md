# TaraNote UI Documentation

**Application Version:** Beta 0.9.80
**UI Design System:** V6.5 "Zen Glass" (Current)



## The Story
**[READ: The Design Journey (History)](./DESIGN_JOURNEY.md)**  
A chronological narrative of how we went from a basic Bootstrap admin panel to the current "Living Digital Garden".



## Category Index

### 1. The Design System (Current)
Documentation for V6 "Electric Earth".
- **[UI/UX Overview](../UI_UX_OVERVIEW.md)**: Core philosophy & strategy.
- **[Design Tokens](../DESIGN_TOKENS.md)**: Color, Typography, Spacing variables.
- **[Component Library](../COMPONENT_LIBRARY.md)**: Buttons, Cards, Modals.
- **[Accessibility](../ACCESSIBILITY.md)**: WCAG 2.1 compliance.

### 2. The Prototypes (Gallery)
Direct links to functional HTML prototypes from different eras.

| Version | Description | Link |
| :--- | :--- | :--- |
| **V6 (Latest)** | **Dashboard** (Glassmorphism) | [Launch](./v6/dashboard.html) |
| **V6 (Latest)** | **Login Portal** | [Launch](./v6/login.html) |
| **V6 (Latest)** | **Landing Page** | [Launch](./v6/home.html) |
| **V4 (Legacy)** | **Notebook Interface** (Vue Prototype) | [Launch](./v4/notebook_legacy.html) |
| **V4 (Legacy)** | **Dashboard Study** (Static) | [Launch](./v4/dashboard_legacy.html) |
| **V5.2** | **Static Export** | [Launch](./v5.2/dashboard.html) |
| **V1** | **Bootstrap Admin** | [Launch](./v1/admin_dashboard.html) |

### 3. Research & Lab
Behind-the-scenes work.
- **[The Workshop](./workshop/README.md)**: Rough experiments and code snippets.
- **[The Design Lab](./research/README.md)**: Theoretical foundations.



## Version Archives
Deep dive into specific version folders.
- **[V6 (Electric Earth)](./v6/README.md)**
- **[V5 (Vue Transition)](./v5/README.md)**
- **[V4 (TaraKreasi Pivot)](./v4/README.md)**
- **[V3 (OneNote Clone)](./v3/README.md)**
- **[V2 (Bento Grid)](./v2/README.md)**
- **[V1 (Bootstrap)](./v1/README.md)**



## How to Run Prototypes

**V4/V1/V2 Prototypes:**
Run directly in browser. Click the "Launch" links above.

**V6 Prototypes:**
Best viewed with a local server due to asset loading.
```bash
# Inside docs/UI/
python3 -m http.server
# Open http://localhost:8000
```
