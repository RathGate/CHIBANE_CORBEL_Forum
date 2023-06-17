var registerValues = 
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

let registerForm = document.querySelector("#register-form")
if (registerForm) {
    registerForm.addEventListener("submit", function (e) {
        e.preventDefault()
        if (validateForm()) {
            axios.post('http://localhost:8080/register', {
                username: document.getElementById(registerValues.username.inputID).value,
                password: document.getElementById(registerValues.password.inputID).value,
                email: document.getElementById(registerValues.email.inputID).value,
                }, {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
            }})
                .then(response => {
                console.log(response)
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

let loginForm = document.querySelector("#login-form")
if (loginForm) {
    loginForm.addEventListener("submit", function (e) {
        e.preventDefault()
        axios.post('http://localhost:8080/login', {
            username: loginForm.querySelector("#username").value,
            password: loginForm.querySelector("#password").value,
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


let test
const toggleForm = (formClass) => {
    const formCtn = document.querySelector(formClass);
    formCtn.style.display = 'flex';
    setTimeout(() => {
        formCtn.classList.add('show');
        document.querySelector(".hide").classList.add("visible")
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
        document.querySelector(".hide").classList.remove("visible")
        formCtn.style.display = 'none';
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
        const registerForm = document.querySelector('.register-form');
        const loginForm = document.querySelector('.login-form');

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
    "Username must be 3-20 char. long. Only characters A-Z, a-z, 0-9, and '_' are allowed.",
    'usernameError'
);

const passwordValidation = validateInput(
    'password',
    /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&_-])[A-Za-z\d@$!%*?&_-]{8,}$/,
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

function changeInnerText(divID, content) {
    
    let div = document.getElementById(registerValues[divID].errorID)
    console.log(div)
    if (div) {
        div.innerHTML = content
        div.classList.add("visible")
    }
}