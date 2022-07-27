"use strict"

function qs(elm) {
    return document.querySelector(elm)
}

function ce(elm) {
    return document.createElement(elm)
}

var email  = document.querySelector("#email_login")
var password = document.querySelector("#password_login")
var btnLogin = document.querySelector("#btn-login")

var xml = new XMLHttpRequest

console.log(window.location.origin)

btnLogin.addEventListener("click", function() {
    xml.open("POST", window.location.origin + "/login", true)
    xml.setRequestHeader("content-type", "application/json")
    xml.send(JSON.stringify({
        user_email : email.value,
        user_password : password.value
    }))
})

xml.onload = function() {
    try {
        var message = JSON.parse(this.responseText)
        console.log(message)
        // if (message && message["status"] === "ok") return window.location.reload()
    } catch (error) {
        console.log(message)
        // window.location.reload()
    }
}

// swicth login register
var loginBox = document.querySelector("#login-box")
var registerBox = document.querySelector("#register-box")
registerBox.style.display = "none"
var formRegister = document.querySelector("#form-register")
var formLogin = document.querySelector("#form-login")

formRegister.addEventListener("click", function() {
    registerBox.style.display = "block"
    loginBox.style.display = "none"
})

formLogin.addEventListener("click", function() {
    registerBox.style.display = "none"
    loginBox.style.display = "block"
})

// register section
var regisFullName = qs("#regis-fullname")
var regisEmail = qs("#regis-email")
var regisPassword1 = qs("#regis-password1")
var regisPassword2 = qs("#regis-password2")
var regisTerm = qs("#regis-term")
var btnReset = qs("#regis-reset")
var btnRegister = qs("#regis-register")

btnRegister.addEventListener("click", function() {
    if (!regisFullName.value || !regisEmail.value || !regisPassword1.value || !regisPassword2.value || regisPassword1.value !== regisPassword2.value || !regisTerm.checked) return alert("Please corect your input")
    var formRegister = JSON.stringify({
        "full_name":regisFullName.value,
        "email":regisEmail.value,
        "password_1":regisPassword1.value,
        "password_2":regisPassword2.value,
        "term":regisTerm.checked
    })

    xml.open("POST", window.location.origin + "/register", true)
    xml.setRequestHeader("content-type", "application/json")
    xml.send(formRegister)
})