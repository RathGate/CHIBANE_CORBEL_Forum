function stringIsValid(string) {
    return (/\S/.test(string))
}

let sendInput = document.querySelector("textarea.answer")
const sendBtn = document.querySelector(".send-btn")
const errorDiv = document.querySelector(".a-error")
if (sendBtn && sendInput) {
    sendBtn.addEventListener("click", function () {
        if (checkStringValidity()) {
            sendAnswer(sendInput.value)
        }
    })
    sendInput.addEventListener("input", function () {
        errorDiv.innerHTML = ""
    })
}

function checkStringValidity() {
    if (!stringIsValid(sendInput.value)) {
        errorDiv.innerHTML = "Your answer must contain at least one valid non white-space character"
        return false
    }
    return true
}

function sendAnswer(string) {
    let url = "http://localhost:8080" + window.location.pathname
    axios.post(url, {
        content: string,
    }, {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
    }})
    .then(response => {
            console.log(response)
            if (response.data.status == 200) {
                location.reload()
            } else {
                errorDiv.innerHTML = response.data.error_msg
            }
    });
}