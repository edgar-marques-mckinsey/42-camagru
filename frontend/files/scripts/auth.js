const PAGES_TO_SKIP_AUTH = [
    '/',
    '/sign-up',
    '/sign-in',
];

window.routeNeedsAuth = () => {
    return !PAGES_TO_SKIP_AUTH.includes(window.location.pathname);
}

window.isUserAuthenticated = async () => {
    return apiFetch('/users/auth', {
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

window.handleLogout = (event) => {
    event.preventDefault();

    localStorage.removeItem('X-User-ID');
    localStorage.removeItem('X-Auth-Token');

    window.location.href = '/';
}
