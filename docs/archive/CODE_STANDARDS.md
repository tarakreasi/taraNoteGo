# Code Standards & Best Practices

Coding conventions and best practices for the TaraNote project.

## PHP/Laravel Standards

### PSR-12 Compliance

Follow [PSR-12](https://www.php-fig.org/psr/psr-12/) coding style:

```php
<?php

namespace App\Http\Controllers;

class ExampleController extends Controller
{
    public function index(): JsonResponse
    {
        $data = Model::query()
            ->where('status', 'active')
            ->get();

        return response()->json($data);
    }
}
```

Use Laravel Pint to auto-format:
```bash
./vendor/bin/pint
```

### Naming Conventions

**Classes:** PascalCase
```php
class AdminNoteController {}
class NotePolicy {}
```

**Methods:** camelCase
```php
public function storeNotebook() {}
public function getPublishedNotes() {}
```

**Variables:** camelCase
```php
$userName = 'John';
$noteCount = 10;
```

**Database:** snake_case
```php
'notebook_id', 'published_at', 'is_featured'
```

**Routes:** kebab-case
```php
Route::get('/admin/note-settings');
```

### Model Best Practices

**Use explicit $fillable (not $guarded):**
```php
protected $fillable = [
    'user_id',
    'title',
    'content',
    'status',
];
```

**Add query scopes for common filters:**
```php
public function scopePublished($query)
{
    return $query->where('status', 'PUBLISHED')
                 ->whereNotNull('published_at');
}
```

**Use constants for enums:**
```php
public const STATUS_DRAFT = 'DRAFT';
public const STATUS_PUBLISHED = 'PUBLISHED';
public const STATUS_ARCHIVED = 'ARCHIVED';
```

### Controller Best Practices

**Single Responsibility:**
```php
// ✅ Good - specific purpose
class NoteController {}
class NotebookController {}

// ❌ Avoid - too broad
class DashboardController {} // handles everything
```

**Use Form Requests for validation:**
```php
public function update(NoteUpdateRequest $request, Note $note)
{
    $note->update($request->validated());
}
```

**Use Authorization:**
```php
$this->authorize('update', $note);
```

### Service Layer Pattern

For complex business logic, extract to services:

```php
// app/Services/NotePublisher.php
class NotePublisher
{
    public function publish(Note $note): Note
    {
        $note->update([
            'status' => Note::STATUS_PUBLISHED,
            'published_at' => now(),
        ]);

        // Additional logic (notifications, indexing, etc.)

        return $note;
    }
}

// Controller
public function publish(Note $note, NotePublisher $publisher)
{
    $publisher->publish($note);
}
```



## Vue.js/TypeScript Standards

### Component Structure

```vue
<script setup lang="ts">
// Imports
import { ref, computed } from 'vue';
import { router } from '@inertiajs/vue3';

// Props
interface Props {
    note: Note;
    editable?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
    editable: false,
});

// State
const isEditing = ref(false);

// Computed
const formattedDate = computed(() => {
    return new Date(props.note.created_at).toLocaleDateString();
});

// Methods
const saveNote = () => {
    router.put(`/notes/${props.note.id}`, formData);
};
</script>

<template>
    <div class="note-card">
        <h2>{{ note.title }}</h2>
        <p>{{ formattedDate }}</p>
    </div>
</template>

<style scoped>
.note-card {
    padding: 1rem;
}
</style>
```

### Component Naming

**PascalCase for components:**
```vue
<NoteCard />
<ArticleList />
<SettingsPanel />
```

### Props Validation

Always define prop types:
```typescript
interface Props {
    title: string;
    count?: number;
    user: User;
}
```

### Emit Events

Use TypeScript for emit definitions:
```typescript
const emit = defineEmits<{
    save: [note: Note];
    delete: [id: number];
}>();

emit('save', note);
```



## Git Workflow

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add cover image upload for notes
fix: resolve published_at not updating on status change
docs: update API reference for note endpoints
refactor: extract slug generation to helper
test: add tests for note query scopes
chore: update dependencies
```

### Branch Naming

```
feature/cover-image-upload
fix/published-at-bug
docs/api-reference-update
refactor/note-model-scopes
```



## File Organization

```
app/
├── Http/
│   ├── Controllers/
│   │   ├── Api/              # API controllers
│   │   ├── Auth/             # Authentication
│   │   └── *.php             # Web controllers
│   ├── Middleware/
│   ├── Requests/             # Form validation
│   └── Resources/            # API resources
├── Models/                   # Eloquent models
├── Policies/                 # Authorization
├── Services/                 # Business logic
└── Providers/

resources/
├── js/
│   ├── Components/           # Reusable Vue components
│   ├── Layouts/              # Layout templates
│   ├── Pages/                # Inertia pages
│   ├── composables/          # Vue composables
│   └── types/                # TypeScript types
└── css/

tests/
├── Feature/                  # Integration tests
└── Unit/                     # Unit tests
```



## Testing Standards

### Test Naming

```php
// Feature tests
public function test_user_can_create_draft_note(): void {}
public function test_guest_cannot_access_admin_api(): void {}

// Unit tests
public function test_it_belongs_to_a_notebook(): void {}
public function test_excerpt_is_auto_generated(): void {}
```

### AAA Pattern

```php
public function test_example(): void
{
    // Arrange
    $user = User::factory()->create();
    $note = Note::factory()->create(['user_id' => $user->id]);

    // Act
    $response = $this->actingAs($user)->deleteJson("/api/notes/{$note->id}");

    // Assert
    $response->assertOk();
    $this->assertDatabaseMissing('notes', ['id' => $note->id]);
}
```



## Documentation Standards

### Inline Comments

```php
// ✅ Good - explains WHY
// Don't auto-update slug on title change to preserve SEO
if (!$request->filled('slug')) {
    $validated['slug'] = $existingSlug;
}

// ❌ Avoid - explains WHAT (code is self-explanatory)
// Set the slug
$validated['slug'] = $slug;
```

### PHPDoc

```php
/**
 * Publish a draft note and set published_at timestamp.
 *
 * @param Note $note The note to publish
 * @return Note The published note
 * @throws \Illuminate\Auth\Access\AuthorizationException
 */
public function publish(Note $note): Note
{
    $this->authorize('update', $note);
    
    $note->update([
        'status' => Note::STATUS_PUBLISHED,
        'published_at' => now(),
    ]);

    return $note;
}
```



## Security Standards

### Input Validation

**Always validate user input:**
```php
$validated = $request->validate([
    'title' => 'required|string|max:255',
    'content' => 'nullable|string',
    'status' => 'in:DRAFT,PUBLISHED,ARCHIVED',
]);
```

### Use Authorization

```php
// Check ownership before operations
$this->authorize('update', $note);
```

### Mass Assignment Protection

```php
// ✅ Use fillable
protected $fillable = ['title', 'content'];

// ❌ Avoid guarded
protected $guarded = ['id'];
```

### File Upload Security

```php
'cover_image' => 'required|image|mimes:jpeg,png,webp,gif|max:2048'
```



## Performance Standards

### Eager Loading

```php
// ✅ Prevent N+1 queries
$notes = Note::with('notebook', 'user')->get();

// ❌ Causes N+1
$notes = Note::all();
foreach ($notes as $note) {
    echo $note->notebook->name; // Query for each iteration
}
```

### Query Optimization

```php
// ✅ Use specific columns
Note::select('id', 'title', 'slug')->get();

// ❌ Select all columns unnecessarily
Note::all();
```

### Caching

```php
use Illuminate\Support\Facades\Cache;

$notes = Cache::remember('featured-notes', 3600, function () {
    return Note::featured()->published()->get();
});
```



## Code Review Checklist

- [ ] PSR-12 compliant (run `./vendor/bin/pint`)
- [ ] All inputs validated
- [ ] Authorization checks in place
- [ ] Tests added/updated
- [ ] No N+1 queries
- [ ] Explicit `$fillable` on models
- [ ] Proper error handling
- [ ] Documentation updated
- [ ] No sensitive data in logs
- [ ] Environment variables for secrets



**Last Updated:** December 27, 2025 | Beta 0.9.80
