// Custom JS for Ponto Eletrônico

document.addEventListener('DOMContentLoaded', () => {
    // Handle HTMX errors
    document.body.addEventListener('htmx:responseError', (event) => {
        const status = event.detail.xhr.status;
        if (status === 401) {
            window.location.href = '/login';
        } else if (status === 403) {
            showToast('error', 'Acesso negado. Você não tem permissão para realizar esta ação.');
        } else {
            showToast('error', 'Ocorreu um erro na requisição. Tente novamente.');
        }
    });

    // Handle HTMX beforeSwap to handle errors with custom templates
    document.body.addEventListener('htmx:beforeSwap', (event) => {
        if (event.detail.xhr.status === 422) {
            // Allow 422 status to swap content (validation errors)
            event.detail.shouldSwap = true;
            event.detail.isError = false;
        } else if (event.detail.xhr.status >= 400) {
            // Log other errors
            console.error('HTMX error:', event.detail.xhr.status, event.detail.xhr.responseText);
        }
    });

    // Handle hx-boosted links to close sidebar on mobile
    document.body.addEventListener('htmx:afterOnLoad', () => {
        // If sidebar is open (Alpine.js state), we could close it here
        // But since we use Alpine.js, it's better to handle it there if needed
    });

    // Global HTMX headers (e.g. CSRF token if needed)
    document.body.addEventListener('htmx:configRequest', (event) => {
        // Example: event.detail.headers['X-CSRF-Token'] = '...';
    });
});

// Toast notification helper
function showToast(type, message) {
    const container = document.getElementById('toast-container');
    if (!container) return;

    // This is a simple implementation. In a real app, you might want to 
    // fetch the toast template from the server or use a client-side template.
    const toast = document.createElement('div');
    toast.className = `p-4 rounded-md shadow-lg border-l-4 flex items-center justify-between mb-2 fade-in ${
        type === 'success' ? 'bg-white border-green-500 text-green-800' : 
        type === 'error' ? 'bg-white border-red-500 text-red-800' : 
        'bg-white border-blue-500 text-blue-800'
    }`;
    
    toast.innerHTML = `
        <div class="flex items-center">
            <span class="font-medium">${message}</span>
        </div>
        <button onclick="this.parentElement.remove()" class="ml-4 text-gray-400 hover:text-gray-600 focus:outline-none">
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="6 18L18 6M6 6l12 12"></path></svg>
        </button>
    `;

    container.appendChild(toast);
    setTimeout(() => {
        if (toast.parentElement) toast.remove();
    }, 5000);
}

// Preview image before upload
function previewImage(input, previewId) {
    if (input.files && input.files[0]) {
        const file = input.files[0];
        
        // Basic validation
        if (file.size > 2 * 1024 * 1024) { // 2MB
            showToast('error', 'O arquivo é muito grande. O limite é 2MB.');
            input.value = '';
            return;
        }

        const allowedTypes = ['image/jpeg', 'image/png', 'image/webp'];
        if (!allowedTypes.includes(file.type)) {
            showToast('error', 'Formato de arquivo não suportado. Use JPG, PNG ou WebP.');
            input.value = '';
            return;
        }

        const reader = new FileReader();
        reader.onload = (e) => {
            const preview = document.getElementById(previewId);
            if (preview) {
                preview.src = e.target.result;
                preview.classList.remove('hidden');
            }
        };
        reader.readAsDataURL(input.files[0]);
    }
}

// Global function to handle image load errors
function handleImageError(img) {
    // Prevent infinite loop if the default image also fails
    if (img.getAttribute('data-tried-default')) return;
    
    img.setAttribute('data-tried-default', 'true');
    // Default SVG avatar
    img.src = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke='%239CA3AF'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z'%3E%3C/path%3E%3C/svg%3E";
    img.classList.add('bg-gray-100'); // Add background to the SVG
}
