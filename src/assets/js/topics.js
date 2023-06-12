let defaultValues = { "order": "newest", "date": 7, "page": 1 }

let dateSelect = document.querySelector("select#date")
let orderSelect = document.querySelector("select#order")
let pageRadio = document.querySelector('input[name="page"]')
let allFilterInputs = document.querySelectorAll(".filter")

allFilterInputs.forEach(element => {
    element.addEventListener("change", function () {
        updateQueryStringParameter(element.name, element.value)
        refreshTopics()
    })
})

function addPaginationListeners() {
    let pageInputs = document.querySelectorAll("input[type='radio'].filter")
    pageInputs.forEach(element => {
        element.addEventListener("change", function () {
            updateQueryStringParameter(element.name, element.value)
            refreshTopics()
        })
    })
}

function refreshTopics(pageIncrementation = 0) {
    console.log(getPageValue())
    axios.post('http://localhost:8080/topics',
        {
            timePeriod: dateSelect.value,
            order: orderSelect.value,
            page: getPageValue()
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

function getPageValue() {
    let temp = document.querySelector('input[name="page"]:checked')

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
            console.log("existing key", key, "found")
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
    console.log(".....\nENTERING REMOVEPARAM FOR", key)
    var rtn = sourceURL.split("?")[0],
        param,
        params_arr = [],
        queryString = (sourceURL.indexOf("?") !== -1) ? sourceURL.split("?")[1] : "";
    console.log("query string:", queryString)
    if (queryString !== "") {
        params_arr = queryString.split("&");
        console.log("parameters:", params_arr)
        for (var i = params_arr.length - 1; i >= 0; i -= 1) {
            param = params_arr[i].split("=")[0];
            console.log("parameter", key,  "is default:", key === param)
            if (param === key) {
                params_arr.splice(i, 1);
                console.log(params_arr)
            }
        }
        if (params_arr.length) rtn = rtn + "?" + params_arr.join("&");
    }
    console.log(".....EXITING REMOVEPARM FOR", key)
    console.log("Results:",rtn)
    return rtn;
}

function checkURL(url = window.location.href, replaceURL=true) {
    let newurl = url.slice()

    Array.from(Object.keys(defaultValues)).forEach(element => {
        var re = new RegExp("([?&])" + element + "=" + defaultValues[element] + "(&|#|$)", "i");
        if (re.test(url)) {
            console.log(element, "should be deleted. Sending to removeParam")
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