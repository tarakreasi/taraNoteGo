# Security Policy

**TaraNote Security Guidelines**

> **Portfolio Project Disclaimer**: TaraNote is a **personal portfolio project** maintained on a **best-effort basis**. While security best practices have been followed, this project has **NOT undergone professional security audits** and is **NOT suitable for enterprise production use**. There are **NO WARRANTIES** or **SLAs** for security responses. Use at your own risk.

We take the security of TaraNote seriously **within the constraints of a hobby project**. This document outlines our security practices and how to report vulnerabilities.



## Supported Versions

We provide security updates for the following versions:

| Version | Supported |
| ------- | --------- |
| Beta 0.9.x | Yes |
| < 0.9.0 | No |



## Reporting a Vulnerability

**Please do NOT report security vulnerabilities through public GitHub issues.**

### Responsible Disclosure

If you discover a security vulnerability, please email us at:

**ajarsinau@gmail.com**

Include:
1. **Description** of the vulnerability
2. **Steps to reproduce** (with example code if applicable)
3. **Potential impact** (what an attacker could do)
4. **Affected versions**
5. **Your contact information** (for follow-up)

### Best Effort Response Targets

> **DISCLAIMER**: The following timelines are **aspirational targets** for this personal portfolio project. They are **NOT guaranteed**. Response times may vary significantly depending on the maintainer's availability.

- **Acknowledgment**: Target: Within 48 hours (Best Effort)
- **Initial Assessment**: Target: Within 5 business days
- **Regular Updates**: As available
- **Fix Timeline**: No guaranteed SLA

### Coordinated Disclosure

We follow a **90-day coordinated disclosure** policy:
1. You report the vulnerability privately
2. We verify and develop a fix
3. We release a security patch
4. After 90 days (or when patched), you may publish your findings

### Recognition

Security researchers who responsibly disclose vulnerabilities will be:
- Credited in our Security Hall of Fame (with your permission)
- Mentioned in release notes (if you wish)
- Sent a thank-you email and TaraNote swag



## Security Measures

### Application Security

#### 1. Authentication
- **Laravel Breeze**: Industry-standard authentication
- **Password hashing**: Bcrypt with work factor 12
- **Session management**: Secure, HTTP-only cookies
- **CSRF Protection**: Enabled on all forms

#### 2. Authorization
- **Policy-based**: Laravel Policies for resource access
- **User Isolation**: Users can only access their own notes
- **Role-based** (future): Admin, Editor, Viewer roles

#### 3. Input Validation
- **Server-side validation**: All inputs validated via FormRequests
- **Content sanitization**: DOMPurify for WYSIWYG content
- **File upload restrictions**: 
  - Max size: 5MB
  - Allowed types: `jpg`, `png`, `gif`, `webp`
  - Stored outside web root

#### 4. SQL Injection Prevention
- **Eloquent ORM**: Parameterized queries by default
- **Query Builder**: Bindings for raw queries
- **No dynamic queries**: All queries are statically typed

#### 5. XSS Protection
- **Blade escaping**: `{{ }}` auto-escapes output
- **DOMPurify**: Sanitizes rich text content
- **Content Security Policy** (CSP): Configured headers (future)

#### 6. HTTPS Enforcement
- **Production**: Force HTTPS via middleware
- **HSTS Headers**: `Strict-Transport-Security` enabled
- **Secure Cookies**: Set in production environment



### Infrastructure Security

#### 1. Server Hardening
- **Firewall**: Only ports 80, 443, and 22 (SSH) open
- **SSH Keys**: Password authentication disabled
- **Fail2Ban**: Enabled for brute-force protection
- **Auto-updates**: Security patches applied weekly

#### 2. Database Security
- **Restricted Access**: Database not publicly accessible
- **Strong Passwords**: 20+ character passwords
- **Backups**: Daily encrypted backups, stored off-site
- **Encryption at Rest**: Database files encrypted (VPS provider)

#### 3. File Storage
- **Private Storage**: Uploads stored in `storage/` (not publicly accessible)
- **Symlink**: Only `storage/app/public` linked to `public/storage`
- **File Permissions**: `www-data:www-data` with `755` directories, `644` files



### Code Security

#### 1. Dependency Management
- **Composer**: Regular updates via `composer update`
- **NPM**: `npm audit` run before deployments
- **Automated Scans**: GitHub Dependabot enabled

#### 2. Secret Management
- **`.env` file**: Never committed to git
- **Environment Variables**: Secrets stored as env vars
- **API Keys**: Rotated quarterly

#### 3. Code Quality
- **Static Analysis**: Laravel Pint for code standards
- **No Debug Code**: `APP_DEBUG=false` in production
- **Error Logging**: Errors logged, not displayed



## Security Best Practices for Users

### For Administrators

1. **Use Strong Passwords**
   - Minimum 12 characters
   - Mix of uppercase, lowercase, numbers, symbols
   - Use a password manager

2. **Enable Two-Factor Authentication** (when available)
   - Authenticator apps (Google Authenticator, Authy)
   - Backup codes stored securely

3. **Regular Backups**
   - Export notes monthly
   - Download database backups

4. **Update Regularly**
   - Check for updates weekly
   - Apply security patches immediately

5. **Monitor Access Logs**
   - Review `storage/logs/laravel.log` for suspicious activity
   - Check failed login attempts

### For Content Creators

1. **Don't Share Credentials**
   - Each user should have their own account
   - Never share passwords via email or chat

2. **Be Careful with External Links**
   - Verify URLs before linking
   - Avoid URL shorteners (they can be hijacked)

3. **Review Published Content**
   - Check for accidentally published private information
   - Use the Draft feature for work-in-progress

4. **Sanitize Embedded Content**
   - The editor auto-sanitizes, but review pasted HTML
   - Avoid embedding untrusted iframes



## Vulnerability Disclosure History

| CVE ID | Date | Severity | Description | Status |
| ------ | ---- | -------- | ----------- | ------ |
| INTERNAL | 2025-12-22 | Critical | Broken Access Control (Admin Settings) | Fixed |
| INTERNAL | 2025-12-22 | High | Stored XSS in Presentation Mode | Fixed |
| INTERNAL | 2025-12-22 | High | Timing Attack on Access Token | Fixed |



## Security Checklist for Deployment

Before deploying to production:

- [ ] `APP_ENV=production` in `.env`
- [ ] `APP_DEBUG=false` in `.env`
- [ ] Strong `APP_KEY` generated (`php artisan key:generate`)
- [ ] Database credentials are strong and unique
- [ ] HTTPS is enforced
- [ ] File permissions are correct (`755` for directories, `644` for files)
- [ ] `.env` file is not in git
- [ ] `storage/` and `bootstrap/cache/` are writable
- [ ] Auto-updates are configured (or manual update schedule)
- [ ] Backups are automated
- [ ] Firewall is configured
- [ ] Fail2Ban is installed



## Compliance

- **GDPR**: See `docs/PRIVACY_POLICY.md`
- **CCPA**: User data can be exported/deleted on request
- **Data Retention**: Subscriber emails stored indefinitely (can be deleted)



## Contact

- **Security Issues**: ajarsinau@gmail.com
- **General Support**: ajarsinau@gmail.com
- **Website**: https://tarakreasi.com



**Last Updated:** December 27, 2025 (Beta 0.9.80 (UI Design V6.5))
