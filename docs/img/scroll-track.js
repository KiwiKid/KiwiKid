class ScrollTracker extends HTMLElement {
constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.appendChild(
    document.getElementById('scroll-tracker-template').content.cloneNode(true)
    );

    this.observer = null;
    this.debounceTimer = null;
}

connectedCallback() {
    this.initIntersectionObserver();
}

disconnectedCallback() {
    if (this.observer) {
    this.observer.disconnect();
    }
    clearTimeout(this.debounceTimer);
}

initIntersectionObserver() {
    const options = {
    root: null, // observing changes to visibility in the viewport
    rootMargin: '0px',
    threshold: 0.5 // trigger when 50% of the target is visible
    };

    this.observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
        this.debounceUpdateUrl(entry.target.id);
        }
    });
    }, options);

    this.querySelectorAll('a').forEach(anchor => {
    if (anchor.href && anchor.id) {
        this.observer.observe(anchor);
    }
    });
}

debounceUpdateUrl(id) {
    clearTimeout(this.debounceTimer);
    this.debounceTimer = setTimeout(() => {
    history.pushState({}, '', `#${id}`);
    }, 300); // Debounce time of 300ms
}
}

customElements.define('scroll-tracker', ScrollTracker);