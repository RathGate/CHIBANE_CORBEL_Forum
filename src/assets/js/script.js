const toggleForm = (formClass) => {
    const formCtn = document.querySelector(formClass);
    formCtn.style.display = 'flex';
    setTimeout(() => {
        formCtn.classList.add('show');
    }, 0);

    const otherFormClass = formClass === '.register-ctn' ? '.login-ctn' : '.register-ctn';
    const otherFormCtn = document.querySelector(otherFormClass);
    otherFormCtn.classList.remove('show');
    setTimeout(() => {
        otherFormCtn.style.display = 'none';
    }, 300);
};

const handleToggle = (buttonClass, formClass) => {
    document.querySelector(buttonClass).addEventListener('click', (event) => {
        event.stopPropagation();
        toggleForm(formClass);
    });
};

const handleClose = (closeBtnClass, formClass) => {
    const closeBtn = document.querySelector(closeBtnClass);
    closeBtn.addEventListener('click', () => {
        const formCtn = document.querySelector(formClass);
        formCtn.classList.remove('show');
        formCtn.style.display = 'none';
        const errorElements = document.querySelectorAll('.error');
        errorElements.forEach((error) => {
            error.innerHTML = '';
        }
        );
    });
};

const handleOutsideClick = (formCtns) => {
    document.addEventListener('click', (event) => {
        const registerForm = document.querySelector('.register-form');
        const loginForm = document.querySelector('.login-form');

        if (!registerForm.contains(event.target) && !loginForm.contains(event.target)) {
            formCtns.forEach((formCtn) => {
                formCtn.classList.remove('show');
                setTimeout(() => {
                    formCtn.style.display = 'none';
                    const errorElements = document.querySelectorAll('.error');
                    errorElements.forEach((error) => {
                        error.innerHTML = '';
                    });
                }, 300);
            });
        }
    });
};



const validateField = (value, regex, errorMessage, errorElem) => {
    const isValid = regex.test(value);
    errorElem.textContent = isValid ? '' : errorMessage;
    errorElem.style.display = isValid ? 'none' : 'flex';
    return isValid;
};

const validateInput = (inputId, regex, errorMessage, errorElemId) => {
    const input = document.getElementById(inputId);
    const errorElem = document.getElementById(errorElemId);

    const validate = () => validateField(input.value, regex, errorMessage, errorElem);

    input.addEventListener('input', validate);
    return validate;
};

const toggleFormClass = (event) => {
    const formClass = event.target.classList.contains('register') ? '.register-ctn' : '.login-ctn';
    toggleForm(formClass);
};

document.querySelectorAll('.creds-btn.register, .creds-btn.login').forEach((button) => {
    button.addEventListener('click', toggleFormClass);
});

handleToggle('.creds-btn.register', '.register-ctn');
handleToggle('.creds-btn.login', '.login-ctn');

handleClose('.register-ctn .close-btn', '.register-ctn');
handleClose('.login-ctn .close-btn', '.login-ctn');

const formCtns = document.querySelectorAll('.register-ctn, .login-ctn');
handleOutsideClick(formCtns);

const usernameValidation = validateInput(
    'username',
    /^(?=.*[a-zA-Z])[a-zA-Z0-9_-]{3,20}$/,
    "Username is not valid. Only characters A-Z, a-z, 0-9, and '_' are allowed, and must be 3-20 characters long.",
    'usernameError'
);

const passwordValidation = validateInput(
    'password',
    /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/,
    'Password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character',
    'passwordError'
);

const emailValidation = validateInput(
    'email',
    /^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,
    'Invalid email address',
    'emailError'
);

const validateForm = () => {
    return usernameValidation() && passwordValidation() && emailValidation();
};
