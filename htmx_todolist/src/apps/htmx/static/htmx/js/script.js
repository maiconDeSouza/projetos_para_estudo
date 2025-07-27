document.addEventListener('htmx:afterRequest', function (evt) {
    const form = evt.target;
    if (form.tagName === "FORM") {
        form.reset();
    }
})



const filterButtons = document.querySelectorAll('#task-filter button')

filterButtons.forEach(btn => {
  btn.addEventListener('click', () => {
 
    document.querySelector('#task-filter button.active')?.classList.remove('active')
    
    btn.classList.add('active')
  })
})


document.body.addEventListener('htmx:configRequest', e => {
  const token = document.querySelector('meta[name="csrf-token"]').content;
  e.detail.headers['X-CSRFToken'] = token;
})