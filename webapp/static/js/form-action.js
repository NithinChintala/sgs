function setForm(method, action) {
    document.getElementById("form").action = action
    document.getElementById("form").method = method
    document.getElementById("form").submit()
}
