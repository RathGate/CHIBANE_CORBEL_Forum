document.querySelector('.register-btn').addEventListener('click', function() {
    const registerCtn = document.querySelector('.register-ctn');
    registerCtn.style.display = 'flex';
    setTimeout(function() {
        registerCtn.classList.add('show');
    }, 0);
});

const closeBtn = document.querySelector('.close-btn');
closeBtn.addEventListener('click', function() {
    const registerCtn = document.querySelector('.register-ctn');
    registerCtn.classList.remove('show');
    setTimeout(function() {
        registerCtn.style.display = 'none';
    }, 300);
});