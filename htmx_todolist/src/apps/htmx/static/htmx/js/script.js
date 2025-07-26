document.addEventListener('htmx:afterRequest', function (evt) {
    const form = evt.target;
    if (form.tagName === "FORM") {
        form.reset();
    }
})
