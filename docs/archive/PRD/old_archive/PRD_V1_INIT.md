# Product Requirements Document - V1.0 (Initial Concept)

**Date:** September 23, 2025  
**Version:** 1.0 (MVP)  
**Status:** Archived (See V6)

## 1. Project Overview
TaraNote is a simple, web-based note-taking application designed to help the developer learn the Laravel 12 ecosystem. The goal is to create a functional CRUD application with user authentication.

## 2. Goals
-   Build a functional secure login system.
-   Allow users to creation, read, update, and delete text notes.
-   Implement a basic folder structure ("Notebooks").

## 3. Core Features
-   **Authentication:** Register, Login, Logout (Laravel Breeze).
-   **Notes:**
    -   Title and Body fields.
    -   Basic WYSIWYG editor (Quill.js).
-   **Notebooks:** Simple categorization for notes.

## 4. Technical Constraints
-   **Backend:** Laravel 12
-   **Frontend:** Vue 3 + Inertia
-   **Database:** MySQL

## 5. Roadmap (Sep - Oct)
-   Week 1: Setup & Auth
-   Week 2: CRUD Operations
-   Week 3: Notebook logic
