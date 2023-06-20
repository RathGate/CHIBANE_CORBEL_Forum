let mobileMenu = document.querySelector(".mobile-menu")

// Function linked to the burger menu icon.
document.querySelector(".nav-btn.menu").addEventListener("click", function() {
    mobileMenu.classList.add("visible")
    document.body.classList.add("no-scroll")
})

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