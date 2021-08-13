var t = document.getElementById("code")
t.innerHTML = "123123"

var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    lineNumbers: true,
    styleActiveLine: true,
    matchBrackets: true,
});
var input = document.getElementById("select");

// setDefaultTheme("abbott")

function setDefaultTheme(themeName){
   editor.setOption("theme",themeName)
    location.hash = "#" + themeName
}

function selectTheme() {
    var theme = input.options[input.selectedIndex].textContent;
    editor.setOption("theme", theme);
    location.hash = "#" + theme;
}

var choice = (location.hash && location.hash.slice(1)) ||
    (document.location.search &&
        decodeURIComponent(document.location.search.slice(1)));
if (choice) {
    input.value = choice;
    editor.setOption("theme", choice);
}
CodeMirror.on(window, "hashchange", function () {
    var theme = location.hash.slice(1);
    if (theme) {
        input.value = theme;
        selectTheme();
    }
});