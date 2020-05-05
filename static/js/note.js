function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, '\\$&');
    var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, ' '));
}

function isValidDate(dateString) {
    var regEx = /^\d{4}-\d{2}-\d{2}$/;
    if (!dateString.match(regEx)) return false;  // Invalid format
    var d = new Date(dateString);
    var dNum = d.getTime();
    if (!dNum && dNum !== 0) return false; // NaN value, Invalid date
    return d.toISOString().slice(0, 10) === dateString;
}

function isValidLang(lang) {
    return ["Java", "MySQL"].indexOf(lang) > -1;
}


var day = getParameterByName('day');
var eDay = document.getElementById("day");
var eLang = document.getElementById("lang");
var eScore = document.getElementById("score");
if (day && !eDay.value) {
    eDay.value = day;
}

function validateForm() {
    var ok = true;
    if (eDay && eDay.value) {
        ok = isValidDate(eDay.value);
        if (!ok) alert("day not ok");
    }
    if (eLang && eLang.value) {
        ok = isValidLang(eLang.value);
        if (!ok) alert("lang not ok");
    }
    if (eScore.value < 1 || eScore.value > 5) {
        ok = false;
        if (!ok) alert("score not ok");
    }
    return ok;

}