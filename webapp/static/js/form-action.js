function setForm(method, action) {
    document.getElementById("form").action = action
    document.getElementById("form").method = method
    document.getElementById("form").submit()
}

(function(window, document, undefined){
    window.onload = init;
    function init(){
        searchBar =  document.getElementById("search")
        if (searchBar != null) {
            searchBar.addEventListener("keydown", function(event) {
            if (event.keyCode == 13) {
                location.href = "/papers?tag=" + searchBar.value
            }
            })
        }
    }
})(window, document, undefined);