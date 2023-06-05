function toggleForm(formClass) {
    const formCtn = document.querySelector(formClass);
    formCtn.style.display = 'flex';
    setTimeout(function() {
        formCtn.classList.add('show');
    }, 0);

    const otherFormClass = formClass === '.register-ctn' ? '.login-ctn' : '.register-ctn';
    const otherFormCtn = document.querySelector(otherFormClass);
    otherFormCtn.classList.remove('show');
    setTimeout(function() {
        otherFormCtn.style.display = 'none';
    }, 300);
}

document.querySelector('.creds-btn.register').addEventListener('click', function() {
    toggleForm('.register-ctn');
});

document.querySelector('.creds-btn.login').addEventListener('click', function() {
    toggleForm('.login-ctn');
});

const closeBtnRegister = document.querySelector('.register-ctn .close-btn');
closeBtnRegister.addEventListener('click', function() {
    const registerCtn = document.querySelector('.register-ctn');
    registerCtn.classList.remove('show');
    setTimeout(function() {
        registerCtn.style.display = 'none';
    }, 300);
});

const closeBtnLogin = document.querySelector('.login-ctn .close-btn');
closeBtnLogin.addEventListener('click', function() {
    const loginCtn = document.querySelector('.login-ctn');
    loginCtn.classList.remove('show');
    setTimeout(function() {
        loginCtn.style.display = 'none';
    }, 300);
});

/*function validateUsername(username) {
    const usernameRegex = /^(?=.*[a-zA-Z])[a-z0-9_-]{3,20}$/;
    if (!usernameRegex.test(username)) {
        return "Your username is not valid. Only characters A-Z, a-z, 0-9 and '-' are valid, and must contain at least";
    }

    return "";
}*/
function validatePassword(password) {
    if (password.length < 8) {
        return "Password must be at least 8 characters long";
    }

    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    if (!passwordRegex.test(password)) {
        return "Password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character";
    }

    return "";
}

function validateEmail(email) {
    const emailRegex = /^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
    if (!emailRegex.test(email)) {
        return "Invalid email address";
    }

    return "";
}

function passwordValidation() {
    const password = document.getElementById("password").value;
    const passwordError = validatePassword(password);
    const passwordErrorElem = document.getElementById("passwordError");

    if (passwordError) {
        passwordErrorElem.textContent = passwordError;
        passwordErrorElem.style.display = "flex";
        return false;
    }

    passwordErrorElem.style.display = "none";
    return true;
}

function emailValidation() {
    const email = document.getElementById("email").value;
    const emailError = validateEmail(email);
    const emailErrorElem = document.getElementById("emailError");

    if (emailError) {
        emailErrorElem.textContent = emailError;
        emailErrorElem.style.display = "flex";
        return false;
    }

    emailErrorElem.style.display = "none";
    return true;
}