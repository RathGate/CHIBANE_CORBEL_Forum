const topicForm = document.querySelector("form#new-tp-form")
const mandatoryInputs = document.querySelectorAll(".new.mandatory")
const errorDiv = document.querySelector(".a-error")

if (topicForm) {
    topicForm.addEventListener("submit", function (e) {
        e.preventDefault()
        if (!document.querySelector(`select[name="category"]`).value ||
            typeof parseInt(document.querySelector(`select[name="category"]`).value) != "number") {
            errorDiv.innerHTML = "You must choose a valid category"
            return
        }
        if (!(/\S/.test(document.querySelector(`input[name="title"]`).value)) ||
            !(/\S/.test(document.querySelector(`textarea[name="post"]`).value))) {
            errorDiv.innerHTML = "Title and content must both have at least one non white-space character"
            return
        }
        axios.post("http://localhost:8080/topic/new", {
            category_id: document.querySelector(`select[name="category"]`).value,
            title: document.querySelector(`input[name="title"]`).value,
            content: document.querySelector(`textarea[name="post"]`).value,
        }, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
        }})
        .then(response => {
            console.log(response)
            if (response.data.status == 200) {
                window.location.href = `http://localhost:8080/topic/${response.data.topic_id}`
            } else {
                errorDiv.innerHTML = response.data.error_msg
            }
    });
    })
}

mandatoryInputs.forEach(element => {
    element.addEventListener("input", function () {
        errorDiv.innerHTML = ""
    })
})