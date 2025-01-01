const navbar = `
    <nav class="fixed top-0 left-0 flex items-center w-full bg-gray-800 text-white h-24 py-6 px-4">
        <ul class="container flex items-center justify-end gap-6 mx-auto">
            <li class="mr-auto">
                <a href="/">
                    <img src="/images/logo-42.svg" alt="Logo of School 42" class="h-10" />
                </a>
            </li>
            <li>
                <a href="/sign-up" class="hover:underline">
                    Sign up
                </a>
            </li>
            <li>
                <a href="/sign-in" class="hover:underline">
                    Sign in
                </a>
            </li>
        </ul>
    </nav>
`;

document.addEventListener('DOMContentLoaded', () => {
    if (routeNeedsAuth() && !isUserAuthenticated()) {
        window.location.href = '/sign-in';
    }

    document.body.insertAdjacentHTML('afterbegin', navbar);
});

window.apiFetch = (path, options) => {
    const userID = localStorage.getItem("X-User-ID");
    const authToken = localStorage.getItem("X-Auth-Token");

    return fetch('http://localhost:3000/api' + path, {
        ...options,
        headers: {
            'Content-Type': 'application/json',
            'X-User-ID': userID,
            'X-Auth-Token': authToken,
            ...(options?.headers || {}),
        },
    });
}

const SKIP_AUTH_PAGES = [
    '/',
    '/sign-up',
    '/sign-in',
];

window.routeNeedsAuth = () => {
    return !SKIP_AUTH_PAGES.includes(window.location.pathname);
}

window.isUserAuthenticated = () => {
    apiFetch('/users/auth', {
        method: 'POST',
    }).then((response) => {
        if (response.status === 200) {
            return true;
        } else {
            return false;
        }
    })
    .catch(() => {
        return false;
    });
}

window.handleFormSubmit = async (event, path, callback) => {
    event.preventDefault();

    const formData = new FormData(event.target);
    const data = Object.fromEntries(formData.entries());
    let isValid = false;

    const response = await apiFetch(path, {
        method: 'POST',
        body: JSON.stringify(data),
    }).then((response) => {
        isValid = response.status >= 200 && response.status <= 299;
        return response.json()
    });

    if (!isValid) {
        setFormError(response.message);
        return;
    }

    setFormError("");

    callback(response);
}

window.setFormError = (message = "") => {
    const formError = document.querySelector('.form-error');
    formError.style.display = message ? 'block' : 'none';
    formError.innerHTML = message;
}
