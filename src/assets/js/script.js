let mobileMenu = document.querySelector(".mobile-menu")
let mobileMenuBtn = document.querySelector(".nav-btn.menu")
let profileMenu = document.querySelector(".smol-profile")
let profileMenuBtn = document.querySelector(".nav-profile")
let profileMenuBtnMobile = document.querySelector(".nav-btn.profile")

// Function linked to the burger menu icon.
mobileMenuBtn.addEventListener("click", function() {
    mobileMenu.classList.add("visible")
    document.body.classList.add("no-scroll")
})
mobileMenu.querySelectorAll(".btn").forEach(element => { 
    element.addEventListener("click", function () {
        mobileMenu.classList.remove("visible")
        document.body.classList.remove("no-scroll")
    })
})
if (profileMenuBtn) {
    
    profileMenuBtn.addEventListener("click", function () {
        profileMenu.classList.toggle("visible")
    })
    profileMenuBtnMobile.addEventListener("click", function () {
        profileMenu.classList.toggle("visible")
    })
    CloseWhenOutside(profileMenu, profileMenuBtn, profileMenuBtnMobile)
}
// Removes the visible attribute when resizing over the max size of 
// the burger menu.
window.addEventListener("resize", function() {
    if (this.innerWidth > 992) {
        mobileMenu.classList.remove("visible")
        document.body.classList.remove("no-scroll")
    } 
})

// Removes the `className` class from all elements in the `divArr` array.
function removeClassFromAll(divArr, className) {
    if (divArr.length == 0) {
        return
    }
    divArr.forEach(div => {
        div.classList.remove(className)
    })
}

mobileMenu.querySelector(".close-btn").addEventListener("click", function () {
    mobileMenu.classList.remove("visible")
    document.body.classList.remove("no-scroll")
})

function CloseWhenOutside(div, btn, btn2=null) {
    document.addEventListener('click', (event) => {
            if (!div.contains(event.target) && !btn.contains(event.target) && (btn2 && !btn2.contains(event.target))) {
                div.classList.remove('visible');
                // document.querySelector(".hide").classList.remove("visible")
                document.body.classList.remove("no-scroll")
            }

    });
};
function CloseWhenOutsideWithoutBtn(div) {
    document.addEventListener('click', (event) => {
        if (!div.contains(event.target)) {
            div.classList.remove('visible');
        }
    });
};
CloseWhenOutside(mobileMenu, mobileMenuBtn)