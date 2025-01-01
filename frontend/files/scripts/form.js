window.handleFormSubmit = async (event, path, method, callback) => {
    event.preventDefault();

    if (typeof method === 'function') {
        callback = method;
        method = 'POST';
    }

    const formData = new FormData(event.target);
    const data = Object.fromEntries(formData.entries());
    let isValid = false;

    const response = await apiFetch(path, {
        method,
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
