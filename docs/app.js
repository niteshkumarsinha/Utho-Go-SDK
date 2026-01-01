// Navigation
const navItems = document.querySelectorAll('.nav-item');
const sections = document.querySelectorAll('section');
const mobileToggle = document.getElementById('mobile-toggle');
const sidebar = document.querySelector('.sidebar');

function showSection(sectionId) {
    // Update Nav
    navItems.forEach(item => {
        item.classList.toggle('active', item.dataset.section === sectionId);
    });

    // Update Section
    sections.forEach(section => {
        section.classList.toggle('active', section.id === sectionId);
    });

    // Scroll to top
    window.scrollTo({ top: 0, behavior: 'smooth' });

    // Close mobile sidebar
    if (sidebar.classList.contains('open')) {
        sidebar.classList.remove('open');
    }
}

navItems.forEach(item => {
    item.addEventListener('click', () => {
        showSection(item.dataset.section);
    });
});

// Mobile Toggle
mobileToggle.addEventListener('click', () => {
    sidebar.classList.toggle('open');
});

// Simple Search (mock)
const searchInput = document.getElementById('search-input');
searchInput.addEventListener('input', (e) => {
    const term = e.target.value.toLowerCase();
    if (term.length < 2) return;
    
    // In a real app, we'd search through a manifest
    console.log('Searching for:', term);
});

// Syntax Highlighting (Primitive for demo, would use Prism/Highlight.js in prod)
function highlightCode() {
    const codeBlocks = document.querySelectorAll('pre code');
    codeBlocks.forEach(block => {
        let text = block.innerHTML;
        // Simple regex-based highlighting for Go
        text = text.replace(/(\/\/.+)/g, '<span class="token comment">$1</span>');
        text = text.replace(/\b(func|package|import|type|struct|var|return|if|err|range)\b/g, '<span class="token keyword">$1</span>');
        text = text.replace(/("[^"]*")/g, '<span class="token string">$1</span>');
        block.innerHTML = text;
    });
}

document.addEventListener('DOMContentLoaded', highlightCode);
window.showSection = showSection;
