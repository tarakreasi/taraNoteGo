# Developer Handbook

How to work on this codebase (notes to future me / other devs).



## Project Structure

```
app/
├── Models/              # Database stuff
│   ├── User.php        # Users (auth)
│   ├── Note.php        # Notes/articles
│   ├── Notebook.php    # Categories
│   └── Setting.php     # Site config (key-value store)
│
├── Http/Controllers/
│   ├── ArticleController.php      # Public articles (/articles)
│   ├── WelcomeController.php      # Landing page (/)
│   ├── PortfolioController.php    # User portfolios (/portfolio/{user})
│   ├── SubscriptionController.php # Newsletter signups
│   └── Api/
│       ├── AdminNoteController.php     # Note CRUD (API)
│       └── AdminSettingController.php  # Settings (API)
│
resources/js/
├── Pages/              # Full page views (Inertia)
│   ├── Home.vue       # Landing page
│   ├── Articles.vue   # Public articles list
│   ├── Article.vue    # Single article view
│   ├── Portfolio.vue  # User portfolio
│   ├── TaraNote.vue   # Private workspace (complex!)
│   └── Dashboard/
│       └── SettingsPanel.vue  # Admin settings
│
├── Components/         # Reusable UI
│   ├── AppNavbar.vue
│   ├── Footer.vue
│   └── Editor.vue     # Quill.js wrapper
│
└── Layouts/
    └── AppLayout.vue  # Public layout (navbar + footer)
```



## Tech Stack

**Backend**:
- Laravel 12 (PHP 8.3)
- MySQL for production, SQLite for dev
- Inertia.js (connects Laravel to Vue without REST API)

**Frontend**:
- Vue 3 (Composition API)
- Tailwind CSS (utility-first)
- Vite (build tool)

**Other**:
- Quill.js (rich text editor)
- Material Symbols (icons)
- DOMPurify (sanitize HTML)



## Common Tasks

### Add a new public page

Example: `/about` page

```bash
# 1. Create controller
php artisan make:controller AboutController

# 2. Add route (routes/web.php)
Route::get('/about', [AboutController::class, 'index']);

# 3. Create Vue page (resources/js/Pages/About.vue)
<template>
  <AppLayout title="About">
    <div class="container mx-auto p-8">
      <h1>About TaraNote</h1>
    </div>
  </AppLayout>
</template>

# 4. Controller returns view
public function index() {
    return Inertia::render('About');
}
```

### Add a new setting

Settings are stored as key-value pairs in the `settings` table.

```javascript
// 1. Add field in SettingsPanel.vue
<input v-model="localSettings.new_setting" />

// 2. That's it! Controller auto-saves any key you send.
```

To display on public pages:

```php
// In controller (e.g., WelcomeController)
$settings = Setting::whereIn('key', [
    'site_title',
    'new_setting', // Add here
])->pluck('value', 'key');
```

### Debug editor issues

Editor is at `resources/js/Components/Editor.vue`.

It wraps Quill.js. Config is in the `onMounted` hook.

**Common issues**:
- Auto-save not working? Check `debounce` in parent component
- Lost formatting? Check `content_sanitized` prop
- Toolbar missing? quill config broken



## Code Style

**PHP** (Laravel Pint):
```bash
./vendor/bin/pint
```

**Vue** (ESLint):
```bash
npm run lint
```

**General rules**:
- Use camelCase for JS, snake_case for PHP
- No inline styles (use Tailwind classes)
- Components start with capital letter (AppNavbar, not appNavbar)
- Keep files under 300 lines (split if bigger)



## Testing

Run all tests:
```bash
php artisan test
```

Run specific test:
```bash
php artisan test --filter=AdminNoteTest
```

**Test structure**:
- `tests/Feature/` - Full request/response cycles
- `tests/Unit/` - Individual functions

**Coverage**: 64 tests, 100% critical path



## Database

**Migrations**:
```bash
php artisan migrate        # Run migrations
php artisan migrate:fresh  # Reset DB
php artisan db:seed        # Add fake data
```

**Schema**:
- `users` - Auth
- `notes` - Articles (has `status`, `published_at`, `slug`)
- `notebooks` - Categories
- `settings` - Site config (key/value)
- `subscribers` - Newsletter emails

See `docs/DATABASE_SCHEMA.md` for full ERD.



## Deployment

**Local dev**:
```bash
npm run dev &
php artisan serve
```

**Production**:
```bash
composer install --no-dev
npm run build
php artisan config:cache
php artisan route:cache
php artisan view:cache
```

See `docs/DEPLOYMENT.md` for VPS setup.



## Troubleshooting

**"Page won't load"**:
- Check `storage/logs/laravel.log`
- Clear cache: `php artisan view:clear`
- Restart Vite

**"Editor not saving"**:
- Check Network tab for failed API calls
- Look for 422 (validation error) or 500 (server error)

**"Tests failing"**:
- Run `php artisan migrate:fresh` to reset test DB
- Check `.env.testing` exists

**"CSS not updating"**:
- Vite cache. Hard refresh (Ctrl+Shift+R)
- Run `npm run build` to force rebuild



## Design System

**V6 "Electric Earth"**:
- Glassmorphism (frosted glass effect)
- Noise texture overlay
- Ambient blob backgrounds
- Electric blue + earthy green palette

See `docs/DESIGN_TOKENS.md` for full spec.

**Icons**: Material Symbols (Google)
```html
<span class="material-symbols-outlined">home</span>
```



## Architecture Decisions

**Why Inertia?**  
Don't want to build a separate API. Inertia lets Vue talk to Laravel controllers directly.

**Why not TypeScript?**  
Solo project. Overkill. Maybe V2.

**Why Quill over TipTap?**  
Tried TipTap first. Too complex for what we need. Quill is simpler.

**Why no state management (Vuex/Pinia)?**  
App is small enough. Props and emits work fine.



## Resources

- [Laravel Docs](https://laravel.com/docs)
- [Vue 3 Docs](https://vuejs.org)
- [Inertia Docs](https://inertiajs.com)
- [Quill API](https://quilljs.com/docs)



Last updated: Dec 27, 2025
