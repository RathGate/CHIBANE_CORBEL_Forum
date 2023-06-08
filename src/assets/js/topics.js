let dateSelect = document.querySelector("select#date")
let orderSelect = document.querySelector("select#order-by")

dateSelect.addEventListener("change", function () {
    updateQueryStringParameter("date", dateSelect.value)
    axios.post('http://localhost:8080/topics',
        { timePeriod: dateSelect.value, order: orderSelect.value }, {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    }
    )
    .then(response => {
        if (response.status == 200) {
            document.querySelector("#test").innerHTML = response.data
        }
    });
})
orderSelect.addEventListener("change", function () {
    updateQueryStringParameter("order", orderSelect.value)
    axios.post('http://localhost:8080/topics',
        { timePeriod: dateSelect.value, order: orderSelect.value }, {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    }
    )
    .then(response => {
        if (response.status == 200) {
            document.querySelector("#test").innerHTML = response.data
        }
    });
})

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
    window.history.pushState({path:newurl},'',newurl);
}