// snippy.js - UI interactivity for Snippy
// (No major JS needed for static look, but this is ready for future UI features)
document.addEventListener('DOMContentLoaded', function() {
    // Highlight active nav link if needed
    const navLinks = document.querySelectorAll('nav a');
    navLinks.forEach(link => {
        if (link.href === window.location.href) {
            link.classList.add('active');
        }
    });

    // Animate main content fade-in
    const main = document.querySelector('main');
    if (main) {
        main.classList.add('fade-in');
    }

    // Animate header/logo
    const headerH1 = document.querySelector('header h1');
    if (headerH1) {
        headerH1.style.opacity = 0;
        setTimeout(() => headerH1.style.opacity = '', 300);
    }
    const logo = document.querySelector('.logo');
    if (logo) {
        logo.style.opacity = 0;
        setTimeout(() => logo.style.opacity = '', 350);
    }

    // Animate rainbow bar
    const rainbow = document.querySelector('.rainbow-bar');
    if (rainbow) {
        rainbow.style.opacity = 0;
        setTimeout(() => rainbow.style.opacity = '', 100);
    }
});
