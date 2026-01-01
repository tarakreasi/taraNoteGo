import '../css/app.css';
import 'remixicon/fonts/remixicon.css';
import './bootstrap';

import { createInertiaApp } from '@inertiajs/vue3';
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers';
import { createApp, h } from 'vue';

const appName = import.meta.env.VITE_APP_NAME || 'TaraNote';

// Mock Ziggy route() function for Go backend
const routeMap = {
    'home': '/',
    'login.view': '/login',
    'login.post': '/login',
    'logout': '/logout',
    'dashboard': '/dashboard',
    'taranote': '/taranote',
    'api.notebooks.index': '/api/v1/admin/notebooks',
    'articles.show': '/articles',
    'docs.index': '/docs',
    'docs.show': '/docs',
};

window.route = (name, params) => {
    let url = routeMap[name];

    if (!url) {
        console.warn(`Route not found: ${name}`);
        return `/${name.replace(/\./g, '/')}`;
    }

    // Handle params
    if (params) {
        if (typeof params === 'object') {
            // Replace {param} in URL if it exists (e.g. /articles/{slug})
            // For now, our simple map doesn't have placeholders, so we append or replace logic
            if (name === 'articles.show' && params.slug) {
                return `${url}/${params.slug}`;
            }
            if (name === 'docs.show' && params.path) {
                return `${url}/${params.path}`;
            }
        } else {
            // If string/number param, append it
            return `${url}/${params}`;
        }
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

