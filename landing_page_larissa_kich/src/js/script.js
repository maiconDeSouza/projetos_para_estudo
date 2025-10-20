const buttonMobile = document.querySelector('.btn-mobile')

buttonMobile.addEventListener('click', e => {
    const navbarMobile = document.querySelector('#id-navbar-mobile')
    const icon = buttonMobile.querySelector('i')

    const classIcon = icon.classList.value


    if(classIcon === 'fa-solid fa-bars'){
        icon.classList.replace('fa-bars', 'fa-xmark')
    }

    if(classIcon === 'fa-solid fa-xmark'){
        icon.classList.replace('fa-xmark', 'fa-bars')
    }

    navbarMobile.classList.toggle('on')

})

