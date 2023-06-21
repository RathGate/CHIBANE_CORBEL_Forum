let defaultValues = { "order": "newest", "date": "all", "page": 1, category: "all" }

let dateSelect = document.querySelector("select#date")
let orderSelect = document.querySelector("select#order")
let pageRadio = document.querySelector('input[name="page"]')
let allFilterInputs = document.querySelectorAll(".filter")
let currentCategory

allFilterInputs.forEach(element => {
    element.addEventListener("change", function () {
        updateQueryStringParameter(element.name, element.value)
        refreshTopics()
    })
})

function urlContainsMe() {
    const urlparam = new URLSearchParams(window.location.search)
    return (urlparam.has("me") && !urlparam.get('me'))
}


function addPaginationListeners() {
    let pageInputs = document.querySelectorAll('input[name="page"]')
    pageInputs.forEach(element => {
        element.addEventListener("change", function () {
            console.log(element, "changed! Refreshing")
            updateQueryStringParameter(element.name, element.value)
            refreshTopics()
        })
    })
}

function refreshTopics(pageIncrementation = 0) {
    axios.post('http://localhost:8080/topics',
        {
            timePeriod: dateSelect.value,
            order: orderSelect.value,
            page: getInputValueByName("page"),
            category: retrieveCategory(),
            useuserid: urlContainsMe()
        }, {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }})
    .then(response => {
        if (response.status == 200) {
            document.querySelector("#test").innerHTML = response.data
            addPaginationListeners()
        }
    });
}

function getInputValueByName(inputName) {
    let temp = document.querySelector(`input[name="${inputName}"]:checked`)

    if (!temp) {
        return 1
    }
    temp = parseInt(temp.value)
    if (typeof temp === NaN) {
        return 1
    }
    return temp
}

function updateQueryStringParameter(key, value) {
    let uri = window.location.search
    let newurl;
    var re = new RegExp("([?&])" + key + "=.*?(&|#|$)", "i");
    if( value === undefined ) {
        if (uri.match(re)) {
            newurl = uri.replace(re, '$1$2').replace(/[?&]$/, '').replaceAll(/([?&])&+/g, '$1').replace(/[?&]#/, '#');
        } else {
            newurl = uri;
        }
    } else {
        if (uri.match(re)) {
            newurl = uri.replace(re, '$1' + key + "=" + value + '$2');
        } else {
            var hash =  '';
            if( uri.indexOf('#') !== -1 ){
                hash = uri.replace(/.*#/, '#');
                uri = uri.replace(/#.*/, '');
            }
            var separator = uri.indexOf('?') !== -1 ? "&" : "?";    
            newurl = uri + separator + key + "=" + value + hash;
        }
    } 
    newurl = window.location.href.split('?')[0] + checkURL(newurl, false)
    window.history.pushState({path:newurl},'',newurl);
}

function removeParam(key, sourceURL) {
    var rtn = sourceURL.split("?")[0],
        param,
        params_arr = [],
        queryString = (sourceURL.indexOf("?") !== -1) ? sourceURL.split("?")[1] : "";
    if (queryString !== "") {
        params_arr = queryString.split("&");
        for (var i = params_arr.length - 1; i >= 0; i -= 1) {
            param = params_arr[i].split("=")[0];
            if (param === key) {
                params_arr.splice(i, 1);
            }
        }
        if (params_arr.length) rtn = rtn + "?" + params_arr.join("&");
    }
    return rtn;
}

function checkURL(url = window.location.href, replaceURL=true) {
    let newurl = url.slice()

    Array.from(Object.keys(defaultValues)).forEach(element => {
        var re = new RegExp("([?&])" + element + "=" + defaultValues[element] + "(&|#|$)", "i");
        if (re.test(url)) {
            newurl = removeParam(element, newurl)
        }
    })

    if (!replaceURL) {
        return newurl
    }

    if (newurl != url) {
        window.history.pushState({path:newurl},'',newurl);
    }
}

checkURL()

function retrieveCategory() {
    const urlParams = new URLSearchParams(window.location.search);
    let categoryParam = urlParams.get('category');
    if (!categoryParam) {
        categoryParam = 0
    }
    return categoryParam
}
// const popupContainer = document.getElementById("pop-up-container");
// const newTopicBtn = document.querySelector(".btn-new-topic");
// const closePopupBtn = document.querySelector("#pop-up-container .pop-up-close-btn");
// const popupForm = document.querySelector("#pop-up-container form");

// function togglePopup() {
//     if (popupContainer.classList.contains("show")) {
//         popupContainer.classList.remove("show");
//         setTimeout(() => {
//             popupContainer.style.display = "none";
//         }, 300);
//     } else {
//         popupContainer.style.display = "flex";
//         setTimeout(() => {
//             popupContainer.classList.add("show");
//         }, 0);
//     }
// }

// function closePopup() {
//     popupContainer.classList.remove("show");
//     setTimeout(() => {
//         popupContainer.style.display = "none";
//     }, 300);
// }

// function handleNewTopicOutsideClick(event) {
//     if (event.target === popupContainer) {
//         closePopup();
//     }
// }

// function handleNewTopicClick() {
//     togglePopup();
// }

// function handleNewTopicClose() {
//     closePopup();
// }

// if (newTopicBtn && popupContainer && closePopupBtn && popupForm) {
//     newTopicBtn.addEventListener("click", handleNewTopicClick);
//     closePopupBtn.addEventListener("click", handleNewTopicClose);
//     document.addEventListener("click", handleNewTopicOutsideClick);
// }

