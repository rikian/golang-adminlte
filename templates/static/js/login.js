"use strict"

function qs(elm) {
    return document.querySelector(elm)
}

function ce(elm) {
    return document.createElement(elm)
}

var xml = new XMLHttpRequest

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

// login section
var email  = document.querySelector("#email_login")
var password = document.querySelector("#password_login")
var btnLogin = document.querySelector("#btn-login")

btnLogin.addEventListener("click", function() {
    var data = JSON.stringify({
        user_email : email.value,
        user_password : password.value
    })
    login(data)
})

function login(data) {
    xml.open("POST", window.location.origin + "/login", true)
    xml.setRequestHeader("content-type", "application/json")
    xml.send(data)
}

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
        "user_name":regisFullName.value,
        "user_email":regisEmail.value,
        "user_password_1":regisPassword1.value,
        "user_password_2":regisPassword2.value,
        "user_term":regisTerm.checked
    })

    console.log(formRegister)

    xml.open("POST", window.location.origin + "/register", true)
    xml.setRequestHeader("content-type", "application/json")
    xml.send(formRegister)
})

// Ajax response
xml.onload = function() {
    try {
        var message = JSON.parse(this.responseText)
        // login
        if (message && message["method"] && message["method"] === "auth" && message["status"] === "ok") return window.location = window.location.origin
        // register
        if (message && message["method"] && message["method"] === "register") {
            console.log(message)
        }
    } catch (error) {
        console.log(this.responseText)
        // window.location.reload()
    }
}