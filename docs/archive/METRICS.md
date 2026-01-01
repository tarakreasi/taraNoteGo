# Product Metrics Strategy (Planned)

**TaraNote Success Framework (Draft Strategy for V1.0)**

> **Portfolio Context**: This document represents the **Analytics Strategy** for the future V1.0 release.
> **Current Status**: **NOT IMPLEMENTED**. No tracking code (GA4/Plausible) is currently installed in the codebase.
> **Purpose**: This file demonstrates **Product Management** and **Data Strategy** skills by defining *what* should be measured when the product goes live.



## 1. North Star Metric (Strategic Goal)

**Target**: **Active Content Creators**
*Definition*: Users who publish at least 1 article per month.

**Why**: This metric captures the core value of TaraNote – enabling people to share knowledge publicly. It moves beyond vanity metrics (views) to measure actual value creation.



## 2. Planned KPI Targets (V1.0 Launch)

These are the **hypothetical targets** we would aim for upon public release:

### Objective 1: User Growth
| Metric | Definition | V1.0 Target |
| :--- | :--- | :--- |
| **Signups** | Total registered accounts | 100 first month |
| **Activation Rate** | % of signups who create a note | > 60% |
| **CAC** | Cost Acquisition Customer | $0 (Organic only) |

### Objective 2: Engagement
| Metric | Definition | V1.0 Target |
| :--- | :--- | :--- |
| **Weekly Active Users** | Users visiting >1 time/week | > 20% of install base |
| **Publish Rate** | Articles published / Notes created | > 10% |
| **Session Duration** | Avg time spent in Editor | > 5 minutes |



## 3. Analytics Implementation Plan (Future)

**Status**: Not Started
**Target Release**: V1.0

### Strategy: Privacy-First
We plan to use **Plausible Analytics** or a self-hosted solution to respect user privacy (GDPR compliance) instead of intrusive trackers.

### Events to Track (Spec)
The following events need to be instrumented in the code:

1.  **`signup_completed`**
    *   *Trigger*: User finishes registration flow.
    *   *Goal*: Measure conversion rate.
2.  **`note_created`**
    *   *Trigger*: User clicks "New Note".
    *   *Goal*: Measure core feature usage.
3.  **`article_published`**
    *   *Trigger*: Status changed to `PUBLISHED`.
    *   *Goal*: Measure "North Star" activity.
4.  **`newsletter_subscribed`**
    *   *Trigger*: Visitor submits email form.
    *   *Goal*: Measure audience growth.



## 4. Technical Performance Metrics (Active)

Unlike user analytics, these **technical metrics** can be measured right now via developer tools and active tests.

| Metric | Current Status | Tool / Proof |
| :--- | :--- | :--- |
| **Test Coverage** | **73 Tests Passing** | `php artisan test` |
| **Backend Latency** | ~20ms (Local) | Laravel Debugbar |
| **Asset Size** | < 300kb (Gzipped) | `npm run build` |
| **Database Queries** | Optimized (N+1 free) | Eager Loading implementation |



## 5. Reporting Dashboard Idea (Mockup)

*Idea for a future Admin Dashboard Widget:*

```
+--------------------------------------------------+
|  Analytics Overview (Last 30 Days)            |
+--------------------------------------------------+
|  Active Users: [ 0 ] (No data)                   |
|  Published:    [ 0 ] (No data)                   |
|  Views:        [ 0 ] (No data)                   |
|                                                  |
|  [Connect Plausible API to see real data]        |
+--------------------------------------------------+
```



**Last Updated**: December 27, 2025
**Status**: Strategy Draft Only (No active tracking)
