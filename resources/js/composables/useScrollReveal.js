import { onMounted, onUnmounted } from 'vue';

/**
 * Composable for scroll-triggered reveal animations
 * Uses IntersectionObserver to animate elements when they enter viewport
 */
export function useScrollReveal(options = {}) {
    const {
        threshold = 0.1,
        rootMargin = '0px',
        once = true
    } = options;

    let observer = null;

    const setupObserver = () => {
        if (typeof window === 'undefined') return;

        const observerOptions = {
            root: null,
            rootMargin,
            threshold
        };

        observer = new IntersectionObserver((entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    entry.target.classList.add('is-visible');

                    // Unobserve if animation should only happen once
                    if (once) {
                        observer.unobserve(entry.target);
                    }
                } else if (!once) {
                    // Remove class if animation should repeat
                    entry.target.classList.remove('is-visible');
                }
            });
        }, observerOptions);

        // Observe all elements with .reveal-on-scroll class
        const elements = document.querySelectorAll('.reveal-on-scroll');
        elements.forEach(el => observer.observe(el));
    };

    const cleanup = () => {
        if (observer) {
            observer.disconnect();
            observer = null;
        }
    };

    onMounted(() => {
        setupObserver();
    });

    onUnmounted(() => {
        cleanup();
    });

    return {
        setupObserver,
        cleanup
    };
}
