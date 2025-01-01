import "./auth.js";
import "./form.js";
import "./utils.js";

import "./components/index.js";

document.addEventListener('DOMContentLoaded', async () => {
    const isAuth = await window.isUserAuthenticated();
    if (routeNeedsAuth() && !isAuth) {
        window.location.href = '/sign-in';
    }

    const navbar = generateNavbar(isAuth)
    document.body.insertAdjacentHTML('afterbegin', navbar);
});
