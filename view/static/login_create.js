var loginBox = document.getElementById("login");
var regBox = document.getElementById("register");
var forgetBox = document.getElementById("forgot");

var loginTab = document.getElementById("lt");
var regTab = document.getElementById("rt");

function regTabFun() {
    event.preventDefault();

    regBox.style.visibility = "visible";
    loginBox.style.visibility = "hidden";
    forgetBox.style.visibility = "hidden";

    regTab.style.backgroundColor = "rgb(12, 132, 189)";
    loginTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";
}
function loginTabFun() {
    event.preventDefault();

    regBox.style.visibility = "hidden";
    loginBox.style.visibility = "visible";
    forgetBox.style.visibility = "hidden";

    loginTab.style.backgroundColor = "rgb(12, 132, 189)";
    regTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";
}
function forTabFun() {
    event.preventDefault();

    regBox.style.visibility = "hidden";
    loginBox.style.visibility = "hidden";
    forgetBox.style.visibility = "visible";

    regTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";
    loginTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";

}

$(document).ready(function () {
    $("#register_button").click(function () {

        var username = document.getElementById("us").value;
        var email = document.getElementById("re").value;
        var password = document.getElementById("rp").value;
        var passwordRetype = document.getElementById("rrp").value;

        if (email == "") {
            $("#message").html("Email required.");
            return;
        }
        else if (password == "") {
            $("#message").html("Password required.");
            return;
        }
        else if (username == "") {
            $("#message").html("Username required.");
            return;
        }
        else if (passwordRetype == "") {
            $("#message").html("Password Retype required.");
            return;
        }
        else if (password != passwordRetype) {
            $("#message").html("Passwords don't match, retype your Password.");
            return;
        }
        else {
            $.ajax({
                url: '/createuser',
                type: 'post',
                dataType: 'json',
                data: { username: username, password: password, email: email },
                success: function (response) {
                    console.log(response);
                    var msg = response.message;
                    if (response.status === "true") {
                        window.location = "/";  // window.location = "/homepage";
                    } else {

                        document.getElementById("rp").value = "";
                        document.getElementById("rrp").value = "";
                        //msg = "Username or email already registered!";
                    }
                    $("#message").html(msg);
                }
            })
        }
    });
    $("#login_button").click(function () {
        var email = document.getElementById("se").value;
        var password = document.getElementById("sp").value;

        if (email == "") {
            $("#message").html("Email required.");
            return;
        }
        else if (password == "") {
            $("#message").html("Password required.");
            return;
        }
        else {
            $.ajax({
                url: '/userlogin',
                type: 'post',
                data: { email: email, password: password },
                dataType: 'json',
                //async: false,
                success: function (response) {
                    console.log(response);
                    //var obj = jQuery.parseJSON(response);
                    var msg = "";
                    //alert(response);

                    if (response.status === true) {
                        window.location = "/";  // window.location = "/homepage";
                    } else {
                        document.getElementById("sp").value = "";
                        msg = "Invalid username or password!";
                    }
                    $("#message").html(msg);
                }
            });
        }
});
});



function forgot() {
    event.preventDefault();

    var email = document.getElementById("fe").value;

    if (emailArray.indexOf(email) == -1) {
        if (email == "") {
            alert("Email required.");
            return;
        }
        alert("Email does not exist.");
        return;
    }

    alert("email is send to your email check it in 24hr. \n Thanks");
    document.getElementById("fe").value = "";
}