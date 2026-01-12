// Custom JS for Ponto Eletrônico

// We use a self-invoking function or check to prevent multiple initialization
if (typeof window.epAppInitialized === 'undefined') {
    window.epAppInitialized = true;

    // Global HTMX error handler
    document.addEventListener('htmx:responseError', (event) => {
        const xhr = event.detail.xhr;
        const status = xhr.status;
        const responseText = xhr.responseText;
        
        console.log(`[HTMX Error ${status}]`, responseText);
        
        if (status === 401) {
            window.location.href = '/login';
            return;
        } 
        
        if (status === 403) {
            showToast('error', 'Acesso negado.');
            return;
        }

        // If the server sent a custom trigger (like show-toast), HTMX will handle it via 
        // the dedicated 'show-toast' event listener below. We just return here to avoid duplicates.
        const triggerHeader = xhr.getResponseHeader('HX-Trigger');
        if (triggerHeader && triggerHeader.includes('show-toast')) {
            console.log('Custom show-toast trigger detected in header, letting dedicated listener handle it.');
            return;
        }
        
        // Use response text as message if it's short and looks like a string
        let message = 'Ocorreu um erro na requisição. Tente novamente.';
        if (responseText && responseText.trim().length > 0 && responseText.length < 600 && !responseText.includes('<html')) {
            message = responseText.trim();
        }
        
        showToast('error', message);
    });

    // Support for custom events from HX-Trigger header (for successful responses)
    document.addEventListener('show-toast', (event) => {
        console.log('[HTMX Trigger] show-toast', event.detail);
        const data = event.detail;
        if (data && data.message) {
            showToast(data.type || 'info', data.message);
        }
    });

    // HTMX beforeSwap logic
    document.addEventListener('htmx:beforeSwap', (event) => {
        const status = event.detail.xhr.status;
        if (status === 422) {
            event.detail.shouldSwap = true;
            event.detail.isError = false;
        }
    });

    // HTMX Progress Bar logic
    let progressInterval;
    document.addEventListener('htmx:configRequest', (event) => {
        event.detail.withCredentials = true;
        
        const progressBar = document.querySelector('#htmx-progress div');
        if (!progressBar) return;
        
        progressBar.style.width = '0%';
        progressBar.style.opacity = '1';
        
        let width = 0;
        clearInterval(progressInterval);
        progressInterval = setInterval(() => {
            if (width < 90) {
                width += Math.random() * 2;
                progressBar.style.width = `${width}%`;
            }
        }, 100);
    });

    document.addEventListener('htmx:afterRequest', (event) => {
        const progressBar = document.querySelector('#htmx-progress div');
        if (progressBar) {
            clearInterval(progressInterval);
            progressBar.style.width = '100%';
            setTimeout(() => {
                progressBar.style.opacity = '0';
                setTimeout(() => {
                    progressBar.style.width = '0%';
                }, 300);
            }, 500);
        }

        const xhr = event.detail.xhr;
        const requestConfig = event.detail.requestConfig;
        const method = (requestConfig && requestConfig.method ? requestConfig.method : '').toUpperCase();
        const stateChangingMethods = ['POST', 'PUT', 'DELETE', 'PATCH'];
        
        if (xhr.status >= 200 && xhr.status < 300 && stateChangingMethods.includes(method)) {
            let message = 'Operação realizada com sucesso!';
            if (method === 'DELETE') message = 'Registro removido com sucesso!';
            if (method === 'POST') message = 'Registro criado com sucesso!';
            if (method === 'PUT' || method === 'PATCH') message = 'Registro atualizado com sucesso!';
            
            const target = event.target;
            const skipToast = (target && typeof target.closest === 'function' && target.closest('[data-no-toast]')) || 
                             xhr.getResponseHeader('HX-Redirect') ||
                             xhr.getResponseHeader('X-Skip-Toast');

            if (!skipToast) {
                showToast('success', message);
            }
        }
    });

    document.addEventListener('htmx:afterOnLoad', () => {
        applyInputMasks();
    });

    // Custom Confirmation Modal Logic
    document.addEventListener('htmx:confirm', (event) => {
        const target = event.target;
        if (!target || typeof target.closest !== 'function') return;

        const elt = target.closest('[hx-confirm]');
        if (!elt || !event.detail.question) return;

        event.preventDefault();

        const title = elt.getAttribute('data-confirm-title') || 'Confirmar Ação';
        const originalIssueRequest = event.detail.issueRequest;

        const modifiedIssueRequest = () => {
            const originalHxConfirm = elt.getAttribute('hx-confirm');
            elt.removeAttribute('hx-confirm');
            originalIssueRequest();
            setTimeout(() => {
                if (originalHxConfirm) elt.setAttribute('hx-confirm', originalHxConfirm);
            }, 100);
        };

        window.dispatchEvent(new CustomEvent('open-confirm-modal', {
            detail: {
                title: title,
                question: event.detail.question,
                target: elt,
                issueRequest: modifiedIssueRequest
            }
        }));
    });
}

// These functions need to be outside the initialization check because they might be 
// called by inline event handlers or after HTMX swaps
function applyInputMasks() {
    const cpfInputs = document.querySelectorAll('input[name="cpf"]');
    cpfInputs.forEach(input => {
        if (input.dataset.maskApplied) return;
        input.dataset.maskApplied = 'true';
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
        if (input.dataset.maskApplied) return;
        input.dataset.maskApplied = 'true';
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

function showToast(type, message) {
    console.log(`[Toast] ${type.toUpperCase()}: ${message}`);
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
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
        </button>
    `;

    container.appendChild(toast);
    requestAnimationFrame(() => {
        toast.classList.remove('translate-x-full', 'opacity-0');
    });

    const duration = type === 'error' ? 8000 : 4000;
    setTimeout(() => removeToast(toast), duration);
}

function removeToast(toast) {
    if (!toast || !toast.parentElement) return;
    toast.classList.add('translate-x-full', 'opacity-0');
    setTimeout(() => {
        if (toast.parentElement) toast.remove();
    }, 500);
}

function previewImage(input, previewId) {
    if (input.files && input.files[0]) {
        const file = input.files[0];
        if (file.size > 2 * 1024 * 1024) {
            showToast('error', 'O arquivo é muito grande. O limite é 2MB.');
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

// Initial mask application
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', applyInputMasks);
} else {
    applyInputMasks();
}
