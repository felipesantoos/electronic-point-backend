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

    // HTMX Progress Bar logic
    const progressBar = document.querySelector('#htmx-progress div');
    let progressInterval;

    document.body.addEventListener('htmx:configRequest', () => {
        if (!progressBar) return;
        
        // Reset and show progress bar
        progressBar.style.width = '0%';
        progressBar.style.opacity = '1';
        
        // Animate to 90% over some time
        let width = 0;
        progressInterval = setInterval(() => {
            if (width < 90) {
                width += Math.random() * 2; // Slower progress
                progressBar.style.width = `${width}%`;
            }
        }, 100);
    });

    document.body.addEventListener('htmx:afterRequest', () => {
        if (!progressBar) return;
        
        clearInterval(progressInterval);
        
        // Complete the bar
        progressBar.style.width = '100%';
        
        // Hide after a small delay
        setTimeout(() => {
            progressBar.style.opacity = '0';
            setTimeout(() => {
                progressBar.style.width = '0%';
            }, 300);
        }, 500);
    });

    // Initialize on load or HTMX swap
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', initApp);
    } else {
        initApp();
    }

    function initApp() {
        applyInputMasks();
    }

    // Global HTMX headers
    document.body.addEventListener('htmx:configRequest', (event) => {
        // Ensure cookies are sent with every request
        event.detail.withCredentials = true;
    });

    document.body.addEventListener('htmx:afterOnLoad', () => {
        applyInputMasks();
    });

    // Custom Confirmation Modal Logic for HTMX
    document.body.addEventListener('htmx:confirm', (event) => {
        // Only intercept if there's a confirmation message (hx-confirm)
        if (!event.detail.question) return;

        // Prevent immediate execution
        event.preventDefault();

        // Dispatch an event that our Alpine.js modal will listen to
        const title = event.target.getAttribute('data-confirm-title') || 'Confirmar Ação';
        window.dispatchEvent(new CustomEvent('open-confirm-modal', {
            detail: {
                title: title,
                question: event.detail.question,
                issueRequest: event.detail.issueRequest
            }
        }));
    });
});

function applyInputMasks() {
    const cpfInputs = document.querySelectorAll('input[name="cpf"]');
    cpfInputs.forEach(input => {
        input.addEventListener('input', (e) => {
            let value = e.target.value.replace(/\D/g, '');
            if (value.length > 11) value = value.slice(0, 11);
            
            value = value.replace(/(\d{3})(\d)/, '$1.$2');
            value = value.replace(/(\d{3})(\d)/, '$1.$2');
            value = value.replace(/(\d{3})(\d{1,2})$/, '$1-$2');
            
            e.target.value = value;
        });
    });

    const phoneInputs = document.querySelectorAll('input[name="phone"]');
    phoneInputs.forEach(input => {
        input.addEventListener('input', (e) => {
            let value = e.target.value.replace(/\D/g, '');
            if (value.length > 11) value = value.slice(0, 11);
            
            if (value.length > 10) {
                value = value.replace(/^(\d{2})(\d{5})(\d{4})/, '($1) $2-$3');
            } else if (value.length > 5) {
                value = value.replace(/^(\d{2})(\d{4})(\d{0,4})/, '($1) $2-$3');
            } else if (value.length > 2) {
                value = value.replace(/^(\d{2})(\d{0,5})/, '($1) $2');
            } else if (value.length > 0) {
                value = value.replace(/^(\d*)/, '($1');
            }
            
            e.target.value = value;
        });
    });
}

// Toast notification helper
function showToast(type, message) {
    const container = document.getElementById('toast-container');
    if (!container) return;

    const toast = document.createElement('div');
    toast.className = `p-4 rounded-md shadow-lg border-l-4 flex items-center justify-between mb-2 transition-all duration-500 ease-in-out transform translate-x-full opacity-0 ${
        type === 'success' ? 'bg-white border-green-500 text-green-800' : 
        type === 'error' ? 'bg-white border-red-500 text-red-800' : 
        'bg-white border-blue-500 text-blue-800'
    }`;
    
    toast.innerHTML = `
        <div class="flex items-center">
            <span class="font-medium">${message}</span>
        </div>
        <button onclick="removeToast(this.parentElement)" class="ml-4 text-gray-400 hover:text-gray-600 focus:outline-none">
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="6 18L18 6M6 6l12 12"></path></svg>
        </button>
    `;

    container.appendChild(toast);

    // Animate in
    requestAnimationFrame(() => {
        toast.classList.remove('translate-x-full', 'opacity-0');
    });

    // Auto-hide
    const duration = type === 'error' ? 8000 : 4000;
    setTimeout(() => {
        removeToast(toast);
    }, duration);
}

function removeToast(toast) {
    if (!toast || !toast.parentElement) return;
    
    // Animate out
    toast.classList.add('translate-x-full', 'opacity-0');
    
    // Remove from DOM after animation
    setTimeout(() => {
        if (toast.parentElement) toast.remove();
    }, 500);
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
