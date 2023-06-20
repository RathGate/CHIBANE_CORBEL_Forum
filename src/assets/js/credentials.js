const registerValues = 
    {
    email: {
        inputID: "email",
        errorID: "emailError",
        regex: /^(?=.*[a-zA-Z])[a-zA-Z0-9_-]{3,20}$/ 
    }, 
    username: {
        inputID: "username",
        errorID: "usernameError",
        regex: /^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/ 
    }, 
    password: {
        inputID: "password",
        errorID: "passwordError",
        regex: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&_-])[A-Za-z\d@$!%*?&_-]{8,}$/ 
    }
}
const loginBtn = document.querySelector(".nav-btn.login")
const loginBtnMobile = document.querySelector(".mobile-menu .nav-btn.login")
const registerBtn = document.querySelector(".nav-btn.register")
const registerBtnMobile = document.querySelector(".mobile-menu .nav-btn.register")
const registerForm = document.querySelector("#register-form")
const loginForm = document.querySelector("#login-form")
const loginCtn = document.querySelector(".login-ctn")
const registerCtn = document.querySelector(".register-ctn")

function registerFormSubmitListener(registerFormDiv) {
    if (!registerFormDiv) {
        return
    }
    registerFormDiv.addEventListener("submit", function (e) {
        e.preventDefault()
        if (validateForm()) {
            axios.post('http://localhost:8080/register', {
                username: registerFormDiv.querySelector(".username.input").value,
                password: registerFormDiv.querySelector(".password.input").value,
                email: registerFormDiv.querySelector(".email.input").value,
            }, {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
            }})
                .then(response => {
                if (response.data.status == 200) {
                    location.reload()
                }

                else if (response.data.fields) {
                    
                    response.data.fields.forEach(element => {
                        changeInnerText(element.name, element.errorMsg)
                    })
                }
            });
        }
    })
}

function loginFormSubmitListener(loginFormDiv) {
    if (!loginFormDiv) {
        return
    }

    loginFormDiv.addEventListener("submit", function (e) {
        e.preventDefault()
        axios.post('http://localhost:8080/login', {
            username: loginForm.querySelector(".username.input").value,
            password: loginForm.querySelector(".password.input").value,
        }, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        })
            .then(response => {
                if (response.data.status == 200) {
                    location.reload()
                }
                else if (response.data.fields) {
                    response.data.fields.forEach(element => {
                        changeInnerText(element.name, element.errorMsg)
                    })
                }
            });
    })
}

function toggleForm(formDiv) {
    if (!formDiv) {
        return
    }
    formDiv.style.display = 'flex';
    setTimeout(() => {
        formDiv.classList.add('show');
        document.querySelector(".hide").classList.add("visible")
    }, 0);

    const otherFormDiv = formDiv.classList.contains("register-ctn") ? loginCtn : registerCtn;
    otherFormDiv.classList.remove('show');
    setTimeout(() => {
        otherFormDiv.style.display = 'none';
    }, 300);
}

function handleToggle(btnDiv, formDiv) {
    if (!btnDiv) {
        return
    }
    btnDiv.addEventListener('click', (event) => {
        event.stopPropagation();
        toggleForm(formDiv);
    });
};

const handleClose = (formDiv) => {
    if (!formDiv || !formDiv.querySelector(".close-btn")) {
        return
    }
    const closeBtn = formDiv.querySelector(".close-btn");
    closeBtn.addEventListener('click', () => {
        formDiv.classList.remove('show');
        document.querySelector(".hide").classList.remove("visible")
        formDiv.style.display = 'none';
        const errorElements = document.querySelectorAll('.error');
        errorElements.forEach((error) => {
            error.innerHTML = '';
            error.classList.remove("visible")
        }
        );
    });
};

const handleOutsideClick = (formCtns) => {
    document.addEventListener('click', (event) => {
        console.log(formCtns, event.target)
        if (!registerForm.contains(event.target) && !loginForm.contains(event.target)) {
            formCtns.forEach((formCtn) => {
                formCtn.classList.remove('show');
                document.querySelector(".hide").classList.remove("visible")
                setTimeout(() => {
                    formCtn.style.display = 'none';
                    const errorElements = document.querySelectorAll('.error');
                    errorElements.forEach((error) => {
                        error.innerHTML = '';
                        error.classList.remove("visible")
                    });
                    formCtn.querySelectorAll(".input").forEach(input => {
                        input.value = ""
                    })
                }, 300);
            });
        }
    });
};

const validateField = (value, regex, errorMessage, errorElem) => {
    const isValid = regex.test(value);
    errorElem.textContent = isValid ? '' : errorMessage;
    isValid ? errorElem.classList.remove("visible") : errorElem.classList.add("visible")
    return isValid;
};

const validateInput = (inputClass, regex, errorMessage, errorElemId) => {
    const input = document.querySelector(`input.${inputClass}`);
    const errorElem = document.getElementById(errorElemId);

    const validate = () => validateField(input.value, regex, errorMessage, errorElem);

    input.addEventListener('input', validate);
    return validate;
};

const toggleFormClass = (event) => {
    const formDiv = event.target.classList.contains('register') ? registerCtn : loginCtn;
    toggleForm(formDiv);
};

document.querySelectorAll('.nav-btn.register, .nav-btn.login').forEach((button) => {
    button.addEventListener('click', toggleFormClass);
});



function usernameValidation() {
    return validateInput(
        'username',
        /^(?=.*[a-zA-Z])[a-zA-Z0-9_-]{3,20}$/,
        "Username must be 3-20 char. long. Only characters A-Z, a-z, 0-9, and '_' are allowed.",
        'usernameError'
    )
};

function passwordValidation() {
    return validateInput(
    'password',
    /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&_-])[A-Za-z\d@$!%*?&_-]{8,}$/,
    'Password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character',
    'passwordError'
)};

function emailValidation() {
    return validateInput(
    'email',
    /^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,
    'Invalid email address',
    'emailError'
)};

function validateForm() {
    return usernameValidation() && passwordValidation() && emailValidation();
};

function changeInnerText(divID, content) {
    let div = document.getElementById(registerValues[divID].errorID)
    if (div) {
        div.innerHTML = content
        div.classList.add("visible")
    }
}

if (registerForm && loginForm) {
    registerFormSubmitListener(registerForm)
    loginFormSubmitListener(loginForm)
    
    handleToggle(registerBtn, registerCtn);
    handleToggle(loginBtn, loginCtn);

    handleToggle(loginBtnMobile, loginCtn);
    handleToggle(registerBtnMobile, registerCtn);

    handleClose(registerCtn);
    handleClose(loginCtn);

    handleOutsideClick([loginCtn, registerCtn]);
}