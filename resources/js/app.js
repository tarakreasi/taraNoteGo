import '../css/app.css';
import 'remixicon/fonts/remixicon.css';
import './bootstrap';

import { createInertiaApp } from '@inertiajs/vue3';
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers';
import { createApp, h } from 'vue';

const appName = import.meta.env.VITE_APP_NAME || 'TaraNote';

// Mock Ziggy route() function for Go backend
const routeMap = {
    'login': '/login',
    'logout': '/logout',
    'register': '/register',
    'dashboard': '/dashboard',
    'taranote': '/taranote',
    'profile.edit': '/profile',
    'profile.update': '/api/v1/admin/profile',
    'profile.destroy': '/api/v1/admin/profile',
    'password.update': '/api/v1/admin/password',
    'password.request': '/forgot-password',
    'password.email': '/forgot-password',
    'password.store': '/reset-password',
    'password.confirm': '/confirm-password',
    'verification.send': '/email/verification-notification',
    'articles.show': '/articles',
    'docs.view': '/docs',
};

window.route = (name, params) => {
    let url = routeMap[name] || `/${name.replace(/\./g, '/')}`;

    // Handle params (simple slug replacement)
    if (params && typeof params === 'string') {
        url = `${url}/${params}`;
    } else if (params && typeof params === 'object') {
        // Handle object params like { path: 'foo', category: 'bar' }
        Object.values(params).forEach(val => {
            if (val) url = `${url}/${val}`;
        });
    }

    return url;
};

createInertiaApp({
    title: (title) => `${title} - ${appName}`,
    resolve: (name) =>
        resolvePageComponent(
            `./Pages/${name}.vue`,
            import.meta.glob('./Pages/**/*.vue'),
        ),
    setup({ el, App, props, plugin }) {
        return createApp({ render: () => h(App, props) })
            .use(plugin)
            .mount(el);
    },
    progress: {
        color: '#4B5563',
    },
});

