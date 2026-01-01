# Output Guide

This document defines how to structure your audit results and sprint plans.

---

## Sprint Plan Format

Create Sprint Plan Documents in `docs/sprint/sprintN.md` using this format:

```markdown
# Sprint [N]: [Sprint Name]

**Objective:** [What this sprint accomplishes]

**Target Files:**
- [ ] `path/to/file1.md`
- [ ] `path/to/file2.md`

**Verification Criteria:**
- Code alignment: [What code files were checked]
- Consistency: [What cross-file checks were done]

**Deliverables:**
- [ ] File 1 refactored
- [ ] File 2 refactored
```

---

## Phase 1: Analysis & Planning Output

1. List verified code paths (e.g., "I checked `TaskController`...").
2. List discrepancies found (e.g., "Docs said JSON, Code returns Inertia").
3. List consistency issues (e.g., "Email in README is X, but in USER_GUIDE is Y").
4. Propose the **Sprint Plan** (Sprint 1 to Sprint N) using the sprint format above.

**⛔ CRITICAL INSTRUCTION:**
**STOP** after proposing the Sprint Plan. Do **NOT** edit any files yet. Wait for the user to approve the plan.

---

## Phase 2: After User Approval

### If User Says "Proceed" or "OK":
*   Execute Sprint 1 only.
*   After completing Sprint 1, STOP again and report completion.
*   Wait for user to say "Continue" before moving to Sprint 2.

### If User Says "Revise the plan" or provides feedback:
*   Update the Sprint Plan based on feedback.
*   Present the revised plan.
*   Wait for approval again.

### If User Says "Skip to Sprint X":
*   Jump directly to that sprint and execute it.

---

## For Antigravity/Cursor

**Audit Results File:** `docs/sprint/.audit_results.md`

Format:
```markdown
# Documentation Audit - [Date]

## Files Found (Total: XX)
**Root:**
- README.md
- GITHUB_UPLOAD.md

**Documentation:**
- docs/ARCHITECTURE.md
- docs/API_REFERENCE.md
...

**Sprint Folder:**
- docs/sprint/prompt_ref.md
- docs/sprint/readme_template.md

## Discrepancies
1. **ARCHITECTURE.md** - PHP version says 8.3, composer.json says 8.2+
2. **readme_template.md** - Still shows Laravel 10, should be Laravel 12

## Consistency Check
- ✅ Email: ajarsinau@gmail.com (consistent)
- ✅ LinkedIn: linkedin.com/in/twantoro (consistent)
- ✅ GitHub: github.com/tarakreasi (consistent)

## Proposed Sprint Plan
[Insert sprint plan here using format above]
```
