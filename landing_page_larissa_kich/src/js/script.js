const buttonMobile = document.querySelector('.btn-mobile')
const navbarMobile = document.querySelector('#id-navbar-mobile')
const icon = buttonMobile.querySelector('i')


buttonMobile.addEventListener('click', e => {
    navbarMobile.classList.toggle('on')
    const isNavbarOpen = navbarMobile.classList.contains('on')

    icon.classList.toggle('fa-bars', !isNavbarOpen)
    icon.classList.toggle('fa-xmark', isNavbarOpen)
})

