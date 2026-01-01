# Feature Flags

**TaraNote Feature Flag System**

Feature flags allow us to control feature rollout, A/B testing, and gradual releases without deploying new code.



## 1. What are Feature Flags?

Feature flags (or feature toggles) are configuration settings that enable or disable features at runtime.

**Benefits**:
- **Gradual Rollout**: Enable features for 10% of users, then 50%, then 100%
- **A/B Testing**: Test variant A vs variant B
- **Kill Switch**: Quickly disable a buggy feature
- **Beta Access**: Give early access to select users



## 2. Current Feature Flags

### 2.1 Newsletter Subscription
**Flag Name**: `newsletter_enabled`  
**Default**: `true`  
**Description**: Controls visibility of newsletter subscription forms

**Usage**:
```vue
<div v-if="$page.props.features.newsletter_enabled">
    <!-- Newsletter form -->
</div>
```

**Rollout Plan**:
- **Beta 0.9.80 (UI Design V6.5)**: Enabled for all (100%)
- **Future**: May disable if email service is down



### 2.2 Comments System (Future)
**Flag Name**: `comments_enabled`  
**Default**: `false`  
**Description**: Enables comments on published articles

**Rollout Plan**:
- **V6.1**: Beta users only (10%)
- **V6.2**: All users (100%)



### 2.3 Analytics Dashboard (Future)
**Flag Name**: `analytics_dashboard_enabled`  
**Default**: `false`  
**Description**: Shows analytics tab in Dashboard

**Rollout Plan**:
- **V6.2**: Premium users only
- **V6.3**: All users



### 2.4 Dark Mode (Existing)
**Flag Name**: `dark_mode_enabled`  
**Default**: `true`  
**Description**: Allows users to toggle dark mode

**Status**: Fully rolled out



### 2.5 Search (Client-Side)
**Flag Name**: `search_client_side`  
**Default**: `true`  
**Description**: Uses client-side search instead of server-side

**Rollout Plan**:
- **UI Design V6**: Client-side (current)
- **V6.1**: Server-side (Laravel Scout) - will flip flag to `false`



## 3. Implementation

### 3.1 Database Storage

**Table**: `feature_flags`

```php
Schema::create('feature_flags', function (Blueprint $table) {
    $table->id();
    $table->string('key')->unique();
    $table->boolean('enabled')->default(false);
    $table->integer('rollout_percentage')->default(0); // 0-100
    $table->json('allowed_users')->nullable(); // User IDs for beta
    $table->timestamps();
});
```

### 3.2 Model

```php
// app/Models/FeatureFlag.php
class FeatureFlag extends Model
{
    protected $fillable = ['key', 'enabled', 'rollout_percentage', 'allowed_users'];
    protected $casts = ['allowed_users' => 'array'];

    public static function isEnabled(string $key, ?User $user = null): bool
    {
        $flag = static::where('key', $key)->first();
        
        if (!$flag) {
            return false; // Default: disabled
        }

        // Fully enabled
        if ($flag->enabled && $flag->rollout_percentage === 100) {
            return true;
        }

        // Beta users
        if ($user && in_array($user->id, $flag->allowed_users ?? [])) {
            return true;
        }

        // Percentage rollout (deterministic based on user ID)
        if ($user && $flag->rollout_percentage > 0) {
            return ($user->id % 100) < $flag->rollout_percentage;
        }

        return false;
    }
}
```

### 3.3 Middleware

```php
// app/Http/Middleware/InjectFeatureFlags.php
public function handle(Request $request, Closure $next)
{
    $user = $request->user();
    
    $flags = [
        'newsletter_enabled' => FeatureFlag::isEnabled('newsletter_enabled', $user),
        'comments_enabled' => FeatureFlag::isEnabled('comments_enabled', $user),
        'analytics_dashboard_enabled' => FeatureFlag::isEnabled('analytics_dashboard_enabled', $user),
    ];

    Inertia::share('features', $flags);

    return $next($request);
}
```

**Register in `app/Http/Kernel.php`:**
```php
'web' => [
    // ...
    \App\Http\Middleware\InjectFeatureFlags::class,
],
```

### 3.4 Seeder

```php
// database/seeders/FeatureFlagSeeder.php
public function run()
{
    FeatureFlag::create([
        'key' => 'newsletter_enabled',
        'enabled' => true,
        'rollout_percentage' => 100,
    ]);

    FeatureFlag::create([
        'key' => 'comments_enabled',
        'enabled' => false,
        'rollout_percentage' => 0,
    ]);
}
```



## 4. Usage Examples

### 4.1 Frontend (Vue)

```vue
<template>
    <!-- Show feature only if enabled -->
    <div v-if="$page.props.features.comments_enabled">
        <CommentSection :article="article" />
    </div>
</template>
```

### 4.2 Backend (Controller)

```php
use App\Models\FeatureFlag;

public function index()
{
    $user = auth()->user();

    if (FeatureFlag::isEnabled('analytics_dashboard_enabled', $user)) {
        return Inertia::render('Dashboard/Analytics');
    }

    return Inertia::render('Dashboard/Standard');
}
```

### 4.3 Blade (Email Templates)

```blade
@if(FeatureFlag::isEnabled('newsletter_footer_enabled'))
    <p>Subscribe to our newsletter!</p>
@endif
```



## 5. Rollout Strategies

### 5.1 Beta Users

Enable for specific users:

```php
FeatureFlag::updateOrCreate(['key' => 'comments_enabled'], [
    'enabled' => true,
    'rollout_percentage' => 0,
    'allowed_users' => [1, 5, 12], // User IDs
]);
```

### 5.2 Percentage Rollout

Enable for 25% of users:

```php
FeatureFlag::updateOrCreate(['key' => 'comments_enabled'], [
    'enabled' => true,
    'rollout_percentage' => 25,
]);
```

**How it works**: User ID `% 100` determines if they're in the rollout.
- User ID 1 -> 1 < 25 [Enabled]
- User ID 42 -> 42 >= 25 [Disabled]
- User ID 99 -> 99 >= 25 [Disabled]

### 5.3 Full Rollout

Enable for everyone:

```php
FeatureFlag::updateOrCreate(['key' => 'comments_enabled'], [
    'enabled' => true,
    'rollout_percentage' => 100,
]);
```

### 5.4 Kill Switch

Disable immediately:

```php
FeatureFlag::where('key', 'comments_enabled')->update(['enabled' => false]);
```



## 6. A/B Testing

### Example: Test Two Newsletter CTAs

**Flag Setup**:
```php
FeatureFlag::create([
    'key' => 'newsletter_cta_variant',
    'enabled' => true,
    'rollout_percentage' => 50, // 50% see Variant A, 50% see Variant B
]);
```

**Frontend**:
```vue
<button v-if="showVariantA">
    Subscribe
</button>
<button v-else>
    Get Weekly Seeds
</button>

<script>
export default {
    computed: {
        showVariantA() {
            // Deterministic: same user always sees same variant
            return this.$page.props.auth.user.id % 2 === 0;
        }
    }
}
</script>
```

**Track Results**: See `docs/METRICS.md` for conversion tracking.



## 7. Admin Panel (Future)

**Feature Flags Manager UI**:
- List all flags with current status
- Toggle enabled/disabled
- Adjust rollout percentage slider
- Add/remove beta users

**Mockup**:
```
┌─────────────────────────────────────────────┐
│ Feature Flags                               │
├─────────────────────────────────────────────┤
│ newsletter_enabled         [x] 100%         │
│ comments_enabled           [ ] 0%    ^ 25%  │
│ analytics_dashboard_enabled[ ] Beta Users   │
└─────────────────────────────────────────────┘
```



## 8. Testing

### Test Feature Flag Logic

```php
// tests/Unit/FeatureFlagTest.php
public function test_feature_enabled_for_all()
{
    FeatureFlag::create([
        'key' => 'test_feature',
        'enabled' => true,
        'rollout_percentage' => 100,
    ]);

    $user = User::factory()->create();
    $this->assertTrue(FeatureFlag::isEnabled('test_feature', $user));
}

public function test_feature_disabled()
{
    FeatureFlag::create([
        'key' => 'test_feature',
        'enabled' => false,
    ]);

    $user = User::factory()->create();
    $this->assertFalse(FeatureFlag::isEnabled('test_feature', $user));
}

public function test_beta_user_access()
{
    $betaUser = User::factory()->create(['id' => 5]);
    $normalUser = User::factory()->create(['id' => 10]);

    FeatureFlag::create([
        'key' => 'beta_feature',
        'enabled' => true,
        'allowed_users' => [5],
    ]);

    $this->assertTrue(FeatureFlag::isEnabled('beta_feature', $betaUser));
    $this->assertFalse(FeatureFlag::isEnabled('beta_feature', $normalUser));
}
```



## 9. Best Practices

### 9.1 Naming Convention
- Use `snake_case`
- Be descriptive: `comments_enabled`, not `feature_1`
- Include context: `newsletter_footer_enabled` vs `newsletter_enabled`

### 9.2 Default to Disabled
New features should default to `enabled: false` until verified.

### 9.3 Clean Up Old Flags
Once a feature is fully rolled out (100% for 3+ months), remove the flag:
1. Replace conditional checks with permanent code
2. Delete flag from database
3. Remove flag from middleware

### 9.4 Document Rollout Plan
Every flag should have a clear rollout plan (see Section 2).

### 9.5 Monitor Feature Usage
Track how many users see/use a feature:
```php
if (FeatureFlag::isEnabled('comments_enabled', $user)) {
    Log::info('User accessed comments', ['user_id' => $user->id]);
}
```



## 10. Roadmap

### Q1 2026
- [ ] Implement `FeatureFlag` model and migration
- [ ] Create Admin UI for managing flags
- [ ] Add `comments_enabled` flag

### Q2 2026
- [ ] Implement percentage rollout for new search feature
- [ ] A/B test newsletter CTA variants

### Q3 2026
- [ ] Create analytics dashboard for flag usage
- [ ] Add user-level flag overrides



## 11. Alternative: Laravel Pennant

Laravel 12+ includes **Laravel Pennant** for feature flags.

**Migration Path**:
```php
use Laravel\Pennant\Feature;

Feature::define('comments', fn (User $user) => $user->isAdmin());

// In code
if (Feature::active('comments')) {
    // Show comments
}
```

**Recommendation**: Adopt Pennant in V7.0 for native support.



## Contact

**Feature Flag Owner**: ajarsinau@gmail.com



**Last Updated:** December 27, 2025 (Beta 0.9.80 (UI Design V6.5))
